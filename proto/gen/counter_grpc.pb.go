// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: counter.proto

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

// CounterServiceClient is the client API for CounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CounterServiceClient interface {
	GetListOrderFulfillments(ctx context.Context, in *GetListOrderFulfillmentRequests, opts ...grpc.CallOption) (*GetListOrderFulfillmentResponses, error)
	PlaceOrder(ctx context.Context, in *PlaceOrderRequests, opts ...grpc.CallOption) (*PlaceOrderResponses, error)
}

type counterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCounterServiceClient(cc grpc.ClientConnInterface) CounterServiceClient {
	return &counterServiceClient{cc}
}

func (c *counterServiceClient) GetListOrderFulfillments(ctx context.Context, in *GetListOrderFulfillmentRequests, opts ...grpc.CallOption) (*GetListOrderFulfillmentResponses, error) {
	out := new(GetListOrderFulfillmentResponses)
	err := c.cc.Invoke(ctx, "/go.coffeeshop.proto.counterapi.CounterService/GetListOrderFulfillments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterServiceClient) PlaceOrder(ctx context.Context, in *PlaceOrderRequests, opts ...grpc.CallOption) (*PlaceOrderResponses, error) {
	out := new(PlaceOrderResponses)
	err := c.cc.Invoke(ctx, "/go.coffeeshop.proto.counterapi.CounterService/PlaceOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServiceServer is the server API for CounterService service.
// All implementations should embed UnimplementedCounterServiceServer
// for forward compatibility
type CounterServiceServer interface {
	GetListOrderFulfillments(context.Context, *GetListOrderFulfillmentRequests) (*GetListOrderFulfillmentResponses, error)
	PlaceOrder(context.Context, *PlaceOrderRequests) (*PlaceOrderResponses, error)
}

// UnimplementedCounterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedCounterServiceServer struct {
}

func (UnimplementedCounterServiceServer) GetListOrderFulfillments(context.Context, *GetListOrderFulfillmentRequests) (*GetListOrderFulfillmentResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOrderFulfillments not implemented")
}
func (UnimplementedCounterServiceServer) PlaceOrder(context.Context, *PlaceOrderRequests) (*PlaceOrderResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrder not implemented")
}

// UnsafeCounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CounterServiceServer will
// result in compilation errors.
type UnsafeCounterServiceServer interface {
	mustEmbedUnimplementedCounterServiceServer()
}

func RegisterCounterServiceServer(s grpc.ServiceRegistrar, srv CounterServiceServer) {
	s.RegisterService(&CounterService_ServiceDesc, srv)
}

func _CounterService_GetListOrderFulfillments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOrderFulfillmentRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).GetListOrderFulfillments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.coffeeshop.proto.counterapi.CounterService/GetListOrderFulfillments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).GetListOrderFulfillments(ctx, req.(*GetListOrderFulfillmentRequests))
	}
	return interceptor(ctx, in, info, handler)
}

func _CounterService_PlaceOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderRequests)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServiceServer).PlaceOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.coffeeshop.proto.counterapi.CounterService/PlaceOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServiceServer).PlaceOrder(ctx, req.(*PlaceOrderRequests))
	}
	return interceptor(ctx, in, info, handler)
}

// CounterService_ServiceDesc is the grpc.ServiceDesc for CounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go.coffeeshop.proto.counterapi.CounterService",
	HandlerType: (*CounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetListOrderFulfillments",
			Handler:    _CounterService_GetListOrderFulfillments_Handler,
		},
		{
			MethodName: "PlaceOrder",
			Handler:    _CounterService_PlaceOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}