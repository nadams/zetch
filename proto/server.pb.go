// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	server.proto

It has these top-level messages:
	Server
	ListRequest
	ListResponse
	AttachRequest
	ServerOutput
	StopResponse
	StopRequest
	StartResponse
	StartRequest
	RestartResponse
	RestartRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Server struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	HostName string   `protobuf:"bytes,2,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
	Address  string   `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Port     string   `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	Iwad     string   `protobuf:"bytes,5,opt,name=iwad" json:"iwad,omitempty"`
	Pwads    []string `protobuf:"bytes,6,rep,name=pwads" json:"pwads,omitempty"`
	GameType string   `protobuf:"bytes,7,opt,name=game_type,json=gameType" json:"game_type,omitempty"`
	Status   string   `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto1.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ListRequest struct {
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto1.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListResponse struct {
	Servers []*Server `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto1.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListResponse) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

type AttachRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *AttachRequest) Reset()                    { *m = AttachRequest{} }
func (m *AttachRequest) String() string            { return proto1.CompactTextString(m) }
func (*AttachRequest) ProtoMessage()               {}
func (*AttachRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ServerOutput struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *ServerOutput) Reset()                    { *m = ServerOutput{} }
func (m *ServerOutput) String() string            { return proto1.CompactTextString(m) }
func (*ServerOutput) ProtoMessage()               {}
func (*ServerOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type StopResponse struct {
}

func (m *StopResponse) Reset()                    { *m = StopResponse{} }
func (m *StopResponse) String() string            { return proto1.CompactTextString(m) }
func (*StopResponse) ProtoMessage()               {}
func (*StopResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type StopRequest struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *StopRequest) Reset()                    { *m = StopRequest{} }
func (m *StopRequest) String() string            { return proto1.CompactTextString(m) }
func (*StopRequest) ProtoMessage()               {}
func (*StopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StartResponse struct {
}

func (m *StartResponse) Reset()                    { *m = StartResponse{} }
func (m *StartResponse) String() string            { return proto1.CompactTextString(m) }
func (*StartResponse) ProtoMessage()               {}
func (*StartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type StartRequest struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (m *StartRequest) String() string            { return proto1.CompactTextString(m) }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type RestartResponse struct {
}

func (m *RestartResponse) Reset()                    { *m = RestartResponse{} }
func (m *RestartResponse) String() string            { return proto1.CompactTextString(m) }
func (*RestartResponse) ProtoMessage()               {}
func (*RestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type RestartRequest struct {
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *RestartRequest) Reset()                    { *m = RestartRequest{} }
func (m *RestartRequest) String() string            { return proto1.CompactTextString(m) }
func (*RestartRequest) ProtoMessage()               {}
func (*RestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto1.RegisterType((*Server)(nil), "proto.Server")
	proto1.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto1.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto1.RegisterType((*AttachRequest)(nil), "proto.AttachRequest")
	proto1.RegisterType((*ServerOutput)(nil), "proto.ServerOutput")
	proto1.RegisterType((*StopResponse)(nil), "proto.StopResponse")
	proto1.RegisterType((*StopRequest)(nil), "proto.StopRequest")
	proto1.RegisterType((*StartResponse)(nil), "proto.StartResponse")
	proto1.RegisterType((*StartRequest)(nil), "proto.StartRequest")
	proto1.RegisterType((*RestartResponse)(nil), "proto.RestartResponse")
	proto1.RegisterType((*RestartRequest)(nil), "proto.RestartRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Daemon service

type DaemonClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Attach(ctx context.Context, opts ...grpc.CallOption) (Daemon_AttachClient, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error)
}

type daemonClient struct {
	cc *grpc.ClientConn
}

func NewDaemonClient(cc *grpc.ClientConn) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/proto.Daemon/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Attach(ctx context.Context, opts ...grpc.CallOption) (Daemon_AttachClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Daemon_serviceDesc.Streams[0], c.cc, "/proto.Daemon/Attach", opts...)
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

func (c *daemonClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := grpc.Invoke(ctx, "/proto.Daemon/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := grpc.Invoke(ctx, "/proto.Daemon/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) Restart(ctx context.Context, in *RestartRequest, opts ...grpc.CallOption) (*RestartResponse, error) {
	out := new(RestartResponse)
	err := grpc.Invoke(ctx, "/proto.Daemon/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Daemon service

type DaemonServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Attach(Daemon_AttachServer) error
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	Restart(context.Context, *RestartRequest) (*RestartResponse, error)
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

func _Daemon_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Daemon/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Daemon/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Daemon/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).Restart(ctx, req.(*RestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Daemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Daemon_List_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Daemon_Stop_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Daemon_Start_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Daemon_Restart_Handler,
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
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x0e, 0xd2, 0x40,
	0x10, 0xc6, 0x29, 0xa5, 0x2d, 0x1d, 0x5a, 0xd0, 0xb5, 0x92, 0x0d, 0x5e, 0x9a, 0xd5, 0x68, 0x4f,
	0x44, 0x51, 0x63, 0xf4, 0x66, 0xe2, 0xd1, 0x68, 0x52, 0xbc, 0x93, 0xd5, 0x6e, 0x80, 0x43, 0xe9,
	0xda, 0xdd, 0x4a, 0x78, 0x16, 0xdf, 0xc7, 0xe7, 0x32, 0xfb, 0xa7, 0xff, 0x8c, 0xe1, 0xd4, 0x9d,
	0xef, 0x9b, 0x6f, 0xa6, 0xf3, 0x83, 0x48, 0xb0, 0xfa, 0x17, 0xab, 0xb7, 0xbc, 0xae, 0x64, 0x85,
	0x3c, 0xfd, 0x21, 0x7f, 0x1c, 0xf0, 0xf7, 0x5a, 0x47, 0x08, 0x66, 0x17, 0x5a, 0x32, 0xec, 0xa4,
	0x4e, 0x16, 0xe6, 0xfa, 0x8d, 0x9e, 0x40, 0x78, 0xaa, 0x84, 0x3c, 0x68, 0x63, 0xaa, 0x8d, 0xb9,
	0x12, 0xbe, 0x28, 0x13, 0x43, 0x40, 0x8b, 0xa2, 0x66, 0x42, 0x60, 0x57, 0x5b, 0x6d, 0xa9, 0x46,
	0xf1, 0xaa, 0x96, 0x78, 0x66, 0x46, 0xa9, 0xb7, 0xd2, 0xce, 0x57, 0x5a, 0x60, 0xcf, 0x68, 0xea,
	0x8d, 0x12, 0xf0, 0xf8, 0x95, 0x16, 0x02, 0xfb, 0xa9, 0x9b, 0x85, 0xb9, 0x29, 0xd4, 0xd2, 0x23,
	0x2d, 0xd9, 0x41, 0xde, 0x38, 0xc3, 0x81, 0x59, 0xaa, 0x84, 0x6f, 0x37, 0xce, 0xd0, 0x1a, 0x7c,
	0x21, 0xa9, 0x6c, 0x04, 0x9e, 0x6b, 0xc7, 0x56, 0x24, 0x86, 0xc5, 0xe7, 0xb3, 0x90, 0x39, 0xfb,
	0xd9, 0x30, 0x21, 0xc9, 0x3b, 0x88, 0x4c, 0x29, 0x78, 0x75, 0x11, 0x0c, 0xbd, 0x80, 0xc0, 0x9c,
	0x2f, 0xb0, 0x93, 0xba, 0xd9, 0x62, 0x17, 0x1b, 0x0e, 0x5b, 0x73, 0x7c, 0xde, 0xba, 0xe4, 0x2d,
	0xc4, 0x1f, 0xa5, 0xa4, 0x3f, 0x4e, 0x76, 0xd2, 0x7f, 0xb1, 0x3c, 0x00, 0xb7, 0x14, 0x47, 0x0b,
	0x44, 0x3d, 0x49, 0x0a, 0x91, 0x99, 0xf4, 0xb5, 0x91, 0xbc, 0x91, 0x6d, 0x87, 0xd3, 0x77, 0x2c,
	0x21, 0xda, 0xcb, 0x8a, 0xb7, 0x7f, 0x44, 0x9e, 0xc2, 0xc2, 0xd4, 0x66, 0x4d, 0x02, 0x9e, 0x1a,
	0x6d, 0x7e, 0x2f, 0xcc, 0x4d, 0x41, 0x56, 0x10, 0xef, 0x25, 0xad, 0xbb, 0x3b, 0xc8, 0x33, 0x35,
	0x45, 0x0b, 0xf7, 0x62, 0x0f, 0x61, 0x95, 0x33, 0x31, 0x0a, 0x3e, 0x87, 0x65, 0x27, 0xdd, 0x89,
	0xee, 0x7e, 0x4f, 0xc1, 0xff, 0x44, 0x59, 0x59, 0x5d, 0xd0, 0x2b, 0x98, 0x29, 0x86, 0x08, 0x59,
	0x54, 0x03, 0xbe, 0x9b, 0x47, 0x23, 0xcd, 0xee, 0x98, 0xa0, 0xf7, 0xe0, 0x1b, 0x7a, 0x28, 0xb1,
	0x0d, 0x23, 0x98, 0x5d, 0x6c, 0xc8, 0x8a, 0x4c, 0x32, 0xe7, 0xa5, 0xa3, 0xb6, 0x29, 0x1e, 0xdd,
	0xb6, 0x01, 0x9c, 0x3e, 0x36, 0x04, 0x38, 0x41, 0x6f, 0xc0, 0xd3, 0x30, 0x50, 0xef, 0xf7, 0xf7,
	0x6d, 0x92, 0xb1, 0xd8, 0xa5, 0x3e, 0x40, 0x60, 0x49, 0xa0, 0xc7, 0xb6, 0x65, 0x4c, 0x66, 0xb3,
	0xfe, 0x57, 0x6e, 0xb3, 0xdf, 0x7d, 0x6d, 0xbc, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x99,
	0xc2, 0x16, 0x4c, 0x03, 0x00, 0x00,
}
