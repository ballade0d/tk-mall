package service

import (
	"context"
	v2 "mall/api/mall/service/v1"
	"mall/app/payment/internal/data"
)

type PaymentService struct {
	v2.UnimplementedPaymentServiceServer
	paymentRepo data.PaymentRepo
}

func NewPaymentService(paymentRepo data.PaymentRepo) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentService) PayOrder(ctx context.Context, req *v2.PayOrderRequest) (*v2.PayOrderResponse, error) {
	p, err := s.paymentRepo.PayOrder(ctx, int(req.OrderId))
	if err != nil {
		return nil, err
	}

	// TODO: send payment to payment gateway
	return &v2.PayOrderResponse{
		Payment: &v2.Payment{
			Id:      int64(p.ID),
			OrderId: int64(p.Edges.Order.ID),
			Amount:  p.Amount,
			Status:  string(p.Status),
		},
	}, nil
}
