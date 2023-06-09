// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: user.proto

package gen

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetItemTypes(ctx context.Context, in *GetItemTypesRequest, opts ...grpc.CallOption) (*GetItemTypesResponse, error)
	GetItemsByType(ctx context.Context, in *GetItemsByTypeRequest, opts ...grpc.CallOption) (*GetItemsByTypeResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	GetDeletedOrders(ctx context.Context, in *GetDeletedOrdersRequest, opts ...grpc.CallOption) (*GetDeletedOrdersResponse, error)
	DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetItemTypes(ctx context.Context, in *GetItemTypesRequest, opts ...grpc.CallOption) (*GetItemTypesResponse, error) {
	out := new(GetItemTypesResponse)
	err := c.cc.Invoke(ctx, "/platform.proto.productapi.UserService/GetItemTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetItemsByType(ctx context.Context, in *GetItemsByTypeRequest, opts ...grpc.CallOption) (*GetItemsByTypeResponse, error) {
	out := new(GetItemsByTypeResponse)
	err := c.cc.Invoke(ctx, "/platform.proto.productapi.UserService/GetItemsByType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := c.cc.Invoke(ctx, "/platform.proto.productapi.UserService/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetDeletedOrders(ctx context.Context, in *GetDeletedOrdersRequest, opts ...grpc.CallOption) (*GetDeletedOrdersResponse, error) {
	out := new(GetDeletedOrdersResponse)
	err := c.cc.Invoke(ctx, "/platform.proto.productapi.UserService/GetDeletedOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteOrders(ctx context.Context, in *DeleteOrdersRequest, opts ...grpc.CallOption) (*DeleteOrdersResponse, error) {
	out := new(DeleteOrdersResponse)
	err := c.cc.Invoke(ctx, "/platform.proto.productapi.UserService/DeleteOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetItemTypes(context.Context, *GetItemTypesRequest) (*GetItemTypesResponse, error)
	GetItemsByType(context.Context, *GetItemsByTypeRequest) (*GetItemsByTypeResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	GetDeletedOrders(context.Context, *GetDeletedOrdersRequest) (*GetDeletedOrdersResponse, error)
	DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error)
}

// UnimplementedUserServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetItemTypes(context.Context, *GetItemTypesRequest) (*GetItemTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemTypes not implemented")
}
func (UnimplementedUserServiceServer) GetItemsByType(context.Context, *GetItemsByTypeRequest) (*GetItemsByTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItemsByType not implemented")
}
func (UnimplementedUserServiceServer) GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserServiceServer) GetDeletedOrders(context.Context, *GetDeletedOrdersRequest) (*GetDeletedOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeletedOrders not implemented")
}
func (UnimplementedUserServiceServer) DeleteOrders(context.Context, *DeleteOrdersRequest) (*DeleteOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrders not implemented")
}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetItemTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetItemTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.proto.productapi.UserService/GetItemTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetItemTypes(ctx, req.(*GetItemTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetItemsByType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemsByTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetItemsByType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.proto.productapi.UserService/GetItemsByType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetItemsByType(ctx, req.(*GetItemsByTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.proto.productapi.UserService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetDeletedOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeletedOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetDeletedOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.proto.productapi.UserService/GetDeletedOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetDeletedOrders(ctx, req.(*GetDeletedOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.proto.productapi.UserService/DeleteOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteOrders(ctx, req.(*DeleteOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "platform.proto.productapi.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetItemTypes",
			Handler:    _UserService_GetItemTypes_Handler,
		},
		{
			MethodName: "GetItemsByType",
			Handler:    _UserService_GetItemsByType_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserService_GetUsers_Handler,
		},
		{
			MethodName: "GetDeletedOrders",
			Handler:    _UserService_GetDeletedOrders_Handler,
		},
		{
			MethodName: "DeleteOrders",
			Handler:    _UserService_DeleteOrders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
