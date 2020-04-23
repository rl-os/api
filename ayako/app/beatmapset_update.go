package app

import (
	"github.com/rs/zerolog/log"
	"time"
)

func (s *App) DoBeatmapSetUpdateMark() {
	log.Info().
		Str("job", "DoBeatmapSetUpdateMark").
		Msg("start beatmapset update check")

	ids, err := s.Store.BeatmapSet().GetBeatmapSetIdForUpdate(100)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdateMark").
			Msg("getting ids for update")
		return
	}

	for _, id := range ids {
		log.Info().
			Str("job", "DoBeatmapSetUpdateMark").
			Uint("beatmap_set_id", id).
			Msg("start update beatmapset")

		data, err := s.Store.BeatmapSet().Fetch(id)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetUpdateMark").
				Uint("beatmapset_id", 1).
				Msg("fetching beatmapset")
			return
		}

		data.LastChecked = time.Now()

		_, err = s.Store.BeatmapSet().UpdateBeatmapSet(id, *data)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetUpdateMark").
				Uint("beatmapset_id", 1).
				Msg("creating/updating beatmapset")
			return
		}
	}
}
