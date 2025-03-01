package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/admin/internal/data"
)

type ItemService struct {
	pb.UnimplementedItemServiceServer
	itemRepo data.ItemRepo
}

func NewItemService(itemRepo data.ItemRepo) *ItemService {
	return &ItemService{itemRepo: itemRepo}
}

func (s *ItemService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	item, err := s.itemRepo.CreateItem(ctx, req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.CreateItemResponse{
		Item: &pb.Item{
			Id:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       0,
		},
	}, nil
}

func (s *ItemService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	err := s.itemRepo.DeleteItem(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteItemResponse{}, nil
}

func (s *ItemService) EditItem(ctx context.Context, req *pb.EditItemRequest) (*pb.EditItemResponse, error) {
	item, err := s.itemRepo.EditItem(ctx, int(req.Id), req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &pb.EditItemResponse{
		Item: &pb.Item{
			Id:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       0,
		},
	}, nil
}

func (s *ItemService) AddStock(ctx context.Context, req *pb.AddStockRequest) (*pb.AddStockResponse, error) {
	// TODO: Lock item
	item, err := s.itemRepo.AddStock(ctx, int(req.Id), int(req.Stock))
	if err != nil {
		return nil, err
	}
	return &pb.AddStockResponse{
		Item: &pb.Item{
			Id:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       int64(item.Stock),
		},
	}, nil
}
