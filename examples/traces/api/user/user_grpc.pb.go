// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc-b"
	codes "google.golang.org/grpc-b/codes"
	status "google.golang.org/grpc-b/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc-b/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	GetMyMessages(ctx context.Context, in *GetMyMessagesRequest, opts ...grpc.CallOption) (*GetMyMessagesReply, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetMyMessages(ctx context.Context, in *GetMyMessagesRequest, opts ...grpc.CallOption) (*GetMyMessagesReply, error) {
	out := new(GetMyMessagesReply)
	err := c.cc.Invoke(ctx, "/api.user.v1.User/GetMyMessages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) GetMyMessages(context.Context, *GetMyMessagesRequest) (*GetMyMessagesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyMessages not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_GetMyMessages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyMessagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetMyMessages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.user.v1.User/GetMyMessages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetMyMessages(ctx, req.(*GetMyMessagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.user.v1.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyMessages",
			Handler:    _User_GetMyMessages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traces/api/user/user.proto",
}
