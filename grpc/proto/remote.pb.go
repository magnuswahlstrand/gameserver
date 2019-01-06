// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/remote.proto

// Web exposes a backend server over gRPC.

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ActionRequest_Action int32

const (
	ActionRequest_Move ActionRequest_Action = 0
)

var ActionRequest_Action_name = map[int32]string{
	0: "Move",
}

var ActionRequest_Action_value = map[string]int32{
	"Move": 0,
}

func (x ActionRequest_Action) String() string {
	return proto.EnumName(ActionRequest_Action_name, int32(x))
}

func (ActionRequest_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{2, 0}
}

type PlayerID struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerID) Reset()         { *m = PlayerID{} }
func (m *PlayerID) String() string { return proto.CompactTextString(m) }
func (*PlayerID) ProtoMessage()    {}
func (*PlayerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{0}
}

func (m *PlayerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerID.Unmarshal(m, b)
}
func (m *PlayerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerID.Marshal(b, m, deterministic)
}
func (m *PlayerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerID.Merge(m, src)
}
func (m *PlayerID) XXX_Size() int {
	return xxx_messageInfo_PlayerID.Size(m)
}
func (m *PlayerID) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerID.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerID proto.InternalMessageInfo

func (m *PlayerID) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

type Entity struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	X                    int32    `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	Y                    int32    `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	Theta                int32    `protobuf:"varint,4,opt,name=Theta,proto3" json:"Theta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{1}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Entity) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Entity) GetTheta() int32 {
	if m != nil {
		return m.Theta
	}
	return 0
}

type ActionRequest struct {
	Entity               *Entity              `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Action               ActionRequest_Action `protobuf:"varint,2,opt,name=action,proto3,enum=web.ActionRequest_Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ActionRequest) Reset()         { *m = ActionRequest{} }
func (m *ActionRequest) String() string { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()    {}
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{2}
}

func (m *ActionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionRequest.Unmarshal(m, b)
}
func (m *ActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionRequest.Marshal(b, m, deterministic)
}
func (m *ActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionRequest.Merge(m, src)
}
func (m *ActionRequest) XXX_Size() int {
	return xxx_messageInfo_ActionRequest.Size(m)
}
func (m *ActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActionRequest proto.InternalMessageInfo

func (m *ActionRequest) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ActionRequest) GetAction() ActionRequest_Action {
	if m != nil {
		return m.Action
	}
	return ActionRequest_Move
}

type ActionResponse struct {
	Entity               *Entity  `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionResponse) Reset()         { *m = ActionResponse{} }
func (m *ActionResponse) String() string { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()    {}
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{3}
}

func (m *ActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResponse.Unmarshal(m, b)
}
func (m *ActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResponse.Marshal(b, m, deterministic)
}
func (m *ActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResponse.Merge(m, src)
}
func (m *ActionResponse) XXX_Size() int {
	return xxx_messageInfo_ActionResponse.Size(m)
}
func (m *ActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResponse proto.InternalMessageInfo

func (m *ActionResponse) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type WorldResponse struct {
	Tiles                []byte   `protobuf:"bytes,1,opt,name=tiles,proto3" json:"tiles,omitempty"`
	Width                int32    `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorldResponse) Reset()         { *m = WorldResponse{} }
func (m *WorldResponse) String() string { return proto.CompactTextString(m) }
func (*WorldResponse) ProtoMessage()    {}
func (*WorldResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{5}
}

func (m *WorldResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorldResponse.Unmarshal(m, b)
}
func (m *WorldResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorldResponse.Marshal(b, m, deterministic)
}
func (m *WorldResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorldResponse.Merge(m, src)
}
func (m *WorldResponse) XXX_Size() int {
	return xxx_messageInfo_WorldResponse.Size(m)
}
func (m *WorldResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorldResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorldResponse proto.InternalMessageInfo

func (m *WorldResponse) GetTiles() []byte {
	if m != nil {
		return m.Tiles
	}
	return nil
}

func (m *WorldResponse) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *WorldResponse) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type EntityResponse struct {
	Payload              []byte   `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityResponse) Reset()         { *m = EntityResponse{} }
func (m *EntityResponse) String() string { return proto.CompactTextString(m) }
func (*EntityResponse) ProtoMessage()    {}
func (*EntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b59cb7d8fa836ac0, []int{6}
}

func (m *EntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResponse.Unmarshal(m, b)
}
func (m *EntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResponse.Marshal(b, m, deterministic)
}
func (m *EntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResponse.Merge(m, src)
}
func (m *EntityResponse) XXX_Size() int {
	return xxx_messageInfo_EntityResponse.Size(m)
}
func (m *EntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResponse proto.InternalMessageInfo

func (m *EntityResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("web.ActionRequest_Action", ActionRequest_Action_name, ActionRequest_Action_value)
	proto.RegisterType((*PlayerID)(nil), "web.PlayerID")
	proto.RegisterType((*Entity)(nil), "web.Entity")
	proto.RegisterType((*ActionRequest)(nil), "web.ActionRequest")
	proto.RegisterType((*ActionResponse)(nil), "web.ActionResponse")
	proto.RegisterType((*Empty)(nil), "web.Empty")
	proto.RegisterType((*WorldResponse)(nil), "web.WorldResponse")
	proto.RegisterType((*EntityResponse)(nil), "web.EntityResponse")
}

func init() { proto.RegisterFile("proto/remote.proto", fileDescriptor_b59cb7d8fa836ac0) }

var fileDescriptor_b59cb7d8fa836ac0 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xd3, 0xc6, 0x69, 0xa7, 0x49, 0x84, 0xa6, 0x15, 0x32, 0x39, 0x55, 0xcb, 0x25, 0xaa,
	0x90, 0x5d, 0x52, 0x71, 0x81, 0x13, 0x55, 0x7a, 0x08, 0x12, 0xa8, 0x72, 0x91, 0x68, 0xb9, 0xad,
	0xed, 0xc1, 0xb6, 0x6a, 0x7b, 0xcd, 0x7a, 0x5b, 0xcb, 0xe2, 0x3f, 0xf9, 0x1e, 0x94, 0xdd, 0x75,
	0x55, 0xc3, 0x85, 0xdb, 0xbe, 0xd9, 0x7d, 0xef, 0xcd, 0xbc, 0x59, 0xc0, 0x5a, 0x0a, 0x25, 0x02,
	0x49, 0xa5, 0x50, 0xe4, 0x6b, 0x80, 0x7b, 0x2d, 0x45, 0x6c, 0x09, 0x07, 0xd7, 0x05, 0xef, 0x48,
	0x6e, 0x37, 0xb8, 0x80, 0xf1, 0x76, 0xe3, 0x39, 0xa7, 0xce, 0xea, 0x30, 0x1c, 0x6f, 0x37, 0xec,
	0x13, 0xb8, 0x57, 0x95, 0xca, 0x55, 0xf7, 0xf7, 0x0d, 0xce, 0xc0, 0xb9, 0xf5, 0xc6, 0xa7, 0xce,
	0x6a, 0x12, 0x3a, 0xb7, 0x3b, 0x74, 0xe7, 0xed, 0x19, 0x74, 0x87, 0x27, 0x30, 0xf9, 0x9a, 0x91,
	0xe2, 0xde, 0xbe, 0xae, 0x18, 0xc0, 0x7e, 0xc1, 0xfc, 0x63, 0xac, 0x72, 0x51, 0x85, 0xf4, 0xf3,
	0x81, 0x1a, 0x85, 0xaf, 0xc1, 0x25, 0x2d, 0xae, 0x65, 0x8f, 0xd6, 0x47, 0x7e, 0x4b, 0x91, 0x6f,
	0xfc, 0x42, 0x7b, 0x85, 0x6f, 0xc1, 0xe5, 0x9a, 0xa5, 0xcd, 0x16, 0xeb, 0x57, 0xfa, 0xd1, 0x40,
	0xa8, 0x47, 0xf6, 0x21, 0x43, 0x70, 0x4d, 0x05, 0x0f, 0x60, 0xff, 0xb3, 0x78, 0xa4, 0x17, 0x23,
	0xf6, 0x0e, 0x16, 0x3d, 0xa7, 0xa9, 0x45, 0xd5, 0xd0, 0x7f, 0xb9, 0xb3, 0x29, 0x4c, 0xae, 0xca,
	0x5a, 0x75, 0xec, 0x06, 0xe6, 0xdf, 0x84, 0x2c, 0x92, 0x27, 0xfa, 0x09, 0x4c, 0x54, 0x5e, 0x50,
	0xa3, 0xd9, 0xb3, 0xd0, 0x80, 0x5d, 0xb5, 0xcd, 0x13, 0x95, 0xd9, 0x64, 0x0c, 0xc0, 0x97, 0xe0,
	0x66, 0x94, 0xa7, 0x99, 0xb2, 0x11, 0x59, 0xc4, 0xce, 0x60, 0x61, 0xfd, 0x7a, 0x55, 0x0f, 0xa6,
	0x35, 0xef, 0x0a, 0xc1, 0x13, 0xab, 0xdb, 0xc3, 0xf5, 0x6f, 0x07, 0xa6, 0x97, 0x3c, 0xbe, 0xa7,
	0x2a, 0xc1, 0x15, 0x1c, 0x7e, 0xa1, 0xd6, 0x2c, 0x0d, 0xc1, 0xf4, 0xbd, 0xeb, 0x72, 0x39, 0xd7,
	0xe7, 0x7e, 0x9b, 0x6c, 0x84, 0xef, 0x61, 0x7e, 0x4d, 0xf2, 0x87, 0x90, 0xa5, 0x4d, 0x04, 0xff,
	0x8d, 0x6f, 0x79, 0x3c, 0xa8, 0x99, 0x4e, 0xd8, 0x08, 0xcf, 0x61, 0x66, 0x47, 0x36, 0xeb, 0x7a,
	0x6e, 0x64, 0x64, 0x06, 0x89, 0xb0, 0x11, 0x5e, 0xc0, 0xcc, 0xcc, 0x73, 0xa3, 0x24, 0xf1, 0x72,
	0xc0, 0x38, 0x7e, 0x1e, 0xef, 0x13, 0xe5, 0xdc, 0xb9, 0x7c, 0xf3, 0xfd, 0x2c, 0xcd, 0x55, 0xf6,
	0x10, 0xf9, 0xb1, 0x28, 0x83, 0xfb, 0x8e, 0x48, 0xa9, 0x20, 0xe5, 0x25, 0x35, 0x24, 0x1f, 0x49,
	0x06, 0xa9, 0xac, 0xe3, 0x40, 0x7f, 0xd7, 0x0f, 0x75, 0x14, 0xb9, 0xfa, 0x74, 0xf1, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x42, 0x63, 0x21, 0x96, 0xce, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackendClient is the client API for Backend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackendClient interface {
	NewPlayer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerID, error)
	PerformAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error)
	WorldRequest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorldResponse, error)
	EntityStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Backend_EntityStreamClient, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) NewPlayer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PlayerID, error) {
	out := new(PlayerID)
	err := c.cc.Invoke(ctx, "/web.Backend/NewPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) PerformAction(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*ActionResponse, error) {
	out := new(ActionResponse)
	err := c.cc.Invoke(ctx, "/web.Backend/PerformAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) WorldRequest(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*WorldResponse, error) {
	out := new(WorldResponse)
	err := c.cc.Invoke(ctx, "/web.Backend/WorldRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) EntityStream(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Backend_EntityStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Backend_serviceDesc.Streams[0], "/web.Backend/EntityStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &backendEntityStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Backend_EntityStreamClient interface {
	Recv() (*EntityResponse, error)
	grpc.ClientStream
}

type backendEntityStreamClient struct {
	grpc.ClientStream
}

func (x *backendEntityStreamClient) Recv() (*EntityResponse, error) {
	m := new(EntityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BackendServer is the server API for Backend service.
type BackendServer interface {
	NewPlayer(context.Context, *Empty) (*PlayerID, error)
	PerformAction(context.Context, *ActionRequest) (*ActionResponse, error)
	WorldRequest(context.Context, *Empty) (*WorldResponse, error)
	EntityStream(*Empty, Backend_EntityStreamServer) error
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_NewPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).NewPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/NewPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).NewPlayer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_PerformAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).PerformAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/PerformAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).PerformAction(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_WorldRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).WorldRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/web.Backend/WorldRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).WorldRequest(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_EntityStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BackendServer).EntityStream(m, &backendEntityStreamServer{stream})
}

type Backend_EntityStreamServer interface {
	Send(*EntityResponse) error
	grpc.ServerStream
}

type backendEntityStreamServer struct {
	grpc.ServerStream
}

func (x *backendEntityStreamServer) Send(m *EntityResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "web.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPlayer",
			Handler:    _Backend_NewPlayer_Handler,
		},
		{
			MethodName: "PerformAction",
			Handler:    _Backend_PerformAction_Handler,
		},
		{
			MethodName: "WorldRequest",
			Handler:    _Backend_WorldRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EntityStream",
			Handler:       _Backend_EntityStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/remote.proto",
}
