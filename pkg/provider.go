package pkg

import (
	"github.com/google/wire"
	"github.com/rl-os/api/pkg/bancho"
	"github.com/rl-os/api/pkg/config"
	"github.com/rl-os/api/pkg/log"
	"github.com/rl-os/api/pkg/validator"
)

var ProviderSet = wire.NewSet(
	validator.ProviderSet,
	log.ProviderSet,
	config.ProviderSet,
	bancho.ProviderSet,
)
