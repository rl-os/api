package jobs

import (
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/config"
	"github.com/rs/zerolog/log"
)

type JobServer struct {
	App    *app.App
	Config *config.Config
}

func NewJobServer(
	app *app.App,
	config *config.Config,
) *JobServer {
	return &JobServer{
		app,
		config,
	}
}

// Start, setup workers and
// also enable scheduler with prebuild jobs
func (s *JobServer) Start() error {
	log.Info().Msg("Starting job server...")

	return nil
}

// Shutdown server before global graceful shutdown
func (s *JobServer) Shutdown() error {
	log.Info().Msg("Stopping job server...")

	return nil
}
