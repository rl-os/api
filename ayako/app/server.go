package app

import (
	"context"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/middlewares/customerror"
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

const WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN = time.Second * 2

type Server struct {
	// Server instance
	Server *echo.Echo

	Config *config.Config
	// Store contains active implementation
	Store store.Store
	// AppInitializedOnce will allow run pre-start initialization only once
	AppInitializedOnce sync.Once

	// RootRouter is the point for all the external HTTP requests
	RootRouter *echo.Group

	// LocalRouter is the starting point for all
	// the local UNIX socket requests
	// For example internal api for creating new maps,
	// users and etc.
	LocalRouter *echo.Group

	// goroutineCount and goroutineExitSignal use for wait
	// all necessary tasks before shutdown
	goroutineCount      int32
	goroutineExitSignal chan struct{}

	// didFinishListen channel notif about stopping server
	didFinishListen chan struct{}
}

func NewServer(config *config.Config) (*Server, error) {
	srv := echo.New()
	srv.HidePort = true
	srv.HideBanner = true
	srv.HTTPErrorHandler = customerror.CustomHTTPErrorHandler

	s := &Server{
		Server:              srv,
		Config:              config,
		goroutineCount:      0,
		goroutineExitSignal: make(chan struct{}, 1),
	}

	log.Info().Msg("Server is initializing...")

	{
		log.Info().Msg("Setting up config callback")
		s.Config.AutoReloadCallback = s.onConfigReload
	}

	return s, nil
}

func (s *Server) Shutdown() error {
	log.Info().Msg("Stopping server...")

	// stopping all created services
	// external http server
	s.StopHTTPServer()

	// wait all necessary goroutines
	s.WaitForGoroutines()

	// remove callback
	s.Config.AutoReloadCallback = nil

	if s.Store != nil {
		s.Store.Close()
	}

	log.Info().Msg("Server stopped")

	return nil
}

func (s *Server) StartHTTPServer() {
	log.Info().Msg("Starting HTTP server...")
}

func (s *Server) StopHTTPServer() {
	if s.Server == nil {
		return
	}

	log.Info().Msg("Shutdown HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN)
	defer cancel()

	didShutdown := false
	for s.didFinishListen != nil && !didShutdown {
		if err := s.Server.Shutdown(ctx); err != nil {
			log.Warn().
				Err(err).
				Msg("Unable to shutdown HTTP server")
		}

		timer := time.NewTimer(time.Millisecond * 50)
		select {
		case <-s.didFinishListen:
			didShutdown = true
		case <-timer.C:
		}

		timer.Stop()
	}
	_ = s.Server.Close()
	s.Server = nil
}
