package main

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	"github.com/rs/zerolog/log"
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
