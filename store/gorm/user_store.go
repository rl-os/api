package sql

import (
	"context"
	"fmt"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
	"gorm.io/gorm"
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
		Where("id = ?", userId).
		First(&user).
		Preload(clause.Associations).
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
		Where("id = ?", userId).
		First(&user).
		Error

	if err != nil {
		return nil, errors.WithCause("user_get", 404, "user not found", err)
	}

	return &user, nil
}

func (u UserStore) Create(ctx context.Context, name, email, pwd string) (*entity.User, error) {
	hashed, err := utils.GetHash(pwd)
	if err != nil {
		return nil, errors.WithCause("user_create", 500, "hashing password", err)
	}

	user := entity.UserAuthBase{
		UserShort: entity.UserShort{
			Username: name,
		},
		Email:        email,
		PasswordHash: hashed,
	}

	now := time.Now()

	err = u.GetMaster().Transaction(func(tx *gorm.DB) error {
		err = tx.WithContext(ctx).
			Create(&user).
			Error

		if err != nil {
			return err
		}

		// creating default records
		tx.WithContext(ctx).
			Create(entity.MonthlyPlaycounts{
				UserId:    user.ID,
				StartDate: fmt.Sprintf("%02d-%02d-01", now.Year(), now.Month()),
				Count:     0,
			})

		tx.WithContext(ctx).
			Create(entity.Statistics{
				UserId: user.ID,
				Level: entity.Level{
					Current:  1,
					Progress: 0,
				},
			})

		return nil
	})

	if err != nil {
		return nil, errors.WithCause("user_create", 500, "creating user", err)
	}

	return u.User().Get(ctx, user.ID, "")
}

func (u UserStore) Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error) {
	panic("implement me")
}

func (u UserStore) Activate(ctx context.Context, userId uint) error {
	err := u.GetMaster().
		WithContext(ctx).
		Model(&entity.User{}).
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
		Model(&entity.User{}).
		Where("id = ?", userId).
		Update("last_visit", currentTime).
		Error
	if err != nil {
		return errors.WithCause("user_update_online", 404, "User not found", err)
	}

	return nil
}

func (u UserStore) ComputeFields(ctx context.Context, user entity.User) (*entity.User, error) {
	//tx, _ := u.GetMaster().BeginTxx(ctx, nil)
	//defer tx.Commit()
	//
	//// =========================
	//// followers
	//_ = tx.GetContext(
	//	ctx,
	//	&user.FollowerCount,
	//	`SELECT count(*) FROM user_relation WHERE target_id = $1`,
	//	user.ID,
	//)
	//
	//// =========================
	//// favourite beatmapsets count
	//_ = tx.GetContext(
	//	ctx,
	//	&user.FavouriteBeatmapsetCount,
	//	`SELECT count(*) FROM favouritemaps WHERE user_id = $1`,
	//	user.ID,
	//)
	//
	//// =========================
	//// getting MonthlyPlayCounts
	//user.MonthlyPlaycounts = make([]entity.MonthlyPlaycounts, 0)
	//err := tx.SelectContext(
	//	ctx,
	//	&user.MonthlyPlaycounts,
	//	`SELECT playcount, year_month FROM user_month_playcount WHERE user_id = $1`,
	//	user.ID,
	//)
	//if err != nil {
	//	return nil, errors.WithCause("user_compute", 500, "getting MonthlyPlayCounts", err)
	//}
	//
	//// =========================
	//// getting RankHistory
	//// ranks := make([]int, 50)
	//user.RankHistory = entity.RankHistory{
	//	Mode: user.Mode,
	//	// todo: https://github.com/ppy/osu-web/blob/7d14d741454e2c8ef5c90b9bfa90213f61020b06/app/Models/RankHistory.php#L119
	//	// очень странный формат, но нужно как разобраться
	//	// сейчас оставил так, когда будет время исправить
	//	Data: []int{1, 1, 2, 3, 1, 1, 1, 1, 4, 4, 5, 1, 1, 1, 1, 1, 11, 1, 1, 1, 2, 1, 1, 1},
	//}
	//
	//// =========================
	//// getting UserAchievements
	//user.UserAchievements = make([]entity.UserAchievements, 0)
	//err = tx.SelectContext(
	//	ctx,
	//	&user.UserAchievements,
	//	`SELECT achievement_id, created_at FROM user_achievements WHERE user_id = $1`,
	//	user.ID,
	//)
	//if err != nil {
	//	return nil, errors.WithCause("user_compute", 500, "getting UserAchievements", err)
	//}
	//
	//// =========================
	//// getting UserStatistics
	//err = tx.GetContext(
	//	ctx,
	//	&user.Statistics,
	//	`SELECT
	//   		json_build_object('current', level_current, 'progress', level_progress) as level,
	//   		json_build_object('ss', grade_counts_ss, 'ssh', grade_counts_ssh, 'sh', grade_counts_sh,
	//   		    			  's', grade_counts_s, 'a', grade_counts_a) as grade_counts,
	//   		pp, ranked_score, hit_accuracy, play_count, play_time, total_score,
	//   		total_hits, maximum_combo, replays_watched_by_others, is_ranked
	//	FROM user_statistics
	//	WHERE user_id = $1`,
	//	user.ID,
	//)
	//if err != nil {
	//	return nil, errors.WithCause("user_compute", 500, "getting UserStatistics", err)
	//}
	//
	//// =========================
	//// getting UserRank
	//err = tx.GetContext(
	//	ctx,
	//	&user.Statistics.Rank,
	//	`SELECT country, global
	//	FROM (
	//		SELECT row_number() over (PARTITION BY t.country_code ORDER BY t.pp DESC) as country,
	//			row_number() over (ORDER BY t.pp DESC) as global,
	//			t.user_id
	//		FROM (
	//			SELECT us.user_id, us.pp, u.country_code
	//			FROM user_statistics us
	//			JOIN users u on us.user_id = u.id
	//			WHERE u.is_active = true
	//		) as t
	//	) as rt
	//	WHERE rt.user_id = $1;`,
	//	user.ID,
	//)
	//if err != nil {
	//	return nil, errors.WithCause("user_compute", 500, "getting UserRank", err)
	//}
	//user.Statistics.PpRank = user.Statistics.Rank.Global
	//
	//return &user, err
	return &user, nil
}
