// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package data

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActionType int32

const (
	ActionType_ANY_VALUE ActionType = 0
)

var ActionType_name = map[int32]string{
	0: "ANY_VALUE",
}

var ActionType_value = map[string]int32{
	"ANY_VALUE": 0,
}

func (x ActionType) String() string {
	return proto.EnumName(ActionType_name, int32(x))
}

func (ActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

type CustomObject struct {
	NetworkObject        *NetworkObject `protobuf:"bytes,1,opt,name=networkObject,proto3" json:"networkObject,omitempty"`
	Param1               int32          `protobuf:"varint,2,opt,name=param1,proto3" json:"param1,omitempty"`
	Param2               int32          `protobuf:"varint,3,opt,name=param2,proto3" json:"param2,omitempty"`
	Param3               int32          `protobuf:"varint,4,opt,name=param3,proto3" json:"param3,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CustomObject) Reset()         { *m = CustomObject{} }
func (m *CustomObject) String() string { return proto.CompactTextString(m) }
func (*CustomObject) ProtoMessage()    {}
func (*CustomObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *CustomObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomObject.Unmarshal(m, b)
}
func (m *CustomObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomObject.Marshal(b, m, deterministic)
}
func (m *CustomObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomObject.Merge(m, src)
}
func (m *CustomObject) XXX_Size() int {
	return xxx_messageInfo_CustomObject.Size(m)
}
func (m *CustomObject) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomObject.DiscardUnknown(m)
}

var xxx_messageInfo_CustomObject proto.InternalMessageInfo

func (m *CustomObject) GetNetworkObject() *NetworkObject {
	if m != nil {
		return m.NetworkObject
	}
	return nil
}

func (m *CustomObject) GetParam1() int32 {
	if m != nil {
		return m.Param1
	}
	return 0
}

func (m *CustomObject) GetParam2() int32 {
	if m != nil {
		return m.Param2
	}
	return 0
}

func (m *CustomObject) GetParam3() int32 {
	if m != nil {
		return m.Param3
	}
	return 0
}

type Player struct {
	StartPosition        uint32         `protobuf:"varint,3,opt,name=startPosition,proto3" json:"startPosition,omitempty"`
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	NetworkObject        *NetworkObject `protobuf:"bytes,2,opt,name=networkObject,proto3" json:"networkObject,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetStartPosition() uint32 {
	if m != nil {
		return m.StartPosition
	}
	return 0
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetNetworkObject() *NetworkObject {
	if m != nil {
		return m.NetworkObject
	}
	return nil
}

type Fish struct {
	Hp                   uint32         `protobuf:"varint,1,opt,name=Hp,proto3" json:"Hp,omitempty"`
	Maxhp                uint32         `protobuf:"varint,2,opt,name=Maxhp,proto3" json:"Maxhp,omitempty"`
	CurveInx             uint32         `protobuf:"varint,7,opt,name=CurveInx,proto3" json:"CurveInx,omitempty"`
	FishType             uint32         `protobuf:"varint,3,opt,name=FishType,proto3" json:"FishType,omitempty"`
	StartTime            uint64         `protobuf:"varint,6,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	CurveTime            uint64         `protobuf:"varint,5,opt,name=CurveTime,proto3" json:"CurveTime,omitempty"`
	IsBoss               bool           `protobuf:"varint,8,opt,name=IsBoss,proto3" json:"IsBoss,omitempty"`
	NetworkObject        *NetworkObject `protobuf:"bytes,4,opt,name=networkObject,proto3" json:"networkObject,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Fish) Reset()         { *m = Fish{} }
func (m *Fish) String() string { return proto.CompactTextString(m) }
func (*Fish) ProtoMessage()    {}
func (*Fish) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *Fish) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Fish.Unmarshal(m, b)
}
func (m *Fish) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Fish.Marshal(b, m, deterministic)
}
func (m *Fish) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Fish.Merge(m, src)
}
func (m *Fish) XXX_Size() int {
	return xxx_messageInfo_Fish.Size(m)
}
func (m *Fish) XXX_DiscardUnknown() {
	xxx_messageInfo_Fish.DiscardUnknown(m)
}

var xxx_messageInfo_Fish proto.InternalMessageInfo

func (m *Fish) GetHp() uint32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *Fish) GetMaxhp() uint32 {
	if m != nil {
		return m.Maxhp
	}
	return 0
}

func (m *Fish) GetCurveInx() uint32 {
	if m != nil {
		return m.CurveInx
	}
	return 0
}

func (m *Fish) GetFishType() uint32 {
	if m != nil {
		return m.FishType
	}
	return 0
}

func (m *Fish) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Fish) GetCurveTime() uint64 {
	if m != nil {
		return m.CurveTime
	}
	return 0
}

func (m *Fish) GetIsBoss() bool {
	if m != nil {
		return m.IsBoss
	}
	return false
}

func (m *Fish) GetNetworkObject() *NetworkObject {
	if m != nil {
		return m.NetworkObject
	}
	return nil
}

type NetworkObject struct {
	ID                   uint32   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Type                 uint32   `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkObject) Reset()         { *m = NetworkObject{} }
func (m *NetworkObject) String() string { return proto.CompactTextString(m) }
func (*NetworkObject) ProtoMessage()    {}
func (*NetworkObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{3}
}

func (m *NetworkObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkObject.Unmarshal(m, b)
}
func (m *NetworkObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkObject.Marshal(b, m, deterministic)
}
func (m *NetworkObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkObject.Merge(m, src)
}
func (m *NetworkObject) XXX_Size() int {
	return xxx_messageInfo_NetworkObject.Size(m)
}
func (m *NetworkObject) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkObject.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkObject proto.InternalMessageInfo

func (m *NetworkObject) GetID() uint32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *NetworkObject) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type Action struct {
	Type                 ActionType `protobuf:"varint,1,opt,name=Type,proto3,enum=ActionType" json:"Type,omitempty"`
	Value                float32    `protobuf:"fixed32,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Value2               float32    `protobuf:"fixed32,4,opt,name=Value2,proto3" json:"Value2,omitempty"`
	TargetID             uint32     `protobuf:"varint,3,opt,name=TargetID,proto3" json:"TargetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{4}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetType() ActionType {
	if m != nil {
		return m.Type
	}
	return ActionType_ANY_VALUE
}

func (m *Action) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Action) GetValue2() float32 {
	if m != nil {
		return m.Value2
	}
	return 0
}

func (m *Action) GetTargetID() uint32 {
	if m != nil {
		return m.TargetID
	}
	return 0
}

type ConnectionData struct {
	RoomID               uint32   `protobuf:"varint,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	PlayerID             uint32   `protobuf:"varint,2,opt,name=playerID,proto3" json:"playerID,omitempty"`
	ConMsg               uint32   `protobuf:"varint,3,opt,name=conMsg,proto3" json:"conMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionData) Reset()         { *m = ConnectionData{} }
func (m *ConnectionData) String() string { return proto.CompactTextString(m) }
func (*ConnectionData) ProtoMessage()    {}
func (*ConnectionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{5}
}

func (m *ConnectionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionData.Unmarshal(m, b)
}
func (m *ConnectionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionData.Marshal(b, m, deterministic)
}
func (m *ConnectionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionData.Merge(m, src)
}
func (m *ConnectionData) XXX_Size() int {
	return xxx_messageInfo_ConnectionData.Size(m)
}
func (m *ConnectionData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionData.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionData proto.InternalMessageInfo

func (m *ConnectionData) GetRoomID() uint32 {
	if m != nil {
		return m.RoomID
	}
	return 0
}

func (m *ConnectionData) GetPlayerID() uint32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *ConnectionData) GetConMsg() uint32 {
	if m != nil {
		return m.ConMsg
	}
	return 0
}

type LocationData struct {
	LocationName         string   `protobuf:"bytes,2,opt,name=locationName,proto3" json:"locationName,omitempty"`
	RoomId               uint32   `protobuf:"varint,3,opt,name=roomId,proto3" json:"roomId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationData) Reset()         { *m = LocationData{} }
func (m *LocationData) String() string { return proto.CompactTextString(m) }
func (*LocationData) ProtoMessage()    {}
func (*LocationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{6}
}

func (m *LocationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationData.Unmarshal(m, b)
}
func (m *LocationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationData.Marshal(b, m, deterministic)
}
func (m *LocationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationData.Merge(m, src)
}
func (m *LocationData) XXX_Size() int {
	return xxx_messageInfo_LocationData.Size(m)
}
func (m *LocationData) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationData.DiscardUnknown(m)
}

var xxx_messageInfo_LocationData proto.InternalMessageInfo

func (m *LocationData) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *LocationData) GetRoomId() uint32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

type AccountGeneral struct {
	Money                float32  `protobuf:"fixed32,1,opt,name=Money,proto3" json:"Money,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountGeneral) Reset()         { *m = AccountGeneral{} }
func (m *AccountGeneral) String() string { return proto.CompactTextString(m) }
func (*AccountGeneral) ProtoMessage()    {}
func (*AccountGeneral) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{7}
}

func (m *AccountGeneral) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountGeneral.Unmarshal(m, b)
}
func (m *AccountGeneral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountGeneral.Marshal(b, m, deterministic)
}
func (m *AccountGeneral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountGeneral.Merge(m, src)
}
func (m *AccountGeneral) XXX_Size() int {
	return xxx_messageInfo_AccountGeneral.Size(m)
}
func (m *AccountGeneral) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountGeneral.DiscardUnknown(m)
}

var xxx_messageInfo_AccountGeneral proto.InternalMessageInfo

func (m *AccountGeneral) GetMoney() float32 {
	if m != nil {
		return m.Money
	}
	return 0
}

func (m *AccountGeneral) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type ServerData struct {
	ConData              *ConnectionData `protobuf:"bytes,2,opt,name=conData,proto3" json:"conData,omitempty"`
	AccountGeneral       *AccountGeneral `protobuf:"bytes,3,opt,name=accountGeneral,proto3" json:"accountGeneral,omitempty"`
	LocationData         *LocationData   `protobuf:"bytes,4,opt,name=locationData,proto3" json:"locationData,omitempty"`
	Curves               []string        `protobuf:"bytes,5,rep,name=curves,proto3" json:"curves,omitempty"`
	CustomObjects        []*CustomObject `protobuf:"bytes,6,rep,name=customObjects,proto3" json:"customObjects,omitempty"`
	Fishes               []*Fish         `protobuf:"bytes,7,rep,name=fishes,proto3" json:"fishes,omitempty"`
	Players              []*Player       `protobuf:"bytes,8,rep,name=players,proto3" json:"players,omitempty"`
	Actions              []*Action       `protobuf:"bytes,9,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ServerData) Reset()         { *m = ServerData{} }
func (m *ServerData) String() string { return proto.CompactTextString(m) }
func (*ServerData) ProtoMessage()    {}
func (*ServerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{8}
}

func (m *ServerData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerData.Unmarshal(m, b)
}
func (m *ServerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerData.Marshal(b, m, deterministic)
}
func (m *ServerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerData.Merge(m, src)
}
func (m *ServerData) XXX_Size() int {
	return xxx_messageInfo_ServerData.Size(m)
}
func (m *ServerData) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerData.DiscardUnknown(m)
}

var xxx_messageInfo_ServerData proto.InternalMessageInfo

func (m *ServerData) GetConData() *ConnectionData {
	if m != nil {
		return m.ConData
	}
	return nil
}

func (m *ServerData) GetAccountGeneral() *AccountGeneral {
	if m != nil {
		return m.AccountGeneral
	}
	return nil
}

func (m *ServerData) GetLocationData() *LocationData {
	if m != nil {
		return m.LocationData
	}
	return nil
}

func (m *ServerData) GetCurves() []string {
	if m != nil {
		return m.Curves
	}
	return nil
}

func (m *ServerData) GetCustomObjects() []*CustomObject {
	if m != nil {
		return m.CustomObjects
	}
	return nil
}

func (m *ServerData) GetFishes() []*Fish {
	if m != nil {
		return m.Fishes
	}
	return nil
}

func (m *ServerData) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *ServerData) GetActions() []*Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type Command struct {
	CommandId            int32    `protobuf:"varint,1,opt,name=CommandId,proto3" json:"CommandId,omitempty"`
	Params               []int32  `protobuf:"varint,2,rep,packed,name=Params,proto3" json:"Params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Command) Reset()         { *m = Command{} }
func (m *Command) String() string { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()    {}
func (*Command) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{9}
}

func (m *Command) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Command.Unmarshal(m, b)
}
func (m *Command) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Command.Marshal(b, m, deterministic)
}
func (m *Command) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Command.Merge(m, src)
}
func (m *Command) XXX_Size() int {
	return xxx_messageInfo_Command.Size(m)
}
func (m *Command) XXX_DiscardUnknown() {
	xxx_messageInfo_Command.DiscardUnknown(m)
}

var xxx_messageInfo_Command proto.InternalMessageInfo

func (m *Command) GetCommandId() int32 {
	if m != nil {
		return m.CommandId
	}
	return 0
}

func (m *Command) GetParams() []int32 {
	if m != nil {
		return m.Params
	}
	return nil
}

type ClientData struct {
	RoomId               int32      `protobuf:"varint,2,opt,name=roomId,proto3" json:"roomId,omitempty"`
	Commands             []*Command `protobuf:"bytes,3,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ClientData) Reset()         { *m = ClientData{} }
func (m *ClientData) String() string { return proto.CompactTextString(m) }
func (*ClientData) ProtoMessage()    {}
func (*ClientData) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{10}
}

func (m *ClientData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientData.Unmarshal(m, b)
}
func (m *ClientData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientData.Marshal(b, m, deterministic)
}
func (m *ClientData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientData.Merge(m, src)
}
func (m *ClientData) XXX_Size() int {
	return xxx_messageInfo_ClientData.Size(m)
}
func (m *ClientData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientData.DiscardUnknown(m)
}

var xxx_messageInfo_ClientData proto.InternalMessageInfo

func (m *ClientData) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *ClientData) GetCommands() []*Command {
	if m != nil {
		return m.Commands
	}
	return nil
}

func init() {
	proto.RegisterEnum("ActionType", ActionType_name, ActionType_value)
	proto.RegisterType((*CustomObject)(nil), "CustomObject")
	proto.RegisterType((*Player)(nil), "Player")
	proto.RegisterType((*Fish)(nil), "Fish")
	proto.RegisterType((*NetworkObject)(nil), "NetworkObject")
	proto.RegisterType((*Action)(nil), "Action")
	proto.RegisterType((*ConnectionData)(nil), "ConnectionData")
	proto.RegisterType((*LocationData)(nil), "LocationData")
	proto.RegisterType((*AccountGeneral)(nil), "AccountGeneral")
	proto.RegisterType((*ServerData)(nil), "ServerData")
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterType((*ClientData)(nil), "ClientData")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x3e, 0x76, 0x12, 0x27, 0x99, 0xd6, 0x39, 0x47, 0xab, 0x23, 0x64, 0xf1, 0x23, 0x8c, 0xd5,
	0x8b, 0xc0, 0x85, 0xa5, 0x26, 0x48, 0x5c, 0xa2, 0x34, 0x01, 0xea, 0xaa, 0x2d, 0xd5, 0xf6, 0x47,
	0x42, 0x42, 0x42, 0x5b, 0x67, 0x69, 0x03, 0xf1, 0xae, 0xe5, 0x75, 0x4a, 0xfb, 0x0e, 0xbc, 0x1e,
	0xaf, 0xc1, 0x33, 0xa0, 0x9d, 0x5d, 0x3b, 0x31, 0xbd, 0xe9, 0xdd, 0x7e, 0xdf, 0x8c, 0xbf, 0x9d,
	0x6f, 0x3c, 0x3b, 0x00, 0x73, 0x56, 0xb2, 0x38, 0x2f, 0x64, 0x29, 0xa3, 0x9f, 0x0e, 0x6c, 0x4f,
	0x57, 0xaa, 0x94, 0xd9, 0xc7, 0xcb, 0x6f, 0x3c, 0x2d, 0xc9, 0x6b, 0xf0, 0x05, 0x2f, 0x7f, 0xc8,
	0xe2, 0xbb, 0x21, 0x02, 0x27, 0x74, 0x86, 0x5b, 0xa3, 0x41, 0x7c, 0xbc, 0xc9, 0xd2, 0x66, 0x12,
	0x79, 0x04, 0x5e, 0xce, 0x0a, 0x96, 0xed, 0x06, 0x6e, 0xe8, 0x0c, 0x3b, 0xd4, 0xa2, 0x9a, 0x1f,
	0x05, 0xad, 0x0d, 0x7e, 0x54, 0xf3, 0xe3, 0xa0, 0xbd, 0xc1, 0x8f, 0xa3, 0x5b, 0xf0, 0x4e, 0x96,
	0xec, 0x8e, 0x17, 0x64, 0x07, 0x7c, 0x55, 0xb2, 0xa2, 0x3c, 0x91, 0x6a, 0x51, 0x2e, 0xa4, 0x40,
	0x01, 0x9f, 0x36, 0x49, 0x42, 0xa0, 0x2d, 0x58, 0xc6, 0xb1, 0xc8, 0x3e, 0xc5, 0xf3, 0x7d, 0x07,
	0xee, 0x03, 0x1c, 0x44, 0xbf, 0x1d, 0x68, 0xbf, 0x5f, 0xa8, 0x6b, 0x32, 0x00, 0x77, 0x3f, 0x47,
	0x41, 0x9f, 0xba, 0xfb, 0x39, 0xf9, 0x1f, 0x3a, 0x47, 0xec, 0xf6, 0x3a, 0x47, 0x19, 0x9f, 0x1a,
	0x40, 0x1e, 0x43, 0x6f, 0xba, 0x2a, 0x6e, 0x78, 0x22, 0x6e, 0x83, 0x2e, 0x06, 0x6a, 0xac, 0x63,
	0x5a, 0xe9, 0xec, 0x2e, 0xe7, 0xb6, 0xea, 0x1a, 0x93, 0xa7, 0xd0, 0x3f, 0xd5, 0x0e, 0xce, 0x16,
	0x19, 0x0f, 0xbc, 0xd0, 0x19, 0xb6, 0xe9, 0x9a, 0xd0, 0x51, 0x54, 0xc1, 0x68, 0xc7, 0x44, 0x6b,
	0x42, 0x37, 0x2d, 0x51, 0x7b, 0x52, 0xa9, 0xa0, 0x17, 0x3a, 0xc3, 0x1e, 0xb5, 0xe8, 0xbe, 0xe1,
	0xf6, 0x43, 0x0c, 0x8f, 0xc1, 0x6f, 0xc4, 0xb5, 0xf1, 0x64, 0x56, 0x19, 0x4f, 0x66, 0xba, 0xb7,
	0x68, 0xc1, 0xf8, 0xc6, 0x73, 0xa4, 0xc0, 0x9b, 0xa4, 0xd8, 0xf9, 0xe7, 0x36, 0xaa, 0xf3, 0x07,
	0xa3, 0xad, 0xd8, 0xd0, 0x9a, 0x32, 0xa9, 0xba, 0x6f, 0x17, 0x6c, 0xb9, 0x32, 0xdf, 0xbb, 0xd4,
	0x00, 0xed, 0x01, 0x0f, 0x23, 0x2c, 0xd2, 0xa5, 0x16, 0xe9, 0x9e, 0x9d, 0xb1, 0xe2, 0x8a, 0x97,
	0xc9, 0xac, 0xea, 0x59, 0x85, 0xa3, 0xcf, 0x30, 0x98, 0x4a, 0x21, 0x38, 0xde, 0x30, 0x63, 0x25,
	0xd3, 0x2a, 0x85, 0x94, 0x59, 0x5d, 0xae, 0x45, 0x5a, 0x25, 0xc7, 0xf1, 0x49, 0x66, 0xb6, 0xec,
	0x1a, 0xeb, 0x6f, 0x52, 0x29, 0x8e, 0xd4, 0x95, 0xd5, 0xb7, 0x28, 0x3a, 0x80, 0xed, 0x43, 0x99,
	0xb2, 0x5a, 0x3b, 0x82, 0xed, 0xa5, 0xc5, 0xc7, 0x7a, 0xb4, 0x5c, 0x1c, 0xad, 0x06, 0x57, 0xdf,
	0x3f, 0xaf, 0xb4, 0x0c, 0x8a, 0xf6, 0x60, 0x30, 0x49, 0x53, 0xb9, 0x12, 0xe5, 0x07, 0x2e, 0x78,
	0xc1, 0x96, 0x38, 0x3d, 0x52, 0xf0, 0x3b, 0x2c, 0xd4, 0xa5, 0x06, 0xe8, 0x3a, 0xcf, 0x15, 0x2f,
	0xc4, 0x5a, 0xbf, 0xc6, 0xd1, 0x2f, 0x17, 0xe0, 0x94, 0x17, 0x37, 0xbc, 0xc0, 0x72, 0x5e, 0x42,
	0x37, 0x35, 0x95, 0xd9, 0x39, 0xfe, 0x37, 0x6e, 0x36, 0x83, 0x56, 0x71, 0xf2, 0x06, 0x06, 0xac,
	0x71, 0x3b, 0x56, 0xa7, 0xbf, 0x68, 0x16, 0x45, 0xff, 0x4a, 0x23, 0xbb, 0x6b, 0xcb, 0x78, 0x91,
	0x99, 0x1f, 0x3f, 0xde, 0xec, 0x0b, 0x6d, 0xa4, 0x60, 0x37, 0xf5, 0x60, 0xaa, 0xa0, 0x13, 0xb6,
	0x86, 0x7d, 0x6a, 0x11, 0x19, 0x83, 0x9f, 0x6e, 0xac, 0x13, 0x15, 0x78, 0x61, 0x0b, 0xb5, 0x36,
	0x97, 0x0c, 0x6d, 0xe6, 0x90, 0x67, 0xe0, 0x7d, 0x5d, 0xa8, 0x6b, 0xae, 0x82, 0x2e, 0x66, 0x77,
	0x62, 0xfd, 0x5e, 0xa8, 0x25, 0xc9, 0x0b, 0xe8, 0x9a, 0xbf, 0xa8, 0x07, 0x5f, 0xc7, 0xbb, 0xb1,
	0x59, 0x12, 0xb4, 0xe2, 0x75, 0x0a, 0xc3, 0x8e, 0xa8, 0xa0, 0x6f, 0x53, 0xcc, 0x40, 0xd2, 0x8a,
	0x8f, 0xde, 0x42, 0x77, 0x2a, 0xb3, 0x8c, 0x89, 0x39, 0x3e, 0x33, 0x73, 0x4c, 0xe6, 0xf8, 0x63,
	0x3a, 0x74, 0x4d, 0x68, 0x6b, 0x27, 0x7a, 0x1b, 0xa9, 0xc0, 0x0d, 0x5b, 0x7a, 0x37, 0x19, 0x14,
	0x1d, 0x00, 0x4c, 0x97, 0x0b, 0x2e, 0xca, 0xc6, 0x08, 0xce, 0xab, 0x8d, 0x67, 0x10, 0xd9, 0x81,
	0x5e, 0x6a, 0xa4, 0x54, 0xd0, 0xc2, 0x52, 0x7a, 0xb1, 0xd5, 0xa6, 0x75, 0xe4, 0xd5, 0x13, 0x80,
	0xf5, 0x83, 0x21, 0x3e, 0xf4, 0x27, 0xc7, 0x9f, 0xbe, 0x5c, 0x4c, 0x0e, 0xcf, 0xdf, 0xfd, 0xf7,
	0xcf, 0xa5, 0x87, 0xab, 0x79, 0xfc, 0x27, 0x00, 0x00, 0xff, 0xff, 0xff, 0x73, 0xf0, 0xa1, 0xa8,
	0x05, 0x00, 0x00,
}
