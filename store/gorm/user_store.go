package sql

import (
	"context"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
	"gorm.io/gorm/clause"
	"time"
)

type UserStore struct {
	SqlStore
}

func newSqlUserStore(sqlStore SqlStore) store.User {
	return &UserStore{sqlStore}
}

func (u UserStore) GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error) {
	var baseUser entity.UserAuthBase

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("username = ? OR email = ?", login).
		First(&baseUser).
		Error
	if err != nil {
		return nil, errors.WithCause("user_get", 401, "user credentials were incorrect", err)
	}

	if ok := utils.CompareHash(baseUser.PasswordHash, pwd); !ok {
		return nil, errors.WithCause("user_get", 401, "user credentials were incorrect", err)
	}

	return &baseUser.UserShort, nil
}

func (u UserStore) Get(ctx context.Context, userId uint, mode string) (*entity.User, error) {
	user := entity.User{Mode: mode}

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Preload(clause.Associations).
		First(&user).
		Error

	if err != nil {
		return nil, errors.WithCause("user_get", 404, "user not found", err)
	}

	return u.ComputeFields(ctx, user)
}

func (u UserStore) GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error) {
	var user entity.UserShort

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		First(&user).
		Error

	if err != nil {
		return nil, errors.WithCause("user_get", 404, "user not found", err)
	}

	return &user, nil
}

func (u UserStore) Create(ctx context.Context, name, email, pwd string) (*entity.User, error) {
	return nil, nil
}

func (u UserStore) Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error) {
	panic("implement me")
}

func (u UserStore) Activate(ctx context.Context, userId uint) error {
	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Update("is_active", true).
		Error
	if err != nil {
		return errors.WithCause("user_activate", 404, "User not found", err)
	}

	return nil
}

func (u UserStore) Deactivate(ctx context.Context, userId uint) error {
	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Model(&entity.User{}).
		Where("id = ?", userId).
		Update("is_active", false).
		Error
	if err != nil {
		return errors.WithCause("user_activate", 404, "User not found", err)
	}

	return nil
}

func (u UserStore) Ban(ctx context.Context, userId uint, time time.Duration) error {
	panic("implement me")
}

func (u UserStore) UnBan(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (u UserStore) Mute(ctx context.Context, userId uint, time time.Duration) error {
	panic("implement me")
}

func (u UserStore) UnMute(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (u UserStore) UpdateLastVisit(ctx context.Context, userId uint) error {
	currentTime := time.Now().UTC()

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Update("last_visit", currentTime).
		Error
	if err != nil {
		return errors.WithCause("user_update_online", 404, "User not found", err)
	}

	return nil
}

func (u UserStore) ComputeFields(ctx context.Context, user entity.User) (*entity.User, error) {
	return &user, nil
}
