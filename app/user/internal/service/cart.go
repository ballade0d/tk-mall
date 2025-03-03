package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/user/internal/data"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	cartRepo *data.CartRepo
}

func NewCartService(cartRepo *data.CartRepo) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	c, err := s.cartRepo.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	items, err := c.QueryItems().All(ctx)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*pb.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &pb.GetCartResponse{
		Cart: &pb.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	_, err := s.cartRepo.AddToCart(ctx, int(req.ItemId), int(req.Quantity))
	if err != nil {
		return nil, err
	}
	c, err := s.cartRepo.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	items, err := c.QueryItems().All(ctx)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*pb.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &pb.AddToCartResponse{
		Cart: &pb.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) RemoveFromCart(ctx context.Context, req *pb.RemoveFromCartRequest) (*pb.RemoveFromCartResponse, error) {
	err := s.cartRepo.RemoveFromCart(ctx, int(req.ItemId))
	if err != nil {
		return nil, err
	}
	c, err := s.cartRepo.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	items, err := c.QueryItems().All(ctx)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*pb.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &pb.RemoveFromCartResponse{
		Cart: &pb.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) ClearCart(ctx context.Context, req *pb.ClearCartRequest) (*pb.ClearCartResponse, error) {
	err := s.cartRepo.ClearCart(ctx)
	if err != nil {
		return nil, err
	}
	c, err := s.cartRepo.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	items, err := c.QueryItems().All(ctx)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*pb.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &pb.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &pb.ClearCartResponse{
		Cart: &pb.Cart{
			Items: cartItems,
		},
	}, nil
}
