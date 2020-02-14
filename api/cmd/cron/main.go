package main

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

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

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Msg("Starting cron")

	log.Info().
		Msg("Loaded configuration and logger")
	log.Info().
		Msg("Start initialize database and redis")

	pkg.InitializeDB()
	pkg.InitializeRedis()

	log.Info().
		Msg("Initialize database and redis successful done")

	log.Info().
		Msg("Start initialize cron")

	c := cron.New()

	// every 30 mins
	_, _ = c.AddFunc("* * * * *", ExpireSupporter)

	c.Run()
}
