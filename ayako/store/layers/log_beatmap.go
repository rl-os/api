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

// Create implements store.Beatmap
func (_d BeatmapWithLog) Create(from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.Create: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Create: returned an error")
		} else {
			log.Debug().
				Msg("store.Create: finished")
		}
	}()
	return _d._base.Create(from)
}

// CreateBatch implements store.Beatmap
func (_d BeatmapWithLog) CreateBatch(from interface{}) (bap1 *[]entity.Beatmap, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.CreateBatch: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.CreateBatch: returned an error")
		} else {
			log.Debug().
				Msg("store.CreateBatch: finished")
		}
	}()
	return _d._base.CreateBatch(from)
}

// Delete implements store.Beatmap
func (_d BeatmapWithLog) Delete(id uint) (err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.Delete: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Delete: returned an error")
		} else {
			log.Debug().
				Msg("store.Delete: finished")
		}
	}()
	return _d._base.Delete(id)
}

// Get implements store.Beatmap
func (_d BeatmapWithLog) Get(id uint) (sp1 *entity.SingleBeatmap, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.Get: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Get: returned an error")
		} else {
			log.Debug().
				Msg("store.Get: finished")
		}
	}()
	return _d._base.Get(id)
}

// GetBySetId implements store.Beatmap
func (_d BeatmapWithLog) GetBySetId(beatmapsetId uint) (ba1 []entity.Beatmap) {
	log.Debug().
		Interface("beatmapsetId", beatmapsetId).
		Msg("store.GetBySetId: calling")
	defer func() {
		log.Debug().
			Msg("store.GetBySetId: finished")
	}()
	return _d._base.GetBySetId(beatmapsetId)
}

// Update implements store.Beatmap
func (_d BeatmapWithLog) Update(id uint, from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Debug().
		Interface("id", id).
		Interface("from", from).
		Msg("store.Update: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Update: returned an error")
		} else {
			log.Debug().
				Msg("store.Update: finished")
		}
	}()
	return _d._base.Update(id, from)
}
