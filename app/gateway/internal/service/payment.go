package service

import (
	"context"
	"google.golang.org/grpc"
	pb "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
	client pb.PaymentServiceClient
}

func NewPaymentService(data *data.Data) *PaymentService {
	grpcClient, err := grpc.NewClient(data.GetConfig().Services.PaymentService,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(claimsClientInterceptor),
	)
	if err != nil {
		panic(err)
	}
	return &PaymentService{
		client: pb.NewPaymentServiceClient(grpcClient),
	}
}

func (s *PaymentService) PayOrder(ctx context.Context, req *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	return s.client.PayOrder(ctx, req)
}
