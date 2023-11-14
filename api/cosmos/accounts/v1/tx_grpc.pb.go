// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: cosmos/accounts/v1/tx.proto

package accountsv1

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

const (
	Msg_Init_FullMethodName          = "/cosmos.accounts.v1.Msg/Init"
	Msg_Execute_FullMethodName       = "/cosmos.accounts.v1.Msg/Execute"
	Msg_ExecuteBundle_FullMethodName = "/cosmos.accounts.v1.Msg/ExecuteBundle"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// Init creates a new account in the chain.
	Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgInitResponse, error)
	// Execute executes a message to the target account.
	Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error)
	// ExecuteBundle pertains account abstraction, it is used by the bundler
	// to execute multiple UserOperations in a single transaction message.
	ExecuteBundle(ctx context.Context, in *MsgExecuteBundle, opts ...grpc.CallOption) (*MsgExecuteBundleResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Init(ctx context.Context, in *MsgInit, opts ...grpc.CallOption) (*MsgInitResponse, error) {
	out := new(MsgInitResponse)
	err := c.cc.Invoke(ctx, Msg_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Execute(ctx context.Context, in *MsgExecute, opts ...grpc.CallOption) (*MsgExecuteResponse, error) {
	out := new(MsgExecuteResponse)
	err := c.cc.Invoke(ctx, Msg_Execute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecuteBundle(ctx context.Context, in *MsgExecuteBundle, opts ...grpc.CallOption) (*MsgExecuteBundleResponse, error) {
	out := new(MsgExecuteBundleResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteBundle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// Init creates a new account in the chain.
	Init(context.Context, *MsgInit) (*MsgInitResponse, error)
	// Execute executes a message to the target account.
	Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error)
	// ExecuteBundle pertains account abstraction, it is used by the bundler
	// to execute multiple UserOperations in a single transaction message.
	ExecuteBundle(context.Context, *MsgExecuteBundle) (*MsgExecuteBundleResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) Init(context.Context, *MsgInit) (*MsgInitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedMsgServer) Execute(context.Context, *MsgExecute) (*MsgExecuteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedMsgServer) ExecuteBundle(context.Context, *MsgExecuteBundle) (*MsgExecuteBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteBundle not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Init(ctx, req.(*MsgInit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Execute(ctx, req.(*MsgExecute))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecuteBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteBundle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteBundle(ctx, req.(*MsgExecuteBundle))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.accounts.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Msg_Init_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _Msg_Execute_Handler,
		},
		{
			MethodName: "ExecuteBundle",
			Handler:    _Msg_ExecuteBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/accounts/v1/tx.proto",
}
