package server

import (
	"google.golang.org/grpc"
	"log"
	v1 "mall/api/mall/service/v1"
	"mall/app/user/internal/service"
	"net"
)

func NewGRPCServer(cartService *service.CartService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	v1.RegisterCartServiceServer(grpcServer, cartService)

	log.Println("grpc server start at :50020")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
