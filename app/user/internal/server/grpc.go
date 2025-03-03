package server

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	v1 "mall/api/mall/service/v1"
	"mall/app/user/internal/service"
	"mall/pkg/util"
	"net"
)

func claimsServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	claimsData := md["x-claims"]
	if len(claimsData) == 0 {
		return nil, fmt.Errorf("missing claims data")
	}

	var claims util.Claims
	if err := json.Unmarshal([]byte(claimsData[0]), &claims); err != nil {
		return nil, fmt.Errorf("failed to decode claims: %v", err)
	}

	newCtx := context.WithValue(ctx, "claims", claims)

	return handler(newCtx, req)
}

func NewGRPCServer(cartService *service.CartService, itemService *service.ItemService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(claimsServerInterceptor))

	v1.RegisterCartServiceServer(grpcServer, cartService)
	v1.RegisterItemServiceServer(grpcServer, itemService)

	log.Println("grpc server start at :50040")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
