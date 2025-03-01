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

// /api.mall.service.v1.UserService/Register
var routesNeedAuth = []string{
	"/api.mall.service.v1.UserService/GetUser",
}

func jwtAuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	if !slices.Contains(routesNeedAuth, info.FullMethod) {
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

	ctx = context.WithValue(ctx, "claims", claims)

	return handler(ctx, req)
}

func NewGRPCServer(userService *service.UserService, itemService *service.ItemService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(jwtAuthInterceptor))

	v1.RegisterUserServiceServer(grpcServer, userService)
	v1.RegisterItemServiceServer(grpcServer, itemService)

	go func() {
		log.Println("grpc server start at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	NewHTTPServer()
	return grpcServer
}
