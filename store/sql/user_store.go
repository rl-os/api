package sql

import (
	"context"
	"fmt"
	"github.com/deissh/go-utils"
	"github.com/rl-os/api/entity"
	"github.com/rl-os/api/errors"
	"github.com/rl-os/api/store"
	"time"
)

var modes = []string{"std", "mania", "catch", "taiko"}

type UserStore struct {
	SqlStore
}

func newSqlUserStore(sqlStore SqlStore) store.User {
	return &UserStore{sqlStore}
}

func (u UserStore) GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error) {
	var baseUser entity.User

	err := u.GetMaster().GetContext(
		ctx,
		&baseUser,
		`SELECT *, check_online(last_visit)
		FROM users
		WHERE username = $1 OR email = $1`,
		login,
	)
	if err != nil {
		return nil, errors.WithCause("user_get", 401, "user credentials were incorrect", err)
	}

	if ok := utils.CompareHash(baseUser.PasswordHash, pwd); !ok {
		return nil, errors.WithCause("user_get", 401, "user credentials were incorrect", err)
	}

	return baseUser.GetShort(), nil
}

func (u UserStore) Get(ctx context.Context, userId uint, mode string) (*entity.User, error) {
	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}
	user := entity.User{Mode: mode}

	err := u.GetMaster().GetContext(
		ctx,
		&user,
		`SELECT u.*, check_online(last_visit),
			json_build_object('code', c.code, 'name', c.name) as country
		FROM users u
		INNER JOIN countries c ON c.code = u.country_code
		WHERE u.id = $1`,
		userId,
	)
	if err != nil {
		return nil, errors.WithCause(
			"user_get", 404, "user not found", err,
		)
	}

	return u.ComputeFields(ctx, user)
}

func (u UserStore) GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error) {
	panic("implement me")
}

func (u UserStore) Create(ctx context.Context, name, email, pwd string) (*entity.User, error) {
	var baseUser entity.User

	now := time.Now()

	hashed, err := utils.GetHash(pwd)
	if err != nil {
		return nil, errors.WithCause(
			"user_create", 500, "hashing password", err,
		)
	}

	tx := u.GetMaster().MustBeginTx(ctx, nil)
	{ // begin transaction
		err = tx.GetContext(
			ctx,
			&baseUser,
			`INSERT INTO users (username, email, password_hash)
			VALUES ($1, $2, $3)
			RETURNING *, check_online(last_visit)`,
			name, email, hashed,
		)
		if err != nil {
			return nil, errors.WithCause("user_create", 500, "creating user", err)
		}

		// creating default records
		tx.MustExecContext(
			ctx,
			`INSERT INTO user_month_playcount (user_id, playcount, year_month) VALUES ($1, 0, $2)`,
			baseUser.ID,
			fmt.Sprintf("%02d-%02d-01", now.Year(), now.Month()),
		)
		tx.MustExecContext(
			ctx,
			`INSERT INTO user_statistics (user_id) VALUES ($1)`,
			baseUser.ID,
		)
	} // end transaction

	err = tx.Commit()
	if err != nil {
		return nil, errors.WithCause("user_create", 500, "transaction not complited", err)
	}

	return u.User().ComputeFields(ctx, baseUser)
}

func (u UserStore) Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error) {
	panic("implement me")
}

func (u UserStore) Activate(ctx context.Context, userId uint) error {
	_, err := u.GetMaster().ExecContext(
		ctx,
		`UPDATE users
		SET is_active = true
		WHERE is_active = false AND id = $1`,
		userId,
	)
	if err != nil {
		return errors.WithCause("user_activate", 404, "User not found", err)
	}

	return nil
}

func (u UserStore) Deactivate(ctx context.Context, userId uint) error {
	_, err := u.GetMaster().ExecContext(
		ctx,
		`UPDATE users
		SET is_active = false 
		WHERE is_active = true AND id = $1`,
		userId,
	)
	if err != nil {
		return errors.WithCause("user_deactivate", 404, "User not found", err)
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

	_, err := u.GetMaster().ExecContext(
		ctx,
		`UPDATE users
		SET last_visit = $1
		WHERE id = $2`,
		currentTime,
		userId,
	)
	if err != nil {
		return errors.WithCause("user_update_online", 404, "User not found", err)
	}
	return nil
}

func (u UserStore) ComputeFields(ctx context.Context, user entity.User) (*entity.User, error) {
	tx, _ := u.GetMaster().BeginTxx(ctx, nil)
	defer tx.Commit()

	// =========================
	// followers
	_ = tx.GetContext(
		ctx,
		&user.FollowerCount,
		`SELECT count(*) FROM user_relation WHERE target_id = $1`,
		user.ID,
	)

	// =========================
	// favourite beatmapsets count
	_ = tx.GetContext(
		ctx,
		&user.FavouriteBeatmapsetCount,
		`SELECT count(*) FROM favouritemaps WHERE user_id = $1`,
		user.ID,
	)

	// =========================
	// getting MonthlyPlayCounts
	user.MonthlyPlaycounts = make([]entity.MonthlyPlaycounts, 0)
	err := tx.SelectContext(
		ctx,
		&user.MonthlyPlaycounts,
		`SELECT playcount, year_month FROM user_month_playcount WHERE user_id = $1`,
		user.ID,
	)
	if err != nil {
		return nil, errors.WithCause("user_compute", 500, "getting MonthlyPlayCounts", err)
	}

	// =========================
	// getting RankHistory
	// ranks := make([]int, 50)
	user.RankHistory = entity.RankHistory{
		Mode: user.Mode,
		// todo: https://github.com/ppy/osu-web/blob/7d14d741454e2c8ef5c90b9bfa90213f61020b06/app/Models/RankHistory.php#L119
		// очень странный формат, но нужно как разобраться
		// сейчас оставил так, когда будет время исправить
		Data: []int{1, 1, 2, 3, 1, 1, 1, 1, 4, 4, 5, 1, 1, 1, 1, 1, 11, 1, 1, 1, 2, 1, 1, 1},
	}

	// =========================
	// getting UserAchievements
	user.UserAchievements = make([]entity.UserAchievements, 0)
	err = tx.SelectContext(
		ctx,
		&user.UserAchievements,
		`SELECT achievement_id, created_at FROM user_achievements WHERE user_id = $1`,
		user.ID,
	)
	if err != nil {
		return nil, errors.WithCause("user_compute", 500, "getting UserAchievements", err)
	}

	// =========================
	// getting UserStatistics
	err = tx.GetContext(
		ctx,
		&user.Statistics,
		`SELECT
       		json_build_object('current', level_current, 'progress', level_progress) as level,
       		json_build_object('ss', grade_counts_ss, 'ssh', grade_counts_ssh, 'sh', grade_counts_sh,
       		    			  's', grade_counts_s, 'a', grade_counts_a) as grade_counts,
       		pp, ranked_score, hit_accuracy, play_count, play_time, total_score,
       		total_hits, maximum_combo, replays_watched_by_others, is_ranked
		FROM user_statistics
		WHERE user_id = $1`,
		user.ID,
	)
	if err != nil {
		return nil, errors.WithCause("user_compute", 500, "getting UserStatistics", err)
	}

	// =========================
	// getting UserRank
	err = tx.GetContext(
		ctx,
		&user.Statistics.Rank,
		`SELECT country, global
		FROM (
			SELECT row_number() over (PARTITION BY t.country_code ORDER BY t.pp DESC) as country,
				row_number() over (ORDER BY t.pp DESC) as global,
				t.user_id
			FROM (
				SELECT us.user_id, us.pp, u.country_code
				FROM user_statistics us
				JOIN users u on us.user_id = u.id
				WHERE u.is_active = true
			) as t
		) as rt
		WHERE rt.user_id = $1;`,
		user.ID,
	)
	if err != nil {
		return nil, errors.WithCause("user_compute", 500, "getting UserRank", err)
	}
	user.Statistics.PpRank = user.Statistics.Rank.Global

	return &user, err
}
