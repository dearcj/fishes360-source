package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dearcj/golangproj/msutil"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"time"
)

var cachedPoints map[string][]*Vec3 = make(map[string][]*Vec3)
type JSONDuration time.Duration

func (d JSONDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *JSONDuration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = JSONDuration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = JSONDuration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func SpawnFish(run *Run, curve *Curve, fishId uint32, fishConfig *FishConfig, curveTime time.Duration, silent bool) {
	run.team.NeedToKnow(run.factory.AddFish(uint32(fishId), run, curve, 0, fishConfig, curveTime, silent))
}

type Vec3 [3]float64


type CurvesList []*Curve

type TimelineScene struct {
	StartTime   *time.Time `json:"-"`
	TotalLength uint32
	Curves      CurvesList
	Name        string
	LastCheck   time.Time `json:"-"`
	MaxFishes   uint32
}

func (c CurvesList) Insert(msg *msutil.XServerDataMsg) {
	str := []string{}
	for _, x := range c {
		str = append(str, x.RawPoints)
	}
	msg.WriteToMsg().Curves = str
}

type Gun struct {
	Id     uint32 `json:"-"`
	Damage float64
}

func (gun Gun) CalcProb(fishConfig *FishConfig, RTP float64) float64 {
	val := (RTP / 100.) / float64(fishConfig.BaseValue) // float64(gun.Damage)
	if val > 1 {
		val = 1
	}

	return val
}

func (gun Gun) CalcHpDelta(fishConfig *FishConfig, hp float64, maxhp float64, RTP float64) float64 {
	dmg := -hp * gun.CalcProb(fishConfig, RTP) * (1 + (server.Rand()-0.5)*0.5)
	return dmg
}

type FishConfig struct {
	Id          uint32
	Hp          uint32
	Speed       float64
	BaseValue   uint32
	Dispersion  float64
	Active      bool
	IsBoss      bool
	Name        string
	SpawnChance float64
}

type FishTimeline struct {
	Guns         []*Gun
	AvailIds     []uint32 `json:"-"`
	RTP          float64
	Scenes       []*TimelineScene
	CurrentScene *TimelineScene `json:"-"`
	Fishes       []*FishConfig
	TimeDelta    time.Duration `json:"-"`
}

func (timeline *FishTimeline) checkRTP(num int, RTP_PERCENTAGE float64) float64 {
	spent := 0.
	win := 0.
	RoomCoef := 100.
	winApprox := []float64{}
	for x := 0; x < num; x ++ {
		f := timeline.Fishes[rand.Intn(len(timeline.Fishes))]
		gun := timeline.Guns[rand.Intn(len(timeline.Guns))]
		betSize := gun.Damage * RoomCoef
		prob := gun.CalcProb(f, RTP_PERCENTAGE)
		killed := prob > server.Rand()
		if killed {
			ff := &Fish{
				FishConfig: f,
			}
			_, money := ff.getMoneyPrize(nil, gun.Damage, RoomCoef, RTP_PERCENTAGE)
			winApprox = append(winApprox, money)
			win += money
		}

		spent -= betSize
	}
	prop := win / (-spent)
	sum := 0.0
	for _, x := range winApprox {
		sum += x
	}
	sum /= float64(len(winApprox))
	server.logger.Info("EV from bet", zap.Float64("ev", sum))


	server.logger.Info("RTP TEST GET BALANCE", zap.Float64("spent", spent), zap.Float64("win", win), zap.Float64("rtp", prop))

	return spent
}



func (timeline *FishTimeline) TimeNow() time.Time {
	return time.Now().Add(timeline.TimeDelta)
}

var lastSpawn = time.Now().Add(-300 * time.Hour)

func (tf *TimelineScene)  FishUpdate(delta time.Duration, r *Run, silent bool, spawnCheckDelay time.Duration, respawnDeadFishes bool, excludeBosses bool) {
	for _, x := range r.team.sessions {
		if x.player != nil {
			pl := x.player.FindByComponent(config.Components.Player).(*Player)
			if pl.needNotifyAngles {
				pl.needNotifyAngles = false

				for _, exceptMe := range r.team.sessions {
					if exceptMe != x {
						exceptMe.NeedToKnow(pl.AngleChangedFx())
					}
				}
			}

			if pl.needNotifyShoot {
				pl.needNotifyShoot = false
				s := pl.parent.session
				obj := r.factory.Find(pl.shootTarget)
				if obj != nil {
					fish := obj.FindByComponent(confComponents.Fish)
					if fish != nil {
						bc := obj.BaseCharacter()
						bet := r.getBet(pl.currentGun)
						x.NeedToKnow(pl.MakeBet(float32(bet)))
						go func(s *Session) {
							err := server.MakeBet(x.account.Token, x.account.Puuid, bet)
							if err != nil {
								server.logger.Error("Failed to make bet", zap.Error(err))
								if s.player != nil {
									s.NeedToKnow(s.player.Effect(confActions.NoMoneyMode))
								}
							}
						}(s)




						r.team.NeedToKnow(bc.DealProbDamage(pl.parent.BaseCharacter(), pl.currentGun))
					}
				}
				for _, exceptMe := range r.team.sessions {
					if exceptMe != x {
						exceptMe.NeedToKnow(pl.ShootFx(pl.shootTarget, int(pl.currentGun.Id)))
					}
				}
			}
		}
	}


	fishes := r.factory.FilterObjects(confComponents.Fish)
	for _, f := range fishes {
		fish := f.FindByComponent(confComponents.Fish).(*Fish)
		fish.CurrentTime += delta
		fish.CachedPos = fish.Curve.GetCachedPoint(fish.GetCurveProp())

		if fish.CurrentTime > (time.Duration(fish.CurveTime)) {

			r.team.NeedToKnow(r.factory.RemoveObject(fish.parent, false, silent))
			if respawnDeadFishes {
				tf.SpawnFishToMax(r, excludeBosses)
			}
		}
	}

	if r.timeline.CurrentScene != nil && r.doFishUpdate {
		r.timeline.CurrentScene.Update(r, r.timeline, silent, spawnCheckDelay, excludeBosses)
	}
}

func (tf *TimelineScene) SpawnFishToMax(run *Run, excludeBosses bool) {
	for x := 0; x < 3; x ++ {
		allfishes := run.factory.FilterObjects(confComponents.Fish)
		if len(allfishes) < int(tf.MaxFishes) {
			tf.TryToSpawnFish(run, allfishes, run.timeline, false, excludeBosses)
		}
	}

}


func (timeline *FishTimeline) AddStartFishes(run *Run) {
	run.timeline.TimeDelta = 0
	delta := time.Millisecond * 100

	for x := 0; x < 4500; x++ {
		run.timeline.TimeDelta += delta
		run.timeline.CurrentScene.FishUpdate(delta, run, true, delta, false, true)
		run.factory.Update(delta)
		if len(run.factory.FilterObjects(confComponents.Fish)) >= int(run.timeline.CurrentScene.MaxFishes) {
			break
		}
	}

	allfishes := run.factory.FilterObjects(confComponents.Fish)

	run.timeline.CurrentScene.LogCurves(run, allfishes)
	for _, f := range allfishes {
		run.team.NeedToKnow(f.Effect(confActions.Appear), f.FindByComponent(confComponents.Fish).(*Fish))
	}
}

func GetCurveTime(curve *Curve, fish *FishConfig) time.Duration {
	const CURVE_COEF = 1 / 10.
	return time.Duration((curve.CurveLen/fish.Speed)*CURVE_COEF) * time.Second
}

type PosSpeed struct {
	Pos float64
	Speed float64
	Curve *Curve
}

func (tf *TimelineScene) Simulate(fishes []*PosSpeed, iterations int) bool {
	delta := 1. / float64(iterations)
	for i := 0; i < iterations; i++ {
		pos2 := fishes[0].Curve.GetCachedPoint(fishes[0].Pos)
		for inx := 1; inx < len(fishes); inx ++ {
			if fishes[inx].Pos <= 1.05 {
				pos1 := fishes[inx].Curve.GetCachedPoint(fishes[inx].Pos)

				if pos1.DistSq(pos2) < OVERLAP_SQUARED_DIST {
					return false
				}
			}
		}

		for _, f := range fishes {
			f.Pos += delta * (fishes[0].Curve.CurveLen / fishes[0].Speed) / (f.Curve.CurveLen / f.Speed)
		}
	}

	return true
}

func (tf *TimelineScene) CanSpawnFishAtCurve(fishToSpawn *FishConfig, curve *Curve, allfishes []*Object) bool {
	//fishesOnCurve := []*Fish{}
	posSpeed := []*PosSpeed{}

	posSpeed = append(posSpeed, &PosSpeed{
		Curve: curve,
		Pos: 0.,
		Speed: fishToSpawn.Speed,
	})

	for _, x := range allfishes {
		f := x.FindByComponent(confComponents.Fish).(*Fish)
		//if f.Curve.Inx == curve.Inx {
		//	fishesOnCurve = append(fishesOnCurve, f)
		//}

		posSpeed = append(posSpeed, &PosSpeed{
			Curve: f.Curve,
			Pos: f.GetCurveProp(),
			Speed: f.FishConfig.Speed,
		})
	}

	//if len(fishesOnCurve) == 0 {
	//	return true
	//}

	override := tf.Simulate(posSpeed, int(125. ))
	return override
/*	myTime := GetCurveTime(curve, fishToSpawn)

	for _, x := range fishesOnCurve {
		totalTime := time.Duration((1 - x.GetCurveProp()) * float64(x.CurveTime))
		if x.GetCurveProp() < 0.2 || myTime < time.Duration(float64(totalTime)) + time.Second*10 {
			return false
		}
	}*/

	//return true
}

var testCheck time.Time

func (tf *TimelineScene) Update(run *Run, allSettings *FishTimeline, silent bool, spawnCheckDelay time.Duration, excludeBosses bool) {
	if tf.StartTime == nil {
		t := allSettings.TimeNow()
		tf.StartTime = &t
	}

	allfishes := run.factory.FilterObjects(confComponents.Fish)


	if len(allfishes) < int(tf.MaxFishes) {
		if len(allSettings.AvailIds) > 0 && allSettings.TimeNow().Sub(tf.LastCheck) > spawnCheckDelay {
			tf.TryToSpawnFish(run, allfishes, allSettings, silent, excludeBosses)
			tf.LastCheck = allSettings.TimeNow()
		}
	}

}

func (tf *TimelineScene) Reset() {
	tf.StartTime = nil
	tf.LastCheck = time.Now()
}

func (tf *TimelineScene) LogCurves(r *Run, fishes []*Object) {
	println("------------------------------------------------------")
	for _, x := range fishes {
		f := x.FindByComponent(confComponents.Fish).(*Fish)
		fmt.Println("fish total time", f.CurveTime)
		fmt.Println("fish time elapsed", f.CurrentTime)
		server.logger.Info("", zap.Uint32("type", f.FishId),  zap.Uint32("hp", uint32(x.BaseCharacter().HP)), zap.Uint32("curve", f.Curve.Inx), zap.Float64("prop", f.GetCurveProp()))
	}
}

func (tf *TimelineScene) TryToSpawnFish(run *Run, allfishes []*Object, allSettings *FishTimeline, silent bool, excludeBosses bool) {
	for it := 0; it < 15; it++ {


		randomCurve := tf.Curves[rand.Intn(len(tf.Curves))]

		randomFishId := tf.GetFishType(allSettings, excludeBosses)




		fish := allSettings.Fishes[randomFishId-1]

		if tf.CanSpawnFishAtCurve(fish, randomCurve, allfishes) {
			curveTime := GetCurveTime(randomCurve, fish)

			SpawnFish(run, randomCurve, randomFishId, fish, curveTime, silent)
			break
		}
	}
}

func (tf *TimelineScene) GetFishType(allSettings *FishTimeline, excludeBosses bool) uint32 {
	all := allSettings.AvailIds

	var filtered []uint32
	if excludeBosses {
		for _, id := range all {
			if !allSettings.Fishes[id - 1].IsBoss {
				filtered = append(filtered, id)
			}
		}
	} else {
		filtered = all
	}

	fullProb := 0.
	for _, f := range filtered {
		fullProb += allSettings.Fishes[f - 1].SpawnChance
	}

	pos := rand.Float64() * fullProb

	left := 0.
	for _, f := range filtered {
		if pos >= left && pos < left + allSettings.Fishes[f - 1].SpawnChance {
			return allSettings.Fishes[f - 1].Id
		} else {
			left +=  allSettings.Fishes[f - 1].SpawnChance
		}
	}
	return allSettings.Fishes[0].Id
//	.AvailIds[rand.Intn(len(allSettings.AvailIds))
}

func TimelineFromFile(file []byte, ft *FishTimeline) error {
	err := json.Unmarshal(file, ft)
	if err != nil {
		return err
	}

	for _, x := range ft.Scenes {
		for inx, c := range x.Curves {
			c.Inx = uint32(inx)
			err := c.Parse()
			if err != nil {
				return err
			}

			if points, ok := cachedPoints[c.RawPoints]; ok {
				c.CachedPoints = points
				c.CachedSize = len(c.CachedPoints)
			} else {
				cachedPoints[c.RawPoints] = c.Cache()
				c.CachedSize = len(c.CachedPoints)
			}

		}
	}

	for inx, g := range ft.Guns {
		g.Id = uint32(inx)
	}

	for _, f := range ft.Fishes {
		if f.IsBoss {
			log.Print("Found boss fish: ", f.Id)
		}


		if f.Active {
			ft.AvailIds = append(ft.AvailIds, f.Id)
		}
	}

	return nil
}
