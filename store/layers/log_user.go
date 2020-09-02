package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"
	"time"

	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/store"
	"github.com/rs/zerolog/log"
)

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
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.Activate: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Activate: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Activate: finished")
		}
	}()
	return _d._base.Activate(ctx, userId)
}

// Ban implements store.User
func (_d UserWithLog) Ban(ctx context.Context, userId uint, time time.Duration) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("time", time).
		Msg("store.User.Ban: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Ban: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Ban: finished")
		}
	}()
	return _d._base.Ban(ctx, userId, time)
}

// ComputeFields implements store.User
func (_d UserWithLog) ComputeFields(ctx context.Context, user entity.User) (up1 *entity.User, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("user", user).
		Msg("store.User.ComputeFields: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.ComputeFields: returned an error")
		} else {
			log.Trace().
				Msg("store.User.ComputeFields: finished")
		}
	}()
	return _d._base.ComputeFields(ctx, user)
}

// Create implements store.User
func (_d UserWithLog) Create(ctx context.Context, name string, email string, pwd string) (up1 *entity.User, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("name", name).
		Interface("email", email).
		Interface("pwd", pwd).
		Msg("store.User.Create: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Create: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Create: finished")
		}
	}()
	return _d._base.Create(ctx, name, email, pwd)
}

// Deactivate implements store.User
func (_d UserWithLog) Deactivate(ctx context.Context, userId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.Deactivate: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Deactivate: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Deactivate: finished")
		}
	}()
	return _d._base.Deactivate(ctx, userId)
}

// Get implements store.User
func (_d UserWithLog) Get(ctx context.Context, userId uint, mode string) (up1 *entity.User, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("mode", mode).
		Msg("store.User.Get: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Get: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Get: finished")
		}
	}()
	return _d._base.Get(ctx, userId, mode)
}

// GetByBasic implements store.User
func (_d UserWithLog) GetByBasic(ctx context.Context, login string, pwd string) (up1 *entity.UserShort, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("login", login).
		Interface("pwd", pwd).
		Msg("store.User.GetByBasic: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.GetByBasic: returned an error")
		} else {
			log.Trace().
				Msg("store.User.GetByBasic: finished")
		}
	}()
	return _d._base.GetByBasic(ctx, login, pwd)
}

// GetShort implements store.User
func (_d UserWithLog) GetShort(ctx context.Context, userId uint, mode string) (up1 *entity.UserShort, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("mode", mode).
		Msg("store.User.GetShort: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.GetShort: returned an error")
		} else {
			log.Trace().
				Msg("store.User.GetShort: finished")
		}
	}()
	return _d._base.GetShort(ctx, userId, mode)
}

// Mute implements store.User
func (_d UserWithLog) Mute(ctx context.Context, userId uint, time time.Duration) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("time", time).
		Msg("store.User.Mute: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Mute: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Mute: finished")
		}
	}()
	return _d._base.Mute(ctx, userId, time)
}

// UnBan implements store.User
func (_d UserWithLog) UnBan(ctx context.Context, userId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UnBan: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.UnBan: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UnBan: finished")
		}
	}()
	return _d._base.UnBan(ctx, userId)
}

// UnMute implements store.User
func (_d UserWithLog) UnMute(ctx context.Context, userId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UnMute: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.UnMute: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UnMute: finished")
		}
	}()
	return _d._base.UnMute(ctx, userId)
}

// Update implements store.User
func (_d UserWithLog) Update(ctx context.Context, userId uint, from interface{}) (up1 *entity.UserShort, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("from", from).
		Msg("store.User.Update: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.Update: returned an error")
		} else {
			log.Trace().
				Msg("store.User.Update: finished")
		}
	}()
	return _d._base.Update(ctx, userId, from)
}

// UpdateLastVisit implements store.User
func (_d UserWithLog) UpdateLastVisit(ctx context.Context, userId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.User.UpdateLastVisit: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.User.UpdateLastVisit: returned an error")
		} else {
			log.Trace().
				Msg("store.User.UpdateLastVisit: finished")
		}
	}()
	return _d._base.UpdateLastVisit(ctx, userId)
}
