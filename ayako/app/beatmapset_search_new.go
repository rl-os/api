package app

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *App) DoBeatmapSetSearchNew() {
	log.Info().
		Str("job", "DoBeatmapSetSearchNew").
		Msg("start beatmapset search")

	//fixme: aaa!!!
	ctx := context.TODO()

	id, err := s.Store.BeatmapSet().GetLatestId(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetSearchNew").
			Msg("getting id latest beatmapset id")
		return
	}

	// trying get 10 beatmaps with id + i
	for i := 1; i <= 10; i++ {
		data, err := s.Store.BeatmapSet().FetchFromBancho(ctx, id+uint(i))
		if err != nil {
			log.Debug().
				Err(err).
				Str("job", "DoBeatmapSetSearchNew").Send()
			continue
		}

		data.LastChecked = time.Now()
		_, err = s.Store.BeatmapSet().Create(ctx, data)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetSearchNew").
				Uint("beatmap_set_id", uint(data.ID)).
				Msg("creating new beatmapsets")
			continue
		}
		log.Info().
			Str("job", "DoBeatmapSetSearchNew").
			Uint("beatmap_set_id", uint(data.ID)).
			Msg("added new beatmapset")
	}
}
