package server

import "github.com/rl-os/api/config"

func (s *Server) onConfigReload(config *config.Config) {
	s.Config = config
}
