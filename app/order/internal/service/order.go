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
	orderRepo *data.OrderRepo
}

func NewOrderService(orderRepo *data.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	usr := ctx.Value("claims").(*util.Claims).UserId
	o, err := s.orderRepo.CreateOrder(ctx, usr, req.Address, req.Items)
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
	// TODO: Lock item and reduce stock
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
