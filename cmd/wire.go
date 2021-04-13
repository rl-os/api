//go:generate wire
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rl-os/api/api"
	"github.com/rl-os/api/app"
	oldCfg "github.com/rl-os/api/config"
	"github.com/rl-os/api/pkg/config"
	"github.com/rl-os/api/pkg/log"
	"github.com/rl-os/api/pkg/transports/http"
	"github.com/rl-os/api/services"
	sql "github.com/rl-os/api/store/gorm"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	api.ProviderSet,
	services.ProviderSet,
	app.ProviderSet,

	sql.Init,
	oldCfg.Init,
)

func Injector(configPath string) (*http.Server, error) {
	panic(wire.Build(providerSet))
}
