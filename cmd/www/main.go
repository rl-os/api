package main

import (
	"github.com/deissh/osu-api-server/pkg"
	v1 "github.com/deissh/osu-api-server/pkg/v1"
	v2 "github.com/deissh/osu-api-server/pkg/v2"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	pkg.InitializeDB()
	pkg.InitializeRedis()

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.RequestID())
	app.Use(middleware.Recover())

	api := app.Group("/api")
	{
		v1.ApplyRoutes(api)
		v2.ApplyRoutes(api)
	}

	err = app.Start(config.String("server.host") + ":" + config.String("server.port"))
	if err != nil {
		log.Fatal().
			Err(err).
			Send()
	}
}
