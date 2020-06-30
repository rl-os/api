package main

import (
	"fmt"
	"github.com/deissh/rl/api/pkg"
	"github.com/deissh/rl/api/pkg/entity"
	"github.com/rs/zerolog/log"
	"time"
)

// ExpireSupporter and send notification to user
func ExpireSupporter() {
	users := make([]entity.User, 0)
	log.Info().Msg("Searching and update user")

	err := pkg.Db.
		Select(&users,
			`UPDATE users
			SET is_supporter = false
			WHERE is_supporter = true AND support_expired_at < CURRENT_TIMESTAMP
			RETURNING *;`,
		)
	if err != nil {
		log.Error().
			Err(err).
			Send()
	}

	log.Info().
		Int("user_count", len(users)).
		Msg("Update expired supporters done")
}

func NewMonthPlaycount() {
	userIds := make([]int, 0)

	err := pkg.Db.
		Select(&userIds,
			`SELECT id
			FROM users`,
		)
	if err != nil {
		log.Error().
			Err(err).
			Send()
	}

	now := time.Now()
	yearMonth := fmt.Sprintf("%02d-%02d-01", now.Year(), now.Month())

	tx := pkg.Db.MustBegin()
	{
		// remove if already created (bugs or user just registered)
		tx.MustExec(`DELETE FROM user_month_playcount WHERE year_month = $1`, yearMonth)

		// creat new month_playcount with each user_id
		for id := range userIds {
			tx.MustExec(
				`INSERT INTO user_month_playcount (user_id, playcount, year_month)
				VALUES ($1, 0, $2)`,
				id,
				yearMonth,
			)
		}
	}

	if err = tx.Commit(); err != nil {
		log.Error().
			Err(err).
			Send()
	}
}
