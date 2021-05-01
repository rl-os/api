package app

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
)

type FriendUseCase struct {
	*App
	FriendRepository repository.Friend
}

func NewFriendUseCase(app *App, rep repository.Friend) *FriendUseCase {
	return &FriendUseCase{app, rep}
}

// Get by user id
func (a *FriendUseCase) Get(ctx context.Context, userID uint) (*[]entity.UserShort, error) {
	data, err := a.FriendRepository.GetSubscriptions(ctx, userID)
	if err != nil {
		return nil, ErrNotFoundUser.WithCause(err)
	}

	return data, nil
}

// Add id as a subscription to use userId
func (a *FriendUseCase) Add(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
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

// Remove from subscriptions
func (a *FriendUseCase) Remove(ctx context.Context, userID, targetID uint) (*[]entity.UserShort, error) {
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
