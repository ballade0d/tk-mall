package server

import (
	"google.golang.org/grpc"
	"log"
	v1 "mall/api/mall/service/v1"
	"mall/app/user/internal/service"
	"net"
)

func NewGRPCServer(cartService *service.CartService, itemService *service.ItemService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50040")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	v1.RegisterCartServiceServer(grpcServer, cartService)
	v1.RegisterItemServiceServer(grpcServer, itemService)

	log.Println("grpc server start at :50040")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
