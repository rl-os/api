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

var friendDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_store_friend_duration_seconds",
		Help:       "friend runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// FriendWithLog implements store.Friend that is instrumented with zerolog
type FriendWithLog struct {
	_base store.Friend
}

func NewFriendWithLog(base store.Friend) store.Friend {
	return FriendWithLog{
		_base: base,
	}
}

// Add implements store.Friend
func (_d FriendWithLog) Add(ctx context.Context, userId uint, targetId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Friend.Add: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Friend.Add: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.Add: finished")
		}
		friendDurationSummaryVec.WithLabelValues("Friend", "Add", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Add(ctx, userId, targetId)
}

// GetSubscriptions implements store.Friend
func (_d FriendWithLog) GetSubscriptions(ctx context.Context, userId uint) (uap1 *[]entity.UserShort, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.GetSubscriptions: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Friend.GetSubscriptions: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.GetSubscriptions: finished")
		}
		friendDurationSummaryVec.WithLabelValues("Friend", "GetSubscriptions", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetSubscriptions(ctx, userId)
}

// Remove implements store.Friend
func (_d FriendWithLog) Remove(ctx context.Context, userId uint, targetId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Friend.Remove: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.Friend.Remove: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.Remove: finished")
		}
		friendDurationSummaryVec.WithLabelValues("Friend", "Remove", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Remove(ctx, userId, targetId)
}
