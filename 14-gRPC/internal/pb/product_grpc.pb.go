// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/product.proto

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
	CreateProductStream(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductStreamClient, error)
	CreateProductStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductStreamBidirectionalClient, error)
	ListProducts(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProductList, error)
	GetCategory(ctx context.Context, in *ProductGetRequest, opts ...grpc.CallOption) (*Product, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pb.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) CreateProductStream(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/pb.ProductService/CreateProductStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceCreateProductStreamClient{stream}
	return x, nil
}

type ProductService_CreateProductStreamClient interface {
	Send(*CreateProductRequest) error
	CloseAndRecv() (*ProductList, error)
	grpc.ClientStream
}

type productServiceCreateProductStreamClient struct {
	grpc.ClientStream
}

func (x *productServiceCreateProductStreamClient) Send(m *CreateProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceCreateProductStreamClient) CloseAndRecv() (*ProductList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) CreateProductStreamBidirectional(ctx context.Context, opts ...grpc.CallOption) (ProductService_CreateProductStreamBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[1], "/pb.ProductService/CreateProductStreamBidirectional", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceCreateProductStreamBidirectionalClient{stream}
	return x, nil
}

type ProductService_CreateProductStreamBidirectionalClient interface {
	Send(*CreateProductRequest) error
	Recv() (*Product, error)
	grpc.ClientStream
}

type productServiceCreateProductStreamBidirectionalClient struct {
	grpc.ClientStream
}

func (x *productServiceCreateProductStreamBidirectionalClient) Send(m *CreateProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *productServiceCreateProductStreamBidirectionalClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) ListProducts(ctx context.Context, in *Blank, opts ...grpc.CallOption) (*ProductList, error) {
	out := new(ProductList)
	err := c.cc.Invoke(ctx, "/pb.ProductService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetCategory(ctx context.Context, in *ProductGetRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/pb.ProductService/GetCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	CreateProductStream(ProductService_CreateProductStreamServer) error
	CreateProductStreamBidirectional(ProductService_CreateProductStreamBidirectionalServer) error
	ListProducts(context.Context, *Blank) (*ProductList, error)
	GetCategory(context.Context, *ProductGetRequest) (*Product, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) CreateProductStream(ProductService_CreateProductStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProductStream not implemented")
}
func (UnimplementedProductServiceServer) CreateProductStreamBidirectional(ProductService_CreateProductStreamBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateProductStreamBidirectional not implemented")
}
func (UnimplementedProductServiceServer) ListProducts(context.Context, *Blank) (*ProductList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductServiceServer) GetCategory(context.Context, *ProductGetRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_CreateProductStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).CreateProductStream(&productServiceCreateProductStreamServer{stream})
}

type ProductService_CreateProductStreamServer interface {
	SendAndClose(*ProductList) error
	Recv() (*CreateProductRequest, error)
	grpc.ServerStream
}

type productServiceCreateProductStreamServer struct {
	grpc.ServerStream
}

func (x *productServiceCreateProductStreamServer) SendAndClose(m *ProductList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceCreateProductStreamServer) Recv() (*CreateProductRequest, error) {
	m := new(CreateProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductService_CreateProductStreamBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProductServiceServer).CreateProductStreamBidirectional(&productServiceCreateProductStreamBidirectionalServer{stream})
}

type ProductService_CreateProductStreamBidirectionalServer interface {
	Send(*Product) error
	Recv() (*CreateProductRequest, error)
	grpc.ServerStream
}

type productServiceCreateProductStreamBidirectionalServer struct {
	grpc.ServerStream
}

func (x *productServiceCreateProductStreamBidirectionalServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func (x *productServiceCreateProductStreamBidirectionalServer) Recv() (*CreateProductRequest, error) {
	m := new(CreateProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProductService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blank)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProducts(ctx, req.(*Blank))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductService/GetCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetCategory(ctx, req.(*ProductGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "ListProducts",
			Handler:    _ProductService_ListProducts_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ProductService_GetCategory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateProductStream",
			Handler:       _ProductService_CreateProductStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "CreateProductStreamBidirectional",
			Handler:       _ProductService_CreateProductStreamBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}