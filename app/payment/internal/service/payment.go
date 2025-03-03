package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/payment/internal/data"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
	paymentRepo *data.PaymentRepo
}

func NewPaymentService(paymentRepo *data.PaymentRepo) *PaymentService {
	return &PaymentService{
		paymentRepo: paymentRepo,
	}
}

func (s *PaymentService) PayOrder(ctx context.Context, req *pb.PayOrderRequest) (*pb.PayOrderResponse, error) {
	p, err := s.paymentRepo.CreatePayment(ctx, int(req.OrderId))
	if err != nil {
		return nil, err
	}

	// TODO: send payment to payment gateway

	var success = true
	if success {
		err := s.PaySuccessful(ctx, p.ID)
		if err != nil {
			return nil, err
		}
	} else {
		err := s.PayFailed(ctx, p.ID)
		if err != nil {
			return nil, err
		}
	}
	p, err = s.paymentRepo.GetPayment(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	err = s.paymentRepo.NotifyPayment(ctx, p)
	if err != nil {
		return nil, err
	}
	return &pb.PayOrderResponse{
		Payment: &pb.Payment{
			Id:      int64(p.ID),
			OrderId: int64(p.Edges.Order.ID),
			Amount:  p.Amount,
			Status:  string(p.Status),
		},
	}, nil
}

func (s *PaymentService) PaySuccessful(ctx context.Context, id int) error {
	return s.paymentRepo.PaySuccessful(ctx, id)
}

func (s *PaymentService) PayFailed(ctx context.Context, id int) error {
	return s.paymentRepo.PayFailed(ctx, id)
}
