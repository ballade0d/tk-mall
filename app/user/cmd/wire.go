//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"mall/app/user/internal/config"
	"mall/app/user/internal/data"
	"mall/app/user/internal/server"
	"mall/app/user/internal/service"
)

func wireApp() (*grpc.Server, error) {
	panic(wire.Build(config.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet))
}
