package server

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Server) DoBeatmapSetSearchNew() {
	log.Info().
		Str("job", "DoBeatmapSetSearchNew").
		Msg("start beatmapset search")

	ctx, cancel := context.WithTimeout(s.Context, time.Minute)
	defer cancel()

	id, err := s.GetStore().BeatmapSet().GetLatestId(ctx)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetSearchNew").
			Msg("getting id latest beatmapset id")
		return
	}

	search := func(id uint) {
		data, err := s.GetStore().BeatmapSet().FetchFromBancho(ctx, id)
		if err != nil {
			log.Debug().
				Err(err).
				Str("job", "DoBeatmapSetSearchNew").Send()
			return
		}

		data.LastChecked = time.Now()
		_, err = s.GetStore().BeatmapSet().Create(ctx, data)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetSearchNew").
				Uint("beatmap_set_id", uint(data.ID)).
				Msg("creating new beatmapsets")
			return
		}
		log.Info().
			Str("job", "DoBeatmapSetSearchNew").
			Uint("beatmap_set_id", uint(data.ID)).
			Msg("added new beatmapset")
	}

	// trying get 10 beatmaps with id + i
	for i := 1; i <= 10; i++ {
		select {
		case <-s.GoroutineExitSignal:
		case <-ctx.Done():
			break
		default:
			search(id + uint(i))
		}
	}
}
