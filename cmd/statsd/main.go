package main

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/common/datadog"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
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

	if !config.Bool("server.datadog.enable") {
		log.Fatal().Msg("Datadog disabled")
	}

	log.Info().
		Msg("Loaded configuration and logger")
	log.Info().
		Msg("Start initialize database and redis")

	pkg.InitializeDB()
	pkg.InitializeRedis()

	log.Info().
		Msg("Initialize database and redis successful done")

	log.Info().
		Msg("Start initialize datadog")

	datadog.InitializeClient()
	datadog.SetPrefix(config.String("server.datadog.prefix", "LAZER_"))
	datadog.AddTag("debug", config.String("debug"))

	log.Info().
		Msg("Initialize datadog successful done")

	log.Info().
		Msg("Creating datadog tasks")

	datadog.RunGaugeTask(
		"user_online",
		time.Minute,
		datadog.Tags{},
		GetUsersOnline,
	)
	datadog.RunGaugeTask(
		"users",
		time.Hour,
		datadog.Tags{},
		GetActiveUsers,
	)

	datadog.Start()
}
