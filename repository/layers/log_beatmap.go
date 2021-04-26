package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"github.com/rs/zerolog/log"
)

// BeatmapWithLog implements repository.Beatmap that is instrumented with zerolog
type BeatmapWithLog struct {
	_base repository.Beatmap
}

func NewBeatmapWithLog(base repository.Beatmap) repository.Beatmap {
	return BeatmapWithLog{
		_base: base,
	}
}

// Create implements repository.Beatmap
func (_d BeatmapWithLog) Create(ctx context.Context, from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("from", from).
		Msg("store.Beatmap.Create: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.Create: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Create: finished")
		}
	}()
	return _d._base.Create(ctx, from)
}

// CreateBatch implements repository.Beatmap
func (_d BeatmapWithLog) CreateBatch(ctx context.Context, from interface{}) (bap1 *[]entity.Beatmap, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("from", from).
		Msg("store.Beatmap.CreateBatch: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.CreateBatch: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.CreateBatch: finished")
		}
	}()
	return _d._base.CreateBatch(ctx, from)
}

// Delete implements repository.Beatmap
func (_d BeatmapWithLog) Delete(ctx context.Context, id uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.Beatmap.Delete: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.Delete: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Delete: finished")
		}
	}()
	return _d._base.Delete(ctx, id)
}

// Get implements repository.Beatmap
func (_d BeatmapWithLog) Get(ctx context.Context, id uint) (sp1 *entity.SingleBeatmap, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.Beatmap.Get: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Get: finished")
		}
	}()
	return _d._base.Get(ctx, id)
}

// GetBySetId implements repository.Beatmap
func (_d BeatmapWithLog) GetBySetId(ctx context.Context, beatmapsetId uint) (bap1 *[]entity.Beatmap, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("beatmapsetId", beatmapsetId).
		Msg("store.Beatmap.GetBySetId: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.GetBySetId: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.GetBySetId: finished")
		}
	}()
	return _d._base.GetBySetId(ctx, beatmapsetId)
}

// Update implements repository.Beatmap
func (_d BeatmapWithLog) Update(ctx context.Context, id uint, from interface{}) (bp1 *entity.Beatmap, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Interface("from", from).
		Msg("store.Beatmap.Update: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Beatmap.Update: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Update: finished")
		}
	}()
	return _d._base.Update(ctx, id, from)
}
