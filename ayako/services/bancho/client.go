package bancho

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/errors"
	"github.com/rs/zerolog/log"
)

func Init(cfg *config.Config) *osu.OsuAPI {
	client, err := osu.WithBasicAuth(
		cfg.Mirror.Bancho.Username,
		cfg.Mirror.Bancho.Password,
	)
	if err != nil {
		log.Fatal().
			Err(errors.WithCause("service_bancho", 500, "auth in bancho", err)).
			Send()
	}

	return client
}
