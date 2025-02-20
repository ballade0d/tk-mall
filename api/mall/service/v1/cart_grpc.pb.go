// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: service/v1/cart.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CartService_GetCart_FullMethodName        = "/api.mall.service.v1.CartService/GetCart"
	CartService_AddToCart_FullMethodName      = "/api.mall.service.v1.CartService/AddToCart"
	CartService_RemoveFromCart_FullMethodName = "/api.mall.service.v1.CartService/RemoveFromCart"
	CartService_ClearCart_FullMethodName      = "/api.mall.service.v1.CartService/ClearCart"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
	AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error)
	RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddToCart(ctx context.Context, in *AddToCartRequest, opts ...grpc.CallOption) (*AddToCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddToCartResponse)
	err := c.cc.Invoke(ctx, CartService_AddToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) RemoveFromCart(ctx context.Context, in *RemoveFromCartRequest, opts ...grpc.CallOption) (*RemoveFromCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveFromCartResponse)
	err := c.cc.Invoke(ctx, CartService_RemoveFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearCartResponse)
	err := c.cc.Invoke(ctx, CartService_ClearCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility.
type CartServiceServer interface {
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error)
	RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error)
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartServiceServer struct{}

func (UnimplementedCartServiceServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) AddToCart(context.Context, *AddToCartRequest) (*AddToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToCart not implemented")
}
func (UnimplementedCartServiceServer) RemoveFromCart(context.Context, *RemoveFromCartRequest) (*RemoveFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromCart not implemented")
}
func (UnimplementedCartServiceServer) ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}
func (UnimplementedCartServiceServer) testEmbeddedByValue()                     {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	// If the following call pancis, it indicates UnimplementedCartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_AddToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddToCart(ctx, req.(*AddToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_RemoveFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).RemoveFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_RemoveFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).RemoveFromCart(ctx, req.(*RemoveFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.mall.service.v1.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "AddToCart",
			Handler:    _CartService_AddToCart_Handler,
		},
		{
			MethodName: "RemoveFromCart",
			Handler:    _CartService_RemoveFromCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _CartService_ClearCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/v1/cart.proto",
}
