// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Ping struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{0}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (dst *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(dst, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CommandsRequest struct {
	Commands             []string `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommandsRequest) Reset()         { *m = CommandsRequest{} }
func (m *CommandsRequest) String() string { return proto.CompactTextString(m) }
func (*CommandsRequest) ProtoMessage()    {}
func (*CommandsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{1}
}
func (m *CommandsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandsRequest.Unmarshal(m, b)
}
func (m *CommandsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandsRequest.Marshal(b, m, deterministic)
}
func (dst *CommandsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandsRequest.Merge(dst, src)
}
func (m *CommandsRequest) XXX_Size() int {
	return xxx_messageInfo_CommandsRequest.Size(m)
}
func (m *CommandsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CommandsRequest proto.InternalMessageInfo

func (m *CommandsRequest) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

type Cue struct {
	ExpectedDurationMS   int32    `protobuf:"varint,1,opt,name=ExpectedDurationMS,proto3" json:"ExpectedDurationMS,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cue) Reset()         { *m = Cue{} }
func (m *Cue) String() string { return proto.CompactTextString(m) }
func (*Cue) ProtoMessage()    {}
func (*Cue) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{2}
}
func (m *Cue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cue.Unmarshal(m, b)
}
func (m *Cue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cue.Marshal(b, m, deterministic)
}
func (dst *Cue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cue.Merge(dst, src)
}
func (m *Cue) XXX_Size() int {
	return xxx_messageInfo_Cue.Size(m)
}
func (m *Cue) XXX_DiscardUnknown() {
	xxx_messageInfo_Cue.DiscardUnknown(m)
}

var xxx_messageInfo_Cue proto.InternalMessageInfo

func (m *Cue) GetExpectedDurationMS() int32 {
	if m != nil {
		return m.ExpectedDurationMS
	}
	return 0
}

type CuesResponse struct {
	Cues                 []*Cue   `protobuf:"bytes,1,rep,name=cues,proto3" json:"cues,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CuesResponse) Reset()         { *m = CuesResponse{} }
func (m *CuesResponse) String() string { return proto.CompactTextString(m) }
func (*CuesResponse) ProtoMessage()    {}
func (*CuesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{3}
}
func (m *CuesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CuesResponse.Unmarshal(m, b)
}
func (m *CuesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CuesResponse.Marshal(b, m, deterministic)
}
func (dst *CuesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CuesResponse.Merge(dst, src)
}
func (m *CuesResponse) XXX_Size() int {
	return xxx_messageInfo_CuesResponse.Size(m)
}
func (m *CuesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CuesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CuesResponse proto.InternalMessageInfo

func (m *CuesResponse) GetCues() []*Cue {
	if m != nil {
		return m.Cues
	}
	return nil
}

func (m *CuesResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type MarshalledJSON struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarshalledJSON) Reset()         { *m = MarshalledJSON{} }
func (m *MarshalledJSON) String() string { return proto.CompactTextString(m) }
func (*MarshalledJSON) ProtoMessage()    {}
func (*MarshalledJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{4}
}
func (m *MarshalledJSON) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarshalledJSON.Unmarshal(m, b)
}
func (m *MarshalledJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarshalledJSON.Marshal(b, m, deterministic)
}
func (dst *MarshalledJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarshalledJSON.Merge(dst, src)
}
func (m *MarshalledJSON) XXX_Size() int {
	return xxx_messageInfo_MarshalledJSON.Size(m)
}
func (m *MarshalledJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_MarshalledJSON.DiscardUnknown(m)
}

var xxx_messageInfo_MarshalledJSON proto.InternalMessageInfo

func (m *MarshalledJSON) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *MarshalledJSON) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Lights struct {
	Lights               []*Light `protobuf:"bytes,1,rep,name=lights,proto3" json:"lights,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lights) Reset()         { *m = Lights{} }
func (m *Lights) String() string { return proto.CompactTextString(m) }
func (*Lights) ProtoMessage()    {}
func (*Lights) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{5}
}
func (m *Lights) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lights.Unmarshal(m, b)
}
func (m *Lights) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lights.Marshal(b, m, deterministic)
}
func (dst *Lights) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lights.Merge(dst, src)
}
func (m *Lights) XXX_Size() int {
	return xxx_messageInfo_Lights.Size(m)
}
func (m *Lights) XXX_DiscardUnknown() {
	xxx_messageInfo_Lights.DiscardUnknown(m)
}

var xxx_messageInfo_Lights proto.InternalMessageInfo

func (m *Lights) GetLights() []*Light {
	if m != nil {
		return m.Lights
	}
	return nil
}

type Light struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CurrentColor         *RGB     `protobuf:"bytes,3,opt,name=currentColor,proto3" json:"currentColor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Light) Reset()         { *m = Light{} }
func (m *Light) String() string { return proto.CompactTextString(m) }
func (*Light) ProtoMessage()    {}
func (*Light) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{6}
}
func (m *Light) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Light.Unmarshal(m, b)
}
func (m *Light) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Light.Marshal(b, m, deterministic)
}
func (dst *Light) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Light.Merge(dst, src)
}
func (m *Light) XXX_Size() int {
	return xxx_messageInfo_Light.Size(m)
}
func (m *Light) XXX_DiscardUnknown() {
	xxx_messageInfo_Light.DiscardUnknown(m)
}

var xxx_messageInfo_Light proto.InternalMessageInfo

func (m *Light) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Light) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Light) GetCurrentColor() *RGB {
	if m != nil {
		return m.CurrentColor
	}
	return nil
}

type RGB struct {
	R                    int32    `protobuf:"varint,1,opt,name=R,proto3" json:"R,omitempty"`
	G                    int32    `protobuf:"varint,2,opt,name=G,proto3" json:"G,omitempty"`
	B                    int32    `protobuf:"varint,3,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RGB) Reset()         { *m = RGB{} }
func (m *RGB) String() string { return proto.CompactTextString(m) }
func (*RGB) ProtoMessage()    {}
func (*RGB) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{7}
}
func (m *RGB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RGB.Unmarshal(m, b)
}
func (m *RGB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RGB.Marshal(b, m, deterministic)
}
func (dst *RGB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RGB.Merge(dst, src)
}
func (m *RGB) XXX_Size() int {
	return xxx_messageInfo_RGB.Size(m)
}
func (m *RGB) XXX_DiscardUnknown() {
	xxx_messageInfo_RGB.DiscardUnknown(m)
}

var xxx_messageInfo_RGB proto.InternalMessageInfo

func (m *RGB) GetR() int32 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *RGB) GetG() int32 {
	if m != nil {
		return m.G
	}
	return 0
}

func (m *RGB) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ConnectionSettings struct {
	Tick                 string   `protobuf:"bytes,1,opt,name=tick,proto3" json:"tick,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionSettings) Reset()         { *m = ConnectionSettings{} }
func (m *ConnectionSettings) String() string { return proto.CompactTextString(m) }
func (*ConnectionSettings) ProtoMessage()    {}
func (*ConnectionSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8b102077f842c8ec, []int{8}
}
func (m *ConnectionSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionSettings.Unmarshal(m, b)
}
func (m *ConnectionSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionSettings.Marshal(b, m, deterministic)
}
func (dst *ConnectionSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionSettings.Merge(dst, src)
}
func (m *ConnectionSettings) XXX_Size() int {
	return xxx_messageInfo_ConnectionSettings.Size(m)
}
func (m *ConnectionSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionSettings proto.InternalMessageInfo

func (m *ConnectionSettings) GetTick() string {
	if m != nil {
		return m.Tick
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "Ping")
	proto.RegisterType((*CommandsRequest)(nil), "CommandsRequest")
	proto.RegisterType((*Cue)(nil), "Cue")
	proto.RegisterType((*CuesResponse)(nil), "CuesResponse")
	proto.RegisterType((*MarshalledJSON)(nil), "MarshalledJSON")
	proto.RegisterType((*Lights)(nil), "Lights")
	proto.RegisterType((*Light)(nil), "Light")
	proto.RegisterType((*RGB)(nil), "RGB")
	proto.RegisterType((*ConnectionSettings)(nil), "ConnectionSettings")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIClient interface {
	GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error)
	StreamCueMaster(ctx context.Context, in *ConnectionSettings, opts ...grpc.CallOption) (API_StreamCueMasterClient, error)
	StreamGetLights(ctx context.Context, in *ConnectionSettings, opts ...grpc.CallOption) (API_StreamGetLightsClient, error)
	ProcessCommands(ctx context.Context, in *CommandsRequest, opts ...grpc.CallOption) (*CuesResponse, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) GetPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Ping, error) {
	out := new(Ping)
	err := c.cc.Invoke(ctx, "/API/GetPing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) StreamCueMaster(ctx context.Context, in *ConnectionSettings, opts ...grpc.CallOption) (API_StreamCueMasterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[0], "/API/StreamCueMaster", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIStreamCueMasterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_StreamCueMasterClient interface {
	Recv() (*MarshalledJSON, error)
	grpc.ClientStream
}

type aPIStreamCueMasterClient struct {
	grpc.ClientStream
}

func (x *aPIStreamCueMasterClient) Recv() (*MarshalledJSON, error) {
	m := new(MarshalledJSON)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) StreamGetLights(ctx context.Context, in *ConnectionSettings, opts ...grpc.CallOption) (API_StreamGetLightsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_API_serviceDesc.Streams[1], "/API/StreamGetLights", opts...)
	if err != nil {
		return nil, err
	}
	x := &aPIStreamGetLightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type API_StreamGetLightsClient interface {
	Recv() (*Lights, error)
	grpc.ClientStream
}

type aPIStreamGetLightsClient struct {
	grpc.ClientStream
}

func (x *aPIStreamGetLightsClient) Recv() (*Lights, error) {
	m := new(Lights)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aPIClient) ProcessCommands(ctx context.Context, in *CommandsRequest, opts ...grpc.CallOption) (*CuesResponse, error) {
	out := new(CuesResponse)
	err := c.cc.Invoke(ctx, "/API/ProcessCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
type APIServer interface {
	GetPing(context.Context, *Ping) (*Ping, error)
	StreamCueMaster(*ConnectionSettings, API_StreamCueMasterServer) error
	StreamGetLights(*ConnectionSettings, API_StreamGetLightsServer) error
	ProcessCommands(context.Context, *CommandsRequest) (*CuesResponse, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_GetPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).GetPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/GetPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).GetPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_StreamCueMaster_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionSettings)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).StreamCueMaster(m, &aPIStreamCueMasterServer{stream})
}

type API_StreamCueMasterServer interface {
	Send(*MarshalledJSON) error
	grpc.ServerStream
}

type aPIStreamCueMasterServer struct {
	grpc.ServerStream
}

func (x *aPIStreamCueMasterServer) Send(m *MarshalledJSON) error {
	return x.ServerStream.SendMsg(m)
}

func _API_StreamGetLights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionSettings)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).StreamGetLights(m, &aPIStreamGetLightsServer{stream})
}

type API_StreamGetLightsServer interface {
	Send(*Lights) error
	grpc.ServerStream
}

type aPIStreamGetLightsServer struct {
	grpc.ServerStream
}

func (x *aPIStreamGetLightsServer) Send(m *Lights) error {
	return x.ServerStream.SendMsg(m)
}

func _API_ProcessCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).ProcessCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/API/ProcessCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).ProcessCommands(ctx, req.(*CommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPing",
			Handler:    _API_GetPing_Handler,
		},
		{
			MethodName: "ProcessCommands",
			Handler:    _API_ProcessCommands_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCueMaster",
			Handler:       _API_StreamCueMaster_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamGetLights",
			Handler:       _API_StreamGetLights_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}

func init() { proto.RegisterFile("proto/api.proto", fileDescriptor_api_8b102077f842c8ec) }

var fileDescriptor_api_8b102077f842c8ec = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0x49, 0xd3, 0xb2, 0x43, 0x21, 0x2b, 0x73, 0x09, 0x3d, 0xa0, 0xca, 0xa7, 0x5c, 0xf0,
	0xa2, 0x5d, 0x90, 0x10, 0x9c, 0x68, 0x40, 0x11, 0x88, 0x42, 0xe5, 0x9e, 0x38, 0x9a, 0x64, 0xd4,
	0x8d, 0xda, 0xd8, 0xc1, 0x76, 0x24, 0xf8, 0x43, 0x3e, 0x0b, 0x65, 0x9a, 0xac, 0x58, 0x76, 0x4f,
	0x7e, 0xf3, 0x46, 0xf3, 0x32, 0xf3, 0xf2, 0x20, 0x69, 0x9d, 0x0d, 0xf6, 0x42, 0xb7, 0xb5, 0x24,
	0x24, 0x56, 0x30, 0xdd, 0xd6, 0x66, 0xcf, 0x53, 0x98, 0x37, 0xe8, 0xbd, 0xde, 0x63, 0xca, 0x56,
	0x2c, 0x3b, 0x53, 0x63, 0x29, 0x5e, 0x40, 0x92, 0xdb, 0xa6, 0xd1, 0xa6, 0xf2, 0x0a, 0x7f, 0x76,
	0xe8, 0x03, 0x5f, 0xc2, 0xc3, 0x72, 0xa0, 0x52, 0xb6, 0x8a, 0xb2, 0x33, 0x75, 0x53, 0x8b, 0xd7,
	0x10, 0xe5, 0x1d, 0x72, 0x09, 0xfc, 0xe3, 0xaf, 0x16, 0xcb, 0x80, 0xd5, 0x87, 0xce, 0xe9, 0x50,
	0x5b, 0xb3, 0xd9, 0x91, 0x74, 0xac, 0xee, 0xe9, 0x88, 0xb7, 0xb0, 0xc8, 0x3b, 0xf4, 0x0a, 0x7d,
	0x6b, 0x8d, 0x47, 0x9e, 0xc2, 0xb4, 0xec, 0xf0, 0x24, 0xff, 0xe8, 0x72, 0x2a, 0xf3, 0x0e, 0x15,
	0x31, 0xfc, 0x1c, 0x22, 0x74, 0x2e, 0x7d, 0x40, 0x5b, 0xf6, 0x50, 0xbc, 0x81, 0x27, 0x1b, 0xed,
	0xfc, 0xb5, 0x3e, 0x1e, 0xb1, 0xfa, 0xbc, 0xfb, 0xf6, 0x95, 0x73, 0x98, 0x1e, 0x6a, 0x53, 0x0d,
	0xa7, 0x10, 0xee, 0xb9, 0x4a, 0x07, 0x4d, 0x83, 0x0b, 0x45, 0x58, 0x64, 0x30, 0xfb, 0x52, 0xef,
	0xaf, 0x83, 0xe7, 0xcf, 0x61, 0x76, 0x24, 0x34, 0x7c, 0x71, 0x26, 0xa9, 0xa1, 0x06, 0x56, 0x7c,
	0x87, 0x98, 0x88, 0x5e, 0x26, 0xfc, 0x6e, 0x47, 0x97, 0x08, 0xf7, 0x9c, 0xd1, 0x0d, 0x0e, 0x3b,
	0x11, 0xe6, 0x19, 0x2c, 0xca, 0xce, 0x39, 0x34, 0x21, 0xb7, 0x47, 0xeb, 0xd2, 0x68, 0xc5, 0xe8,
	0x10, 0x55, 0xac, 0xd5, 0xad, 0x8e, 0xb8, 0x80, 0x48, 0x15, 0x6b, 0xbe, 0x00, 0xa6, 0x06, 0x83,
	0x98, 0xea, 0xab, 0x82, 0xf4, 0x62, 0xc5, 0x8a, 0xbe, 0x5a, 0x93, 0x42, 0xac, 0xd8, 0x5a, 0x64,
	0xc0, 0x73, 0x6b, 0x0c, 0x96, 0xbd, 0x77, 0x3b, 0x0c, 0xa1, 0x36, 0x7b, 0x4f, 0x8b, 0xd5, 0xe5,
	0xe1, 0x66, 0xb1, 0xba, 0x3c, 0x5c, 0xfe, 0x61, 0x10, 0xbd, 0xdf, 0x7e, 0xe2, 0xcf, 0x60, 0x5e,
	0x60, 0xa0, 0x1f, 0x1d, 0xcb, 0xfe, 0x59, 0x9e, 0x1e, 0x31, 0xe1, 0xef, 0x20, 0xd9, 0x05, 0x87,
	0xba, 0xc9, 0x3b, 0xdc, 0x68, 0x1f, 0xd0, 0xf1, 0xa7, 0xf2, 0xae, 0xfc, 0x32, 0x91, 0xb7, 0x3d,
	0x16, 0x93, 0x97, 0x8c, 0x5f, 0x8d, 0xc3, 0x05, 0x86, 0xc1, 0xc8, 0x7b, 0x87, 0xe7, 0x27, 0x37,
	0x3d, 0x0d, 0xbd, 0x82, 0x64, 0xeb, 0x6c, 0x89, 0xde, 0x8f, 0xb9, 0xe2, 0xe7, 0xf2, 0xbf, 0x88,
	0x2d, 0x1f, 0xcb, 0x7f, 0xe3, 0x20, 0x26, 0x3f, 0x66, 0x94, 0xd7, 0xab, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8f, 0xa9, 0xd7, 0x51, 0xc2, 0x02, 0x00, 0x00,
}
