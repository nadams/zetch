// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Server struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HostName             string   `protobuf:"bytes,2,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Port                 string   `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	Iwad                 string   `protobuf:"bytes,5,opt,name=iwad,proto3" json:"iwad,omitempty"`
	Pwads                []string `protobuf:"bytes,6,rep,name=pwads,proto3" json:"pwads,omitempty"`
	GameType             string   `protobuf:"bytes,7,opt,name=game_type,json=gameType,proto3" json:"game_type,omitempty"`
	Status               string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Server) Reset()         { *m = Server{} }
func (m *Server) String() string { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()    {}
func (*Server) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *Server) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Server.Unmarshal(m, b)
}
func (m *Server) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Server.Marshal(b, m, deterministic)
}
func (m *Server) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Server.Merge(m, src)
}
func (m *Server) XXX_Size() int {
	return xxx_messageInfo_Server.Size(m)
}
func (m *Server) XXX_DiscardUnknown() {
	xxx_messageInfo_Server.DiscardUnknown(m)
}

var xxx_messageInfo_Server proto.InternalMessageInfo

func (m *Server) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Server) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *Server) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Server) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *Server) GetIwad() string {
	if m != nil {
		return m.Iwad
	}
	return ""
}

func (m *Server) GetPwads() []string {
	if m != nil {
		return m.Pwads
	}
	return nil
}

func (m *Server) GetGameType() string {
	if m != nil {
		return m.GameType
	}
	return ""
}

func (m *Server) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Servers              []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type AttachRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachRequest) Reset()         { *m = AttachRequest{} }
func (m *AttachRequest) String() string { return proto.CompactTextString(m) }
func (*AttachRequest) ProtoMessage()    {}
func (*AttachRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *AttachRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachRequest.Unmarshal(m, b)
}
func (m *AttachRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachRequest.Marshal(b, m, deterministic)
}
func (m *AttachRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachRequest.Merge(m, src)
}
func (m *AttachRequest) XXX_Size() int {
	return xxx_messageInfo_AttachRequest.Size(m)
}
func (m *AttachRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttachRequest proto.InternalMessageInfo

func (m *AttachRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttachRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type ServerOutput struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerOutput) Reset()         { *m = ServerOutput{} }
func (m *ServerOutput) String() string { return proto.CompactTextString(m) }
func (*ServerOutput) ProtoMessage()    {}
func (*ServerOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *ServerOutput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerOutput.Unmarshal(m, b)
}
func (m *ServerOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerOutput.Marshal(b, m, deterministic)
}
func (m *ServerOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerOutput.Merge(m, src)
}
func (m *ServerOutput) XXX_Size() int {
	return xxx_messageInfo_ServerOutput.Size(m)
}
func (m *ServerOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerOutput.DiscardUnknown(m)
}

var xxx_messageInfo_ServerOutput proto.InternalMessageInfo

func (m *ServerOutput) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Server)(nil), "proto.Server")
	proto.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto.RegisterType((*AttachRequest)(nil), "proto.AttachRequest")
	proto.RegisterType((*ServerOutput)(nil), "proto.ServerOutput")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x65, 0x2d, 0x14, 0x18, 0x20, 0x31, 0x23, 0x31, 0x1b, 0xbd, 0x34, 0x7b, 0xb1, 0x27, 0xa2,
	0x18, 0x63, 0x3c, 0x9a, 0x78, 0x34, 0x9a, 0x54, 0xef, 0x64, 0xb5, 0x1b, 0xe0, 0x50, 0xba, 0x76,
	0xa6, 0x10, 0xfe, 0xa0, 0xbf, 0xcb, 0xec, 0x6e, 0x4b, 0x20, 0xf1, 0xd4, 0xf7, 0x31, 0x7d, 0xfb,
	0x66, 0x60, 0x4c, 0xa6, 0xda, 0x9a, 0x6a, 0x66, 0xab, 0x92, 0x4b, 0xec, 0xf9, 0x8f, 0xfa, 0x15,
	0x10, 0x7f, 0x78, 0x1d, 0x11, 0xba, 0x1b, 0x5d, 0x18, 0x29, 0x12, 0x91, 0x0e, 0x33, 0x8f, 0xf1,
	0x1a, 0x86, 0xab, 0x92, 0x78, 0xe1, 0x8d, 0x33, 0x6f, 0x0c, 0x9c, 0xf0, 0xe6, 0x4c, 0x09, 0x7d,
	0x9d, 0xe7, 0x95, 0x21, 0x92, 0x91, 0xb7, 0x5a, 0xea, 0xa2, 0x6c, 0x59, 0xb1, 0xec, 0x86, 0x28,
	0x87, 0x9d, 0xb6, 0xde, 0xe9, 0x5c, 0xf6, 0x82, 0xe6, 0x30, 0x4e, 0xa1, 0x67, 0x77, 0x3a, 0x27,
	0x19, 0x27, 0x51, 0x3a, 0xcc, 0x02, 0x71, 0x8f, 0x2e, 0x75, 0x61, 0x16, 0xbc, 0xb7, 0x46, 0xf6,
	0xc3, 0xa3, 0x4e, 0xf8, 0xdc, 0x5b, 0x83, 0x97, 0x10, 0x13, 0x6b, 0xae, 0x49, 0x0e, 0xbc, 0xd3,
	0x30, 0x35, 0x81, 0xd1, 0xeb, 0x9a, 0x38, 0x33, 0x3f, 0xb5, 0x21, 0x56, 0x8f, 0x30, 0x0e, 0x94,
	0x6c, 0xb9, 0x21, 0x83, 0x37, 0xd0, 0x0f, 0xeb, 0x93, 0x14, 0x49, 0x94, 0x8e, 0xe6, 0x93, 0x70,
	0x87, 0x59, 0x58, 0x3e, 0x6b, 0x5d, 0xf5, 0x00, 0x93, 0x67, 0x66, 0xfd, 0xbd, 0x6a, 0x92, 0xfe,
	0x3d, 0xcb, 0x39, 0x44, 0x05, 0x2d, 0x9b, 0x83, 0x38, 0xa8, 0x12, 0x18, 0x87, 0xa4, 0xf7, 0x9a,
	0x6d, 0xcd, 0xed, 0x84, 0x38, 0x4c, 0xcc, 0xb7, 0x10, 0xbf, 0x68, 0x53, 0x94, 0x1b, 0xbc, 0x83,
	0xae, 0xeb, 0x86, 0xd8, 0x54, 0x38, 0xea, 0x7d, 0x75, 0x71, 0xa2, 0x85, 0xf2, 0xaa, 0x83, 0x4f,
	0x10, 0x87, 0x56, 0x38, 0x6d, 0x06, 0x4e, 0x4a, 0x1e, 0x7e, 0x3b, 0xee, 0xa0, 0x3a, 0xa9, 0xb8,
	0x15, 0x5f, 0xb1, 0x77, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb6, 0xa0, 0xe9, 0xff,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DaemonClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Attach(ctx context.Context, opts ...grpc.CallOption) (Daemon_AttachClient, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/proto.Daemon/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Daemon_AttachClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Daemon_serviceDesc.Streams[0], "/proto.Daemon/Attach", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonAttachClient{stream}
	return x, nil
}

type Daemon_AttachClient interface {
	Send(*AttachRequest) error
	Recv() (*ServerOutput, error)
	grpc.ClientStream
}

type daemonAttachClient struct {
	grpc.ClientStream
}

func (x *daemonAttachClient) Send(m *AttachRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *daemonAttachClient) Recv() (*ServerOutput, error) {
	m := new(ServerOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DaemonServer is the server API for Daemon service.
type DaemonServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Attach(Daemon_AttachServer) error
}

// UnimplementedDaemonServer can be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (*UnimplementedDaemonServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDaemonServer) Attach(srv Daemon_AttachServer) error {
	return status.Errorf(codes.Unimplemented, "method Attach not implemented")
}

func RegisterDaemonServer(s *grpc.Server, srv DaemonServer) {
	s.RegisterService(&_Daemon_serviceDesc, srv)
}

func _Daemon_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Daemon/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Attach_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DaemonServer).Attach(&daemonAttachServer{stream})
}

type Daemon_AttachServer interface {
	Send(*ServerOutput) error
	Recv() (*AttachRequest, error)
	grpc.ServerStream
}

type daemonAttachServer struct {
	grpc.ServerStream
}

func (x *daemonAttachServer) Send(m *ServerOutput) error {
	return x.ServerStream.SendMsg(m)
}

func (x *daemonAttachServer) Recv() (*AttachRequest, error) {
	m := new(AttachRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Daemon_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attach",
			Handler:       _Daemon_Attach_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "server.proto",
}
