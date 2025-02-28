package service

import (
	"context"
	"google.golang.org/grpc"
	pb "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
)

type ItemService struct {
	pb.UnimplementedItemServiceServer
	data data.Data
}

func NewItemService(data *data.Data) *ItemService {
	return &ItemService{
		data: *data,
	}
}

func (s *ItemService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {

	host := s.data.GetConfig().Services.AdminService
	grpcClient, err := grpc.NewClient(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	response, err := pb.NewItemServiceClient(grpcClient).CreateItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
