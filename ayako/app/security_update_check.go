package app

import (
	"github.com/rs/zerolog/log"
)

func (s *App) DoSecurityUpdateCheck() {
	if !s.Config.Service.EnableSecurityFixAlert {
		return
	}

	log.Info().
		Str("job", "DoSecurityUpdateCheck").
		Msg("start security update check")

	// todo
}
