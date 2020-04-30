package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"github.com/deissh/osu-lazer/ayako/entity"
	"github.com/deissh/osu-lazer/ayako/store"
	"github.com/rs/zerolog/log"
)

// BeatmapSetWithLog implements store.BeatmapSet that is instrumented with zerolog
type BeatmapSetWithLog struct {
	_base store.BeatmapSet
}

// ComputeBeatmapSet implements store.BeatmapSet
func (_d BeatmapSetWithLog) ComputeBeatmapSet(set entity.BeatmapSetFull) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("set", set).
		Msg("store.ComputeBeatmapSet: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.ComputeBeatmapSet: returned an error")
		} else {
			log.Debug().
				Msg("store.ComputeBeatmapSet: finished")
		}
	}()
	return _d._base.ComputeBeatmapSet(set)
}

// CreateBeatmapSet implements store.BeatmapSet
func (_d BeatmapSetWithLog) CreateBeatmapSet(from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.CreateBeatmapSet: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.CreateBeatmapSet: returned an error")
		} else {
			log.Debug().
				Msg("store.CreateBeatmapSet: finished")
		}
	}()
	return _d._base.CreateBeatmapSet(from)
}

// DeleteBeatmapSet implements store.BeatmapSet
func (_d BeatmapSetWithLog) DeleteBeatmapSet(id uint) (err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.DeleteBeatmapSet: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.DeleteBeatmapSet: returned an error")
		} else {
			log.Debug().
				Msg("store.DeleteBeatmapSet: finished")
		}
	}()
	return _d._base.DeleteBeatmapSet(id)
}

// Fetch implements store.BeatmapSet
func (_d BeatmapSetWithLog) Fetch(id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.Fetch: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Fetch: returned an error")
		} else {
			log.Debug().
				Msg("store.Fetch: finished")
		}
	}()
	return _d._base.Fetch(id)
}

// GetAllBeatmapSets implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetAllBeatmapSets(page int, limit int) (bap1 *[]entity.BeatmapSet, err error) {
	log.Debug().
		Interface("page", page).
		Interface("limit", limit).
		Msg("store.GetAllBeatmapSets: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetAllBeatmapSets: returned an error")
		} else {
			log.Debug().
				Msg("store.GetAllBeatmapSets: finished")
		}
	}()
	return _d._base.GetAllBeatmapSets(page, limit)
}

// GetBeatmapSet implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetBeatmapSet(id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.GetBeatmapSet: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetBeatmapSet: returned an error")
		} else {
			log.Debug().
				Msg("store.GetBeatmapSet: finished")
		}
	}()
	return _d._base.GetBeatmapSet(id)
}

// GetBeatmapSetIdForUpdate implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetBeatmapSetIdForUpdate(limit int) (ua1 []uint, err error) {
	log.Debug().
		Interface("limit", limit).
		Msg("store.GetBeatmapSetIdForUpdate: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetBeatmapSetIdForUpdate: returned an error")
		} else {
			log.Debug().
				Msg("store.GetBeatmapSetIdForUpdate: finished")
		}
	}()
	return _d._base.GetBeatmapSetIdForUpdate(limit)
}

// GetLatestBeatmapId implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetLatestBeatmapId() (u1 uint, err error) {
	log.Debug().Msg("store.GetLatestBeatmapId: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetLatestBeatmapId: returned an error")
		} else {
			log.Debug().
				Msg("store.GetLatestBeatmapId: finished")
		}
	}()
	return _d._base.GetLatestBeatmapId()
}

// UpdateBeatmapSet implements store.BeatmapSet
func (_d BeatmapSetWithLog) UpdateBeatmapSet(id uint, from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Interface("from", from).
		Msg("store.UpdateBeatmapSet: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.UpdateBeatmapSet: returned an error")
		} else {
			log.Debug().
				Msg("store.UpdateBeatmapSet: finished")
		}
	}()
	return _d._base.UpdateBeatmapSet(id, from)
}
