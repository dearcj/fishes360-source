package main

import (
	"github.com/dearcj/golangproj/bitmask"
	"github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
	"go.uber.org/zap"
	"log"
	"time"
)

type Fish struct {
	FishConfig  *FishConfig
	CurveTime   time.Duration
	SkipTime    time.Duration
	BornTime    time.Time
	CurrentTime time.Duration
	Curve       *Curve
	FishId      uint32
	parent      *Object
	CurveInx    uint32
	CachedPos   *Vec3
}

func (fish *Fish) processIteration(me *Object, inum uint32, iprop float32) FList {
	return nil
}

func (a *Fish) onCollide(me *Object, col *Object) FList { return nil }

func (fish *Fish) onDestroy() {
}

func (fish *Fish) Insert(m *msutil.XServerDataMsg) {
	fish.InsertToList(m)
}

func (fish *Fish) InsertToList(m *msutil.XServerDataMsg) *pb.Fish {
	msg := m.WriteToMsg()
	objs := msg.Fishes
	var obj *pb.Fish

	if objs != nil {
		for inx, v := range objs {
			if v.NetworkObject.ID == fish.parent.NetworkObject.ID {
				obj = objs[inx]
				break
			}
		}
	}

	if obj == nil {
		if fish.FishConfig.IsBoss {
			log.Print("Sending boss")
		}

		obj = &pb.Fish{
			IsBoss: fish.FishConfig.IsBoss,
			CurveTime:     uint64(fish.CurveTime),
			CurveInx:      fish.Curve.Inx,
			FishType:      fish.FishId,
			StartTime:     uint64(fish.CurrentTime),
			Hp:            uint32(fish.parent.Unit().BaseCharacter.HP),
			Maxhp:         uint32(fish.parent.Unit().BaseCharacter.MaxHP),
			NetworkObject: &fish.parent.NetworkObject}
		objs = append(objs, obj)
	}

	msg.Fishes = objs
	return obj
}

func (fish *Fish) process(b *Object, dt float64) {}

func (fish *Fish) onInit(o *Object) {

	fish.parent = o
}

func (fish *Fish) getTypeId() bitmask.Bitmask {
	return config.Components.Fish
}

func (fish *Fish) getMoneyPrize(killer *Session, gunCoef float64, roomCoef float64, RTP float64) (win bool, money float64) {
	bv := fish.FishConfig.BaseValue
	//fish.FishConfig.Dispersion = 0.
	//	min := (float64(bv)*(1 - (fish.FishConfig.Dispersion) / 100.))
	//	max := (float64(bv)*(1 + (fish.FishConfig.Dispersion) / 100.))
	//	lose := (2*float64(fish.FishConfig.Hp)*(RTP/100.) - float64(bv) - max) / (min - max)
	win = true
	money = float64(bv) * gunCoef * roomCoef
	/*	if server.Rand() < lose {
			money = roomCoef*gunCoef*(min + server.Rand()*(float64(bv)-min))
			win = false
		} else {
			money = roomCoef*gunCoef*(float64(bv) + server.Rand()*(max-float64(bv)))
			win = true
	*/
	if killer != nil {
		go func(killer *Session) {
			if killer.account != nil {
				err := server.MakeWin(killer.account.Token, killer.account.Puuid, money)
				if err != nil {
					server.logger.Error("Failed to make bet", zap.Error(err))
				}
			}
		}(killer)

	}

	//	logger.Debug("Won money", zap.String("mon", fmt.Sprintf("%6.2f", money)))
	return
}

func (fish *Fish) GetCurveProp() float64 {
	prop := float64(fish.CurrentTime) / float64(fish.CurveTime)

	return prop
}

func (fish *Fish) PosOnCurve() float64 {
	return fish.Curve.CurveLen * fish.GetCurveProp()
}
