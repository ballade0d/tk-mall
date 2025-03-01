package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/user/internal/data"
)

type ItemService struct {
	pb.UnimplementedItemServiceServer
	itemRepo data.ItemRepo
}

func NewItemService(itemRepo data.ItemRepo) *ItemService {
	return &ItemService{itemRepo: itemRepo}
}

func (s *ItemService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := s.itemRepo.FindItemByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{
		Item: &pb.Item{
			Id:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Stock:       int64(item.Stock),
		},
	}, nil
}

func (s *ItemService) SearchItems(ctx context.Context, req *pb.SearchItemsRequest) (*pb.SearchItemsResponse, error) {
	search, err := s.itemRepo.SearchItems(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.Item, 0, len(search))
	for _, item := range search {
		items = append(items, &pb.Item{
			Id:          int64(item.ID),
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
		})
	}
	return &pb.SearchItemsResponse{Items: items}, nil
}
