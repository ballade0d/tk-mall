package service

import (
	pb "mall/api/mall/service/v1"
)

type ItemService struct {
	pb.UnimplementedItemServiceServer
}

func NewItemService() *ItemService {
	return &ItemService{}
}
