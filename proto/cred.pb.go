// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cred.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorMsg struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Desp                 string   `protobuf:"bytes,2,opt,name=desp,proto3" json:"desp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ErrorMsg) Reset()         { *m = ErrorMsg{} }
func (m *ErrorMsg) String() string { return proto.CompactTextString(m) }
func (*ErrorMsg) ProtoMessage()    {}
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19cefea999b106e, []int{0}
}

func (m *ErrorMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorMsg.Unmarshal(m, b)
}
func (m *ErrorMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorMsg.Marshal(b, m, deterministic)
}
func (m *ErrorMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorMsg.Merge(m, src)
}
func (m *ErrorMsg) XXX_Size() int {
	return xxx_messageInfo_ErrorMsg.Size(m)
}
func (m *ErrorMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorMsg.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorMsg proto.InternalMessageInfo

func (m *ErrorMsg) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ErrorMsg) GetDesp() string {
	if m != nil {
		return m.Desp
	}
	return ""
}

type CredReq struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=Ip,proto3" json:"Ip,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CredReq) Reset()         { *m = CredReq{} }
func (m *CredReq) String() string { return proto.CompactTextString(m) }
func (*CredReq) ProtoMessage()    {}
func (*CredReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19cefea999b106e, []int{1}
}

func (m *CredReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredReq.Unmarshal(m, b)
}
func (m *CredReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredReq.Marshal(b, m, deterministic)
}
func (m *CredReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredReq.Merge(m, src)
}
func (m *CredReq) XXX_Size() int {
	return xxx_messageInfo_CredReq.Size(m)
}
func (m *CredReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CredReq.DiscardUnknown(m)
}

var xxx_messageInfo_CredReq proto.InternalMessageInfo

func (m *CredReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CredReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CredReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateCredResp struct {
	Code                 int32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	File                 []byte    `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
	ErrorMsg             *ErrorMsg `protobuf:"bytes,3,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateCredResp) Reset()         { *m = CreateCredResp{} }
func (m *CreateCredResp) String() string { return proto.CompactTextString(m) }
func (*CreateCredResp) ProtoMessage()    {}
func (*CreateCredResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19cefea999b106e, []int{2}
}

func (m *CreateCredResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCredResp.Unmarshal(m, b)
}
func (m *CreateCredResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCredResp.Marshal(b, m, deterministic)
}
func (m *CreateCredResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCredResp.Merge(m, src)
}
func (m *CreateCredResp) XXX_Size() int {
	return xxx_messageInfo_CreateCredResp.Size(m)
}
func (m *CreateCredResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCredResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCredResp proto.InternalMessageInfo

func (m *CreateCredResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CreateCredResp) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *CreateCredResp) GetErrorMsg() *ErrorMsg {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

type CheckCredResp struct {
	ResponseCode         int32     `protobuf:"varint,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	ErrorMsg             *ErrorMsg `protobuf:"bytes,2,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CheckCredResp) Reset()         { *m = CheckCredResp{} }
func (m *CheckCredResp) String() string { return proto.CompactTextString(m) }
func (*CheckCredResp) ProtoMessage()    {}
func (*CheckCredResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_c19cefea999b106e, []int{3}
}

func (m *CheckCredResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckCredResp.Unmarshal(m, b)
}
func (m *CheckCredResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckCredResp.Marshal(b, m, deterministic)
}
func (m *CheckCredResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCredResp.Merge(m, src)
}
func (m *CheckCredResp) XXX_Size() int {
	return xxx_messageInfo_CheckCredResp.Size(m)
}
func (m *CheckCredResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCredResp.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCredResp proto.InternalMessageInfo

func (m *CheckCredResp) GetResponseCode() int32 {
	if m != nil {
		return m.ResponseCode
	}
	return 0
}

func (m *CheckCredResp) GetErrorMsg() *ErrorMsg {
	if m != nil {
		return m.ErrorMsg
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorMsg)(nil), "RoomStatus.ErrorMsg")
	proto.RegisterType((*CredReq)(nil), "RoomStatus.CredReq")
	proto.RegisterType((*CreateCredResp)(nil), "RoomStatus.CreateCredResp")
	proto.RegisterType((*CheckCredResp)(nil), "RoomStatus.CheckCredResp")
}

func init() { proto.RegisterFile("cred.proto", fileDescriptor_c19cefea999b106e) }

var fileDescriptor_c19cefea999b106e = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xd5, 0xba, 0x84, 0x26, 0xdb, 0x36, 0xa0, 0x6d, 0x0f, 0xa9, 0x41, 0x51, 0x64, 0x2e, 0x55,
	0x44, 0xbc, 0x69, 0x7a, 0x40, 0xe4, 0x56, 0x0c, 0x82, 0x1c, 0x2a, 0x05, 0xd3, 0x72, 0x80, 0x43,
	0xe5, 0xd8, 0xd3, 0x8d, 0xd5, 0xd8, 0xbb, 0xec, 0x6e, 0x52, 0xe5, 0x8a, 0xc4, 0x0f, 0xc0, 0x89,
	0x5f, 0xe0, 0x2b, 0x38, 0xf6, 0x88, 0xc4, 0x0f, 0xa0, 0x94, 0x0f, 0x41, 0xde, 0x38, 0x89, 0x8b,
	0x28, 0x82, 0x93, 0xe7, 0xcd, 0x7b, 0xf3, 0x9e, 0x77, 0x34, 0x18, 0x87, 0x12, 0x22, 0x57, 0x48,
	0xae, 0x39, 0xc1, 0x3e, 0xe7, 0xc9, 0x2b, 0x1d, 0xe8, 0xb1, 0xb2, 0xef, 0x31, 0xce, 0xd9, 0x08,
	0xa8, 0x61, 0x06, 0xe3, 0x33, 0x0a, 0x89, 0xd0, 0xd3, 0xb9, 0xd0, 0xbe, 0x9f, 0x93, 0x81, 0x88,
	0x69, 0x90, 0xa6, 0x5c, 0x07, 0x3a, 0xe6, 0xa9, 0xca, 0xd9, 0x87, 0xe6, 0x13, 0xb6, 0x18, 0xa4,
	0x2d, 0x75, 0x11, 0x30, 0x06, 0x92, 0x72, 0x61, 0x14, 0x7f, 0x50, 0xb7, 0x58, 0xac, 0x87, 0xe3,
	0x81, 0x1b, 0xf2, 0x84, 0x32, 0xce, 0xf8, 0x2a, 0x31, 0x43, 0x06, 0x98, 0x6a, 0x2e, 0x77, 0x3a,
	0xb8, 0xfc, 0x4c, 0x4a, 0x2e, 0x8f, 0x14, 0x23, 0x04, 0xdf, 0x0a, 0x79, 0x04, 0x35, 0xd4, 0x40,
	0x7b, 0x25, 0xdf, 0xd4, 0x59, 0x2f, 0x02, 0x25, 0x6a, 0x56, 0x03, 0xed, 0x55, 0x7c, 0x53, 0x3b,
	0x2f, 0xf1, 0xba, 0x27, 0x21, 0xf2, 0xe1, 0x1d, 0xa9, 0x62, 0xab, 0x27, 0xcc, 0x40, 0xc5, 0xb7,
	0x7a, 0x82, 0xd8, 0xb8, 0x7c, 0xa2, 0x40, 0xa6, 0x41, 0x02, 0xf9, 0xc8, 0x12, 0x67, 0x5c, 0x3f,
	0x50, 0xea, 0x82, 0xcb, 0xa8, 0xb6, 0x36, 0xe7, 0x16, 0xd8, 0x39, 0xc7, 0x55, 0x4f, 0x42, 0xa0,
	0x61, 0x6e, 0xac, 0xc4, 0x4d, 0x3f, 0x73, 0x16, 0x8f, 0xe6, 0xce, 0x9b, 0xbe, 0xa9, 0xc9, 0x3e,
	0xae, 0x40, 0xf6, 0x80, 0xd3, 0x44, 0x31, 0x63, 0xbb, 0xd1, 0xd9, 0x71, 0x57, 0x8b, 0x77, 0x17,
	0xaf, 0xf3, 0xcb, 0x90, 0x57, 0x0e, 0xc3, 0x5b, 0xde, 0x10, 0xc2, 0xf3, 0x65, 0xd6, 0x03, 0xbc,
	0x25, 0x41, 0x09, 0x9e, 0x2a, 0x38, 0x2d, 0x84, 0x6e, 0x2e, 0x9a, 0x5e, 0x16, 0x7e, 0x2d, 0xc8,
	0xfa, 0x97, 0xa0, 0xce, 0x67, 0x0b, 0x6f, 0x64, 0x21, 0xb1, 0x56, 0x87, 0x63, 0x3d, 0x24, 0xaf,
	0x71, 0x65, 0x19, 0x4c, 0xb6, 0x8b, 0xc3, 0xf9, 0x3e, 0xed, 0xdd, 0x6b, 0xcd, 0xe2, 0x4f, 0x3a,
	0xbb, 0xef, 0xbf, 0xff, 0xfc, 0x64, 0x6d, 0x3b, 0x55, 0x3a, 0xd9, 0xa7, 0xd9, 0x95, 0xd1, 0x30,
	0xe3, 0xbb, 0xa8, 0x49, 0x4e, 0xf0, 0xfa, 0x73, 0xd0, 0x37, 0xbb, 0xda, 0xbf, 0x35, 0x0b, 0x7b,
	0xfe, 0x9b, 0xed, 0x5b, 0x8c, 0x57, 0xe2, 0xff, 0x77, 0xb6, 0x8d, 0xf3, 0x8e, 0x73, 0x67, 0xe5,
	0x6c, 0x04, 0x5d, 0xd4, 0x6c, 0xa3, 0x27, 0x1f, 0xd0, 0xc7, 0xc3, 0x63, 0x52, 0xea, 0xac, 0xb5,
	0xdd, 0x4e, 0x13, 0x59, 0xf2, 0x05, 0x76, 0x98, 0xdf, 0xf7, 0x1a, 0x0a, 0xe4, 0x04, 0x64, 0xa3,
	0x97, 0x88, 0x11, 0x68, 0x48, 0x20, 0xd5, 0x0d, 0x09, 0x82, 0xab, 0x58, 0x73, 0x39, 0x25, 0xce,
	0x50, 0x6b, 0xa1, 0xba, 0x94, 0x16, 0xae, 0x5c, 0xe9, 0x49, 0x38, 0x4c, 0x1e, 0x3f, 0x6a, 0x1f,
	0xd0, 0xde, 0x53, 0xef, 0xf8, 0x72, 0x56, 0x47, 0xdf, 0x66, 0x75, 0xf4, 0x63, 0x56, 0x47, 0x5f,
	0xaf, 0xea, 0xe8, 0xf2, 0xaa, 0x8e, 0xde, 0x94, 0xcc, 0xb5, 0x7f, 0xb1, 0xee, 0xf6, 0x47, 0xc1,
	0xd4, 0x1b, 0xc5, 0x6e, 0x3f, 0xc3, 0x47, 0x3c, 0x1a, 0xdc, 0x36, 0xcc, 0xc1, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0xe5, 0x6e, 0xe9, 0xb9, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CreditsAuthClient is the client API for CreditsAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CreditsAuthClient interface {
	CheckCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (*CheckCredResp, error)
	GetCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (*CreateCredResp, error)
	CreateCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (CreditsAuth_CreateCredClient, error)
}

type creditsAuthClient struct {
	cc *grpc.ClientConn
}

func NewCreditsAuthClient(cc *grpc.ClientConn) CreditsAuthClient {
	return &creditsAuthClient{cc}
}

func (c *creditsAuthClient) CheckCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (*CheckCredResp, error) {
	out := new(CheckCredResp)
	err := c.cc.Invoke(ctx, "/RoomStatus.CreditsAuth/CheckCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditsAuthClient) GetCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (*CreateCredResp, error) {
	out := new(CreateCredResp)
	err := c.cc.Invoke(ctx, "/RoomStatus.CreditsAuth/GetCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditsAuthClient) CreateCred(ctx context.Context, in *CredReq, opts ...grpc.CallOption) (CreditsAuth_CreateCredClient, error) {
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
	Recv() (*CreateCredResp, error)
	grpc.ClientStream
}

type creditsAuthCreateCredClient struct {
	grpc.ClientStream
}

func (x *creditsAuthCreateCredClient) Recv() (*CreateCredResp, error) {
	m := new(CreateCredResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CreditsAuthServer is the server API for CreditsAuth service.
type CreditsAuthServer interface {
	CheckCred(context.Context, *CredReq) (*CheckCredResp, error)
	GetCred(context.Context, *CredReq) (*CreateCredResp, error)
	CreateCred(*CredReq, CreditsAuth_CreateCredServer) error
}

// UnimplementedCreditsAuthServer can be embedded to have forward compatible implementations.
type UnimplementedCreditsAuthServer struct {
}

func (*UnimplementedCreditsAuthServer) CheckCred(ctx context.Context, req *CredReq) (*CheckCredResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckCred not implemented")
}
func (*UnimplementedCreditsAuthServer) GetCred(ctx context.Context, req *CredReq) (*CreateCredResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCred not implemented")
}
func (*UnimplementedCreditsAuthServer) CreateCred(req *CredReq, srv CreditsAuth_CreateCredServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateCred not implemented")
}

func RegisterCreditsAuthServer(s *grpc.Server, srv CreditsAuthServer) {
	s.RegisterService(&_CreditsAuth_serviceDesc, srv)
}

func _CreditsAuth_CheckCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditsAuthServer).CheckCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.CreditsAuth/CheckCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditsAuthServer).CheckCred(ctx, req.(*CredReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditsAuth_GetCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditsAuthServer).GetCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoomStatus.CreditsAuth/GetCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditsAuthServer).GetCred(ctx, req.(*CredReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditsAuth_CreateCred_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CredReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CreditsAuthServer).CreateCred(m, &creditsAuthCreateCredServer{stream})
}

type CreditsAuth_CreateCredServer interface {
	Send(*CreateCredResp) error
	grpc.ServerStream
}

type creditsAuthCreateCredServer struct {
	grpc.ServerStream
}

func (x *creditsAuthCreateCredServer) Send(m *CreateCredResp) error {
	return x.ServerStream.SendMsg(m)
}

var _CreditsAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RoomStatus.CreditsAuth",
	HandlerType: (*CreditsAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckCred",
			Handler:    _CreditsAuth_CheckCred_Handler,
		},
		{
			MethodName: "GetCred",
			Handler:    _CreditsAuth_GetCred_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateCred",
			Handler:       _CreditsAuth_CreateCred_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cred.proto",
}
