package main

import (
	"context"
	"github.com/dearcj/golangproj/msutil"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"

	//"context"
	ws "github.com/gorilla/websocket"
	"go.opencensus.io/stats"

	//	"go.opencensus.io/stats"

	//	"strconv"
	"sync"
	"sync/atomic"
	"time"
	//	"unsafe"

	pb "github.com/dearcj/golangproj/network"
	"github.com/golang/protobuf/proto"
)

type Connection struct {
	msutil.XServerDataMsg

	msgsToWrite      []*[]byte
	msgsReaded       [][]byte
	queue            *QueuePlayer
	ReadMsg          pb.ClientData
	session          *Session
	conn             *ws.Conn
	ping             bool
	doRemove         bool
	removeAfterWrite int32
	send             chan []byte
	lastSend         time.Time
	id               int
	delay            int
	lastMessage      int
	mu               sync.Mutex
	maxDelay         time.Duration
}

type ConnectionRedirect struct {
	con        *Connection
	DestRoomId int
	Pos        Vec2
}

func (c *Connection) CreateXServerDataMsg(sesId uuid.UUID) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.XServerDataMsg = msutil.CreateXServerDataMsg(sesId.Bytes())
}

func (c *Connection) SetSession(s *Session) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.UniqueID = s.id.Bytes()
	c.session = s
}

func (con *Connection) GetSession() *Session {
	return con.session
}

func (con *Connection) parseCommands(commands *[]*pb.Command) {
	ses := con.GetSession()
	if ses == nil {
		return
	}
	for _, c := range *commands {
		ses.ParseCommand(MapID(c.CommandId), c.Params)
	}
}

func (con *Connection) applyMsg() {
	con.mu.Lock()
	msgs := con.msgsReaded
	con.msgsReaded = make([][]byte, 0)
	con.mu.Unlock()

	for _, msg := range msgs {
		err := proto.Unmarshal(msg, &con.ReadMsg)
		if err != nil {
			server.logger.Error("Can't unmarshal proto message", zap.Error(err))
			return
		}

		commands := con.ReadMsg.Commands
		now := server.loopStartTime
		con.delay = now.Nanosecond() - con.lastMessage
		con.lastMessage = now.Nanosecond()
		con.parseCommands(&commands)
	}
}

func (c *Connection) writeServerDataAndReset() {
	c.writeServerData(c.WriteToMsg(), false)
	c.Reset()
}

func (c *Connection) write(ping bool) {
	if c.conn == nil { //fake Connection for bots for example
		return
	}

	defer c.mu.Unlock()

	c.mu.Lock()
	t := time.Now()

	delay := t.Sub(c.lastSend)
	if c.maxDelay < delay {
		c.maxDelay = delay
	}

	c.lastSend = t
	if len(c.msgsToWrite) == 0 && ping == false {
		return
	}

	c.conn.SetWriteDeadline(t.Add(nodeConfig.Server.WriteWait))

	if ping {
		c.conn.WriteMessage(ws.PingMessage, []byte{})
	}

	for x := len(c.msgsToWrite) - 1; x >= 0; x-- {
		c.conn.SetWriteDeadline(t.Add(nodeConfig.Server.WriteWait))
		w, err := c.conn.NextWriter(ws.BinaryMessage)
		if err != nil {
			return
		}

		stats.Record(context.Background(),
			server.stats.Metrics.PacketSize.M(int64(len(*c.msgsToWrite[x]))))

		w.Write(*c.msgsToWrite[x])
		w.Close()

	}
	c.msgsToWrite = []*[]byte{}
	atomic.AddInt64(&server.stats.Metrics.SocketWriteCounter, 1)
	if atomic.LoadInt32(&c.removeAfterWrite) == 1 {
		server.SafeRemCon(c, CON_REM_BYE)
	}
}

func (c *Connection) writeServerData(serverData *pb.ServerData, lastMessage bool) error {
	strTotal, err := proto.Marshal(serverData)
	if err != nil {
		server.logger.Error("Can't marshal protobuff message", zap.Error(err))
		return err
	} else {
		c.mu.Lock()
		c.msgsToWrite = append(c.msgsToWrite, &strTotal)
		c.mu.Unlock()

		if lastMessage {
			atomic.StoreInt32(&c.removeAfterWrite, 1)
		}

		return nil
	}
}

func (c *Connection) read() {
	defer func() {
		c.conn.Close()
	}()

	for {
		if c.doRemove {
			return
		}

		t := time.Now()
		_, message, err := c.conn.ReadMessage()
		atomic.AddInt64(&server.stats.Metrics.SocketReadCounter, 1)

		if err != nil {
			if ws.IsUnexpectedCloseError(err, ws.CloseGoingAway) {
				server.logger.Info("Connection closed by client")
			}

			server.SafeRemCon(c, CON_REM_DISCONNECT)
			return
		}

		if message != nil {
			c.conn.SetReadDeadline(t.Add(nodeConfig.Server.PongWait))
			c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(t.Add(nodeConfig.Server.PongWait)); return nil })
			c.mu.Lock()
			c.msgsReaded = append(c.msgsReaded, message)
			c.mu.Unlock()
		}
	}
}
