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

func NewBeatmapSetWithLog(base store.BeatmapSet) store.BeatmapSet {
	return BeatmapSetWithLog{
		_base: base,
	}
}

// ComputeFields implements store.BeatmapSet
func (_d BeatmapSetWithLog) ComputeFields(set entity.BeatmapSetFull) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("set", set).
		Msg("store.BeatmapSet.ComputeFields: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.ComputeFields: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.ComputeFields: finished")
		}
	}()
	return _d._base.ComputeFields(set)
}

// Create implements store.BeatmapSet
func (_d BeatmapSetWithLog) Create(from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("from", from).
		Msg("store.BeatmapSet.Create: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.Create: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.Create: finished")
		}
	}()
	return _d._base.Create(from)
}

// Delete implements store.BeatmapSet
func (_d BeatmapSetWithLog) Delete(id uint) (err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.BeatmapSet.Delete: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.Delete: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.Delete: finished")
		}
	}()
	return _d._base.Delete(id)
}

// FetchFromBancho implements store.BeatmapSet
func (_d BeatmapSetWithLog) FetchFromBancho(id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.BeatmapSet.FetchFromBancho: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.FetchFromBancho: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.FetchFromBancho: finished")
		}
	}()
	return _d._base.FetchFromBancho(id)
}

// Get implements store.BeatmapSet
func (_d BeatmapSetWithLog) Get(id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Msg("store.BeatmapSet.Get: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.Get: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.Get: finished")
		}
	}()
	return _d._base.Get(id)
}

// GetAll implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetAll(page int, limit int) (bap1 *[]entity.BeatmapSet, err error) {
	log.Debug().
		Interface("page", page).
		Interface("limit", limit).
		Msg("store.BeatmapSet.GetAll: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.GetAll: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.GetAll: finished")
		}
	}()
	return _d._base.GetAll(page, limit)
}

// GetIdsForUpdate implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetIdsForUpdate(limit int) (ua1 []uint, err error) {
	log.Debug().
		Interface("limit", limit).
		Msg("store.BeatmapSet.GetIdsForUpdate: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.GetIdsForUpdate: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.GetIdsForUpdate: finished")
		}
	}()
	return _d._base.GetIdsForUpdate(limit)
}

// GetLatestId implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetLatestId() (u1 uint, err error) {
	log.Debug().Msg("store.BeatmapSet.GetLatestId: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.GetLatestId: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.GetLatestId: finished")
		}
	}()
	return _d._base.GetLatestId()
}

// SetFavourite implements store.BeatmapSet
func (_d BeatmapSetWithLog) SetFavourite(userId uint, id uint) (u1 uint, err error) {
	log.Debug().
		Interface("userId", userId).
		Interface("id", id).
		Msg("store.BeatmapSet.SetFavourite: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.SetFavourite: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.SetFavourite: finished")
		}
	}()
	return _d._base.SetFavourite(userId, id)
}

// SetUnFavourite implements store.BeatmapSet
func (_d BeatmapSetWithLog) SetUnFavourite(userId uint, id uint) (u1 uint, err error) {
	log.Debug().
		Interface("userId", userId).
		Interface("id", id).
		Msg("store.BeatmapSet.SetUnFavourite: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.SetUnFavourite: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.SetUnFavourite: finished")
		}
	}()
	return _d._base.SetUnFavourite(userId, id)
}

// Update implements store.BeatmapSet
func (_d BeatmapSetWithLog) Update(id uint, from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Debug().
		Interface("id", id).
		Interface("from", from).
		Msg("store.BeatmapSet.Update: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.BeatmapSet.Update: returned an error")
		} else {
			log.Debug().
				Msg("store.BeatmapSet.Update: finished")
		}
	}()
	return _d._base.Update(id, from)
}
