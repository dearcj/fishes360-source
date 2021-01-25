package main

import (
	n "github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
	"github.com/gofrs/uuid"
	"reflect"
	"time"
)

type ConnectionData struct {
	*pb.ConnectionData
}

func (c ConnectionData) Insert(s *n.XServerDataMsg) {
	s.WriteToMsg().ConData = c.ConnectionData
}

func ParamsSize(p []int32, x int) bool {
	return len(p) == x
}

type Session struct {
	ConnectionData  ConnectionData
	ServerListIndex int
	Killed          bool
	account         *Account
	id              uuid.UUID
	lastCID         int
	sessionState    SessionStateType
	con             *Connection
	progress        *Progress //GLOBAL PROGRESS
	player          *Object
	run             *Run //run where player located

	finalDisconAt     *time.Time
	RequestedRoomType uint32
}

func (s *Session) IsBot() bool {
	return s.con != nil && s.con.conn == nil
}

func (s *Session) Event(actionName *string) {
	server.analytics.AddAction(actionName, nil, s.account.Token, 0)
}

func (s *Session) Track(actionName *string, param string) {
	x := &param
	server.analytics.AddAction(actionName, x, s.account.Token, 0)
}

func (s *Session) Create(c *Connection, id int) *Session {
	server.lastSessionId++
	ses := &Session{
		con: c,
		id:  uuid.Must(uuid.NewV4())}
	ses.progress = &Progress{}

	ses.progress.resetProgress()

	ses.ConnectionData = ConnectionData{ConnectionData: &pb.ConnectionData{}}
	return ses
}

func (s *Session) SetPlayer(pl *Object) {
	pl.session = s
	s.player = pl
}

func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func (s *Session) NeedToKnow(l ...n.Insertable) {
	if s.con == nil {
		return
	}

	for _, v := range l {
		if v != nil && !isNil(v) {
			v.Insert(&s.con.XServerDataMsg)
		}
	}
}

func (s *Session) ParseCommand(cmd MapID, Params []int32) {
	//PRE RUN COMMANDS
	switch cmd {
	case config.Commands.CMD_REQUEST_ROOM:
		if len(Params) == 1 {
			if Params[0] >= 1 && Params[0] <= 6 {
				s.RequestedRoomType = uint32(Params[0])
			}
		}
		break
	}


	//AFTER RUN COMMANDS

	if s.run == nil || s.run.doRemove {
		return
	}
	_, ok := confCommandsMap[cmd]
	//loggerInner := server.logger.With(zap.String("Command use", cmdname))
	//loggerInner.Debug("Client command"+cmdname, zap.Int32s("param1", Params))
	if ok && s.run != nil && s.run.pause == false {
		switch cmd {

		case config.Commands.CMD_FINAL_EXIT_GAME:
			s.run.Byebye(s)
			break

		}
	}

	switch cmd {

	case config.Commands.CMD_NEXT_SCENE_HACK:
	/*	if len(Params) > 0 && Params[0] == -1 {
			s.run.StartScene("aquaman")
		} else {
			s.run.StartScene("deadfish")
		}*/
		break

	case config.Commands.CMD_ANGLE_CHANGE:
		if s.player != nil && len(Params) == 2 {
			player := s.player.FindByComponent(config.Components.Player).(*Player)
			player.SetAngle(Params[0], Params[1])
		}

		break
	case config.Commands.CMD_CHANGE_GUN:
		if s.player != nil && len(Params) == 1 {
			gunid := int(Params[0])
			if gunid >= 0 && gunid < len(s.run.timeline.Guns) {
				newgun := s.run.timeline.Guns[gunid]
				pl := s.player.FindByComponent(confComponents.Player).(*Player)
				pl.currentGun = newgun
			}

		}
		break
	case config.Commands.CMD_SHOOT:
		if s.player != nil && len(Params) == 1 {
			id := int(Params[0])
			obj := s.player.factory.Find(id)
			if obj != nil {
				fish := obj.FindByComponent(confComponents.Fish)
				if fish != nil {
					//bc := obj.BaseCharacter()
					pl := s.player.FindByComponent(confComponents.Player).(*Player)
					pl.shootTarget = id
					pl.Shoot()
				}
			}
		}

		break

	case config.Commands.CMD_RECONNECT_YES:
		server.TryReconnect(s)
		break

	case config.Commands.CMD_RECONNECT_NO:
		server.DisconnectFin(s)
		s.con.SetSession(server.AddSession(s.account, s.con))
		s.con.writeServerData(&pb.ServerData{ConData: &pb.ConnectionData{ConMsg: config.ConnectionMessage.CON_SUCCESS_LOGIN}}, false)
		break
	}
}

func (session *Session) SendAllObjects(r *Run, exceptThis *Object) (insertables []n.Insertable) {
	//insertables = append(insertables,
	//	session.ConnectionData,
	//	r.Location,
	//	session.account,
	//	r.timeline.CurrentScene.Curves)

	for _, x := range r.factory.actors {
		if x == exceptThis {
			//we already spawned ourself durinng Join / Run create

			continue
		}

		mon := x.FindByComponent(confComponents.Fish)
		if mon != nil {
			insertables = append(insertables, mon.(*Fish))
		}

		pl := x.FindByComponent(confComponents.Player)
		if pl != nil {
			insertables = append(insertables, pl.(*Player))
		}
		insertables = append(insertables, x.Effect(confActions.Appear))
	}

	return
}
func (s *Session) FinalDisconnectAfter(duration time.Duration) {
	time := server.loopStartTime.Add(duration)
	s.finalDisconAt = &time
}

func (s *Session) OnFinishCurrentState() {
	s.run.FinishSessionState(s)
}
