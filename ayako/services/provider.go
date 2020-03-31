package services

import (
	"github.com/deissh/osu-lazer/ayako/services/bancho"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	bancho.Init,
)
