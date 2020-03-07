package app

import (
	"context"
	"github.com/deissh/osu-lazer/ayako/api"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/gookit/config/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Store *store.Store
	Echo  *echo.Echo

	goroutineCount      int32
	goroutineExitSignal chan struct{}
}

// NewApp with DI
// expect store.Store
func NewApp(store store.Store) *App {
	app := echo.New()
	app.HidePort = true
	app.HideBanner = true

	app.Use(middleware.RequestID())

	api.New(store, app.Group("/v2"))

	return &App{
		Store: nil,
		Echo:  app,
	}
}

// Start http server and setup graceful shutdown
func (s *App) Start() error {
	for _, route := range s.Echo.Routes() {
		log.Debug().
			Str("method", route.Method).
			Str("path", route.Path).
			Msg("Route loaded")
	}

	log.Info().Msg("Starting App...")

	addr := config.String("host") + ":" + config.String("port")

	// Graceful start and stop HTTP server
	go func() {
		err := s.Echo.Start(addr)
		if err != nil {
			log.Error().
				Err(err).
				Msg("shutting down the server")
		}
	}()

	log.Info().Msg("App started on http://" + addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.Echo.Shutdown(ctx)
}
