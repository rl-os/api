package main

import (
	// "github.com/deissh/osu-api-server/pkg"
	oauth "github.com/deissh/osu-api-server/pkg/oauth"
	v1 "github.com/deissh/osu-api-server/pkg/v1"
	v2 "github.com/deissh/osu-api-server/pkg/v2"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	// loading configuration
	config.WithOptions(config.ParseEnv, config.Readonly)
	config.AddDriver(yaml.Driver)
	config.LoadOSEnv([]string{"config"}, true)

	err := config.LoadFiles(config.String("config", "config.yaml"))
	if err != nil {
		panic(err)
	}

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if config.Bool("debug", false) {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: false,
			},
		).With().Caller().Logger()

	}

	// pkg.InitializeDB()
	// pkg.InitializeRedis()

	// Seting up Echo
	app := echo.New()
	app.HideBanner = true

	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: config.String("server.http_log_format") + "\n",
	}))
	app.Use(middleware.RequestID())
	app.Use(middleware.Recover())

	if config.Bool("server.cors.enable") {
		app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Strings("server.cors.allow_origins"),
			AllowHeaders: config.Strings("server.cors.allow_headers"),
		}))
	}

	// Mount routes
	api := app.Group("/api")
	{
		oauth.ApplyRoutes(api)
		v1.ApplyRoutes(api)
		v2.ApplyRoutes(api)
	}

	// Graceful start and stop HTTP server
	go func() {
		if err := app.Start(config.String("server.host") + ":" + config.String("server.port")); err != nil {
			log.Info().Msg("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatal().
			Err(err).
			Send()
	}
}
