package service

import (
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
