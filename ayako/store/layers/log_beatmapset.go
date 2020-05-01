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

// ComputeFields implements store.BeatmapSet
func (_d BeatmapSetWithLog) ComputeFields(set entity.BeatmapSetFull) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("set", set).
		Msg("store.ComputeFields: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.ComputeFields: returned an error")
		} else {
			log.Debug().
				Msg("store.ComputeFields: finished")
		}
	}()
	return _d._base.ComputeFields(set)
}

// Create implements store.BeatmapSet
func (_d BeatmapSetWithLog) Create(from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
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

// Delete implements store.BeatmapSet
func (_d BeatmapSetWithLog) Delete(id uint) (err error) {
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

// FetchFromBancho implements store.BeatmapSet
func (_d BeatmapSetWithLog) FetchFromBancho(id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.FetchFromBancho: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.FetchFromBancho: returned an error")
		} else {
			log.Debug().
				Msg("store.FetchFromBancho: finished")
		}
	}()
	return _d._base.FetchFromBancho(id)
}

// Get implements store.BeatmapSet
func (_d BeatmapSetWithLog) Get(id uint) (bp1 *entity.BeatmapSetFull, err error) {
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

// GetAll implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetAll(page int, limit int) (bap1 *[]entity.BeatmapSet, err error) {
	log.Debug().
		Interface("page", page).
		Interface("limit", limit).
		Msg("store.GetAll: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetAll: returned an error")
		} else {
			log.Debug().
				Msg("store.GetAll: finished")
		}
	}()
	return _d._base.GetAll(page, limit)
}

// GetIdsForUpdate implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetIdsForUpdate(limit int) (ua1 []uint, err error) {
	log.Debug().
		Interface("limit", limit).
		Msg("store.GetIdsForUpdate: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetIdsForUpdate: returned an error")
		} else {
			log.Debug().
				Msg("store.GetIdsForUpdate: finished")
		}
	}()
	return _d._base.GetIdsForUpdate(limit)
}

// GetLatestId implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetLatestId() (u1 uint, err error) {
	log.Debug().Msg("store.GetLatestId: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.GetLatestId: returned an error")
		} else {
			log.Debug().
				Msg("store.GetLatestId: finished")
		}
	}()
	return _d._base.GetLatestId()
}

// Update implements store.BeatmapSet
func (_d BeatmapSetWithLog) Update(id uint, from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
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
