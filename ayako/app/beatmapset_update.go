package app

import (
	"github.com/rs/zerolog/log"
	"time"
)

func (s *App) DoBeatmapSetUpdate() {
	log.Info().
		Str("job", "DoBeatmapSetUpdate").
		Uint("batch_size", 100).
		Msg("start beatmapset update check")

	ids, err := s.Store.BeatmapSet().GetBeatmapSetIdForUpdate(100)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdate").
			Msg("getting ids for update")
		return
	}

	for _, id := range ids {
		log.Debug().
			Str("job", "DoBeatmapSetUpdate").
			Uint("beatmap_set_id", id).
			Msg("fetching")

		data, err := s.Store.BeatmapSet().Fetch(id)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetUpdate").
				Uint("beatmap_set_id", id).
				Msg("not fetched")
			return
		}

		data.LastChecked = time.Now()

		_, err = s.Store.BeatmapSet().UpdateBeatmapSet(id, *data)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetUpdate").
				Uint("beatmap_set_id", id).
				Msg("not updated")
			return
		}

		log.Debug().
			Str("job", "DoBeatmapSetUpdate").
			Uint("beatmap_set_id", id).
			Msg("updated")
	}
}
