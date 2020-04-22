package app

import (
	"github.com/rs/zerolog/log"
	"time"
)

func (s *App) DoBeatmapSetUpdateMark() {
	log.Info().
		Str("job", "DoBeatmapSetUpdateMark").
		Msg("start beatmapset update check")

	log.Info().
		Str("job", "DoBeatmapSetUpdateMark").
		Msg("start update 1 beatmapset")

	data, err := s.Store.BeatmapSet().Fetch(23416, true)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdateMark").
			Uint("beatmapset_id", 1).
			Msg("fetching beatmapset")
		return
	}

	data.LastUpdated = time.Now()
	_, err = s.Store.BeatmapSet().CreateBeatmapSet(*data)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdateMark").
			Uint("beatmapset_id", 1).
			Msg("creating beatmapset")
		return
	}
}
