package user

import (
	"github.com/deissh/osu-lazer/api/pkg"
	"github.com/deissh/osu-lazer/api/pkg/common/utils"
	"github.com/deissh/osu-lazer/api/pkg/entity"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

var modes = []string{"std", "mania", "catch", "taiko"}

// GetUser and compute some fields
func GetUser(id uint, mode string) (*entity.User, error) {
	var user entity.User

	if !utils.ContainsString(modes, mode) {
		mode = "std"
	}

	err := pkg.Db.Get(
		&user,
		`SELECT *, check_online(last_visit)
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
	var user entity.User

	err := pkg.Db.Get(
		&user,
		`SELECT *, check_online(last_visit)
		FROM users
		WHERE username = $1 OR email = $1`,
		username,
	)
	if err != nil {
		log.Debug().
			Err(err).
			Msg("login uncorrect")
		return nil, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
	}

	if ok := utils.CompareHash(user.PasswordHash, password); !ok {
		log.Debug().Msg("password uncorrect")
		return nil, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
	}

	return user.GetShort(), nil
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
			`INSERT INTO users (username, email, password_hash)
			VALUES ($1, $2, $3)
			RETURNING *, check_online(last_visit)`,
			username, email, hashed,
		)
		if err != nil {
			log.Err(err).Send()
			return nil, pkg.NewHTTPError(http.StatusBadRequest, "create_user_error", "Registration info is are incorrect.")
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, pkg.NewHTTPError(http.StatusBadRequest, "create_user_error", "Registration info is are incorrect.")
	}

	return &baseUser, nil
}

// UpdateLastVisit and set current time
func UpdateLastVisit(userId uint) (*entity.User, error) {
	var user entity.User
	currentTime := time.Now().UTC()

	err := pkg.Db.Get(
		&user,
		`UPDATE users
		SET last_visit = $1
		WHERE id = $2
		RETURNING *, check_online(last_visit)`,
		currentTime,
		userId,
	)

	return &user, err
}
