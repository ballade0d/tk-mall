package service

import (
	"context"
	pb "mall/api/mall/service/v1"
	"mall/app/callback/internal/data"
)

type CallbackService struct {
	pb.UnimplementedCallbackServiceServer
	paymentRepo *data.PaymentRepo
}

func NewCallbackService(paymentRepo *data.PaymentRepo) *CallbackService {
	return &CallbackService{
		paymentRepo: paymentRepo,
	}
}

func (s *CallbackService) Callback(ctx context.Context, req *pb.CallbackRequest) (*pb.CallbackResponse, error) {
	// TODO:
	err := s.paymentRepo.PaymentCallback(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.CallbackResponse{
		Code:    "200",
		Message: "success",
	}, nil
}
