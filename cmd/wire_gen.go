// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"flag"
	"github.com/rl-os/api/app"
	"github.com/rl-os/api/config"
	"github.com/rl-os/api/server"
	"github.com/rl-os/api/services"
	"github.com/rl-os/api/services/bancho"
	"github.com/rl-os/api/store/gorm"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	log2 "log"
	"os"
	"os/signal"
	"time"
)

// Injectors from main.go:

func Injector(configPath string) *server.Server {
	configConfig := config.Init(configPath)
	osuAPI := bancho.Init(configConfig)
	servicesServices := services.NewServices(osuAPI)
	store := sql.Init(configConfig, servicesServices)
	appApp := app.NewApp(store, configConfig, servicesServices)
	serverServer := server.NewServer(configConfig, appApp)
	return serverServer
}

// main.go:

var Version string

var Commit string

var Branch string

var BuildTimestamp string

func main() {
	logLevel := flag.String("log", "info", "sets log level")
	configFile := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	setupLogger(*logLevel)
	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Send()
	log.Debug().Msg("Start initialize dependencies")

	srv := Injector(*configFile)
	log.Debug().Msg("Initialize dependencies successful done")

	if err := srv.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := srv.Shutdown(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func setupLogger(logLevel string) {
	if level, err := zerolog.ParseLevel(logLevel); err != nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(level)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	},
	).With().Caller().Logger()
	log2.SetFlags(0)
	log2.SetOutput(log.Logger)
}
