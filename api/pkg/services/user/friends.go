package user

import (
	"github.com/deissh/rl/api/pkg"
	"github.com/deissh/rl/api/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetSubscriptions and all subscription to other users
func GetSubscriptions(userId uint) (*[]entity.UserShort, error) {
	users := make([]entity.UserShort, 0)

	err := pkg.Db.Select(
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
		log.Error().
			Err(err).
			Msg("friends not founded")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "user_friends", "Friends not founded.")
	}

	return &users, nil
}

// GetFriends function return all user friends
func GetFriends(userId uint) (*[]entity.UserShort, error) {
	users := make([]entity.UserShort, 0)

	err := pkg.Db.Select(
		&users,
		`SELECT u.id, u.username, u.email, u.is_bot, u.is_active,
       				u.is_supporter, u.has_supported, u.support_level,
       				u.pm_friends_only, u.avatar_url, u.country_code,
       				u.default_group, u.last_visit, u.created_at,
       				u.support_expired_at, check_online(u.last_visit)
				FROM user_relation ur
				JOIN users u on ur.target_id = u.id
				WHERE ur.user_id = $1
				    AND EXISTS(SELECT 1 from user_relation ur2 WHERE ur2.user_id = ur.target_id)`,
		userId,
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("friends not founded")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "user_friends", "Friends not founded.")
	}

	return &users, nil
}

// SetFriend to user
func SetFriend(userId uint, targetId uint) error {
	_, err := pkg.Db.Query(
		`INSERT INTO user_relation (user_id, target_id) VALUES ($1, $2)`,
		userId, targetId,
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("friend not added")
		return pkg.NewHTTPError(http.StatusNotFound, "user_friends", "Friend not added.")
	}

	return nil
}

// RemoveFriend from subscriptions
func RemoveFriend(userId uint, targetId uint) error {
	_, err := pkg.Db.Query(
		`DELETE FROM user_relation WHERE user_id = $1 AND target_id = $2`,
		userId, targetId,
	)
	if err != nil {
		log.Error().
			Err(err).
			Msg("friend not added")
		return pkg.NewHTTPError(http.StatusNotFound, "user_friends", "Friend not added.")
	}

	return nil
}
