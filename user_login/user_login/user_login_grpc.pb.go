// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user_login.proto

package user_login

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

// UserLoginClient is the client API for UserLogin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserLoginClient interface {
	UserLogin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type userLoginClient struct {
	cc grpc.ClientConnInterface
}

func NewUserLoginClient(cc grpc.ClientConnInterface) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) UserLogin(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/user_login.user_login/userLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserLoginServer is the server API for UserLogin service.
// All implementations must embed UnimplementedUserLoginServer
// for forward compatibility
type UserLoginServer interface {
	UserLogin(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedUserLoginServer()
}

// UnimplementedUserLoginServer must be embedded to have forward compatible implementations.
type UnimplementedUserLoginServer struct {
}

func (UnimplementedUserLoginServer) UserLogin(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserLoginServer) mustEmbedUnimplementedUserLoginServer() {}

// UnsafeUserLoginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserLoginServer will
// result in compilation errors.
type UnsafeUserLoginServer interface {
	mustEmbedUnimplementedUserLoginServer()
}

func RegisterUserLoginServer(s grpc.ServiceRegistrar, srv UserLoginServer) {
	s.RegisterService(&UserLogin_ServiceDesc, srv)
}

func _UserLogin_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_login.user_login/userLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).UserLogin(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// UserLogin_ServiceDesc is the grpc.ServiceDesc for UserLogin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserLogin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_login.user_login",
	HandlerType: (*UserLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "userLogin",
			Handler:    _UserLogin_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_login.proto",
}
