package bancho

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/osu-lazer/ayako/config"
)

func Init(cfg *config.Config) *osu.OsuAPI {
	client := osu.WithAccessToken(
		cfg.Mirror.Bancho.AccessToken,
		cfg.Mirror.Bancho.RefreshToken,
	)
	return &client
}
