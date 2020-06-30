package services

import (
	"github.com/deissh/rl/ayako/services/bancho"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	bancho.Init,
)
