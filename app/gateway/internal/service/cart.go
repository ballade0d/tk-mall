package service

import (
	"context"
	"google.golang.org/grpc"
	pb "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	client pb.CartServiceClient
}

func NewCartService(data *data.Data) *CartService {
	grpcClient, err := grpc.NewClient(data.GetConfig().Services.UserService, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return &CartService{
		client: pb.NewCartServiceClient(grpcClient),
	}
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	return s.client.GetCart(ctx, req)
}

func (s *CartService) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	return s.client.AddToCart(ctx, req)
}

func (s *CartService) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromCartResponse, error) {
	return s.client.RemoveFromCart(ctx, req)
}

func (s *CartService) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	return s.client.ClearCart(ctx, req)
}
