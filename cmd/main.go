package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
)

var Version string
var Commit string
var Branch string
var BuildTimestamp string

var configFile = flag.String("config", "config.yaml", "config file")

func main() {
	flag.Parse()

	app, err := Injector(*configFile)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	log.Info().
		Str("version", Version).
		Str("branch", Branch).
		Str("commit", Commit).
		Str("build_timestamp", BuildTimestamp).
		Msg("by Rhythm Lovers")

	log.Debug().Msg("initialize dependencies successful done")

	if err := app.Start(); err != nil {
		log.Fatal().Err(err).Send()
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	if err := app.Stop(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
