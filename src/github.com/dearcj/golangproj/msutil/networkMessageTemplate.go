package msutil

import "github.com/mauricelam/genny/generic"
import x "github.com/dearcj/golangproj/network"
import "reflect"

type GenericData generic.Type

var _ x.ServerData

type GenericDataMsg struct {
	Changed bool
	data    *GenericData
	backup  *GenericData
}

func (n *GenericDataMsg) WriteToMsg() *GenericData {
	n.Changed = true
	return n.data
}

func (n *GenericDataMsg) Reset() {
	*n.data = *n.backup
}

func CreateGenericData() *GenericData {
	var x GenericData
	g := reflect.New(reflect.TypeOf(x).Elem()).Interface()
	return g.(*GenericData)
}

func CreateGenericDataMsg() GenericDataMsg {
	return GenericDataMsg{
		backup: CreateGenericData(),
		data:   CreateGenericData()}
}
