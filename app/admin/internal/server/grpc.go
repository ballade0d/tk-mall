package server

import (
	"google.golang.org/grpc"
	"log"
	"mall/api/mall/service/v1"
	"mall/app/admin/internal/service"
	"net"
)

func NewGRPCServer(service *service.ItemService) *grpc.Server {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	v1.RegisterItemServiceServer(grpcServer, service)

	go func() {
		log.Println("grpc server start at :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	return grpcServer
}
