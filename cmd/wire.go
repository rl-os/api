//go:generate wire .
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rl-os/api/api"
	"github.com/rl-os/api/app"
	oldCfg "github.com/rl-os/api/config"
	"github.com/rl-os/api/pkg"
	"github.com/rl-os/api/pkg/transports/http" // remove: use universal server starter (http/grpc/rmq/etc)
	sql "github.com/rl-os/api/store/gorm"
)

var providerSet = wire.NewSet(
	pkg.ProviderSet,
	api.ProviderSet,
	app.ProviderSet,

	sql.Init,
	oldCfg.Init,
)

func Injector(configPath string) (*http.Server, error) {
	panic(wire.Build(providerSet))
}
