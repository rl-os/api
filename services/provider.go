package services

import (
	"github.com/google/wire"
	"github.com/rl-os/api/services/bancho"
	"github.com/rl-os/api/services/config"
	"github.com/rl-os/api/services/log"
	"github.com/rl-os/api/services/validator"
)

var ProviderSet = wire.NewSet(
	validator.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	bancho.ProviderSet,
)
