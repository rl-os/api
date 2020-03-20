package main

import (
	"github.com/deissh/osu-lazer/ayako/client"
	"github.com/rs/zerolog/log"
)

func main() {
	api := client.WithAccessToken("", "")

	data, _ := api.BeatmapSet.Get(23416)
	log.Info().
		Int64("id", data.ID).
		Str("title", data.Title).
		Interface("covers", data.Covers).
		Send()
}
