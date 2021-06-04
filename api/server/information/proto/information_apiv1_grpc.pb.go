// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Reply(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Reply(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/proto.NotificationService/Reply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	Reply(context.Context, *HelloRequest) (*HelloResponse, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Reply(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reply not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NotificationService/Reply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Reply(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reply",
			Handler:    _NotificationService_Reply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/information_apiv1.proto",
}

// InquiryServiceClient is the client API for InquiryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InquiryServiceClient interface {
	CreateInquiry(ctx context.Context, in *CreateInquiryRequest, opts ...grpc.CallOption) (*InquiryResponse, error)
}

type inquiryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInquiryServiceClient(cc grpc.ClientConnInterface) InquiryServiceClient {
	return &inquiryServiceClient{cc}
}

func (c *inquiryServiceClient) CreateInquiry(ctx context.Context, in *CreateInquiryRequest, opts ...grpc.CallOption) (*InquiryResponse, error) {
	out := new(InquiryResponse)
	err := c.cc.Invoke(ctx, "/proto.InquiryService/CreateInquiry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InquiryServiceServer is the server API for InquiryService service.
// All implementations must embed UnimplementedInquiryServiceServer
// for forward compatibility
type InquiryServiceServer interface {
	CreateInquiry(context.Context, *CreateInquiryRequest) (*InquiryResponse, error)
	mustEmbedUnimplementedInquiryServiceServer()
}

// UnimplementedInquiryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInquiryServiceServer struct {
}

func (UnimplementedInquiryServiceServer) CreateInquiry(context.Context, *CreateInquiryRequest) (*InquiryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInquiry not implemented")
}
func (UnimplementedInquiryServiceServer) mustEmbedUnimplementedInquiryServiceServer() {}

// UnsafeInquiryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InquiryServiceServer will
// result in compilation errors.
type UnsafeInquiryServiceServer interface {
	mustEmbedUnimplementedInquiryServiceServer()
}

func RegisterInquiryServiceServer(s grpc.ServiceRegistrar, srv InquiryServiceServer) {
	s.RegisterService(&_InquiryService_serviceDesc, srv)
}

func _InquiryService_CreateInquiry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InquiryServiceServer).CreateInquiry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.InquiryService/CreateInquiry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InquiryServiceServer).CreateInquiry(ctx, req.(*CreateInquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InquiryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.InquiryService",
	HandlerType: (*InquiryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateInquiry",
			Handler:    _InquiryService_CreateInquiry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/information_apiv1.proto",
}
