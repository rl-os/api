package app

import (
	"context"
	"github.com/rl-os/api/entity"
)

// GetAllFriends by user id
func (a *App) GetAllFriends(ctx context.Context, userID uint) (*[]entity.UserShort, error) {
	data, err := a.FriendRepository.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}

// AddFriend id as a subscription to use userId
func (a *App) AddFriend(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
	err := a.FriendRepository.Add(ctx, userID, targetID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	data, err := a.FriendRepository.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	return data, nil
}

// RemoveFriend from subscriptions
func (a *App) RemoveFriend(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
	err := a.FriendRepository.Remove(ctx, userID, targetID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	data, err := a.FriendRepository.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser
	}

	return data, nil
}
