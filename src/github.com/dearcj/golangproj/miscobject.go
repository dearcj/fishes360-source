package main

import (
	"github.com/dearcj/golangproj/bitmask"
	pb "github.com/dearcj/golangproj/network"
)

type MiscObject struct {
	gfxId int32
}

func (a *MiscObject) mutateState(d *Object, no *pb.NetworkObject) {

}
func (a *MiscObject) getNetworkObject(o *Object) *pb.CustomObject {

	return &pb.CustomObject{
		NetworkObject: &o.NetworkObject,
		Param1:        a.gfxId,
	}
}

func (a *MiscObject) pProcess(b *Object, dt float64) {

}

func (a *MiscObject) onCollide(parent *Object, col *Object) {

}

func (a *MiscObject) onInit(o *Object) {
}

func (a *MiscObject) onDestroy() {
}

func (a *MiscObject) getTypeId() bitmask.Bitmask {
	return config.Components.MiscObject
}
