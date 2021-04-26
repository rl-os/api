package gorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type UserRepository struct {
	*Supplier
}

func NewUserRepository(supplier *Supplier) repository.User {
	return &UserRepository{supplier}
}

func (u UserRepository) GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error) {
	var baseUser entity.UserShort

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("username = ? OR email = ?", login, login).
		First(&baseUser).
		Error
	if err != nil {
		return nil, err
	}

	if ok := utils.CompareHash(baseUser.PasswordHash, pwd); !ok {
		return nil, errors.New("Invalid password. ")
	}

	return &baseUser, nil
}

func (u UserRepository) Get(ctx context.Context, userId uint, mode string) (*entity.User, error) {
	user := entity.User{}
	user.Mode = mode

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Preload("RankHistory", "mode = ?", mode).
		Preload("Statistics", "mode = ?", mode).
		Preload(clause.Associations).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return u.ComputeFields(ctx, user)
}

func (u UserRepository) GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error) {
	var user entity.UserShort

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) Create(ctx context.Context, name, email, pwd string) (*entity.User, error) {
	now := time.Now()

	hashed, err := utils.GetHash(pwd)
	if err != nil {
		return nil, err
	}

	user := entity.UserBasic{
		Username:     name,
		Email:        email,
		PasswordHash: hashed,
	}

	err = u.GetMaster().Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).
			Table("users").
			Create(&user).
			Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).
			Table("user_month_playcount").
			Create(&entity.MonthlyPlaycounts{
				UserId:    user.ID,
				Count:     0,
				StartDate: fmt.Sprintf("%02d-%02d-01", now.Year(), now.Month()),
			}).Error
		if err != nil {
			return err
		}

		for _, mode := range entity.Modes {
			err = tx.WithContext(ctx).
				Table("user_statistics").
				Create(map[string]interface{}{
					"user_id": user.ID,
					"mode":    mode,
				}).
				Error
			if err != nil {
				return err
			}

			err = tx.WithContext(ctx).
				Table("user_performance_ranks").
				Create(map[string]interface{}{
					"user_id": user.ID,
					"mode":    mode,
				}).
				Error
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return u.Get(ctx, user.ID, "")
}

func (u UserRepository) Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error) {
	panic("implement me")
}

func (u UserRepository) Activate(ctx context.Context, userId uint) error {
	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Update("is_active", true).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Deactivate(ctx context.Context, userId uint) error {
	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Model(&entity.User{}).
		Where("id = ?", userId).
		Update("is_active", false).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Ban(ctx context.Context, userId uint, time time.Duration) error {
	panic("implement me")
}

func (u UserRepository) UnBan(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (u UserRepository) Mute(ctx context.Context, userId uint, time time.Duration) error {
	panic("implement me")
}

func (u UserRepository) UnMute(ctx context.Context, userId uint) error {
	panic("implement me")
}

func (u UserRepository) UpdateLastVisit(ctx context.Context, userId uint) error {
	currentTime := time.Now().UTC()

	err := u.GetMaster().
		WithContext(ctx).
		Table("users").
		Where("id = ?", userId).
		Update("last_visit", currentTime).
		Error
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) ComputeFields(ctx context.Context, user entity.User) (*entity.User, error) {
	return &user, nil
}
