// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: GameCtl.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Status int32

const (
	Status_ON_START     Status = 0
	Status_ON_WAIT      Status = 1
	Status_ON_HOST_TURN Status = 2
	Status_ON_DUEL_TURN Status = 3
	Status_ON_END       Status = 4
)

var Status_name = map[int32]string{
	0: "ON_START",
	1: "ON_WAIT",
	2: "ON_HOST_TURN",
	3: "ON_DUEL_TURN",
	4: "ON_END",
}

var Status_value = map[string]int32{
	"ON_START":     0,
	"ON_WAIT":      1,
	"ON_HOST_TURN": 2,
	"ON_DUEL_TURN": 3,
	"ON_END":       4,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{0}
}

type RoomListRequest struct {
	Requirement          string   `protobuf:"bytes,1,opt,name=requirement,proto3" json:"requirement,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomListRequest) Reset()         { *m = RoomListRequest{} }
func (m *RoomListRequest) String() string { return proto.CompactTextString(m) }
func (*RoomListRequest) ProtoMessage()    {}
func (*RoomListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{0}
}
func (m *RoomListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomListRequest.Unmarshal(m, b)
}
func (m *RoomListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomListRequest.Marshal(b, m, deterministic)
}
func (m *RoomListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomListRequest.Merge(m, src)
}
func (m *RoomListRequest) XXX_Size() int {
	return xxx_messageInfo_RoomListRequest.Size(m)
}
func (m *RoomListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoomListRequest proto.InternalMessageInfo

func (m *RoomListRequest) GetRequirement() string {
	if m != nil {
		return m.Requirement
	}
	return ""
}

type RoomListResponse struct {
	Result               []*Room  `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomListResponse) Reset()         { *m = RoomListResponse{} }
func (m *RoomListResponse) String() string { return proto.CompactTextString(m) }
func (*RoomListResponse) ProtoMessage()    {}
func (*RoomListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{1}
}
func (m *RoomListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomListResponse.Unmarshal(m, b)
}
func (m *RoomListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomListResponse.Marshal(b, m, deterministic)
}
func (m *RoomListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomListResponse.Merge(m, src)
}
func (m *RoomListResponse) XXX_Size() int {
	return xxx_messageInfo_RoomListResponse.Size(m)
}
func (m *RoomListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RoomListResponse proto.InternalMessageInfo

func (m *RoomListResponse) GetResult() []*Room {
	if m != nil {
		return m.Result
	}
	return nil
}

type RoomCreateRequest struct {
	HostId               string   `protobuf:"bytes,1,opt,name=HostId,proto3" json:"HostId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomCreateRequest) Reset()         { *m = RoomCreateRequest{} }
func (m *RoomCreateRequest) String() string { return proto.CompactTextString(m) }
func (*RoomCreateRequest) ProtoMessage()    {}
func (*RoomCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{2}
}
func (m *RoomCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomCreateRequest.Unmarshal(m, b)
}
func (m *RoomCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomCreateRequest.Marshal(b, m, deterministic)
}
func (m *RoomCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomCreateRequest.Merge(m, src)
}
func (m *RoomCreateRequest) XXX_Size() int {
	return xxx_messageInfo_RoomCreateRequest.Size(m)
}
func (m *RoomCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoomCreateRequest proto.InternalMessageInfo

func (m *RoomCreateRequest) GetHostId() string {
	if m != nil {
		return m.HostId
	}
	return ""
}

type RoomRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomRequest) Reset()         { *m = RoomRequest{} }
func (m *RoomRequest) String() string { return proto.CompactTextString(m) }
func (*RoomRequest) ProtoMessage()    {}
func (*RoomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{3}
}
func (m *RoomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomRequest.Unmarshal(m, b)
}
func (m *RoomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomRequest.Marshal(b, m, deterministic)
}
func (m *RoomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomRequest.Merge(m, src)
}
func (m *RoomRequest) XXX_Size() int {
	return xxx_messageInfo_RoomRequest.Size(m)
}
func (m *RoomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RoomRequest proto.InternalMessageInfo

func (m *RoomRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Room struct {
	Key                  string        `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	HostId               string        `protobuf:"bytes,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	DuelerId             string        `protobuf:"bytes,3,opt,name=dueler_id,json=duelerId,proto3" json:"dueler_id,omitempty"`
	Status               Status        `protobuf:"varint,4,opt,name=status,proto3,enum=RoomStatus.Status" json:"status,omitempty"`
	Round                int32         `protobuf:"varint,5,opt,name=round,proto3" json:"round,omitempty"`
	Cell                 int32         `protobuf:"varint,6,opt,name=cell,proto3" json:"cell,omitempty"`
	CellStatus           []*CellStatus `protobuf:"bytes,7,rep,name=cell_status,json=cellStatus,proto3" json:"cell_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Room) Reset()         { *m = Room{} }
func (m *Room) String() string { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()    {}
func (*Room) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{4}
}
func (m *Room) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Room.Unmarshal(m, b)
}
func (m *Room) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Room.Marshal(b, m, deterministic)
}
func (m *Room) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Room.Merge(m, src)
}
func (m *Room) XXX_Size() int {
	return xxx_messageInfo_Room.Size(m)
}
func (m *Room) XXX_DiscardUnknown() {
	xxx_messageInfo_Room.DiscardUnknown(m)
}

var xxx_messageInfo_Room proto.InternalMessageInfo

func (m *Room) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Room) GetHostId() string {
	if m != nil {
		return m.HostId
	}
	return ""
}

func (m *Room) GetDuelerId() string {
	if m != nil {
		return m.DuelerId
	}
	return ""
}

func (m *Room) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_ON_START
}

func (m *Room) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Room) GetCell() int32 {
	if m != nil {
		return m.Cell
	}
	return 0
}

func (m *Room) GetCellStatus() []*CellStatus {
	if m != nil {
		return m.CellStatus
	}
	return nil
}

type CellStatus struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Turn                 int32    `protobuf:"varint,2,opt,name=turn,proto3" json:"turn,omitempty"`
	CellNum              int32    `protobuf:"varint,3,opt,name=cell_num,json=cellNum,proto3" json:"cell_num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellStatus) Reset()         { *m = CellStatus{} }
func (m *CellStatus) String() string { return proto.CompactTextString(m) }
func (*CellStatus) ProtoMessage()    {}
func (*CellStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{5}
}
func (m *CellStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellStatus.Unmarshal(m, b)
}
func (m *CellStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellStatus.Marshal(b, m, deterministic)
}
func (m *CellStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellStatus.Merge(m, src)
}
func (m *CellStatus) XXX_Size() int {
	return xxx_messageInfo_CellStatus.Size(m)
}
func (m *CellStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_CellStatus.DiscardUnknown(m)
}

var xxx_messageInfo_CellStatus proto.InternalMessageInfo

func (m *CellStatus) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CellStatus) GetTurn() int32 {
	if m != nil {
		return m.Turn
	}
	return 0
}

func (m *CellStatus) GetCellNum() int32 {
	if m != nil {
		return m.CellNum
	}
	return 0
}

type CreateCredReq struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCredReq) Reset()         { *m = CreateCredReq{} }
func (m *CreateCredReq) String() string { return proto.CompactTextString(m) }
func (*CreateCredReq) ProtoMessage()    {}
func (*CreateCredReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{6}
}
func (m *CreateCredReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredReq.Unmarshal(m, b)
}
func (m *CreateCredReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredReq.Marshal(b, m, deterministic)
}
func (m *CreateCredReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredReq.Merge(m, src)
}
func (m *CreateCredReq) XXX_Size() int {
	return xxx_messageInfo_CreateCredReq.Size(m)
}
func (m *CreateCredReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredReq proto.InternalMessageInfo

func (m *CreateCredReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateCredReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateCredReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Cred struct {
	File                 string   `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cred) Reset()         { *m = Cred{} }
func (m *Cred) String() string { return proto.CompactTextString(m) }
func (*Cred) ProtoMessage()    {}
func (*Cred) Descriptor() ([]byte, []int) {
	return fileDescriptor_844dd485888a1988, []int{7}
}
func (m *Cred) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cred.Unmarshal(m, b)
}
func (m *Cred) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cred.Marshal(b, m, deterministic)
}
func (m *Cred) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cred.Merge(m, src)
}
func (m *Cred) XXX_Size() int {
	return xxx_messageInfo_Cred.Size(m)
}
func (m *Cred) XXX_DiscardUnknown() {
	xxx_messageInfo_Cred.DiscardUnknown(m)
}

var xxx_messageInfo_Cred proto.InternalMessageInfo

func (m *Cred) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func init() {
	proto.RegisterEnum("RoomStatus.Status", Status_name, Status_value)
	proto.RegisterType((*RoomListRequest)(nil), "RoomStatus.RoomListRequest")
	proto.RegisterType((*RoomListResponse)(nil), "RoomStatus.RoomListResponse")
	proto.RegisterType((*RoomCreateRequest)(nil), "RoomStatus.RoomCreateRequest")
	proto.RegisterType((*RoomRequest)(nil), "RoomStatus.RoomRequest")
	proto.RegisterType((*Room)(nil), "RoomStatus.Room")
	proto.RegisterType((*CellStatus)(nil), "RoomStatus.CellStatus")
	proto.RegisterType((*CreateCredReq)(nil), "RoomStatus.CreateCredReq")
	proto.RegisterType((*Cred)(nil), "RoomStatus.Cred")
}

func init() { proto.RegisterFile("GameCtl.proto", fileDescriptor_844dd485888a1988) }

var fileDescriptor_844dd485888a1988 = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x6e, 0xdb, 0x38,
	0x10, 0x5e, 0xfa, 0x3f, 0xe3, 0x38, 0xd1, 0x12, 0x81, 0xe3, 0xc8, 0x59, 0xac, 0x57, 0x27, 0xc3,
	0x0b, 0xd8, 0x59, 0xe7, 0xb0, 0x80, 0xb1, 0x97, 0xac, 0x13, 0x24, 0x06, 0x12, 0x3b, 0x90, 0x6d,
	0x04, 0xe9, 0xa1, 0x86, 0x62, 0xd1, 0x89, 0x50, 0x49, 0x74, 0x48, 0xaa, 0x45, 0x8e, 0xed, 0x2b,
	0xf4, 0x2d, 0xfa, 0x3a, 0x3d, 0xf4, 0xd4, 0x5b, 0x1f, 0xa4, 0x20, 0x45, 0xff, 0xc5, 0x76, 0x4f,
	0x9a, 0x99, 0x8f, 0xfc, 0x3e, 0xce, 0xe8, 0x23, 0xa1, 0x70, 0xe9, 0x04, 0xa4, 0x2d, 0xfc, 0xfa,
	0x94, 0x51, 0x41, 0x31, 0xd8, 0x94, 0x06, 0x7d, 0xe1, 0x88, 0x88, 0x9b, 0xe5, 0x47, 0x4a, 0x1f,
	0x7d, 0xd2, 0x50, 0xc8, 0x43, 0x34, 0x69, 0x90, 0x60, 0x2a, 0x5e, 0xe2, 0x85, 0xe6, 0xb1, 0x06,
	0x9d, 0xa9, 0xd7, 0x70, 0xc2, 0x90, 0x0a, 0x47, 0x78, 0x34, 0xe4, 0x31, 0x6a, 0x9d, 0xc2, 0xbe,
	0x24, 0xba, 0xf6, 0xb8, 0xb0, 0xc9, 0x73, 0x44, 0xb8, 0xc0, 0x15, 0xc8, 0x33, 0xf2, 0x1c, 0x79,
	0x8c, 0x04, 0x24, 0x14, 0x25, 0x54, 0x41, 0xd5, 0x1d, 0x7b, 0xb9, 0x64, 0xfd, 0x07, 0xc6, 0x62,
	0x13, 0x9f, 0xd2, 0x90, 0x13, 0x5c, 0x85, 0x0c, 0x23, 0x3c, 0xf2, 0xe5, 0x86, 0x64, 0x35, 0xdf,
	0x34, 0xea, 0x8b, 0x03, 0xaa, 0xd0, 0xd6, 0xb8, 0xf5, 0x37, 0xfc, 0x2e, 0xf3, 0x36, 0x23, 0x8e,
	0x20, 0x33, 0xd1, 0x22, 0x64, 0xae, 0x28, 0x17, 0x1d, 0x57, 0xeb, 0xe9, 0xcc, 0xfa, 0x13, 0xf2,
	0x6a, 0xb3, 0x5e, 0x66, 0x40, 0xf2, 0x1d, 0x79, 0xd1, 0x6b, 0x64, 0x68, 0x7d, 0x47, 0x90, 0x92,
	0x2b, 0xd6, 0x21, 0x7c, 0x08, 0xd9, 0x27, 0xca, 0xc5, 0xc8, 0x73, 0x4b, 0x89, 0x98, 0xf4, 0x49,
	0x91, 0xe2, 0x32, 0xec, 0xb8, 0x11, 0xf1, 0x09, 0x93, 0x50, 0x52, 0x41, 0xb9, 0xb8, 0xd0, 0x71,
	0x71, 0x0d, 0x32, 0x5c, 0x9d, 0xba, 0x94, 0xaa, 0xa0, 0xea, 0x5e, 0x13, 0x2f, 0x37, 0x12, 0x7f,
	0x6c, 0xbd, 0x02, 0x1f, 0x40, 0x9a, 0xd1, 0x28, 0x74, 0x4b, 0xe9, 0x0a, 0xaa, 0xa6, 0xed, 0x38,
	0xc1, 0x18, 0x52, 0x63, 0xe2, 0xfb, 0xa5, 0x8c, 0x2a, 0xaa, 0x18, 0xff, 0x0b, 0x79, 0xf9, 0x1d,
	0x69, 0xea, 0xac, 0x9a, 0x51, 0x71, 0x99, 0xba, 0x4d, 0x7c, 0x5f, 0xd3, 0xc3, 0x78, 0x1e, 0x5b,
	0x37, 0x00, 0x0b, 0x64, 0x43, 0x93, 0x18, 0x52, 0x22, 0x62, 0xa1, 0xea, 0x30, 0x6d, 0xab, 0x18,
	0x1f, 0x41, 0x4e, 0x89, 0x85, 0x51, 0xa0, 0xda, 0x4b, 0xdb, 0x59, 0x99, 0x77, 0xa3, 0xc0, 0xba,
	0x83, 0x42, 0x3c, 0xf8, 0x36, 0x23, 0xae, 0x4d, 0x9e, 0xf1, 0x1e, 0x24, 0x3a, 0x53, 0x4d, 0x98,
	0xe8, 0x4c, 0xb1, 0x09, 0xb9, 0x21, 0x27, 0x2c, 0x74, 0x02, 0xa2, 0xa7, 0x36, 0xcf, 0x25, 0x76,
	0xeb, 0x70, 0xfe, 0x81, 0xb2, 0xf9, 0xd8, 0x66, 0xb9, 0x65, 0x42, 0x4a, 0x52, 0xca, 0xf3, 0x4c,
	0x3c, 0x9f, 0x68, 0x46, 0x15, 0xd7, 0xfa, 0x90, 0xd1, 0xe7, 0xdf, 0x85, 0x5c, 0xaf, 0x3b, 0xea,
	0x0f, 0xce, 0xec, 0x81, 0xf1, 0x1b, 0xce, 0x43, 0xb6, 0xd7, 0x1d, 0xdd, 0x9d, 0x75, 0x06, 0x06,
	0xc2, 0x06, 0xec, 0xf6, 0xba, 0xa3, 0xab, 0x5e, 0x7f, 0x30, 0x1a, 0x0c, 0xed, 0xae, 0x91, 0xd0,
	0x95, 0xf3, 0xe1, 0xc5, 0x75, 0x5c, 0x49, 0x62, 0x80, 0x4c, 0xaf, 0x3b, 0xba, 0xe8, 0x9e, 0x1b,
	0xa9, 0xe6, 0xc7, 0x24, 0x2c, 0xdd, 0x01, 0x7c, 0x0f, 0xa0, 0x1d, 0x25, 0xcd, 0xf0, 0xc7, 0x6b,
	0xf7, 0xad, 0xb8, 0xcd, 0x5c, 0x33, 0xa7, 0x65, 0x7e, 0xfa, 0xfa, 0xe3, 0x73, 0xe2, 0xc0, 0xda,
	0x6f, 0xbc, 0xff, 0xa7, 0xc1, 0x28, 0x0d, 0x1a, 0x63, 0xb5, 0xa3, 0x85, 0x6a, 0xf8, 0x01, 0xf2,
	0x97, 0x44, 0xcc, 0x1c, 0x8f, 0xcb, 0xaf, 0x37, 0x2f, 0x5d, 0x1e, 0xf3, 0x78, 0x33, 0x18, 0x5f,
	0x12, 0xab, 0xa4, 0x54, 0xb0, 0x55, 0x98, 0xab, 0xf8, 0x1e, 0x17, 0x52, 0xe3, 0x1e, 0xb0, 0xd6,
	0x68, 0x47, 0x8c, 0x91, 0x50, 0x74, 0xc2, 0x09, 0xc5, 0x87, 0x6b, 0x97, 0x68, 0x6b, 0x03, 0xeb,
	0xd4, 0x5e, 0x38, 0xa1, 0x31, 0x35, 0x9c, 0x13, 0x9f, 0xe8, 0xc9, 0x6c, 0xa5, 0x2c, 0xd6, 0xe3,
	0x87, 0xa2, 0x3e, 0x7b, 0x45, 0xea, 0x17, 0xf2, 0x15, 0xd9, 0x30, 0x19, 0x57, 0xb1, 0xb5, 0x50,
	0xad, 0xf9, 0x0d, 0xcd, 0xfe, 0x01, 0x23, 0x4e, 0x80, 0xdf, 0x42, 0x41, 0x37, 0xa1, 0x0b, 0xbf,
	0x10, 0xdb, 0xe8, 0xfc, 0x0d, 0x62, 0x5c, 0x31, 0xb5, 0x50, 0xed, 0x04, 0x61, 0x0f, 0x8c, 0xe1,
	0xd4, 0xd5, 0xff, 0x58, 0x4b, 0x6c, 0x61, 0xda, 0xaa, 0xf0, 0x97, 0x52, 0x28, 0x5b, 0xc5, 0x57,
	0x0a, 0x8d, 0x48, 0x31, 0xb7, 0x50, 0xad, 0x8a, 0x4e, 0x50, 0x73, 0x02, 0x79, 0x69, 0x67, 0x4f,
	0xf0, 0xb3, 0x48, 0x3c, 0xe1, 0xbb, 0x99, 0xbb, 0x94, 0xc7, 0x8f, 0x56, 0xb8, 0x97, 0xaf, 0xd3,
	0xea, 0x8f, 0x91, 0xc5, 0xd5, 0x96, 0xc6, 0x8c, 0xb8, 0x0b, 0x67, 0x9d, 0xa0, 0xff, 0xcd, 0x37,
	0x69, 0x35, 0xef, 0x2f, 0x09, 0xe3, 0xd6, 0x77, 0x5e, 0xda, 0xbe, 0x57, 0xbf, 0x95, 0xf9, 0x0d,
	0x75, 0x1f, 0x32, 0x0a, 0x39, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x38, 0xec, 0xea, 0xd5, 0xfa,
	0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RoomStatusClient is the client API for RoomStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomStatusClient interface {
	// rpc CreateCred (CreateCredReq) returns (stream Cred) {
	//      option (google.api.http) = {
	//         post: "/v1/cred/create"
	//         body: "*"
	//     };
	// }
	CreateRoom(ctx context.Context, in *RoomCreateRequest, opts ...grpc.CallOption) (*Room, error)
	GetRoomList(ctx context.Context, in *RoomListRequest, opts ...grpc.CallOption) (*RoomListResponse, error)
	GetRoomCurrentInfo(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*Room, error)
	// rpc GetRoomStream (RoomRequest) returns (stream CellStatus){
	//     option (google.api.http) = {
	//         post: "/v1/room/stream"
	//         body: "*"
	//     };
	// };
	// rpc UpdateRoomStatus (CellStatus) returns (google.protobuf.Empty){
	//     option (google.api.http) = {
	//         post: "/v1/room/update"
	//         body: "*"
	//     };
	// };
	DeleteRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type roomStatusClient struct {
	cc *grpc.ClientConn
}

func NewRoomStatusClient(cc *grpc.ClientConn) RoomStatusClient {
	return &roomStatusClient{cc}
}

func (c *roomStatusClient) CreateRoom(ctx context.Context, in *RoomCreateRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/RoomStatus.RoomStatus/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomStatusClient) GetRoomList(ctx context.Context, in *RoomListRequest, opts ...grpc.CallOption) (*RoomListResponse, error) {
	out := new(RoomListResponse)
	err := c.cc.Invoke(ctx, "/RoomStatus.RoomStatus/GetRoomList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomStatusClient) GetRoomCurrentInfo(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*Room, error) {
	out := new(Room)
	err := c.cc.Invoke(ctx, "/RoomStatus.RoomStatus/GetRoomCurrentInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomStatusClient) DeleteRoom(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/RoomStatus.RoomStatus/DeleteRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomStatusServer is the server API for RoomStatus service.
type RoomStatusServer interface {
	// rpc CreateCred (CreateCredReq) returns (stream Cred) {
	//      option (google.api.http) = {
	//         post: "/v1/cred/create"
	//         body: "*"
	//     };
	// }
	CreateRoom(context.Context, *RoomCreateRequest) (*Room, error)
	GetRoomList(context.Context, *RoomListRequest) (*RoomListResponse, error)
	GetRoomCurrentInfo(context.Context, *RoomRequest) (*Room, error)
	// rpc GetRoomStream (RoomRequest) returns (stream CellStatus){
	//     option (google.api.http) = {
	//         post: "/v1/room/stream"
	//         body: "*"
	//     };
	// };
	// rpc UpdateRoomStatus (CellStatus) returns (google.protobuf.Empty){
	//     option (google.api.http) = {
	//         post: "/v1/room/update"
	//         body: "*"
	//     };
	// };
	DeleteRoom(context.Context, *RoomRequest) (*types.Empty, error)
}

// UnimplementedRoomStatusServer can be embedded to have forward compatible implementations.
type UnimplementedRoomStatusServer struct {
}

func (*UnimplementedRoomStatusServer) CreateRoom(ctx context.Context, req *RoomCreateRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (*UnimplementedRoomStatusServer) GetRoomList(ctx context.Context, req *RoomListRequest) (*RoomListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomList not implemented")
}
func (*UnimplementedRoomStatusServer) GetRoomCurrentInfo(ctx context.Context, req *RoomRequest) (*Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomCurrentInfo not implemented")
}
func (*UnimplementedRoomStatusServer) DeleteRoom(ctx context.Context, req *RoomRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoom not implemented")
}

func RegisterRoomStatusServer(s *grpc.Server, srv RoomStatusServer) {
	s.RegisterService(&_RoomStatus_serviceDesc, srv)
}

func _RoomStatus_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomStatusServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.RoomStatus/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomStatusServer).CreateRoom(ctx, req.(*RoomCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomStatus_GetRoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomStatusServer).GetRoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.RoomStatus/GetRoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomStatusServer).GetRoomList(ctx, req.(*RoomListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomStatus_GetRoomCurrentInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomStatusServer).GetRoomCurrentInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.RoomStatus/GetRoomCurrentInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomStatusServer).GetRoomCurrentInfo(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomStatus_DeleteRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomStatusServer).DeleteRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.RoomStatus/DeleteRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomStatusServer).DeleteRoom(ctx, req.(*RoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RoomStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomStatus.RoomStatus",
	HandlerType: (*RoomStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomStatus_CreateRoom_Handler,
		},
		{
			MethodName: "GetRoomList",
			Handler:    _RoomStatus_GetRoomList_Handler,
		},
		{
			MethodName: "GetRoomCurrentInfo",
			Handler:    _RoomStatus_GetRoomCurrentInfo_Handler,
		},
		{
			MethodName: "DeleteRoom",
			Handler:    _RoomStatus_DeleteRoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "GameCtl.proto",
}

// RoomStreamClient is the client API for RoomStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoomStreamClient interface {
	GetRoomStream(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (RoomStream_GetRoomStreamClient, error)
	UpdateRoomStream(ctx context.Context, opts ...grpc.CallOption) (RoomStream_UpdateRoomStreamClient, error)
}

type roomStreamClient struct {
	cc *grpc.ClientConn
}

func NewRoomStreamClient(cc *grpc.ClientConn) RoomStreamClient {
	return &roomStreamClient{cc}
}

func (c *roomStreamClient) GetRoomStream(ctx context.Context, in *RoomRequest, opts ...grpc.CallOption) (RoomStream_GetRoomStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoomStream_serviceDesc.Streams[0], "/RoomStatus.RoomStream/GetRoomStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomStreamGetRoomStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoomStream_GetRoomStreamClient interface {
	Recv() (*CellStatus, error)
	grpc.ClientStream
}

type roomStreamGetRoomStreamClient struct {
	grpc.ClientStream
}

func (x *roomStreamGetRoomStreamClient) Recv() (*CellStatus, error) {
	m := new(CellStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *roomStreamClient) UpdateRoomStream(ctx context.Context, opts ...grpc.CallOption) (RoomStream_UpdateRoomStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RoomStream_serviceDesc.Streams[1], "/RoomStatus.RoomStream/UpdateRoomStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomStreamUpdateRoomStreamClient{stream}
	return x, nil
}

type RoomStream_UpdateRoomStreamClient interface {
	Send(*CellStatus) error
	Recv() (*CellStatus, error)
	grpc.ClientStream
}

type roomStreamUpdateRoomStreamClient struct {
	grpc.ClientStream
}

func (x *roomStreamUpdateRoomStreamClient) Send(m *CellStatus) error {
	return x.ClientStream.SendMsg(m)
}

func (x *roomStreamUpdateRoomStreamClient) Recv() (*CellStatus, error) {
	m := new(CellStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoomStreamServer is the server API for RoomStream service.
type RoomStreamServer interface {
	GetRoomStream(*RoomRequest, RoomStream_GetRoomStreamServer) error
	UpdateRoomStream(RoomStream_UpdateRoomStreamServer) error
}

// UnimplementedRoomStreamServer can be embedded to have forward compatible implementations.
type UnimplementedRoomStreamServer struct {
}

func (*UnimplementedRoomStreamServer) GetRoomStream(req *RoomRequest, srv RoomStream_GetRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRoomStream not implemented")
}
func (*UnimplementedRoomStreamServer) UpdateRoomStream(srv RoomStream_UpdateRoomStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateRoomStream not implemented")
}

func RegisterRoomStreamServer(s *grpc.Server, srv RoomStreamServer) {
	s.RegisterService(&_RoomStream_serviceDesc, srv)
}

func _RoomStream_GetRoomStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoomStreamServer).GetRoomStream(m, &roomStreamGetRoomStreamServer{stream})
}

type RoomStream_GetRoomStreamServer interface {
	Send(*CellStatus) error
	grpc.ServerStream
}

type roomStreamGetRoomStreamServer struct {
	grpc.ServerStream
}

func (x *roomStreamGetRoomStreamServer) Send(m *CellStatus) error {
	return x.ServerStream.SendMsg(m)
}

func _RoomStream_UpdateRoomStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoomStreamServer).UpdateRoomStream(&roomStreamUpdateRoomStreamServer{stream})
}

type RoomStream_UpdateRoomStreamServer interface {
	Send(*CellStatus) error
	Recv() (*CellStatus, error)
	grpc.ServerStream
}

type roomStreamUpdateRoomStreamServer struct {
	grpc.ServerStream
}

func (x *roomStreamUpdateRoomStreamServer) Send(m *CellStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *roomStreamUpdateRoomStreamServer) Recv() (*CellStatus, error) {
	m := new(CellStatus)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RoomStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomStatus.RoomStream",
	HandlerType: (*RoomStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetRoomStream",
			Handler:       _RoomStream_GetRoomStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateRoomStream",
			Handler:       _RoomStream_UpdateRoomStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "GameCtl.proto",
}

// CreditsAuthClient is the client API for CreditsAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreditsAuthClient interface {
	CreateCred(ctx context.Context, in *CreateCredReq, opts ...grpc.CallOption) (CreditsAuth_CreateCredClient, error)
}

type creditsAuthClient struct {
	cc *grpc.ClientConn
}

func NewCreditsAuthClient(cc *grpc.ClientConn) CreditsAuthClient {
	return &creditsAuthClient{cc}
}

func (c *creditsAuthClient) CreateCred(ctx context.Context, in *CreateCredReq, opts ...grpc.CallOption) (CreditsAuth_CreateCredClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CreditsAuth_serviceDesc.Streams[0], "/RoomStatus.CreditsAuth/CreateCred", opts...)
	if err != nil {
		return nil, err
	}
	x := &creditsAuthCreateCredClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CreditsAuth_CreateCredClient interface {
	Recv() (*Cred, error)
	grpc.ClientStream
}

type creditsAuthCreateCredClient struct {
	grpc.ClientStream
}

func (x *creditsAuthCreateCredClient) Recv() (*Cred, error) {
	m := new(Cred)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CreditsAuthServer is the server API for CreditsAuth service.
type CreditsAuthServer interface {
	CreateCred(*CreateCredReq, CreditsAuth_CreateCredServer) error
}

// UnimplementedCreditsAuthServer can be embedded to have forward compatible implementations.
type UnimplementedCreditsAuthServer struct {
}

func (*UnimplementedCreditsAuthServer) CreateCred(req *CreateCredReq, srv CreditsAuth_CreateCredServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCred not implemented")
}

func RegisterCreditsAuthServer(s *grpc.Server, srv CreditsAuthServer) {
	s.RegisterService(&_CreditsAuth_serviceDesc, srv)
}

func _CreditsAuth_CreateCred_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateCredReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CreditsAuthServer).CreateCred(m, &creditsAuthCreateCredServer{stream})
}

type CreditsAuth_CreateCredServer interface {
	Send(*Cred) error
	grpc.ServerStream
}

type creditsAuthCreateCredServer struct {
	grpc.ServerStream
}

func (x *creditsAuthCreateCredServer) Send(m *Cred) error {
	return x.ServerStream.SendMsg(m)
}

var _CreditsAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomStatus.CreditsAuth",
	HandlerType: (*CreditsAuthServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCred",
			Handler:       _CreditsAuth_CreateCred_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "GameCtl.proto",
}
