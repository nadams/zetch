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
	Id       string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
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
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
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

func init() {
	proto1.RegisterType((*Server)(nil), "proto.Server")
	proto1.RegisterType((*ListRequest)(nil), "proto.ListRequest")
	proto1.RegisterType((*ListResponse)(nil), "proto.ListResponse")
	proto1.RegisterType((*AttachRequest)(nil), "proto.AttachRequest")
	proto1.RegisterType((*ServerOutput)(nil), "proto.ServerOutput")
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
	Attach(ctx context.Context, in *AttachRequest, opts ...grpc.CallOption) (Daemon_AttachClient, error)
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

func (c *daemonClient) Attach(ctx context.Context, in *AttachRequest, opts ...grpc.CallOption) (Daemon_AttachClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Daemon_serviceDesc.Streams[0], c.cc, "/proto.Daemon/Attach", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonAttachClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_AttachClient interface {
	Recv() (*ServerOutput, error)
	grpc.ClientStream
}

type daemonAttachClient struct {
	grpc.ClientStream
}

func (x *daemonAttachClient) Recv() (*ServerOutput, error) {
	m := new(ServerOutput)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Daemon service

type DaemonServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Attach(*AttachRequest, Daemon_AttachServer) error
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
	m := new(AttachRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).Attach(m, &daemonAttachServer{stream})
}

type Daemon_AttachServer interface {
	Send(*ServerOutput) error
	grpc.ServerStream
}

type daemonAttachServer struct {
	grpc.ServerStream
}

func (x *daemonAttachServer) Send(m *ServerOutput) error {
	return x.ServerStream.SendMsg(m)
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
		},
	},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0x9b, 0xa6, 0x75, 0xdb, 0x6b, 0xfb, 0xd7, 0x5f, 0x47, 0x85, 0x2c, 0x18, 0x88, 0xbc,
	0xd0, 0xa9, 0x82, 0x32, 0x74, 0x46, 0x62, 0x44, 0x42, 0x0a, 0xec, 0xc8, 0x60, 0xab, 0x64, 0x48,
	0x62, 0x72, 0x17, 0xaa, 0x7e, 0x34, 0xbe, 0x1d, 0xb2, 0x9d, 0xa2, 0x96, 0x29, 0xef, 0xfd, 0xfc,
	0x9c, 0x7b, 0x67, 0x98, 0x91, 0x6d, 0xbe, 0x6c, 0xb3, 0x72, 0x4d, 0xcd, 0x35, 0x0e, 0xc3, 0x47,
	0x7d, 0x27, 0x20, 0x9e, 0x03, 0xc7, 0x7f, 0xd0, 0x2f, 0x8c, 0x4c, 0xb2, 0x64, 0x39, 0xc9, 0xfb,
	0x85, 0x41, 0x84, 0x41, 0xa5, 0x4b, 0x2b, 0xfb, 0x81, 0x04, 0x8d, 0x12, 0x46, 0xda, 0x98, 0xc6,
	0x12, 0xc9, 0x34, 0xe0, 0x83, 0xf5, 0x69, 0x57, 0x37, 0x2c, 0x07, 0x31, 0xed, 0xb5, 0x67, 0xc5,
	0x4e, 0x1b, 0x39, 0x8c, 0xcc, 0x6b, 0x5c, 0xc0, 0xd0, 0xed, 0xb4, 0x21, 0x29, 0xb2, 0x74, 0x39,
	0xc9, 0xa3, 0xc1, 0x4b, 0x98, 0x6c, 0x75, 0x69, 0x5f, 0x79, 0xef, 0xac, 0x1c, 0x85, 0xf8, 0xd8,
	0x83, 0x97, 0xbd, 0xb3, 0x78, 0x0e, 0x82, 0x58, 0x73, 0x4b, 0x72, 0x1c, 0x4e, 0x3a, 0xa7, 0xe6,
	0x30, 0x7d, 0x2c, 0x88, 0x73, 0xfb, 0xd9, 0x5a, 0x62, 0xb5, 0x81, 0x59, 0xb4, 0xe4, 0xea, 0x8a,
	0x2c, 0x5e, 0xc3, 0x28, 0x6e, 0x4c, 0x32, 0xc9, 0xd2, 0xe5, 0x74, 0x3d, 0x8f, 0xab, 0xaf, 0xe2,
	0xbe, 0xf9, 0xe1, 0x54, 0x5d, 0xc1, 0xfc, 0x9e, 0x59, 0xbf, 0x7f, 0x74, 0x7f, 0xfa, 0xfb, 0x12,
	0x2a, 0x83, 0x59, 0xbc, 0xf3, 0xd4, 0xb2, 0x6b, 0x19, 0xff, 0x43, 0x5a, 0xd2, 0xb6, 0x0b, 0x78,
	0xb9, 0x66, 0x10, 0x0f, 0xda, 0x96, 0x75, 0x85, 0xb7, 0x30, 0xf0, 0x2d, 0x10, 0xbb, 0x61, 0x47,
	0x0d, 0x2f, 0xce, 0x4e, 0x58, 0xac, 0xa9, 0x7a, 0xb8, 0x01, 0x11, 0xe7, 0xe3, 0xa2, 0x0b, 0x9c,
	0xd4, 0xf9, 0xbd, 0x76, 0xdc, 0x41, 0xf5, 0x6e, 0x92, 0x37, 0x11, 0xf8, 0xdd, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x05, 0xd9, 0x99, 0xa2, 0xda, 0x01, 0x00, 0x00,
}
