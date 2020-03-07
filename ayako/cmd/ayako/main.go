//go:generate wire
//+build wireinject

package main

import (
	"github.com/deissh/osu-lazer/ayako/app"
	"github.com/deissh/osu-lazer/ayako/config"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if os.Getenv("DEBUG") == "true" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: false,
			},
		).With().Caller().Logger()
	}

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Msg("Starting API")

	log.Debug().Msg("Start initialize dependencies")

	app := Injector()

	log.Debug().Msg("Initialize dependencies successful done")

	if err := app.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func Injector() *app.App {
	wire.Build(
		config.Init,
		store.Init,
		app.ProviderSet,
	)

	return nil
}
