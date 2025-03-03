package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"mall/api/mall/service/v1"
	"mall/app/gateway/internal/service"
	"mall/pkg/util"
	"net"
	"slices"
)

var permissions = map[string][]string{
	"/api.mall.service.v1.ItemService/CreateItem":     {"admin"},
	"/api.mall.service.v1.ItemService/DeleteItem":     {"admin"},
	"/api.mall.service.v1.ItemService/EditItem":       {"admin"},
	"/api.mall.service.v1.ItemService/AddStock":       {"admin"},
	"/api.mall.service.v1.ItemService/ListItems":      {"admin"},
	"/api.mall.service.v1.CartService/GetCart":        {"user", "admin"},
	"/api.mall.service.v1.CartService/AddToCart":      {"user", "admin"},
	"/api.mall.service.v1.CartService/RemoveFromCart": {"user", "admin"},
	"/api.mall.service.v1.CartService/ClearCart":      {"user", "admin"},
	"/api.mall.service.v1.ItemService/GetItem":        {"user", "admin"},
	"/api.mall.service.v1.ItemService/SearchItems":    {"user", "admin"},
	"/api.mall.service.v1.OrderService/CreateOrder":   {"user", "admin"},
	"/api.mall.service.v1.OrderService/GetOrderList":  {"user", "admin"},
	"/api.mall.service.v1.OrderService/GetOrder":      {"user", "admin"},
	"/api.mall.service.v1.PaymentService/PayOrder":    {"user", "admin"},
	"/api.mall.service.v1.UserService/GetUser":        {"user", "admin"},
}

func jwtAuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if permissions[info.FullMethod] == nil {
		return handler(ctx, req)
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, fmt.Errorf("missing authorization token")
	}

	tokenStr := authHeader[0]
	if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
		tokenStr = tokenStr[7:]
	}

	claims, err := util.VerifyJWT(tokenStr)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	if !slices.Contains(permissions[info.FullMethod], claims.Role) {
		return nil, fmt.Errorf("permission denied")
	}

	ctx = context.WithValue(ctx, "claims", *claims)

	return handler(ctx, req)
}

func NewGRPCServer(cartService *service.CartService, itemService *service.ItemService, orderService *service.OrderService, paymentService *service.PaymentService, userService *service.UserService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(jwtAuthInterceptor))

	v1.RegisterCartServiceServer(grpcServer, cartService)
	v1.RegisterItemServiceServer(grpcServer, itemService)
	v1.RegisterOrderServiceServer(grpcServer, orderService)
	v1.RegisterPaymentServiceServer(grpcServer, paymentService)
	v1.RegisterUserServiceServer(grpcServer, userService)

	go func() {
		log.Println("grpc server start at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	NewHTTPServer()
	return grpcServer
}
