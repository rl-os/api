package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog/log"
)

var beatmapDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_store_beatmap_duration_seconds",
		Help:       "beatmap runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// BeatmapWithLog implements store.Beatmap that is instrumented with zerolog
type BeatmapWithLog struct {
	_base store.Beatmap
}

func NewBeatmapWithLog(base store.Beatmap) store.Beatmap {
	return BeatmapWithLog{
		_base: base,
	}
}

// Create implements store.Beatmap
func (_d BeatmapWithLog) Create(ctx context.Context, from interface{}) (bp1 *entity.Beatmap, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("from", from).
		Msg("store.Beatmap.Create: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.Create: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Create: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Create(ctx, from)
}

// CreateBatch implements store.Beatmap
func (_d BeatmapWithLog) CreateBatch(ctx context.Context, from interface{}) (bap1 *[]entity.Beatmap, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("from", from).
		Msg("store.Beatmap.CreateBatch: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.CreateBatch: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.CreateBatch: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "CreateBatch", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.CreateBatch(ctx, from)
}

// Delete implements store.Beatmap
func (_d BeatmapWithLog) Delete(ctx context.Context, id uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.Beatmap.Delete: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.Delete: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Delete: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "Delete", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Delete(ctx, id)
}

// Get implements store.Beatmap
func (_d BeatmapWithLog) Get(ctx context.Context, id uint) (sp1 *entity.SingleBeatmap, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Msg("store.Beatmap.Get: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Get: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "Get", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Get(ctx, id)
}

// GetBySetId implements store.Beatmap
func (_d BeatmapWithLog) GetBySetId(ctx context.Context, beatmapsetId uint) (bap1 *[]entity.Beatmap, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("beatmapsetId", beatmapsetId).
		Msg("store.Beatmap.GetBySetId: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.GetBySetId: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.GetBySetId: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "GetBySetId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetBySetId(ctx, beatmapsetId)
}

// Update implements store.Beatmap
func (_d BeatmapWithLog) Update(ctx context.Context, id uint, from interface{}) (bp1 *entity.Beatmap, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("id", id).
		Interface("from", from).
		Msg("store.Beatmap.Update: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Beatmap.Update: returned an error")
		} else {
			log.Trace().
				Msg("store.Beatmap.Update: finished")
		}
		beatmapDurationSummaryVec.WithLabelValues("Beatmap", "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Update(ctx, id, from)
}
