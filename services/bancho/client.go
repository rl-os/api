package bancho

import (
	osu "github.com/deissh/osu-go-client"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/errors"
	"github.com/rs/zerolog/log"
)

func Init(cfg *config.Config) *osu.OsuAPI {
	client, err := osu.WithBasicAuth(
		cfg.Mirror.Bancho.Username,
		cfg.Mirror.Bancho.Password,
	)
	if err != nil {
		log.Warn().
			Err(errors.WithCause("service_bancho", 500, "auth in bancho", err)).
			Send()

		return nil
	}

	return client
}
