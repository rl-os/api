package main

import (
	"github.com/deissh/osu-api-server/api"
	"os"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if err := godotenv.Load(); err != nil {
		log.Panic().Err(err)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	app := gin.New()
	app.Use(logger.SetLogger())

	api.ApplyRoutes(app)

	if err := app.Run(":" + os.Getenv("PORT")); err != nil {
		log.Err(err)
	}
}
