// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: consensus.proto

package consensus

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
	Consensus_StartFunction_FullMethodName = "/consensus/StartFunction"
	Consensus_PassToken_FullMethodName     = "/consensus/PassToken"
)

// ConsensusClient is the client API for Consensus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConsensusClient interface {
	StartFunction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SuccessStart, error)
	PassToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error)
}

type consensusClient struct {
	cc grpc.ClientConnInterface
}

func NewConsensusClient(cc grpc.ClientConnInterface) ConsensusClient {
	return &consensusClient{cc}
}

func (c *consensusClient) StartFunction(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SuccessStart, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SuccessStart)
	err := c.cc.Invoke(ctx, Consensus_StartFunction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consensusClient) PassToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Consensus_PassToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsensusServer is the server API for Consensus service.
// All implementations must embed UnimplementedConsensusServer
// for forward compatibility.
type ConsensusServer interface {
	StartFunction(context.Context, *Empty) (*SuccessStart, error)
	PassToken(context.Context, *Token) (*Empty, error)
	mustEmbedUnimplementedConsensusServer()
}

// UnimplementedConsensusServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConsensusServer struct{}

func (UnimplementedConsensusServer) StartFunction(context.Context, *Empty) (*SuccessStart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFunction not implemented")
}
func (UnimplementedConsensusServer) PassToken(context.Context, *Token) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PassToken not implemented")
}
func (UnimplementedConsensusServer) mustEmbedUnimplementedConsensusServer() {}
func (UnimplementedConsensusServer) testEmbeddedByValue()                   {}

// UnsafeConsensusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConsensusServer will
// result in compilation errors.
type UnsafeConsensusServer interface {
	mustEmbedUnimplementedConsensusServer()
}

func RegisterConsensusServer(s grpc.ServiceRegistrar, srv ConsensusServer) {
	// If the following call pancis, it indicates UnimplementedConsensusServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Consensus_ServiceDesc, srv)
}

func _Consensus_StartFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServer).StartFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Consensus_StartFunction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServer).StartFunction(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Consensus_PassToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsensusServer).PassToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Consensus_PassToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsensusServer).PassToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

// Consensus_ServiceDesc is the grpc.ServiceDesc for Consensus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Consensus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "consensus",
	HandlerType: (*ConsensusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFunction",
			Handler:    _Consensus_StartFunction_Handler,
		},
		{
			MethodName: "PassToken",
			Handler:    _Consensus_PassToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consensus.proto",
}
