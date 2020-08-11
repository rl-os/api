//go:generate wire
//+build wireinject

package main

import (
	"github.com/deissh/rl/ayako/app"
	"github.com/deissh/rl/ayako/config"
	"github.com/deissh/rl/ayako/server"
	"github.com/deissh/rl/ayako/services"
	"github.com/deissh/rl/ayako/store/sql"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

func main() {
	setupLogger()

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Send()

	log.Debug().Msg("Start initialize dependencies")

	srv := Injector("config.yaml")

	log.Debug().Msg("Initialize dependencies successful done")

	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func setupLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if os.Getenv("env") != "production" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		},
	).With().Caller().Logger()
}

func Injector(configPath string) *server.Server {
	wire.Build(
		config.Init,
		sql.Init,
		app.ProviderSet,
		server.ProviderSet,
		services.ProviderSet,
	)

	return nil
}
