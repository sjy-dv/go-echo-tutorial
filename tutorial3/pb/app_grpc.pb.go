// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// GRpcAppClient is the client API for GRpcApp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRpcAppClient interface {
	SignUp(ctx context.Context, in *ReqProtoUser, opts ...grpc.CallOption) (*Response, error)
	SignIn(ctx context.Context, in *LoginProto, opts ...grpc.CallOption) (*ResToken, error)
	UserInfo(ctx context.Context, in *ResToken, opts ...grpc.CallOption) (*ProtoUser, error)
	AllUser(ctx context.Context, in *QueryPage, opts ...grpc.CallOption) (*ProtoUsers, error)
}

type gRpcAppClient struct {
	cc grpc.ClientConnInterface
}

func NewGRpcAppClient(cc grpc.ClientConnInterface) GRpcAppClient {
	return &gRpcAppClient{cc}
}

func (c *gRpcAppClient) SignUp(ctx context.Context, in *ReqProtoUser, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.gRpcApp/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRpcAppClient) SignIn(ctx context.Context, in *LoginProto, opts ...grpc.CallOption) (*ResToken, error) {
	out := new(ResToken)
	err := c.cc.Invoke(ctx, "/pb.gRpcApp/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRpcAppClient) UserInfo(ctx context.Context, in *ResToken, opts ...grpc.CallOption) (*ProtoUser, error) {
	out := new(ProtoUser)
	err := c.cc.Invoke(ctx, "/pb.gRpcApp/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRpcAppClient) AllUser(ctx context.Context, in *QueryPage, opts ...grpc.CallOption) (*ProtoUsers, error) {
	out := new(ProtoUsers)
	err := c.cc.Invoke(ctx, "/pb.gRpcApp/AllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRpcAppServer is the server API for GRpcApp service.
// All implementations must embed UnimplementedGRpcAppServer
// for forward compatibility
type GRpcAppServer interface {
	SignUp(context.Context, *ReqProtoUser) (*Response, error)
	SignIn(context.Context, *LoginProto) (*ResToken, error)
	UserInfo(context.Context, *ResToken) (*ProtoUser, error)
	AllUser(context.Context, *QueryPage) (*ProtoUsers, error)
	mustEmbedUnimplementedGRpcAppServer()
}

// UnimplementedGRpcAppServer must be embedded to have forward compatible implementations.
type UnimplementedGRpcAppServer struct {
}

func (UnimplementedGRpcAppServer) SignUp(context.Context, *ReqProtoUser) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedGRpcAppServer) SignIn(context.Context, *LoginProto) (*ResToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedGRpcAppServer) UserInfo(context.Context, *ResToken) (*ProtoUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedGRpcAppServer) AllUser(context.Context, *QueryPage) (*ProtoUsers, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUser not implemented")
}
func (UnimplementedGRpcAppServer) mustEmbedUnimplementedGRpcAppServer() {}

// UnsafeGRpcAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRpcAppServer will
// result in compilation errors.
type UnsafeGRpcAppServer interface {
	mustEmbedUnimplementedGRpcAppServer()
}

func RegisterGRpcAppServer(s grpc.ServiceRegistrar, srv GRpcAppServer) {
	s.RegisterService(&GRpcApp_ServiceDesc, srv)
}

func _GRpcApp_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqProtoUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRpcAppServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gRpcApp/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRpcAppServer).SignUp(ctx, req.(*ReqProtoUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRpcApp_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRpcAppServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gRpcApp/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRpcAppServer).SignIn(ctx, req.(*LoginProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRpcApp_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRpcAppServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gRpcApp/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRpcAppServer).UserInfo(ctx, req.(*ResToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRpcApp_AllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRpcAppServer).AllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.gRpcApp/AllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRpcAppServer).AllUser(ctx, req.(*QueryPage))
	}
	return interceptor(ctx, in, info, handler)
}

// GRpcApp_ServiceDesc is the grpc.ServiceDesc for GRpcApp service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRpcApp_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.gRpcApp",
	HandlerType: (*GRpcAppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _GRpcApp_SignUp_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _GRpcApp_SignIn_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _GRpcApp_UserInfo_Handler,
		},
		{
			MethodName: "AllUser",
			Handler:    _GRpcApp_AllUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}
