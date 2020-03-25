package main

import (
	"github.com/deissh/osu-lazer/ayako/client"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	api := client.WithAccessToken(
		os.Getenv("access_token"),
		os.Getenv("refresh_token"),
	)

	data, err := api.BeatmapSet.Get(23416)
	if err != nil {
		log.Error().
			Err(err).
			Send()
		return
	}

	log.Info().
		Int64("id", data.ID).
		Str("title", data.Title).
		Interface("covers", data.Covers).
		Send()
}
