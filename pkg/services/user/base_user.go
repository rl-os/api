package user

import (
	"github.com/deissh/osu-api-server/pkg"
	"github.com/deissh/osu-api-server/pkg/utils"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

// BaseUser data struct
type BaseUser struct {
	ID            uint      `json:"id" db:"id"`
	Username      string    `json:"username" db:"username"`
	Email         string    `json:"email" db:"email"`
	PasswordHash  string    `json:"-" db:"password_hash"`
	IsBot         bool      `json:"is_bot" db:"is_bot"`
	IsActive      bool      `json:"is_active" db:"is_active"`
	IsSupporter   bool      `json:"is_supporter" db:"is_supporter"`
	HasSupported  bool      `json:"has_supported" db:"has_supported"`
	SupportLevel  int       `json:"support_level" db:"support_level"`
	PmFriendsOnly bool      `json:"pm_friends_only" db:"pm_friends_only"`
	AvatarURL     string    `json:"avatar_url" db:"avatar_url"`
	CountryCode   string    `json:"country_code" db:"country_code"`
	DefaultGroup  string    `json:"default_group" db:"default_group"`
	LastVisit     time.Time `json:"last_visit" db:"last_visit"`
	JoinDate      time.Time `json:"join_date" db:"created_at"`

	// computed
	IsOnline bool `json:"is_online"`
}

// Compute fields and return error if not successful
func (this *BaseUser) Compute() error {
	this.IsOnline = pkg.Rb.SIsMember("online_users", this.ID).Val()

	return nil
}

// LoginByPassword and return user data such ID
func LoginByPassword(username string, password string) (BaseUser, error) {
	var baseUser BaseUser

	err := pkg.Db.Get(
		&baseUser,
		`SELECT * FROM users WHERE username = $1 OR email = $1`,
		username,
	)
	if err != nil {
		log.Debug().Msg("login uncorrect")
		return BaseUser{}, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
	}

	if ok := utils.CompareHash(baseUser.PasswordHash, password); !ok {
		log.Debug().Msg("password uncorrect")
		return BaseUser{}, pkg.NewHTTPError(http.StatusUnauthorized, "user_login_error", "The user credentials were incorrect.")
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
