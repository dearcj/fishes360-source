package msutil

import x "github.com/dearcj/golangproj/network"

type XServerDataMsg struct {
	UniqueID []byte
	changed  bool
	data     *x.ServerData
	backup   *x.ServerData
}

func (n *XServerDataMsg) WriteToMsg() *x.ServerData {
	n.changed = true
	return n.data
}

func (n *XServerDataMsg) Reset() {
	*n.data = *n.backup
	n.changed = false
}

func (n *XServerDataMsg) IsChanged() bool {
	return n.changed
}

func CreateXServerData() *x.ServerData {
	return &x.ServerData{}
}

func CreateXServerDataMsg(UniqueID []byte) XServerDataMsg {
	return XServerDataMsg{
		UniqueID: UniqueID,
		backup:   CreateXServerData(),
		data:     CreateXServerData()}
}
