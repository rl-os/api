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
