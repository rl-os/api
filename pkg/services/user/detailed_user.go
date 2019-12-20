package user

import (
	"github.com/deissh/osu-api-server/pkg"
	"net/http"
)

type DetailedUser struct {
	BaseUser
	Detail
}

// GetDetailedUser and compute some fields
func GetDetailedUser(id uint) (DetailedUser, error) {
	var user DetailedUser

	err := pkg.Db.Get(
		&user,
		`SELECT users.id, username, email, last_visit, created_at, is_bot, is_active, is_supporter, has_supported,
       					support_level, pm_friends_only, avatar_url, country_code, default_group, can_moderate,
       					interests, occupation,title, location, twitter, lastfm, skype, website, discord, playmode,
       					playstyle, cover_url, max_blocks, max_friends
				FROM users
    			INNER JOIN user_details ON users.id = user_details.user_id
				WHERE users.id = $1`,
		id,
	)
	if err != nil {
		return DetailedUser{}, pkg.NewHTTPError(http.StatusUnauthorized, "user_not_founded", "User not founded.")
	}

	err = user.Compute()
	if err != nil {
		return DetailedUser{}, err
	}

	return user, nil
}
