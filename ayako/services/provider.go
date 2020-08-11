package services

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/rl/ayako/services/bancho"
	"github.com/google/wire"
)

// Services struct contains all enabled services
type Services struct {
	Bancho *osu.OsuAPI
}

// NewServices returns struct with all services
func NewServices(bancho *osu.OsuAPI) *Services {
	return &Services{
		bancho,
	}
}

var ProviderSet = wire.NewSet(
	NewServices,

	bancho.Init,
)
