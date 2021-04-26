//go:generate wire .
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rl-os/api/api"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/repository/gorm"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/services/transports"
	"github.com/rl-os/api/services/transports/http"
)

var providerSet = wire.NewSet(
	services.ProviderSet,
	api.ProviderSet,
	app.ProviderSet,
	http.ProviderSet,
	gorm.ProviderSet,
)

func Injector(configPath string) (transports.Server, error) {
	panic(wire.Build(providerSet))
}
