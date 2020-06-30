package sql

import (
	"context"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/store"
)

type FriendStore struct {
	SqlStore
}

func newSqlFriendStore(sqlStore SqlStore) store.Friend {
	return &FriendStore{sqlStore}
}

func (f FriendStore) Add(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (f FriendStore) Remove(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (f FriendStore) Check(ctx context.Context, userId, targetId uint) error {
	panic("implement me")
}

func (f FriendStore) GetSubscriptions(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	panic("implement me")
}

func (f FriendStore) GetFriends(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	panic("implement me")
}
