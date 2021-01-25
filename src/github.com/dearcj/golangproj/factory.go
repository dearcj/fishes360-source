package main

import (
	"github.com/dearcj/golangproj/bitmask"
	pb "github.com/dearcj/golangproj/network"
	"go.uber.org/zap"
	"math/rand"
	"time"
	"unsafe"
)

type DelayedCall struct {
	param     unsafe.Pointer
	function  func(unsafe.Pointer)
	duration  time.Duration
	startTime time.Time
}

type ActorF struct {
	run                *Run
	spawners           []*MonsterSpawner
	lastUpdateSpawners time.Time
	actors             []*Object
	maxID              int
	positions          []*Object
	ids                map[int]*Object
}

func (f *ActorF) SpawnPlace(ownerPlace int) int {
	min := 0
	place := -1

	if ownerPlace >= 4 {
		min = 4
	}

	for x := min; x <= min+3; x++ {
		if f.positions[x] == nil {
			place = x
			break
		}
	}

	return place
}

func (f *ActorF) CountType(typeColGroup bitmask.Bitmask) int {
	count := 0
	for _, el := range f.actors {
		if el.FindByComponent(typeColGroup) != nil {
			unit := el.FindByComponent(config.Components.Unit)
			if unit != nil {
				if unit.(*Unit).Killed == false {
					//	println("MONSER HAS ", int32(unit.(*Unit).HP))
					count++
				}
			}
		}
	}
	return count
}

func (f *ActorF) Update(dt time.Duration) {
	var i = 0
	//var prevAddedRemoved bool

	for i = 0; i < len(f.actors); i++ {
		a := f.actors[i]
		if a.doRemove {

			//prevAddedRemoved = true
			f.actors[i] = f.actors[len(f.actors)-1]
			f.actors = f.actors[:len(f.actors)-1]
			i--

		}
	}

	for i = 0; i < len(f.actors); i++ {
		a := f.actors[i]
		a.Process(0)
	}
}

func (f *ActorF) Remove(inx int) {
	f.actors[inx] = f.actors[len(f.actors)-1]
	f.actors = f.actors[:len(f.actors)-1]
}

func (f *ActorF) AddFish(ID uint32, room *Run, curve *Curve, skipTime time.Duration, fish *FishConfig, curveTime time.Duration, silent bool) (*Fish, FList) {
	var list FList
	o, fx := f.Add(&f.actors)

	list = list.AddSingle(fx)
	BaseCharacter := BaseCharacter{
		HP:    float64(fish.Hp),
		MaxHP: float64(fish.Hp)}

	AI := &Fish{
		FishConfig:  fish,
		SkipTime:    skipTime,
		BornTime:    room.timeline.TimeNow(),
		CurrentTime: skipTime,
		FishId:      ID,
		CurveTime:   curveTime,
		CurveInx:    curve.Inx,
		Curve:       curve}

	AI.CachedPos = curve.GetCachedPoint(AI.GetCurveProp())


	unit := &Unit{
		BaseCharacter: BaseCharacter,
	}

	unit.BaseCharacter.Object = o

	o.BindComponent(unit)
	o.BindComponent(AI)
	o.Type = uint32(config.Components.Fish)
	o.InitComponents()
	if silent {
		return nil, nil
	} else {
		return AI, list
	}
}

func (f *ActorF) AddPlayer(s *Session, defaultGun *Gun) (*Player, *ServerEffect) {
	var newActor, fx = f.Add(&f.actors)
	newActor.session = s

	unit, pl := CreateDefaultPlayer(s.progress)
	pl.currentGun = defaultGun
	unit.BaseCharacter.Object = newActor
	newActor.BindComponent(unit)
	newActor.BindComponent(pl)

	s.SetPlayer(newActor)

	//TODO: for deathmatch is current

	pl.NetworkData = &pb.Player{
		StartPosition: 0,
		NetworkObject: &pl.parent.NetworkObject,
	}

	pl.NetworkData.StartPosition = uint32(sessionPos(f, s))
	return pl, fx
}

func sessionPos(f *ActorF, s *Session) int {
	positions := []int{}

	for inx, sess := range f.run.positions {
		if sess == nil {
			positions = append(positions, inx)
		}
	}

	if len(positions) > 0 {
		inx := rand.Intn(len(positions))
		pos := positions[inx]
		f.run.positions[pos] = s
		return pos
	} else {
		return 0
	}
}

func (f *ActorF) CreateObjectWithComponent(c Component) (*Object, *ServerEffect) {
	var newActor, ff = f.Add(&f.actors)
	newActor.BindComponent(c)

	return newActor, ff
}

func (f *ActorF) RemoveObject(obj *Object, death bool, silent bool) FList {
	if obj.doRemove {
		return nil
	}

	if f.ids[int(obj.ID)] != nil {
		f.ids[int(obj.ID)].Remove()
		f.ids[int(obj.ID)] = nil
	} else {
		server.logger.Debug("Object removed but not in map", zap.Uint32("ID", obj.ID))
	}
	//server.logger.Debug("Removing object", zap.Uint32("ID", obj.ID))
	dv := 0.
	if death {
		dv = 1
	}
	if silent {
		return FList{}
	} else {
		return FList{}.AddSingle(cf(confActions.Remove).ID(obj.ID).V(float32(dv)))
	}
}

func (f *ActorF) Add(list *[]*Object) (*Object, *ServerEffect) {
	var newActor = &Object{
		Components:    make(map[bitmask.Bitmask]Component),
		NetworkObject: pb.NetworkObject{ID: uint32(f.getNewId())},
	}
	newActor.factory = f

	f.ids[int(newActor.NetworkObject.ID)] = newActor
	*list = append(*list, newActor)

	if int(newActor.ID) > f.maxID {
		f.maxID = int(newActor.ID)
	}

	return newActor, newActor.Effect(confActions.Appear)
}

func (f *ActorF) FindByClientID(id int, session *Session) *Object {
	inx := f.InxByCID(id, session)
	if inx >= 0 {
		return f.actors[inx]
	} else {
		return nil
	}
}

func (f *ActorF) Find(id int) *Object {
	inx := f.InxById(id)
	if inx >= 0 {
		if !f.actors[inx].doRemove {
			return f.actors[inx]
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func (f *ActorF) InxByCID(id int, s *Session) int {
	for i, el := range f.actors {
		if el.session == s {
			return i
		}
	}

	return -1
}

func (f *ActorF) InxById(id int) int {
	for i, el := range f.actors {
		if int(el.ID) == id {
			return i
		}
	}

	return -1
}

func (f *ActorF) FilterObjects(component bitmask.Bitmask) []*Object {
	var objs []*Object
	for _, x := range f.actors {
		if x.FindByComponent(component) != nil && !x.doRemove {
			objs = append(objs, x)
		}
	}

	return objs
}

func (f *ActorF) FilterComponents(component bitmask.Bitmask) []Component {
	var objs []Component
	for _, x := range f.actors {
		c := x.FindByComponent(component)
		if c != nil && !x.doRemove {
			objs = append(objs, c)
		}
	}

	return objs
}

func (f *ActorF) getNewId() int {
	if len(f.actors) == 0 {
		return 0
	}

	f.maxID = 0
	for _, x := range f.actors {
		if int(x.ID) > f.maxID {
			f.maxID = int(x.ID)
		}
	}

	ids := make([]bool, f.maxID+1)

	for _, x := range f.actors {
		ids[x.ID] = true
	}

	for inx, haveId := range ids {
		if !haveId {
			return inx
		}
	}

	return f.maxID + 1
}

func CreateActorsFabric(room *Run) *ActorF {
	var fabric = &ActorF{}
	fabric.positions = make([]*Object, 10, 10)
	fabric.ids = make(map[int]*Object)
	fabric.run = room
	fabric.maxID = 0
	return fabric
}
