package server

import (
	"google.golang.org/grpc"
	"log"
	v1 "mall/api/mall/service/v1"
	"mall/app/payment/internal/config"
	"mall/app/payment/internal/service"

	"net"
)

func NewGRPCServer(conf *config.Config, paymentService *service.PaymentService) *grpc.Server {
	NewRabbitMQServer(conf, paymentService)
	lis, err := net.Listen("tcp", ":50030")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	v1.RegisterPaymentServiceServer(grpcServer, paymentService)

	log.Println("grpc server start at :50030")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return grpcServer
}
