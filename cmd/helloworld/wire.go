//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/data"
	"helloworld/internal/server"
	"helloworld/internal/service"
	"helloworld/pkg/registry"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry_Nacos, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, registry.ProviderSet, newApp))
}
