package service

import (
	"context"
	v2 "mall/api/mall/service/v1"
	"mall/app/user/internal/data"
)

type CartService struct {
	v2.UnimplementedCartServiceServer
	cartRepo data.CartRepo
}

func NewCartService(cartRepo data.CartRepo) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

func (s *CartService) GetCart(ctx context.Context, req *v2.GetCartRequest) (*v2.GetCartResponse, error) {
	c, err := s.cartRepo.GetCart(ctx)
	if err != nil {
		return nil, err
	}
	items, err := c.QueryItems().All(ctx)
	if err != nil {
		return nil, err
	}
	cartItems := make([]*v2.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &v2.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &v2.GetCartResponse{
		Cart: &v2.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) AddToCart(ctx context.Context, req *v2.AddToCartRequest) (*v2.AddToCartResponse, error) {
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
	cartItems := make([]*v2.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &v2.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &v2.AddToCartResponse{
		Cart: &v2.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) RemoveFromCart(ctx context.Context, req *v2.RemoveFromCartRequest) (*v2.RemoveFromCartResponse, error) {
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
	cartItems := make([]*v2.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &v2.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &v2.RemoveFromCartResponse{
		Cart: &v2.Cart{
			Items: cartItems,
		},
	}, nil
}

func (s *CartService) ClearCart(ctx context.Context, req *v2.ClearCartRequest) (*v2.ClearCartResponse, error) {
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
	cartItems := make([]*v2.CartItem, 0)
	for _, item := range items {
		cartItems = append(cartItems, &v2.CartItem{
			ItemId:   int64(item.ID),
			Quantity: int64(item.Quantity),
		})
	}
	return &v2.ClearCartResponse{
		Cart: &v2.Cart{
			Items: cartItems,
		},
	}, nil
}
