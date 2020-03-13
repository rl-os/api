//go:generate wire
//+build wireinject

package main

import (
	"os"

	"github.com/deissh/osu-lazer/ayako/app"
	"github.com/deissh/osu-lazer/ayako/config"
	"github.com/deissh/osu-lazer/ayako/store/sql"
	"github.com/google/wire"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     os.Stderr,
			NoColor: false,
		},
	).With().Caller().Logger()

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Send()

	log.Debug().Msg("Start initialize dependencies")

	app := Injector("config.yaml")

	log.Debug().Msg("Initialize dependencies successful done")

	if err := app.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func Injector(configPath string) *app.App {
	wire.Build(
		config.Init,
		sql.Init,
		app.ProviderSet,
	)

	return nil
}
