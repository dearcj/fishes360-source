package msutil

import "reflect"

func NewOrSame(x interface{}) interface{} {
	if reflect.ValueOf(x).IsNil() {
		return reflect.New(reflect.TypeOf(x).Elem()).Interface()
	} else {
		return x
	}
}

type Insertable interface {
	Insert(s *XServerDataMsg)
}

type DataLink interface {
	Insertable
}
