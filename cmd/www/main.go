package main

import (
	"github.com/deissh/osu-api-server/pkg/middlewares/customerror"
	"github.com/deissh/osu-api-server/pkg/middlewares/customlogger"
	"github.com/deissh/osu-api-server/pkg/oauth"
	"github.com/deissh/osu-api-server/pkg/v1"
	"github.com/deissh/osu-api-server/pkg/v2"
	"github.com/deissh/osu-api-server/pkg"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"context"
	"os"
	"os/signal"
	"time"
)

func main() {
	// loading configuration
	config.WithOptions(config.ParseEnv, config.Readonly)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("config.yaml")
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

	pkg.InitializeDB()
	pkg.InitializeRedis()

	// Seting up Echo
	app := echo.New()
	app.HideBanner = true
	app.HTTPErrorHandler = customerror.CustomHTTPErrorHandler

	app.Use(middleware.RequestID())
	// app.Use(middleware.Recover())
	app.Use(customlogger.Middleware())

	if config.Bool("server.cors.enable") {
		app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: config.Strings("server.cors.allow_origins"),
			AllowHeaders: config.Strings("server.cors.allow_headers"),
		}))
	}

	// Mount routes
	root := app.Group("")
	{
		oauth.ApplyRoutes(root)
	}
	api := app.Group("/api")
	{
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
