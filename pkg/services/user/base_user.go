package user

import (
	"net/http"
	"time"
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/labstack/echo/v4"
)

// BaseUser data struct
type BaseUser struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email		 string `json:"email" db:"email"`
	PasswordHash string `json:"-" db:"password_hash"`

	// todo: вынести данные которые общие у модели пользователя и текущего пользователя

	LastVisit time.Time `json:"last_visit" db:"last_visit"`
	JoinDate  time.Time `json:"join_date" db:"created_at"`
}

// LoginByPassword and return user data such ID
func LoginByPassword(username string, password string) (BaseUser, error) {
	var baseUser BaseUser

	hashed, err := utils.GetHash(password)
	if err != nil {
		return BaseUser{}, echo.NewHTTPError(500, "Getting hash from password error.")
	}

	err = pkg.Db.Get(
		&baseUser,
		`SELECT * FROM users WHERE username = ? AND password_hash = ?`,
				username, hashed,
	)
	if err != nil {
		return BaseUser{}, echo.NewHTTPError(401, "The user credentials were incorrect.")
	}

	return baseUser, nil
}

// Register and return new user
func Register(username string, email string, password string) (BaseUser, error) {
	var baseUser BaseUser

	hashed, err := utils.GetHash(password)
	if err != nil {
		return BaseUser{}, pkg.NewHTTPError(http.StatusInternalServerError, "internal_error", "Getting hash from password error.")
	}

	err = pkg.Db.Get(
		&baseUser,
		`INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING *`,
		username, email, hashed,
	)
	if err != nil {
		return BaseUser{}, pkg.NewHTTPError(http.StatusBadRequest, "create_user_error", "Registration info is are incorrect.")
	}

	return baseUser, nil
}