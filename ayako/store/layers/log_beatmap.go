package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/rs/zerolog/log"
)

// BeatmapWithLog implements store.Beatmap that is instrumented with zerolog
type BeatmapWithLog struct {
	_base store.Beatmap
}

// CreateBeatmap implements store.Beatmap
func (_d BeatmapWithLog) CreateBeatmap(from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.CreateBeatmap: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.CreateBeatmap: returned an error")
		} else {
			log.Debug().
				Msg("store.CreateBeatmap: finished")
		}
	}()
	return _d._base.CreateBeatmap(from)
}

// CreateBeatmaps implements store.Beatmap
func (_d BeatmapWithLog) CreateBeatmaps(from interface{}) (bap1 *[]entity.Beatmap, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.CreateBeatmaps: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.CreateBeatmaps: returned an error")
		} else {
			log.Debug().
				Msg("store.CreateBeatmaps: finished")
		}
	}()
	return _d._base.CreateBeatmaps(from)
}

// DeleteBeatmap implements store.Beatmap
func (_d BeatmapWithLog) DeleteBeatmap(id uint) (err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.DeleteBeatmap: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.DeleteBeatmap: returned an error")
		} else {
			log.Debug().
				Msg("store.DeleteBeatmap: finished")
		}
	}()
	return _d._base.DeleteBeatmap(id)
}

// GetBeatmap implements store.Beatmap
func (_d BeatmapWithLog) GetBeatmap(id uint) (sp1 *entity.SingleBeatmap, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.GetBeatmap: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetBeatmap: returned an error")
		} else {
			log.Debug().
				Msg("store.GetBeatmap: finished")
		}
	}()
	return _d._base.GetBeatmap(id)
}

// GetBeatmapsBySet implements store.Beatmap
func (_d BeatmapWithLog) GetBeatmapsBySet(beatmapsetId uint) (ba1 []entity.Beatmap) {
	log.Debug().
		Interface("beatmapsetId", beatmapsetId).
		Msg("store.GetBeatmapsBySet: calling")
	defer func() {
		log.Debug().
			Msg("store.GetBeatmapsBySet: finished")
	}()
	return _d._base.GetBeatmapsBySet(beatmapsetId)
}

// UpdateBeatmap implements store.Beatmap
func (_d BeatmapWithLog) UpdateBeatmap(id uint, from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Debug().
		Interface("id", id).
		Interface("from", from).
		Msg("store.UpdateBeatmap: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.UpdateBeatmap: returned an error")
		} else {
			log.Debug().
				Msg("store.UpdateBeatmap: finished")
		}
	}()
	return _d._base.UpdateBeatmap(id, from)
}
