//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"mall/app/order/internal/config"
	"mall/app/order/internal/data"
	"mall/app/order/internal/server"
	"mall/app/order/internal/service"
)

func wireApp() (*grpc.Server, error) {
	panic(wire.Build(config.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet))
}
