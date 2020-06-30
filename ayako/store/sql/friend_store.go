package sql

import (
	"context"
	"github.com/deissh/rl/ayako/entity"
	"github.com/deissh/rl/ayako/errors"
	"github.com/deissh/rl/ayako/store"
)

type FriendStore struct {
	SqlStore
}

func newSqlFriendStore(sqlStore SqlStore) store.Friend {
	return &FriendStore{sqlStore}
}

func (f FriendStore) Add(ctx context.Context, userId, targetId uint) error {
	_, err := f.GetMaster().ExecContext(
		ctx,
		`INSERT INTO user_relation (user_id, target_id) VALUES ($1, $2)`,
		userId, targetId,
	)
	if err != nil {
		return errors.WithCause(500, "creating relationships", err)
	}

	return nil
}

func (f FriendStore) Remove(ctx context.Context, userId, targetId uint) error {
	_, err := f.GetMaster().ExecContext(
		ctx,
		`DELETE FROM user_relation WHERE user_id = $1 AND target_id = $2`,
		userId, targetId,
	)
	if err != nil {
		return errors.WithCause(500, "remove relationships", err)
	}

	return nil
}

func (f FriendStore) Check(ctx context.Context, userId, targetId uint) error {
	panic("implement me")
}

func (f FriendStore) GetSubscriptions(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	users := make([]entity.UserShort, 0)

	err := f.GetMaster().SelectContext(
		ctx,
		&users,
		`SELECT u.id, u.username, u.email, u.is_bot, u.is_active,
       		u.is_supporter, u.has_supported, u.support_level,
       		u.pm_friends_only, u.avatar_url, u.country_code,
       		u.default_group, u.last_visit, u.created_at,
       		u.support_expired_at, check_online(u.last_visit)
		FROM user_relation
		INNER JOIN users u on user_relation.target_id = u.id
		WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return nil, errors.WithCause(500, "get all subscriptions", err)
	}

	return &users, nil
}

func (f FriendStore) GetFriends(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	panic("implement me")
}
