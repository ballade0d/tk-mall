//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"mall/app/gateway/internal/config"
	"mall/app/gateway/internal/data"
	"mall/app/gateway/internal/server"
	"mall/app/gateway/internal/service"
)

func wireApp() (*grpc.Server, error) {
	panic(wire.Build(config.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet))
}
