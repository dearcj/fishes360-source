package main

import (
	"github.com/dearcj/golangproj/msutil"
	"github.com/dearcj/golangproj/network"
	"reflect"
)

type InsertableList []msutil.Insertable
type FList []*data.Action
type ServerEffect data.Action

func IsInt32Chance(i int32) bool {
	return server.Rand() < float64(i)/100.
}

func (s InsertableList) Insert(m *msutil.XServerDataMsg) {
	for _, x := range s {
		x.Insert(m)
	}
}

func (s InsertableList) AddSingle(f msutil.Insertable) InsertableList {
	if f == nil {
		return s
	}
	return append(s, f)
}

func (s InsertableList) Add(ss InsertableList) InsertableList {
	return append(s, ss...)
}

func (s FList) AddSingle(f *ServerEffect) FList {
	if f == nil {
		return s
	}
	return append(s, (*data.Action)(f))
}

func (s FList) Add(ss ...FList) FList {
	for _, x := range ss {
		s = append(s, x...)
	}
	return s
}

func (s FList) Insert(m *msutil.XServerDataMsg) {
	if len(s) > 0 {
		msg := m.WriteToMsg()
		msg.Actions = append(msg.Actions, s...)
	}
}

func (s FList) SetIteration(iterationTime uint32) {
	//for _, x := range s {
	//	x.Iteration = iterationTime
	//}
}

func (s *ServerEffect) Insert(m *msutil.XServerDataMsg) {
	if s != nil {
		msg := m.WriteToMsg()
		msg.Actions = append(msg.Actions, (*data.Action)(s))
	}
}

func (m *ServerEffect) ID(t uint32) *ServerEffect {
	m.TargetID = t
	return m
}

func (m *ServerEffect) V(v float32) *ServerEffect {
	m.Value = v
	return m
}

func (m *ServerEffect) V2(v float32) *ServerEffect {
	m.Value2 = v
	return m
}

func getBaseChar(o *Object) *BaseCharacter {
	unit := o.FindByComponent(config.Components.Unit)
	if unit != nil {
		return &(unit.(*Unit)).BaseCharacter
	}

	return nil
}

func fxlist() FList {
	return FList{}
}

func boolToInt32(b bool) int32 {
	if b {
		return 1
	} else {
		return 0
	}
}

func cf(Type data.ActionType) *ServerEffect {
	return &ServerEffect{
		Type: Type,
	}
}

func UNSAFE_INCREMENT_UINT32_POW_2_STRUCT(s interface{}) interface{} {
	e := reflect.ValueOf(s).Elem()
	p := 1
	for x := 0; x < e.NumField(); x++ {
		e.Field(x).SetUint(uint64(p))
		p = p << 1
	}
	return s
}

func UNSAFE_INCREMENT_UINT32_STRUCT(s interface{}) interface{} {
	e := reflect.ValueOf(s).Elem()
	for x := 0; x < e.NumField(); x++ {
		e.Field(x).SetUint(uint64(x) + 1)
	}
	return s
}

func UNSAFE_INCREMENT_INT32_STRUCT(s interface{}) interface{} {
	e := reflect.ValueOf(s).Elem()
	for x := 0; x < e.NumField(); x++ {
		e.Field(x).SetInt(int64(x) + 1)
	}
	return s
}

func lerp(a, b, alpha float32) float32 {
	return a + (b-a)*alpha
}
