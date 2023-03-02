// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: api/proto/main.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SayClient is the client API for Say service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SayClient interface {
	// 测试HelloWorld
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayClient struct {
	cc grpc.ClientConnInterface
}

func NewSayClient(cc grpc.ClientConnInterface) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.Say/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayServer is the server API for Say service.
// All implementations must embed UnimplementedSayServer
// for forward compatibility
type SayServer interface {
	// 测试HelloWorld
	Hello(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedSayServer()
}

// UnimplementedSayServer must be embedded to have forward compatible implementations.
type UnimplementedSayServer struct {
}

func (UnimplementedSayServer) Hello(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedSayServer) mustEmbedUnimplementedSayServer() {}

// UnsafeSayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SayServer will
// result in compilation errors.
type UnsafeSayServer interface {
	mustEmbedUnimplementedSayServer()
}

func RegisterSayServer(s grpc.ServiceRegistrar, srv SayServer) {
	s.RegisterService(&Say_ServiceDesc, srv)
}

func _Say_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Say/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Say_ServiceDesc is the grpc.ServiceDesc for Say service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Say_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Say_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/main.proto",
}
