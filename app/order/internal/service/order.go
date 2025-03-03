package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/order/internal/data"
	"mall/ent/order"
	"mall/pkg/util"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	itemRepo  *data.ItemRepo
	orderRepo *data.OrderRepo
}

func NewOrderService(itemRepo *data.ItemRepo, orderRepo *data.OrderRepo) *OrderService {
	return &OrderService{
		itemRepo:  itemRepo,
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	id := ctx.Value("claims").(util.Claims).UserId
	err := s.itemRepo.CheckAndReduceStock(ctx, req.Items)
	if err != nil {
		return nil, err
	}
	o, err := s.orderRepo.CreateOrder(ctx, id, req.Address, req.Items)
	if err != nil {
		return nil, err
	}
	items := make([]*pb.OrderItem, 0)
	for _, oi := range o.Edges.Items {
		i, err := oi.QueryItem().Only(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &pb.OrderItem{
			Id:        int64(oi.ID),
			ProductId: int64(i.ID),
			Quantity:  int64(oi.Quantity),
		})
	}
	return &pb.CreateOrderResponse{
		Order: &pb.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		},
	}, nil
}

func (s *OrderService) GetOrderList(ctx context.Context, req *pb.GetOrderListRequest) (*pb.GetOrderListResponse, error) {
	list, err := s.orderRepo.GetOrderList(ctx, int(req.Size), int(req.Page))
	if err != nil {
		return nil, err
	}
	orders := make([]*pb.Order, 0)
	for _, o := range list {
		items := make([]*pb.OrderItem, 0)
		for _, item := range o.Edges.Items {
			i, err := item.QueryItem().Only(ctx)
			if err != nil {
				return nil, err
			}
			items = append(items, &pb.OrderItem{
				Id:        int64(item.ID),
				ProductId: int64(i.ID),
				Quantity:  int64(item.Quantity),
			})
		}
		orders = append(orders, &pb.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		})
	}
	return &pb.GetOrderListResponse{
		Orders: orders,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	o, err := s.orderRepo.GetOrder(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	items := make([]*pb.OrderItem, 0)
	for _, item := range o.Edges.Items {
		i, err := item.QueryItem().Only(ctx)
		if err != nil {
			return nil, err
		}
		items = append(items, &pb.OrderItem{
			Id:        int64(item.ID),
			ProductId: int64(i.ID),
			Quantity:  int64(item.Quantity),
		})
	}
	return &pb.GetOrderResponse{
		Order: &pb.Order{
			Id:      int64(o.ID),
			UserId:  int64(o.Edges.User.ID),
			Status:  string(o.Status),
			Address: o.Address,
			Items:   items,
		},
	}, nil
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, id int, status order.Status) error {
	return s.orderRepo.UpdateOrderStatus(ctx, id, status)
}
