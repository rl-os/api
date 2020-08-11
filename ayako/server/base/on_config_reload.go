package base

import "github.com/deissh/rl/ayako/config"

func (s *Server) onConfigReload(config *config.Config) {
	s.Config = config
}
