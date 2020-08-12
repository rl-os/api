package server

import (
	"context"
	"github.com/deissh/rl/ayako/api"
	"github.com/deissh/rl/ayako/app"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/middlewares/customerror"
	"github.com/deissh/rl/ayako/middlewares/customlogger"
	"github.com/deissh/rl/ayako/middlewares/permission"
	"github.com/deissh/rl/ayako/middlewares/reqest_context"
	"github.com/deissh/rl/ayako/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"sync"
	"time"
)

const WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN = time.Second * 2

type Server struct {
	App *app.App

	// Global context
	Context context.Context

	// Server instance
	Server *echo.Echo

	Config *config.Config
	// AppInitializedOnce will allow run pre-start initialization only once
	AppInitializedOnce sync.Once

	// GetRootRouter is the point for all the external HTTP requests
	RootRouter *echo.Group

	// LocalRouter is the starting point for all
	// the local UNIX socket requests
	// For example internal api for creating new maps,
	// users and etc.
	LocalRouter *echo.Group

	// GoroutineCount and GoroutineExitSignal use for wait
	// all necessary tasks before shutdown
	GoroutineCount      int32
	GoroutineExitSignal chan struct{}

	// DidFinishListen channel notif about stopping server
	DidFinishListen chan struct{}
}

// NewServer create and fill server configuration
func NewServer(
	config *config.Config,
	app *app.App,
) *Server {
	srv := &Server{
		App:                 app,
		Context:             context.Background(),
		Config:              config,
		GoroutineCount:      0,
		GoroutineExitSignal: make(chan struct{}, 1),
	}

	log.Info().Msg("Server is initializing...")

	{
		log.Info().Msg("Setting up config callback")
		srv.Config.AutoReloadCallback = srv.onConfigReload
	}

	return srv
}

func (s *Server) Start() error {
	s.StartHTTPServer()

	return nil
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

	if s.App != nil {
		s.GetStore().Close()
	}

	log.Info().Msg("Server stopped")

	return nil
}

func (s *Server) StartHTTPServer() {
	log.Info().Msg("Starting HTTP server...")

	srv := echo.New()
	srv.HidePort = true
	srv.HideBanner = true
	srv.HTTPErrorHandler = customerror.CustomHTTPErrorHandler

	srv.Use(
		middleware.RequestID(),
		permission.GlobalMiddleware(s.App),
		reqest_context.GlobalMiddleware(s.App),
		customlogger.Middleware(),
	)

	s.Server = srv
	s.RootRouter = srv.Group("")

	api.New(s.App, s.RootRouter)

	// log all routes
	for _, route := range s.Server.Routes() {
		log.Debug().
			Str("name", route.Name).
			Str("method", route.Method).
			Str("path", route.Path).
			Msg("Route loaded")
	}

	addr := s.Config.Server.Host + ":" + s.Config.Server.Port

	s.DidFinishListen = make(chan struct{})
	go func() {
		err := s.Server.Start(addr)
		if err != nil {
			log.Error().
				Err(err).
				Msg("Error starting server")
		}

		close(s.DidFinishListen)
	}()
}

func (s *Server) StopHTTPServer() {
	if s.Server == nil {
		return
	}

	log.Info().Msg("Shutdown HTTP server")

	ctx, cancel := context.WithTimeout(context.Background(), WAIT_FOR_CONNECTIONS_BEFORE_SHUTDOWN)
	defer cancel()

	didShutdown := false
	for s.DidFinishListen != nil && !didShutdown {
		if err := s.Server.Shutdown(ctx); err != nil {
			log.Warn().
				Err(err).
				Msg("Unable to shutdown HTTP server")
		}

		timer := time.NewTimer(time.Millisecond * 50)
		select {
		case <-s.DidFinishListen:
			didShutdown = true
		case <-timer.C:
		}

		timer.Stop()
	}
	_ = s.Server.Close()
	s.Server = nil
}

func (s *Server) GetStore() store.Store      { return s.GetStore() }
func (s *Server) GetRootRouter() *echo.Group { return s.RootRouter }

func runSecurityJob(s *Server) {
	CreateRecurringTask("Security", func() {
		s.DoSecurityUpdateCheck()
	}, time.Hour*6)
}

func runUpdateCheck(s *Server) {
	CreateRecurringTask("UpdateCheck", func() {
		s.DoBeatmapSetUpdate()
	}, time.Minute*30)
}

func runSearchNew(s *Server) {
	CreateRecurringTask("SearchNew", func() {
		s.DoBeatmapSetSearchNew()
	}, time.Hour)
}
