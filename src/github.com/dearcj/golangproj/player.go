package main

import (
	"github.com/dearcj/golangproj/bitmask"
	"github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
)

type Player struct {
	additionalMoves  uint32
	emotion          bitmask.Bitmask
	parent           *Object
	currentGun       *Gun
	angle1           int32
	angle2           int32
	needNotifyAngles bool
	needNotifyShoot  bool
	shootTarget      int

	NetworkData *pb.Player

	PlayerRelative PlayerRelativeInserter
}

type PlayerRelativeInserter struct {
	p *Player
}

func (a Player) onCollide(me *Object, col *Object) FList { return nil }

func (l *Player) processIteration(me *Object, inum uint32, iprop float32) FList {
	return nil
}

func (l *Player) InsertToList(m *msutil.XServerDataMsg) *pb.Player {
	msg := m.WriteToMsg()
	players := msg.Players
	var obj *pb.Player
	if players != nil {
		for inx, v := range players {
			if v.NetworkObject.ID == l.parent.NetworkObject.ID {
				obj = players[inx]
				break
			}
		}
	}

	if obj == nil {
		players = append(players, l.NetworkData)
	}

	msg.Players = players

	return obj
}

func (l *Player) Insert(m *msutil.XServerDataMsg) {
	l.InsertToList(m)
}

//todo: maybe change it later
func (a *Player) process(me *Object, dt float64) {
}

func (a *Player) onInit(o *Object) {
	a.parent = o
	a.PlayerRelative.p = a
}

func (a *Player) onDestroy() {
	a.parent = nil
}

func (a *Player) getTypeId() bitmask.Bitmask {
	return config.Components.Player
}

func (a *Player) SetAngle(i int32, i2 int32) {
	if a.angle1 != i && a.angle2 != i2 {
		a.needNotifyAngles = true
		a.angle1 = i
		a.angle2 = i2
	}
}

func (a *Player) MakeBet(bet float32) msutil.Insertable {
	a.parent.session.account.Balance -= float64(bet)

	return a.parent.Effect(confActions.MakeBet).V(bet)
}


func (a *Player) AddMoney(amount float32) msutil.Insertable {
	a.parent.session.account.Balance += float64(amount)

	return a.parent.Effect(confActions.MoneyChange).V(amount)
}

func (a *Player) ShootFx(target int, gun int) msutil.Insertable {
	return a.parent.Effect(confActions.Shoot).V(float32(target)).V2(float32(gun))
}

func (a *Player) AngleChangedFx() msutil.Insertable {
	return a.parent.Effect(confActions.AngleChange).V(float32(a.angle1)).V2(float32(a.angle2))
}

func (a *Player) Shoot() {
	a.needNotifyShoot = true
}

func CreateDefaultPlayer(p *Progress) (*Unit, *Player) {
	return &Unit{
		BaseCharacter: BaseCharacter{
			HP:    p.MaxHP,
			MaxHP: p.MaxHP,
		},
	}, &Player{}
}
