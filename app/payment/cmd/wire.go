//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"mall/app/payment/internal/config"
	"mall/app/payment/internal/data"
	"mall/app/payment/internal/server"
	"mall/app/payment/internal/service"
)

func wireApp() (*grpc.Server, error) {
	panic(wire.Build(config.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet))
}
