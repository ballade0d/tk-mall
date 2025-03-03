package service

import (
	"context"
	"encoding/json"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"mall/pkg/util"
)

var ProviderSet = wire.NewSet(NewCartService, NewItemService, NewOrderService, NewPaymentService, NewUserService)

func claimsClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	claims := ctx.Value("claims").(util.Claims)

	claimsData, err := json.Marshal(claims)
	if err != nil {
		return err
	}

	md := metadata.Pairs("x-claims", string(claimsData))
	newCtx := metadata.NewOutgoingContext(ctx, md)

	return invoker(newCtx, method, req, reply, cc, opts...)
}
