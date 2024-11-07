// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto.proto

package Client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Client_PassAlong_FullMethodName = "/Client/passAlong"
)

// ClientClient is the client API for Client service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Interface exposed by the client
type ClientClient interface {
	PassAlong(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*Empty, error)
}

type clientClient struct {
	cc grpc.ClientConnInterface
}

func NewClientClient(cc grpc.ClientConnInterface) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) PassAlong(ctx context.Context, in *ClientMessage, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Client_PassAlong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServer is the server API for Client service.
// All implementations must embed UnimplementedClientServer
// for forward compatibility.
//
// Interface exposed by the client
type ClientServer interface {
	PassAlong(context.Context, *ClientMessage) (*Empty, error)
	mustEmbedUnimplementedClientServer()
}

// UnimplementedClientServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClientServer struct{}

func (UnimplementedClientServer) PassAlong(context.Context, *ClientMessage) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassAlong not implemented")
}
func (UnimplementedClientServer) mustEmbedUnimplementedClientServer() {}
func (UnimplementedClientServer) testEmbeddedByValue()                {}

// UnsafeClientServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServer will
// result in compilation errors.
type UnsafeClientServer interface {
	mustEmbedUnimplementedClientServer()
}

func RegisterClientServer(s grpc.ServiceRegistrar, srv ClientServer) {
	// If the following call pancis, it indicates UnimplementedClientServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Client_ServiceDesc, srv)
}

func _Client_PassAlong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).PassAlong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Client_PassAlong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).PassAlong(ctx, req.(*ClientMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Client_ServiceDesc is the grpc.ServiceDesc for Client service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Client_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "passAlong",
			Handler:    _Client_PassAlong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
