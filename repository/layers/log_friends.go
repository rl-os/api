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

// FriendWithLog implements repository.Friend that is instrumented with zerolog
type FriendWithLog struct {
	_base repository.Friend
}

func NewFriendWithLog(base repository.Friend) repository.Friend {
	return FriendWithLog{
		_base: base,
	}
}

// Add implements repository.Friend
func (_d FriendWithLog) Add(ctx context.Context, userId uint, targetId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Friend.Add: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Friend.Add: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.Add: finished")
		}
	}()
	return _d._base.Add(ctx, userId, targetId)
}

// GetSubscriptions implements repository.Friend
func (_d FriendWithLog) GetSubscriptions(ctx context.Context, userId uint) (uap1 *[]entity.UserShort, err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Msg("store.Friend.GetSubscriptions: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Friend.GetSubscriptions: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.GetSubscriptions: finished")
		}
	}()
	return _d._base.GetSubscriptions(ctx, userId)
}

// Remove implements repository.Friend
func (_d FriendWithLog) Remove(ctx context.Context, userId uint, targetId uint) (err error) {
	log.Trace().
		Interface("ctx", ctx).
		Interface("userId", userId).
		Interface("targetId", targetId).
		Msg("store.Friend.Remove: calling")
	defer func() {
		if err != nil {
			log.Trace().Err(err).
				Msg("store.Friend.Remove: returned an error")
		} else {
			log.Trace().
				Msg("store.Friend.Remove: finished")
		}
	}()
	return _d._base.Remove(ctx, userId, targetId)
}
