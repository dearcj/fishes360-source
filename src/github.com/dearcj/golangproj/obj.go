package main

import (
	"github.com/dearcj/golangproj/bitmask"
	pb "github.com/dearcj/golangproj/network"
	"reflect"
)

type Component interface {
	onCollide(me *Object, another *Object) FList
	getTypeId() bitmask.Bitmask
	process(parent *Object, dt float64)
	processIteration(me *Object, inum uint32, iprop float32) FList
	onInit(o *Object)
	onDestroy()
}

type MonsterSpawner struct {
	Monster    *Object
	POS        Vec2
	EnemiesIDs []int
}

const ACTOR_RADIUS float32 = 20.

type Object struct {
	pb.NetworkObject
	Components map[bitmask.Bitmask]Component
	DesiredPos *Vec2

	factory  *ActorF
	session  *Session
	doRemove bool
}

func (a *Object) Move(dx float32, dy float32) {

	//server.logger.Debug("SETTING VEL", zap.Float32("velx", p.Vel[0]), zap.Float32("vely", p.Vel[1]))
}

func (a *Object) Effect(effectType pb.ActionType) *ServerEffect {
	return cf(effectType).ID(a.ID)
}

func (a *Object) ProcessIteration(i uint32, prop float32) (list FList) {
	for _, c := range a.Components {
		list = list.Add(c.processIteration(a, i, prop))
	}

	return
}

func (a *Object) Unit() *Unit {
	u := a.FindByComponent(confComponents.Unit)
	if u != nil {
		return u.(*Unit)
	}
	return nil
}

func (a *Object) BaseCharacter() *BaseCharacter {
	u := a.FindByComponent(confComponents.Unit)
	if u != nil {
		return &u.(*Unit).BaseCharacter
	}
	return nil
}

func (a *Object) ComponentsOnDestroy() {
	for _, c := range a.Components {
		c.onDestroy()
	}
}

func (a *Object) FindByComponent(_type bitmask.Bitmask) Component {
	return a.Components[_type]
}

func (a *Object) InitComponents() {
	for _, c := range a.Components {
		c.onInit(a)
	}
}

func (a *Object) FindByType(_type string) Component {
	for _, c := range a.Components {
		tp := reflect.TypeOf(c).String()
		if tp == "*main."+_type {
			return c
		}
	}
	return nil
}

func (a *Object) IsType(_type bitmask.Bitmask) bool {
	return bitmask.Bitmask(a.Type).HasFlag(_type)
}

func (a *Object) BindComponent(c Component) {
	a.Components[c.getTypeId()] = c
	c.onInit(a)
}

func (a *Object) Remove() {
	a.doRemove = true

	if a.session != nil {
		a.session.player = nil
	}
	a.ComponentsOnDestroy()
}

func (a *Object) Process(dt float64) {
	for _, c := range a.Components {
		c.process(a, dt)
	}
}

func (a *Object) TypeBitmask() bitmask.Bitmask {
	var b bitmask.Bitmask

	for _, c := range a.Components {
		b.AddFlag(c.getTypeId())
	}

	return b
}
