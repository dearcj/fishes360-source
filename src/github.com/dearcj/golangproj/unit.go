package main

import (
	"github.com/dearcj/golangproj/bitmask"
)

type Unit struct {
	BaseCharacter
}

func (a *Unit) processIteration(me *Object, inum uint32, iprop float32) FList {
	return nil
}

func (a *Unit) onCollide(bulletObject *Object, col *Object) FList { return nil }

func (a *Unit) process(me *Object, dt float64) {}

func (a *Unit) resetNetworkData() {}

func (a *Unit) updateNetworkData(o *Object) {}

func (a *Unit) getNetworkDataForSession(s *Session) interface{} { return nil }

func (a *Unit) onInit(o *Object) {

}

func (a *Unit) onDestroy() {
	a.BaseCharacter.OnDestroy()
}

func (a *Unit) getTypeId() bitmask.Bitmask {
	return config.Components.Unit
}
