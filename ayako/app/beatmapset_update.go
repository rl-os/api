package app

import "github.com/rs/zerolog/log"

func (s *App) DoBeatmapSetUpdateMark() {
	log.Info().
		Str("job", "DoBeatmapSetUpdateMark").
		Msg("start beatmapset update check")

	log.Info().
		Str("job", "DoBeatmapSetUpdateMark").
		Msg("start update 1 beatmapset")

	// todo
	_, err := s.Store.BeatmapSet().Fetch(80200, true)
	if err != nil {
		log.Error().
			Err(err).
			Str("job", "DoBeatmapSetUpdateMark").
			Uint("beatmapset_id", 1).
			Msg("fetching beatmapset")
		return
	}
	// todo
}
