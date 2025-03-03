package service

import (
	"context"
	"google.golang.org/grpc"
	pb "mall/api/mall/service/v1"
	"mall/app/gateway/internal/data"
)

type ItemService struct {
	pb.UnimplementedItemServiceServer
	admin pb.ItemServiceClient
	user  pb.ItemServiceClient
}

func NewItemService(data *data.Data) *ItemService {
	adminClient, err := grpc.NewClient(data.GetConfig().Services.AdminService,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(claimsClientInterceptor),
	)
	if err != nil {
		panic(err)
	}
	userClient, err := grpc.NewClient(data.GetConfig().Services.UserService,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(claimsClientInterceptor),
	)
	if err != nil {
		panic(err)
	}
	return &ItemService{
		admin: pb.NewItemServiceClient(adminClient),
		user:  pb.NewItemServiceClient(userClient),
	}
}

func (s *ItemService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	return s.admin.CreateItem(ctx, req)
}

func (s *ItemService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return s.admin.DeleteItem(ctx, req)
}

func (s *ItemService) EditItem(ctx context.Context, req *pb.EditItemRequest) (*pb.EditItemResponse, error) {
	return s.admin.EditItem(ctx, req)
}

func (s *ItemService) AddStock(ctx context.Context, req *pb.AddStockRequest) (*pb.AddStockResponse, error) {
	return s.admin.AddStock(ctx, req)
}

func (s *ItemService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return s.user.GetItem(ctx, req)
}

func (s *ItemService) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	return s.admin.ListItems(ctx, req)
}

func (s *ItemService) SearchItems(ctx context.Context, req *pb.SearchItemsRequest) (*pb.SearchItemsResponse, error) {
	return s.user.SearchItems(ctx, req)
}
