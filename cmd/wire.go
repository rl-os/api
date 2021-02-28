//go:generate wire
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rl-os/api/pkg/config"
	"github.com/rl-os/api/pkg/log"
	"github.com/rl-os/api/pkg/transports/http"
)

var providerSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,

	http.ProviderSet,
)

func Injector(configPath string) (*http.Server, error) {
	panic(wire.Build(providerSet))
}
