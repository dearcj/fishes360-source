package main

import (
	"context"
	"fmt"
	"github.com/dearcj/golangproj/analytics"
	"github.com/dearcj/golangproj/asyncutils"
	"github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
	"github.com/gofrs/uuid"
	ws "github.com/gorilla/websocket"
	"go.opencensus.io/stats"
	"go.uber.org/zap"
	"math"
	"time"

	"math/rand"
	"sync"
)

type QueuePlayer struct {
	connection  *Connection
	startSearch time.Time //time when player added to queue
}

func (q *QueuePlayer) remCon(con *Connection) {
	q.connection = nil
	q.sendActualQueue()
}

func (q *QueuePlayer) sendActualQueue() {
	//for _, c := range q.connections {
	//	c.writeServerData(&pb.ServerData{QueueData: &q.QueueData}, false)
	//}
}

func (q *QueuePlayer) addCon(c *Connection) {
	c.queue = q
	q.connection = c
	q.sendActualQueue()
}

func CreateQueueTeam(con *Connection) *QueuePlayer {
	q := &QueuePlayer{
		startSearch: time.Now(),
	}

	q.addCon(con)

	return q
}

type (
	ConnectionRemove struct {
		Connection *Connection
		Reason     ConnectionRemoveReason
	}

	NewSession struct {
		Session *Session
		Account *Account
	}

	Server struct {
		RunMutex         sync.RWMutex
		Seed             int64
		logger           *zap.Logger
		tickNum          uint64
		lastPing         time.Time
		analytics        analytics.AnalyticsInterface
		stats            *ServerStats
		conQueue         []*QueuePlayer
		newSessChan      chan *NewSession
		newConChan       chan *Connection
		removeConChan    chan *ConnectionRemove
		lastConnectionId int
		lastSessionId    int
		lastDBUpdate     time.Time
		accsToUpdate     map[string]*Account
		conToRedirect    []*ConnectionRedirect
		loopStartTime    *time.Time
		allConnections   map[int]*Connection
		mu               sync.Mutex
		runs             []*Run
		sessions         map[uuid.UUID]*Session
		_rand            *rand.Rand
	}
)

func (s *Server) NeedToKnow(l ...msutil.Insertable) {
	for _, v := range l {
		for _, s := range s.sessions {
			v.Insert(&s.con.XServerDataMsg)
		}
	}
}

func (s *Server) Rand() float64 {
	return s._rand.Float64()
}

func (s *Server) SafeUpdateAccount(a *Account) {
	s.mu.Lock()
	s.accsToUpdate[a.Token] = a
	s.mu.Unlock()
}

func (s *Server) SafeAddCon(con *Connection) {
	s.newConChan <- con
}

func (s *Server) SafeRemCon(con *Connection, reason ConnectionRemoveReason) {
	s.removeConChan <- &ConnectionRemove{Connection: con, Reason: reason}
}

func (s *Server) AddConnection(conn *ws.Conn) *Connection {
	con := &Connection{}
	con.msgsToWrite = []*[]byte{}
	con.msgsReaded = make([][]byte, 0)
	con.removeAfterWrite = 0
	con.conn = conn

	s.SafeAddCon(con)

	return con
}

func (s *Server) HandleConnection(conn *ws.Conn, log string, pass string, address string, acc *Account) {
	con := s.AddConnection(conn)
	go func(con *Connection) {
		con.SetSession(s.AddSession(acc, con))

		servData := &pb.ServerData{
			AccountGeneral: acc.GetAccountGeneral(),
			ConData: &pb.ConnectionData{ConMsg: config.ConnectionMessage.CON_SUCCESS_LOGIN}}
		con.writeServerData(servData, false)
	}(con)

	con.read()
}

func (s *Server) RemoveRun(run *Run) {
	run.disconnected = nil

	for _, s := range run.team.sessions {
		run.team.RemoveSession(s)
		s.run = nil
	}

	for _, r := range s.runs {
		if r == run {
			run.doRemove = true
			s.logger.Info("Run removed", zap.Int("run id", r.id))
		}
	}
}

func intRandom(min int, max int) int {
	if max == min {
		return min
	} else {
		if max > min {
			return min + rand.Intn(max-min+1)
		} else {
			return min - rand.Intn(min-max+1)
		}
	}
}

func (s *Server) CreateRun(RoomTypeNum uint32) *Run {
	RoomTypeToCoef :=  map[uint32]float64{
		1: 10,
		2: 100,
		3: 1000,
		4: 10,
		5: 50,
		6: 100,
	}

	coef, _ := RoomTypeToCoef[RoomTypeNum]

	run := &Run{
		RoomCoef: coef,
		RoomTypeNum: RoomTypeNum,
		doFishUpdate: true,
		RunState: RunState{}, StateMachines: StateMachines{}}

	run.team = &Team{
		sessions: []*Session{},
	}

	run.async = asyncutils.CreateAsycnUtil(server.loopStartTime)
	run.disconnected = []*Session{}
	run.positions = [MaxPlayers]*Session{}
	run.factory = CreateActorsFabric(run)

	run.timeline = &FishTimeline{}
	timelineFile := GetTimelineFile()
	err := TimelineFromFile(timelineFile, run.timeline)
	if err != nil {
		panic(err)
	}

	run.stateFinished = make(map[uuid.UUID]*Session)
	run.id = len(s.runs)

	s.RunMutex.Lock()
	s.runs = append(s.runs, run)
	s.RunMutex.Unlock()

	return run
}

func GetRandNickName() string {
	nicknames := []string{"dadisback", "magez", "undoo", "Dishi", "NachosAreCool", "NumbNuts", "Rendan", "Sliceros", "Rocketson", "LittleMaster", "brich", "ExhaustedDivinity", "TrollBayan", "Quideatanash"}
	return nicknames[intRandom(0, len(nicknames)-1)]
}

func (s *Server) QueueConnection(con *Connection) {
	server.logger.Info("Connection added to queue")
	queueTeam := CreateQueueTeam(con)
	s.conQueue = append(s.conQueue, queueTeam)
}

func CreateServer(logger *zap.Logger) (error, *Server) {
	a := &Server{}
	n := time.Now()
	a.RunMutex = sync.RWMutex{}
	a.logger = logger
	a.loopStartTime = &n
	a.accsToUpdate = make(map[string]*Account)
	a.stats = InitServerStats()

	err, an := analytics.InitTestAnalytics()
	if err != nil {
		return err, nil
	}
	a.analytics = an

	a.sessions = make(map[uuid.UUID]*Session, 0)
	a.newConChan = make(chan *Connection, nodeConfig.Server.ConnectionsBufferSize)
	a.newSessChan = make(chan *NewSession, nodeConfig.Server.ConnectionsBufferSize)
	a.removeConChan = make(chan *ConnectionRemove, nodeConfig.Server.ConnectionsBufferSize)
	a.allConnections = make(map[int]*Connection)
	a.Seed = time.Now().UTC().UnixNano()
	rand.Seed(a.Seed)//time.Now().UTC().UnixNano())
	a._rand = rand.New(rand.NewSource(time.Now().UnixNano()))//rand.NewSource(time.Now().UnixNano()))

	logger.Info("Server created")
	return nil, a
}

func (s *Server) GetDisconnectedAndRemove(sess *Session) (*Session, *Run) {
	var disconnectedSession *Session
	var room *Run
	for _, r := range s.runs {
		for i, d := range r.disconnected {
			if d.account.Token == sess.account.Token {
				disconnectedSession = d
				r.disconnected = append(r.disconnected[:i], r.disconnected[i+1:]...)
				room = r
				break
			}
		}
	}
	return disconnectedSession, room
}

func (s *Server) DisconnectFin(sess *Session) {
	s.logger.Info("Final session disconnect", zap.String("session ID", sess.id.String()))
	disconnectedSession, room := s.GetDisconnectedAndRemove(sess)
	if disconnectedSession != nil {
		room.Byebye(sess)
		s.RemoveSessionById(disconnectedSession.id)

		if len(room.disconnected) == 0 && len(room.team.sessions) == 0 {
			s.RemoveRun(room)
			return
		}

		//if len(room.disconnected) == 0 && sessionsCount > 0 {
		//	room.Resume()
		///}
	}
}

func (s *Server) TryReconnect(temporarySession *Session) {
	server.logger.Info("Trying to reconnect")
	disconnectedSession, room := s.GetDisconnectedAndRemove(temporarySession)
	if disconnectedSession != nil {
		disconnectedSession.finalDisconAt = nil
		server.logger.Info("Found disconnected session")
		temporarySession.con.SetSession(disconnectedSession)
		room.ReconSession(disconnectedSession, temporarySession.con)
	} else {
		newSession := temporarySession
		newSession.con.SetSession(s.AddSession(newSession.account, newSession.con))
		newSession.con.WriteToMsg().ConData = &pb.ConnectionData{ConMsg: config.ConnectionMessage.CON_SUCCESS_LOGIN}
	}
}

func (s *Server) AddSession(acc *Account, con *Connection) *Session {
	var ses *Session
	ses = ses.Create(con, server.lastSessionId)
	s.newSessChan <- &NewSession{Session: ses, Account: acc}
	return ses
}

func (s *Server) AddNewSession(ses *Session, acc *Account) {
	server.logger.Info("Creating session for connection with progress")

	ses.con.CreateXServerDataMsg(ses.id)
	ses.sessionState = config.SessionState.InGame
	ses.account = acc

	server.QueueConnection(ses.con)
}

func (s *Server) RemoveFromQueue(c *Connection) {
	if c.queue != nil {
		c.queue.remCon(c)
	}
}

func (s *Server) spawnPlayers(sessions []*Session, run *Run) {
	for _, x := range sessions {
		x.NeedToKnow(run.Location)

		pl, fx := run.CreatePlayer(x)
		x.NeedToKnow(x.account)
		x.NeedToKnow(x.player.Effect(confActions.StartScene))
		x.ConnectionData.PlayerID = x.player.ID
		x.NeedToKnow(&x.ConnectionData)
		x.NeedToKnow(run.timeline.CurrentScene.Curves)
		run.team.NeedToKnow(pl, fx)
	}
}

func (s *Server) StartRunWithPlayers(player1 *Connection) {
	run := s.CreateRun(player1.session.RequestedRoomType)
	sess := []*Session{player1.session}
	for _, pl := range sess {
		run.AddSession(pl)
	}

	run.SetScene("aquaman")

	s.spawnPlayers(sess, run)

	//run.timeline.AddStartFishes(run)
}

func (s *Server) playerJoinRun(run *Run, playerCon *Connection) {
	server.logger.Info("Player joined run")
	run.AddSession(playerCon.session)

	s.spawnPlayers([]*Session{playerCon.session}, run)

	playerCon.session.NeedToKnow(playerCon.session.SendAllObjects(run, playerCon.session.player)...)
}

func (s *Server) UpdateQueue() {
	lq := len(s.conQueue)
	for i := 0; i < lq; i++ {
		player1 := s.conQueue[i].connection

		if player1 != nil && player1.session.RequestedRoomType > 0 {
			newRun := true
			for _, run := range s.runs {
				if run.timeline.CurrentScene.Name == "aquaman" && run.RoomTypeNum == player1.session.RequestedRoomType &&  run.haveEmptySlots() {
					s.playerJoinRun(run, player1)
					newRun = false
				}
			}

			if newRun {
				s.StartRunWithPlayers(player1)
			}

			s.RemoveFromQueue(player1)
			s.conQueue = append(s.conQueue[:i], s.conQueue[i+1:]...)
			lq--
			i--
		}

		//removing connection anyway
		//removing connection anyway
	}
}

func (s *Server) UpdateDB() {
	s.mu.Lock()
	toUpdate := s.accsToUpdate

	for _, v := range toUpdate {
		v.Reset()
	}

	s.accsToUpdate = make(map[string]*Account)
	s.mu.Unlock()

	if len(toUpdate) > 0 {
		database.SaveAccounts(toUpdate)
	}
}

func (s *Server) RemoveSessionById(id uuid.UUID) {
	delete(server.sessions, id)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func (s *Server) Loop() {
	tickerStats := time.NewTicker(nodeConfig.Server.StatsUpdateTime)
	s.analytics.StartLoop()

	go func(tickerStats *time.Ticker, s *Server) {
		for {
			<-tickerStats.C
			server.stats.EverySec()
		}
	}(tickerStats, s)

	ticker := time.NewTicker(nodeConfig.Server.FrameDelay)
	logticker := time.NewTicker(nodeConfig.Server.RoomLogDelay)

	go func(ticker *time.Ticker) {
		for {
			select {
			case remCon := <-s.removeConChan:
				server.RemoveFromQueue(remCon.Connection)

				session := remCon.Connection.GetSession()
				if session != nil && session.run != nil {
					if remCon.Reason == CON_REM_DISCONNECT {
						session.run.Disconnect(session)
						session.con = nil
					}

					if remCon.Reason == CON_REM_BYE {
						session.run.team.RemoveSession(session)
						s.RemoveSessionById(session.id)
					}

				}

				delete(server.allConnections, remCon.Connection.id)
				break
			case newSess := <-s.newSessChan:
				s.AddNewSession(newSess.Session, newSess.Account)
				newSess.Session.ServerListIndex = len(s.sessions)
				s.sessions[newSess.Session.id] = newSess.Session
				break
			case newCon := <-s.newConChan:
				newCon.id = s.NextConId()
				s.allConnections[newCon.id] = newCon
				break

			case <- logticker.C:
				for _, r := range s.runs {
					r.Log(server.logger)
				}
				break
			case <-ticker.C:
				s.stats.OnServerLoop()
				prev := *s.loopStartTime
				*s.loopStartTime = time.Now()
				delta := (*s.loopStartTime).Sub(prev)

				ping := false
				if time.Now().Sub(s.lastPing) > nodeConfig.Server.PingPeriod {
					ping = true
					s.lastPing = time.Now()
				}

				//TODO: ITERATION OVER MAP VERY SLOW
				for _, c := range s.allConnections {
					c.applyMsg()
				}

				//durStr := fmtDuration(since)
				//fmt.Println(durStr)

				stats.Record(context.Background(),
					server.stats.Metrics.PacketSize.M(int64(math.Round(server.Rand()*500))))

				for i := len(server.runs) - 1; i >= 0; i-- {
					r := server.runs[i]
					if r.doRemove {
						s.RunMutex.Lock()
						server.runs = append(server.runs[:i], server.runs[i+1:]...)
						s.RunMutex.Unlock()
						r.factory.run = nil
						r.factory = nil
						continue
					}

					r.Update(delta)
				}

				for _, c := range s.allConnections {
					if c.IsChanged() {
						c.writeServerDataAndReset()
						go c.write(false)
					} else {
						go c.write(ping)
					}
				}

				server.UpdateQueue()

				break
			}
		}
	}(ticker)
}

func (server *Server) NextConId() int {
	server.lastConnectionId++
	return server.lastConnectionId
}

func (s *Server) TxnId() string {
	uuid := uuid.Must(uuid.NewV1())
	return uuid.String()
}
