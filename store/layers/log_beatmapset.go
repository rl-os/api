package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
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
func (_d BeatmapSetWithLog) ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (bp1 *entity.BeatmapSetFull, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("set", set).
		Msg("store.BeatmapSet.ComputeFields: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.ComputeFields: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.ComputeFields: finished")
		}
	}()
	return _d._base.ComputeFields(ctx, set)
}

// Create implements store.BeatmapSet
func (_d BeatmapSetWithLog) Create(ctx context.Context, from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("from", from).
		Msg("store.BeatmapSet.Create: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.Create: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.Create: finished")
		}
	}()
	return _d._base.Create(ctx, from)
}

// Delete implements store.BeatmapSet
func (_d BeatmapSetWithLog) Delete(ctx context.Context, id uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.BeatmapSet.Delete: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.Delete: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.Delete: finished")
		}
	}()
	return _d._base.Delete(ctx, id)
}

// FetchFromBancho implements store.BeatmapSet
func (_d BeatmapSetWithLog) FetchFromBancho(ctx context.Context, id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.BeatmapSet.FetchFromBancho: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.FetchFromBancho: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.FetchFromBancho: finished")
		}
	}()
	return _d._base.FetchFromBancho(ctx, id)
}

// Get implements store.BeatmapSet
func (_d BeatmapSetWithLog) Get(ctx context.Context, id uint) (bp1 *entity.BeatmapSetFull, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.BeatmapSet.Get: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.Get: finished")
		}
	}()
	return _d._base.Get(ctx, id)
}

// GetAll implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetAll(ctx context.Context, page int, limit int) (bap1 *[]entity.BeatmapSet, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("page", page).
		Interface("limit", limit).
		Msg("store.BeatmapSet.GetAll: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.GetAll: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.GetAll: finished")
		}
	}()
	return _d._base.GetAll(ctx, page, limit)
}

// GetIdsForUpdate implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetIdsForUpdate(ctx context.Context, limit int) (ua1 []uint, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("limit", limit).
		Msg("store.BeatmapSet.GetIdsForUpdate: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.GetIdsForUpdate: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.GetIdsForUpdate: finished")
		}
	}()
	return _d._base.GetIdsForUpdate(ctx, limit)
}

// GetLatestId implements store.BeatmapSet
func (_d BeatmapSetWithLog) GetLatestId(ctx context.Context) (u1 uint, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Msg("store.BeatmapSet.GetLatestId: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.GetLatestId: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.GetLatestId: finished")
		}
	}()
	return _d._base.GetLatestId(ctx)
}

// SetFavourite implements store.BeatmapSet
func (_d BeatmapSetWithLog) SetFavourite(ctx context.Context, userId uint, id uint) (u1 uint, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("id", id).
		Msg("store.BeatmapSet.SetFavourite: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.SetFavourite: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.SetFavourite: finished")
		}
	}()
	return _d._base.SetFavourite(ctx, userId, id)
}

// SetUnFavourite implements store.BeatmapSet
func (_d BeatmapSetWithLog) SetUnFavourite(ctx context.Context, userId uint, id uint) (u1 uint, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("id", id).
		Msg("store.BeatmapSet.SetUnFavourite: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.SetUnFavourite: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.SetUnFavourite: finished")
		}
	}()
	return _d._base.SetUnFavourite(ctx, userId, id)
}

// Update implements store.BeatmapSet
func (_d BeatmapSetWithLog) Update(ctx context.Context, id uint, from interface{}) (bp1 *entity.BeatmapSetFull, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Interface("from", from).
		Msg("store.BeatmapSet.Update: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.BeatmapSet.Update: returned an error")
		} else {
			log.Trace().
				Msg("store.BeatmapSet.Update: finished")
		}
	}()
	return _d._base.Update(ctx, id, from)
}
