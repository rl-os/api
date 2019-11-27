package main

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/database"
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

	_, _ = database.Initialize()

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	)

	app := gin.New()
	app.Use(logger.SetLogger())

	pkg.ApplyRoutes(app)

	if err := app.Run(":" + os.Getenv("PORT")); err != nil {
		log.Err(err)
	}
}
