package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewCartService, NewItemService, NewOrderService, NewPaymentService, NewUserService)
