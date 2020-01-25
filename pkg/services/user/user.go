package user

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/entity"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/rs/zerolog/log"
	"net/http"
)

// GetUser and compute some fields
func GetUser(id uint, mode string) (*entity.User, error) {
	var user entity.User

	log.Debug().
		Uint("id", id).
		Str("mode", mode).
		Msg("Get detailed user")

	err := pkg.Db.Get(
		&user,
		`SELECT users.id, username, email, last_visit, created_at, is_bot, is_active, is_supporter, has_supported,
       					support_level, pm_friends_only, avatar_url, country_code, default_group, can_moderate,
       					interests, occupation,title, location, twitter, lastfm, skype, website, discord, playmode,
       					playstyle, cover_url, max_blocks, max_friends
				FROM users
				WHERE users.id = $1`,
		id,
	)
	if err != nil {
		return nil, pkg.NewHTTPError(http.StatusNotFound, "user_not_founded", "User not founded.")
	}

	// todo: getting stats by mode

	err = user.Compute()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser and compute some fields
func GetUserShort(id uint, mode string) (*entity.UserShort, error) {
	var user entity.UserShort

	log.Debug().
		Uint("id", id).
		Str("mode", mode).
		Msg("Get detailed user")

	err := pkg.Db.Get(
		&user,
		`SELECT users.id, username, email, last_visit, created_at, is_bot, is_active, is_supporter, has_supported,
       					support_level, pm_friends_only, avatar_url, country_code, default_group, last_visit, created_at
				FROM users
				WHERE users.id = $1`,
		id,
	)
	if err != nil {
		return nil, pkg.NewHTTPError(http.StatusNotFound, "user_not_founded", "User not founded.")
	}

	// todo: getting stats by mode

	err = user.Compute()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// LoginByPassword and return user data such ID
func LoginByPassword(username string, password string) (*entity.UserShort, error) {
	var baseUser entity.UserShort

	err := pkg.Db.Get(
		&baseUser,
		`SELECT * FROM users WHERE username = $1 OR email = $1`,
		username,
	)
	if err != nil {
		log.Debug().Msg("login uncorrect")
		return nil, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
	}

	if ok := utils.CompareHash(baseUser.PasswordHash, password); !ok {
		log.Debug().Msg("password uncorrect")
		return nil, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
	}

	return &baseUser, nil
}

// Register and return new user
func Register(username string, email string, password string) (*entity.User, error) {
	var baseUser entity.User

	hashed, err := utils.GetHash(password)
	if err != nil {
		return nil, pkg.NewHTTPError(http.StatusInternalServerError, "internal_error", "Getting hash from password error.")
	}

	tx := pkg.Db.MustBegin()
	{
		err = tx.Get(
			&baseUser,
			`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING *`,
			username, email, hashed,
		)
		if err != nil {
			return nil, pkg.NewHTTPError(http.StatusBadRequest, "create_user_error", "Registration info is are incorrect.")
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "create_user_error", "Registration info is are incorrect.")
	}

	return &baseUser, nil
}
