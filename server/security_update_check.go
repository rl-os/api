package server

import (
	"github.com/rs/zerolog/log"
)

func (s *Server) DoSecurityUpdateCheck() {
	log.Info().
		Str("job", "DoSecurityUpdateCheck").
		Msg("start security update check")

	// todo
}
