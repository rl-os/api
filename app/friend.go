package app

import (
	"context"
	"github.com/rl-os/api/entity"
)

type Friend struct {
	*App
}

// GetAll friends by user id
func (a *Friend) GetAll(ctx context.Context, userID uint) (*[]entity.UserShort, error) {
	data, err := a.Store.Friend().GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}

// Add friend id as a subscription to use userId
func (a *Friend) Add(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
	err := a.Store.Friend().Add(ctx, userID, targetID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	data, err := a.Store.Friend().GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	return data, nil
}

// Remove friend from subscriptions
func (a *Friend) Remove(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
	err := a.Store.Friend().Remove(ctx, userID, targetID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	data, err := a.Store.Friend().GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	return data, nil
}
