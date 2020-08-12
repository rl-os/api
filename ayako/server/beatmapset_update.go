package server

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

func (s *Server) DoBeatmapSetUpdate() {
	log.Info().
		Str("job", "DoBeatmapSetUpdate").
		Uint("batch_size", 100).
		Msg("start beatmapset update check")

	ctx, cancel := context.WithCancel(s.Context)
	defer cancel()

	ids, err := s.GetStore().BeatmapSet().GetIdsForUpdate(ctx, 100)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdate").
			Msg("getting ids for update")
		return
	}

	updater := func(id uint) {
		log.Debug().
			Str("job", "DoBeatmapSetUpdate").
			Uint("beatmap_set_id", id).
			Msg("fetching")

		data, err := s.GetStore().BeatmapSet().FetchFromBancho(ctx, id)
		if err != nil {
			log.Error().
				Err(err).
				Str("job", "DoBeatmapSetUpdate").
				Uint("beatmap_set_id", id).
				Msg("not fetched")
			return
		}

		data.LastChecked = time.Now()

		_, err = s.GetStore().BeatmapSet().Update(ctx, id, *data)
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

	for _, id := range ids {
		select {
		case <-s.GoroutineExitSignal:
		case <-ctx.Done():
			break
		default:
			updater(id)
		}
	}
}
