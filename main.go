package main

import (
	"github.com/deissh/osu-api-server/pkg"
	v2 "github.com/deissh/osu-api-server/pkg/v2"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// setup default mode
	gin.SetMode(gin.ReleaseMode)

	// loading configuration
	config.WithOptions(config.ParseEnv, config.Readonly)
	config.AddDriver(yaml.Driver)
	config.LoadOSEnv([]string{"config"}, true)

	err := config.LoadFiles(config.String("config", "configs/dev.yaml"))
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

		gin.SetMode(gin.DebugMode)
	}

	pkg.InitializeDB()
	pkg.InitializeRedis()

	app := gin.New()
	app.Use(logger.SetLogger())

	api := app.Group("/api")
	{
		// todo: v1.ApplyRoutes(api)
		v2.ApplyRoutes(api)
	}

	err = app.Run(config.String("server.host") + ":" + config.String("server.port"))
	if err != nil {
		log.Fatal().
			Err(err).
			Send()
	}
}
