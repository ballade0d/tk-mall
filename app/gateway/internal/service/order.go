package service

import (
	"context"
	"google.golang.org/grpc"
	pb "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	client pb.OrderServiceClient
}

func NewOrderService(data *data.Data) *OrderService {
	grpcClient, err := grpc.NewClient(data.GetConfig().Services.OrderService,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(claimsClientInterceptor),
	)
	if err != nil {
		panic(err)
	}
	return &OrderService{
		client: pb.NewOrderServiceClient(grpcClient),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return s.client.CreateOrder(ctx, req)
}

func (s *OrderService) GetOrderList(ctx context.Context, req *pb.GetOrderListRequest) (*pb.GetOrderListResponse, error) {
	return s.client.GetOrderList(ctx, req)
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return s.client.GetOrder(ctx, req)
}
