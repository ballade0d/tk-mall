package service

import (
	"context"
	v2 "mall/api/mall/service/v1"
	"mall/app/order/internal/data"
)

type OrderService struct {
	v2.UnimplementedOrderServiceServer
	orderRepo data.OrderRepo
}

func NewOrderService(orderRepo data.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *v2.CreateOrderRequest) (*v2.CreateOrderResponse, error) {
	return nil, nil
}
