// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.29.0
// source: protobuf/user.proto

package _go

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CreateUser_GetUserName_FullMethodName = "/proto.CreateUser/GetUserName"
)

// CreateUserClient is the client API for CreateUser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateUserClient interface {
	GetUserName(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserNameResp, error)
}

type createUserClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateUserClient(cc grpc.ClientConnInterface) CreateUserClient {
	return &createUserClient{cc}
}

func (c *createUserClient) GetUserName(ctx context.Context, in *User, opts ...grpc.CallOption) (*UserNameResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserNameResp)
	err := c.cc.Invoke(ctx, CreateUser_GetUserName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateUserServer is the server API for CreateUser service.
// All implementations must embed UnimplementedCreateUserServer
// for forward compatibility
type CreateUserServer interface {
	GetUserName(context.Context, *User) (*UserNameResp, error)
	mustEmbedUnimplementedCreateUserServer()
}

// UnimplementedCreateUserServer must be embedded to have forward compatible implementations.
type UnimplementedCreateUserServer struct {
}

func (UnimplementedCreateUserServer) GetUserName(context.Context, *User) (*UserNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserName not implemented")
}
func (UnimplementedCreateUserServer) mustEmbedUnimplementedCreateUserServer() {}

// UnsafeCreateUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateUserServer will
// result in compilation errors.
type UnsafeCreateUserServer interface {
	mustEmbedUnimplementedCreateUserServer()
}

func RegisterCreateUserServer(s grpc.ServiceRegistrar, srv CreateUserServer) {
	s.RegisterService(&CreateUser_ServiceDesc, srv)
}

func _CreateUser_GetUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateUserServer).GetUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateUser_GetUserName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateUserServer).GetUserName(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateUser_ServiceDesc is the grpc.ServiceDesc for CreateUser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateUser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreateUser",
	HandlerType: (*CreateUserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserName",
			Handler:    _CreateUser_GetUserName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/user.proto",
}
