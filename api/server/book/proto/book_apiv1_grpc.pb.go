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

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookServiceClient interface {
	ListBookshelf(ctx context.Context, in *ListBookshelfRequest, opts ...grpc.CallOption) (*BookshelfListResponse, error)
	GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	ReadBookshelf(ctx context.Context, in *ReadBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error)
	ReadingBookshelf(ctx context.Context, in *ReadingBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error)
	StackBookshelf(ctx context.Context, in *StackBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error)
	WantBookshelf(ctx context.Context, in *WantBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error)
	ReleaseBookshelf(ctx context.Context, in *ReleaseBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*EmptyBook, error)
	DeleteBookshelf(ctx context.Context, in *DeleteBookshelfRequest, opts ...grpc.CallOption) (*EmptyBook, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) ListBookshelf(ctx context.Context, in *ListBookshelfRequest, opts ...grpc.CallOption) (*BookshelfListResponse, error) {
	out := new(BookshelfListResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/ListBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBook(ctx context.Context, in *GetBookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/GetBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) CreateBook(ctx context.Context, in *CreateBookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/CreateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ReadBookshelf(ctx context.Context, in *ReadBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error) {
	out := new(BookshelfResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/ReadBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ReadingBookshelf(ctx context.Context, in *ReadingBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error) {
	out := new(BookshelfResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/ReadingBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) StackBookshelf(ctx context.Context, in *StackBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error) {
	out := new(BookshelfResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/StackBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) WantBookshelf(ctx context.Context, in *WantBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error) {
	out := new(BookshelfResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/WantBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) ReleaseBookshelf(ctx context.Context, in *ReleaseBookshelfRequest, opts ...grpc.CallOption) (*BookshelfResponse, error) {
	out := new(BookshelfResponse)
	err := c.cc.Invoke(ctx, "/proto.BookService/ReleaseBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*EmptyBook, error) {
	out := new(EmptyBook)
	err := c.cc.Invoke(ctx, "/proto.BookService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) DeleteBookshelf(ctx context.Context, in *DeleteBookshelfRequest, opts ...grpc.CallOption) (*EmptyBook, error) {
	out := new(EmptyBook)
	err := c.cc.Invoke(ctx, "/proto.BookService/DeleteBookshelf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
// All implementations must embed UnimplementedBookServiceServer
// for forward compatibility
type BookServiceServer interface {
	ListBookshelf(context.Context, *ListBookshelfRequest) (*BookshelfListResponse, error)
	GetBook(context.Context, *GetBookRequest) (*BookResponse, error)
	CreateBook(context.Context, *CreateBookRequest) (*BookResponse, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*BookResponse, error)
	ReadBookshelf(context.Context, *ReadBookshelfRequest) (*BookshelfResponse, error)
	ReadingBookshelf(context.Context, *ReadingBookshelfRequest) (*BookshelfResponse, error)
	StackBookshelf(context.Context, *StackBookshelfRequest) (*BookshelfResponse, error)
	WantBookshelf(context.Context, *WantBookshelfRequest) (*BookshelfResponse, error)
	ReleaseBookshelf(context.Context, *ReleaseBookshelfRequest) (*BookshelfResponse, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*EmptyBook, error)
	DeleteBookshelf(context.Context, *DeleteBookshelfRequest) (*EmptyBook, error)
	mustEmbedUnimplementedBookServiceServer()
}

// UnimplementedBookServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (UnimplementedBookServiceServer) ListBookshelf(context.Context, *ListBookshelfRequest) (*BookshelfListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBookshelf not implemented")
}
func (UnimplementedBookServiceServer) GetBook(context.Context, *GetBookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServiceServer) CreateBook(context.Context, *CreateBookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (UnimplementedBookServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServiceServer) ReadBookshelf(context.Context, *ReadBookshelfRequest) (*BookshelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBookshelf not implemented")
}
func (UnimplementedBookServiceServer) ReadingBookshelf(context.Context, *ReadingBookshelfRequest) (*BookshelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadingBookshelf not implemented")
}
func (UnimplementedBookServiceServer) StackBookshelf(context.Context, *StackBookshelfRequest) (*BookshelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StackBookshelf not implemented")
}
func (UnimplementedBookServiceServer) WantBookshelf(context.Context, *WantBookshelfRequest) (*BookshelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WantBookshelf not implemented")
}
func (UnimplementedBookServiceServer) ReleaseBookshelf(context.Context, *ReleaseBookshelfRequest) (*BookshelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseBookshelf not implemented")
}
func (UnimplementedBookServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*EmptyBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedBookServiceServer) DeleteBookshelf(context.Context, *DeleteBookshelfRequest) (*EmptyBook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookshelf not implemented")
}
func (UnimplementedBookServiceServer) mustEmbedUnimplementedBookServiceServer() {}

// UnsafeBookServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServiceServer will
// result in compilation errors.
type UnsafeBookServiceServer interface {
	mustEmbedUnimplementedBookServiceServer()
}

func RegisterBookServiceServer(s grpc.ServiceRegistrar, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_ListBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ListBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/ListBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ListBookshelf(ctx, req.(*ListBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/GetBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBook(ctx, req.(*GetBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_CreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).CreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/CreateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).CreateBook(ctx, req.(*CreateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ReadBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ReadBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/ReadBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ReadBookshelf(ctx, req.(*ReadBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ReadingBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadingBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ReadingBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/ReadingBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ReadingBookshelf(ctx, req.(*ReadingBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_StackBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StackBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).StackBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/StackBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).StackBookshelf(ctx, req.(*StackBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_WantBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WantBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).WantBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/WantBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).WantBookshelf(ctx, req.(*WantBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_ReleaseBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).ReleaseBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/ReleaseBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).ReleaseBookshelf(ctx, req.(*ReleaseBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_DeleteBookshelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookshelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).DeleteBookshelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BookService/DeleteBookshelf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).DeleteBookshelf(ctx, req.(*DeleteBookshelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListBookshelf",
			Handler:    _BookService_ListBookshelf_Handler,
		},
		{
			MethodName: "GetBook",
			Handler:    _BookService_GetBook_Handler,
		},
		{
			MethodName: "CreateBook",
			Handler:    _BookService_CreateBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _BookService_UpdateBook_Handler,
		},
		{
			MethodName: "ReadBookshelf",
			Handler:    _BookService_ReadBookshelf_Handler,
		},
		{
			MethodName: "ReadingBookshelf",
			Handler:    _BookService_ReadingBookshelf_Handler,
		},
		{
			MethodName: "StackBookshelf",
			Handler:    _BookService_StackBookshelf_Handler,
		},
		{
			MethodName: "WantBookshelf",
			Handler:    _BookService_WantBookshelf_Handler,
		},
		{
			MethodName: "ReleaseBookshelf",
			Handler:    _BookService_ReleaseBookshelf_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _BookService_DeleteBook_Handler,
		},
		{
			MethodName: "DeleteBookshelf",
			Handler:    _BookService_DeleteBookshelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/book_apiv1.proto",
}
