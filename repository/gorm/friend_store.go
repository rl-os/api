package gorm

import (
	"context"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
)

type FriendRepository struct {
	*Supplier
}

func NewFriendRepository(supplier *Supplier) repository.Friend {
	return &FriendRepository{supplier}
}

func (f FriendRepository) Add(ctx context.Context, userId, targetId uint) error {
	panic("implement me")
}

func (f FriendRepository) Remove(ctx context.Context, userId, targetId uint) error {
	panic("implement me")
}

func (f FriendRepository) GetSubscriptions(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	panic("implement me")
}
