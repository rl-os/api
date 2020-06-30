package entity

import (
	"database/sql/driver"
	"github.com/deissh/go-utils"
	"time"
)

// UserShort data struct
type UserShort struct {
	ID               uint      `json:"id" db:"id"`
	Username         string    `json:"username" db:"username"`
	Email            string    `json:"email" db:"email"`
	PasswordHash     string    `json:"-" db:"password_hash"`
	IsOnline         bool      `json:"is_online" db:"check_online"`
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
}

type UserShortField UserShort

func (c UserShortField) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *UserShortField) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }
