package server

import (
	"google.golang.org/grpc"
	"log"
	v1 "mall/api/mall/service/v1"
	"mall/app/order/internal/config"
	"mall/app/order/internal/service"

	"net"
)

func NewGRPCServer(conf *config.Config, orderService *service.OrderService) *grpc.Server {
	NewRabbitMQServer(conf, orderService)
	lis, err := net.Listen("tcp", ":50020")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	v1.RegisterOrderServiceServer(grpcServer, orderService)

	log.Println("grpc server start at :50020")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
