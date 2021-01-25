package main

import (
	"github.com/dearcj/golangproj/msutil"
	pb "github.com/dearcj/golangproj/network"
)

type LocationData struct {
	*pb.LocationData
}

func (p *LocationData) Insert(n *msutil.XServerDataMsg) {
	if p.LocationData != nil {
		n.WriteToMsg().LocationData = p.LocationData
	}
}
