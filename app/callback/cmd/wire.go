//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	"mall/app/callback/internal/config"
	"mall/app/callback/internal/data"
	"mall/app/callback/internal/server"
	"mall/app/callback/internal/service"
)

func wireApp() (*grpc.Server, error) {
	panic(wire.Build(config.ProviderSet, data.ProviderSet, service.ProviderSet, server.ProviderSet))
}
