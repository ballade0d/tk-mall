package service

import (
	"context"
	v2 "mall/api/mall/service/v1"
	"mall/app/order/internal/data"
	"mall/pkg/util"
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
	usr := ctx.Value("claims").(*util.Claims).UserId
	o, err := s.orderRepo.CreateOrder(ctx, usr, req.Address, req.Items)
	if err != nil {
		return nil, err
	}
	items := make([]*v2.OrderItem, 0)
	for _, item := range o.Edges.Items {
		i, err := item.QueryItem().Only(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &v2.OrderItem{
			Id:        int64(item.ID),
			ProductId: int64(i.ID),
			Quantity:  int64(item.Quantity),
		})
	}
	// TODO: Reduce stock
	return &v2.CreateOrderResponse{
		Order: &v2.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		},
	}, nil
}

func (s *OrderService) GetOrderList(ctx context.Context, req *v2.GetOrderListRequest) (*v2.GetOrderListResponse, error) {
	list, err := s.orderRepo.GetOrderList(ctx, int(req.Size), int(req.Page))
	if err != nil {
		return nil, err
	}
	orders := make([]*v2.Order, 0)
	for _, o := range list {
		items := make([]*v2.OrderItem, 0)
		for _, item := range o.Edges.Items {
			i, err := item.QueryItem().Only(ctx)
			if err != nil {
				return nil, err
			}
			items = append(items, &v2.OrderItem{
				Id:        int64(item.ID),
				ProductId: int64(i.ID),
				Quantity:  int64(item.Quantity),
			})
		}
		orders = append(orders, &v2.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		})
	}
	return &v2.GetOrderListResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *v2.GetOrderRequest) (*v2.GetOrderResponse, error) {
	o, err := s.orderRepo.GetOrder(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	items := make([]*v2.OrderItem, 0)
	for _, item := range o.Edges.Items {
		i, err := item.QueryItem().Only(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &v2.OrderItem{
			Id:        int64(item.ID),
			ProductId: int64(i.ID),
			Quantity:  int64(item.Quantity),
		})
	}
	return &v2.GetOrderResponse{
		Order: &v2.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		},
	}, nil
}
