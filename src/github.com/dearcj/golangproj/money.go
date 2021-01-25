package main

import (
	"github.com/dearcj/golangproj/bitmask"
	pb "github.com/dearcj/golangproj/network"
)

type Money struct {
	Amount int
}

func (a *Money) mutateState(d *Object, no *pb.NetworkObject) {}

func (a *Money) process(b *Object, dt float64) {
}

func (a *Money) onCollide(parent *Object, col *Object) {
}

func (a *Money) onInit(o *Object) {
}

func (a *Money) onDestroy() {
}

func (a *Money) getTypeId() bitmask.Bitmask {
	return config.Components.Fish
}
