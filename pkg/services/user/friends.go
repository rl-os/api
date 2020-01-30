package user

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetFriends and all subscription to other users
func GetFriends(userId uint) (*[]entity.UserShort, error) {
	users := make([]entity.UserShort, 0)

	err := pkg.Db.Select(
		&users,
		`SELECT u.id, u.username, u.email, u.is_bot, u.is_active,
       				u.is_supporter, u.has_supported, u.support_level,
       				u.pm_friends_only, u.avatar_url, u.country_code,
       				u.default_group, u.last_visit, u.created_at, u.support_expired_at
				FROM user_relation
				INNER JOIN users u on user_relation.target_id = u.id
				WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("friends not founded")
		return nil, pkg.NewHTTPError(http.StatusNotFound, "user_friends", "Friends not founded.")
	}

	return &users, nil
}

// SetFriend to user
func SetFriend(userId uint, targetId uint) (*[]entity.UserShort, error) {

	return nil, nil
}

func RemoveFriend(userId uint, targetId uint) error {

	return nil
}
