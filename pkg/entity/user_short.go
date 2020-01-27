package entity

import (
	"fmt"
	"github.com/deissh/osu-api-server/pkg"
	"github.com/rs/zerolog/log"
	"time"
)

// UserShort data struct
type UserShort struct {
	ID               uint      `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	PasswordHash     string    `json:"-" db:"password_hash"`
	IsBot            bool      `json:"is_bot" db:"is_bot"`
	IsActive         bool      `json:"is_active" db:"is_active"`
	IsSupporter      bool      `json:"is_supporter" db:"is_supporter"`
	HasSupported     bool      `json:"has_supported" db:"has_supported"`
	SupportLevel     int       `json:"support_level" db:"support_level"`
	PmFriendsOnly    bool      `json:"pm_friends_only" db:"pm_friends_only"`
	AvatarURL        string    `json:"avatar_url" db:"avatar_url"`
	CountryCode      string    `json:"country_code" db:"country_code"`
	DefaultGroup     string    `json:"default_group" db:"default_group"`
	LastVisit        time.Time `json:"last_visit" db:"last_visit"`
	JoinDate         time.Time `json:"join_date" db:"created_at"`
	SupportExpiredAt time.Time `json:"-" db:"support_expired_at"`

	// computed
	IsOnline bool `json:"is_online"`
}

// Compute fields and return error if not successful
func (u *UserShort) Compute() error {
	log.Debug().
		Msg("Computing user fields")

	if err := pkg.Rb.Get(fmt.Sprintf("online_users::%d", u.ID)).Err(); err == nil {
		u.IsOnline = true
	}

	return nil
}
