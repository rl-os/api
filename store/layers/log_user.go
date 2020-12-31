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

var userDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "app_store_user_duration_seconds",
		Help:       "user runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// UserWithLog implements store.User that is instrumented with zerolog
type UserWithLog struct {
	_base store.User
}

func NewUserWithLog(base store.User) store.User {
	return UserWithLog{
		_base: base,
	}
}

// Activate implements store.User
func (_d UserWithLog) Activate(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.Activate: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Activate: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Activate: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Activate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Activate(ctx, userId)
}

// Ban implements store.User
func (_d UserWithLog) Ban(ctx context.Context, userId uint, duration time.Duration) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("duration", duration).
		Msg("store.User.Ban: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Ban: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Ban: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Ban", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Ban(ctx, userId, duration)
}

// ComputeFields implements store.User
func (_d UserWithLog) ComputeFields(ctx context.Context, user entity.User) (up1 *entity.User, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("user", user).
		Msg("store.User.ComputeFields: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.ComputeFields: returned an error")
		} else {
			log.Trace().
				Msg("store.User.ComputeFields: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "ComputeFields", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.ComputeFields(ctx, user)
}

// Create implements store.User
func (_d UserWithLog) Create(ctx context.Context, name string, email string, pwd string) (up1 *entity.User, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("name", name).
		Interface("email", email).
		Interface("pwd", pwd).
		Msg("store.User.Create: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Create: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Create: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Create(ctx, name, email, pwd)
}

// Deactivate implements store.User
func (_d UserWithLog) Deactivate(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.Deactivate: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Deactivate: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Deactivate: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Deactivate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Deactivate(ctx, userId)
}

// Get implements store.User
func (_d UserWithLog) Get(ctx context.Context, userId uint, mode string) (up1 *entity.User, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("mode", mode).
		Msg("store.User.Get: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Get: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Get", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Get(ctx, userId, mode)
}

// GetByBasic implements store.User
func (_d UserWithLog) GetByBasic(ctx context.Context, login string, pwd string) (up1 *entity.UserShort, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("login", login).
		Interface("pwd", pwd).
		Msg("store.User.GetByBasic: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.GetByBasic: returned an error")
		} else {
			log.Trace().
				Msg("store.User.GetByBasic: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "GetByBasic", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetByBasic(ctx, login, pwd)
}

// GetShort implements store.User
func (_d UserWithLog) GetShort(ctx context.Context, userId uint, mode string) (up1 *entity.UserShort, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("mode", mode).
		Msg("store.User.GetShort: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.GetShort: returned an error")
		} else {
			log.Trace().
				Msg("store.User.GetShort: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "GetShort", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.GetShort(ctx, userId, mode)
}

// Mute implements store.User
func (_d UserWithLog) Mute(ctx context.Context, userId uint, duration time.Duration) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("duration", duration).
		Msg("store.User.Mute: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Mute: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Mute: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Mute", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Mute(ctx, userId, duration)
}

// UnBan implements store.User
func (_d UserWithLog) UnBan(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UnBan: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.UnBan: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UnBan: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "UnBan", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.UnBan(ctx, userId)
}

// UnMute implements store.User
func (_d UserWithLog) UnMute(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UnMute: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.UnMute: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UnMute: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "UnMute", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.UnMute(ctx, userId)
}

// Update implements store.User
func (_d UserWithLog) Update(ctx context.Context, userId uint, from interface{}) (up1 *entity.UserShort, err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("from", from).
		Msg("store.User.Update: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.Update: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Update: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "Update", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.Update(ctx, userId, from)
}

// UpdateLastVisit implements store.User
func (_d UserWithLog) UpdateLastVisit(ctx context.Context, userId uint) (err error) {
	_since := time.Now()
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UpdateLastVisit: calling")
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
			log.Trace().Err(err).
				Msg("store.User.UpdateLastVisit: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UpdateLastVisit: finished")
		}
		userDurationSummaryVec.WithLabelValues("User", "UpdateLastVisit", result).Observe(time.Since(_since).Seconds())
	}()
	return _d._base.UpdateLastVisit(ctx, userId)
}
