package layers

// DO NOT EDIT!
// This code is generated with http://github.com/hexdigest/gowrap tool
// using log.tmpl template

import (
	"context"

	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/store"
	"github.com/rs/zerolog/log"
)

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
func (_d FriendWithLog) Add(ctx context.Context, userId uint) (err error) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.Add: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Friend.Add: returned an error")
		} else {
			log.Debug().
				Msg("store.Friend.Add: finished")
		}
	}()
	return _d._base.Add(ctx, userId)
}

// Check implements store.Friend
func (_d FriendWithLog) Check(ctx context.Context, userId uint, targetId uint) (err error) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Friend.Check: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Friend.Check: returned an error")
		} else {
			log.Debug().
				Msg("store.Friend.Check: finished")
		}
	}()
	return _d._base.Check(ctx, userId, targetId)
}

// GetFriends implements store.Friend
func (_d FriendWithLog) GetFriends(ctx context.Context, userId uint) (uap1 *[]entity.UserShort, err error) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.GetFriends: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Friend.GetFriends: returned an error")
		} else {
			log.Debug().
				Msg("store.Friend.GetFriends: finished")
		}
	}()
	return _d._base.GetFriends(ctx, userId)
}

// GetSubscriptions implements store.Friend
func (_d FriendWithLog) GetSubscriptions(ctx context.Context, userId uint) (uap1 *[]entity.UserShort, err error) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.GetSubscriptions: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Friend.GetSubscriptions: returned an error")
		} else {
			log.Debug().
				Msg("store.Friend.GetSubscriptions: finished")
		}
	}()
	return _d._base.GetSubscriptions(ctx, userId)
}

// Remove implements store.Friend
func (_d FriendWithLog) Remove(ctx context.Context, userId uint) (err error) {
	log.Debug().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.Remove: calling")
	defer func() {
		if err != nil {
			log.Error().Err(err).
				Msg("store.Friend.Remove: returned an error")
		} else {
			log.Debug().
				Msg("store.Friend.Remove: finished")
		}
	}()
	return _d._base.Remove(ctx, userId)
}
