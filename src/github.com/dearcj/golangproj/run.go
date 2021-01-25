package main

import (
	"errors"
	async "github.com/dearcj/golangproj/asyncutils"
	pb "github.com/dearcj/golangproj/network"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"strings"
	"time"
)

type LevelMiscObject struct {
	Pos Vec2
}

type LevelMiscObjInterface interface {
	ProcessTurn(r *Run, turnNum int32)
}

type PlayerSpawner struct {
	LevelMiscObject
}

func (p *PlayerSpawner) ProcessTurn(r *Run, turnNum int32) {

}

type AABB struct {
	LeftRight Vec2
	Size      Vec2
}

func (aabb AABB) RandomPos() Vec2 {
	p := Vec2{aabb.LeftRight[0] + float32(server.Rand()*float64(aabb.Size[0])),
		aabb.LeftRight[1] + float32(server.Rand()*float64(aabb.Size[1]))}
	return p
}

type ThingSpawner struct {
	Area AABB
	LevelMiscObject
	LastSpawnTurn int32
	SpawnInterval int32
}

const MaxPlayers = 8

type (
	BattleInfo struct {
		playerTurnStarted time.Time
		turnNumber        int
	}

	IFSM interface {
		Run(r *Run) (TurnStateType, time.Duration, NextStateFunc)
	}

	NextStateFunc func(r *Run) (TurnStateType, time.Duration, CurStateFunc, NextStateFunc)
	CurStateFunc func()

	RunState struct {
		stateFinished       map[uuid.UUID]*Session
		state               TurnStateType
		pauseAdditionalTime time.Duration
		maxStateDuration    time.Duration
		stateSet            time.Time
		onStateFinish       NextStateFunc
		pause               bool
		pauseTime           time.Duration
		prevTick            time.Time
		tick                time.Duration
	}

	StateMachines struct {
	}

	Run struct {
	RunState
	StateMachines
	BattleInfo

	RoomCoef      float64
	doFishUpdate  bool
	timeFromStart time.Duration
	positions     [MaxPlayers]*Session
	miscObjs      []LevelMiscObjInterface
	level         *Level
	stateHold     bool
	doRemove      bool
	id            int
	async         *async.AsyncUtil
	disconnected  []*Session
	Location      *LocationData
	pauseStart    time.Time
	currentFSM    IFSM
	factory       *ActorF
	team          *Team
	timeline      *FishTimeline
	RoomTypeNum   uint32
	transition    bool
}
)

func (r *Run) HoldUntilAnimation(f NextStateFunc) NextStateFunc {
	r.stateHold = true
	r.Pause(time.Hour * 10000)
	return f
}

func (r *Run) RunFSM(f IFSM) (TurnStateType, time.Duration, NextStateFunc) {
	r.currentFSM = f
	return r.currentFSM.Run(r)
}

func (r *Run) Pause(dur time.Duration) {
	r.pause = true
	r.pauseTime = dur
	r.async.Pause()
	r.pauseStart = time.Now()
}

func (r *Run) Resume() {
	r.pause = false
	r.async.Unpause()
}

func (r *Run) Disconnect(session *Session) {
	r.disconnected = append(r.disconnected, session)
	//r.Pause(config.Player.DisconnectWait)
	server.logger.Info("Player started disconnecting", zap.String("session ID", session.id.String()))
	session.FinalDisconnectAfter(config.Player.DisconnectWait)
}

func (r *Run) FinishSessionState(s *Session) {
	if r.stateFinished[s.id] != nil {
		return
	}

	r.team.NeedToKnow(
		s.player.Effect(confActions.FinishState).V(float32(r.state)))

	r.stateFinished[s.id] = s

	canFinish := true
	for _, ses := range r.team.sessions {
		if _, ok := r.stateFinished[ses.id]; !ses.Killed && !ok {
			canFinish = false
			break
		}
	}

	if canFinish {
		r.EndCurrentState()
	}
}

func (r *Run) EndCurrentState() {
	if r.onStateFinish != nil {
		r.SetState(r.onStateFinish(r))
	}
}

func (r *Run) UpdateStoryData() {
	var team_hp_pct float32
	count := 0
	for _, s := range r.team.sessions {
		if s.player == nil {
			continue
		}
		u := s.player.FindByComponent(config.Components.Unit).(*Unit)
		team_hp_pct += 100. * (float32(u.HP) / float32(u.MaxHP))
		count++
	}
	if count > 0 {
		team_hp_pct /= float32(count)
	}
}

func (r *Run) SetState(state TurnStateType, maxDuration time.Duration, doState CurStateFunc, onStateFinish NextStateFunc) {
	r.stateSet = time.Now()
	r.pauseAdditionalTime = 0
	r.maxStateDuration = maxDuration
	r.state = state
	r.stateFinished = make(map[uuid.UUID]*Session)
	r.onStateFinish = onStateFinish
	//r.UpdateStoryData()
	//r.team.NeedToKnow(r.StateChangeFx())
	if doState != nil {
		doState()
	}
}

func (r *Run) BotsMove() {
	server.logger.Info("Bots move")
	for _, session := range r.team.sessions {
		if session.con != nil && session.con.conn == nil {
			if session.player == nil {
				continue
			}

			unit := session.player.FindByComponent(config.Components.Unit).(*Unit)
			if unit.Killed {
				continue
			}

			session.OnFinishCurrentState()
		}
	}
}

func (r *Run) Update(delta time.Duration) {
	r.async.Update()

	if r.pause {
		r.pauseAdditionalTime += delta
	}

	if !r.pause && server.loopStartTime.Sub(r.stateSet) > r.maxStateDuration+r.pauseAdditionalTime {
		r.EndCurrentState()
	}

	var toDisc []*Session
	for _, s := range r.disconnected {
		if s.finalDisconAt != nil && server.loopStartTime.After(*s.finalDisconAt) {
			toDisc = append(toDisc, s)
		}
	}

	for _, s := range toDisc {
		server.DisconnectFin(s)
	}
	r.timeFromStart += delta
	r.factory.Update(delta)

	elapsed := time.Since(r.prevTick)
	r.tick = elapsed
	r.prevTick = time.Now()

	r.timeline.CurrentScene.FishUpdate(delta, r, false, time.Millisecond*1400, true, false)

	//sessList := make([]*Session, len(r.team.sessions))
	//copy(sessList, r.team.sessions)
	//for _, s := range r.team.sessions {
	//if len(removed) > 0 && s.con != nil {
	//	s.con.WriteToMsg().Toremove = removed
	//}
	//}

}

func (r *Run) ReconSession(s *Session, newConnection *Connection) {
	s.run = r //RETURN ROOM TO SESSION
	s.con = newConnection
	s.con.SetSession(s)

	s.ConnectionData.ConMsg = config.ConnectionMessage.CON_SUCCESS_LOGIN
	s.NeedToKnow(s.ConnectionData)
	s.NeedToKnow(s.SendAllObjects(r, s.player)...)

	if len(r.disconnected) == 0 {
		r.Resume()
	}
}

func (r *Run) RemoveSession(s *Session) {
	for inx, sess := range r.positions {
		if sess == s {
			r.positions[inx] = nil
		}
	}

	r.team.RemoveSession(s)
	server.logger.Debug("Removing Player", zap.Uint32("player id", s.player.NetworkObject.ID))
	r.team.NeedToKnow(r.factory.RemoveObject(s.player, false, false))
}

func (r *Run) CreatePlayer(session *Session) (player *Player, serverEffect *ServerEffect) {
	player, serverEffect = r.factory.AddPlayer(session, r.timeline.Guns[0])

	//r.factory.SetPlace(session.player, defPos)
	session.player.session = session
	//unit := session.player.FindByComponent(config.Components.Unit).(*Unit)

	//TODO: need to know of all objects at the start
	//worldInsertables := session.SendAllObjects(r)
	//session.NeedToKnow(worldInsertables...)

	return
}

func (r *Run) AddSession(s *Session) {
	s.run = r
	r.team.AddSession(s)
	s.sessionState = config.SessionState.InGame
}

func (r *Run) Byebye(s *Session) {
	pos := s.player.FindByComponent(confComponents.Player).(*Player).NetworkData.StartPosition
	server.logger.Info("Removing session", zap.String("ID", s.id.String()), zap.Uint32("ID", pos))
	r.RemoveSession(s)
	if s.con != nil {
		s.con.SetSession(nil)
	}
	s.con = nil
}

func (run *Run) SetLocation(scenename string, roomid int) {
	run.Location = &LocationData{}
	run.Location.LocationData = &pb.LocationData{
		LocationName: scenename,
		RoomId:       uint32(roomid),
	}
}

func (run *Run) StateChangeFx() *ServerEffect {
	//prop := float64((time.Now()).Sub(run.stateSet)-run.pauseAdditionalTime) / float64(run.maxStateDuration+1)
	//intProp := int32(prop * float64(config.Client.StateElapsedPrecision))
	return cf(confActions.ChangeState).V(float32(run.state))
}

func (run *Run) SkipToNextRoom() {
	//	if run.state == config.TurnState.BATTLE_PREPARE {
	//		for _, s := range run.team.sessions {
	//			run.FinishSessionState(s)
	//		}
	//	}

	//	run.turnActions.actions = []*pb.TurnAction{}
	//	run.team.NeedToKnow(&run.turnActions)

	//	for _, s := range run.team.sessions {
	//		run.FinishSessionState(s)
	//	}
}

func (r *Run) execSpawnersAndNotify(turnNum int32) {
	for _, v := range r.miscObjs {
		v.ProcessTurn(r, turnNum)
	}
}

func (r *Run) haveEmptySlots() bool {
	return MaxPlayers-len(r.team.sessions) > 0
}

func (r *Run) StartScene(sceneName string) {
	r.SetScene(sceneName)
	for _, x := range r.team.sessions {
		x.NeedToKnow(x.player.Effect(confActions.StartScene))
	}
}

func (run *Run) SetScene(name string) {
	run.SetLocation(name, run.id)
	run.Location = &LocationData{&pb.LocationData{
		LocationName: name,
		RoomId:       uint32(run.id)}}

	for _, s := range run.timeline.Scenes {
		if strings.ToLower(s.Name) == strings.ToLower(name) {
			run.timeline.CurrentScene = s
			run.timeline.CurrentScene.Reset()
			break
		}
	}

	if run.timeline.CurrentScene == nil {
		panic(errors.New("No such scene: " + name))
	}

	run.team.NeedToKnow(run.Location, run.timeline.CurrentScene.Curves)
	run.KillAllFishes(false)
	run.doFishUpdate = false
	run.async.DelayedCall(func() {
		run.doFishUpdate = true
		run.timeline.AddStartFishes(run)
	}, time.Millisecond*500)
}

func (r *Run) KillAllFishes(silent bool) {
	fishes := r.factory.FilterObjects(confComponents.Fish)
	for _, x := range fishes {
		r.team.NeedToKnow(r.factory.RemoveObject(x, false, silent))
	}
}

func (r *Run) Log(logger *zap.Logger) {
	logger.Info("Run", zap.Int("#", r.id), zap.Int("players", len(r.team.sessions)))
}

func (r *Run) getBet(gun *Gun) float64 {
	return gun.Damage * r.RoomCoef
}
