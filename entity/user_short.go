package entity

import (
	"database/sql/driver"
	"github.com/deissh/go-utils"
	"time"
)

// UserShort selector from database
type UserShort struct {
	ID              uint      `json:"id"`
	IsActive        bool      `json:"is_active"`
	IsBot           bool      `json:"is_bot"`
	IsOnline        bool      `json:"is_online"`
	IsSupporter     bool      `json:"is_supporter"`
	LastVisit       time.Time `json:"last_visit"`
	PmFriendsOnly   bool      `json:"pm_friends_only"`
	ProfileColour   string    `json:"profile_colour"`
	Username        string    `json:"username"`
	CountryCode     string    `json:"country_code"`
	Country         Country   `json:"country" gorm:"foreignkey:code;references:country_code"`
	Cover           Cover     `json:"cover"`
	CurrentModeRank int       `json:"current_mode_rank"`
	Groups          string    `json:"groups"`
	SupportLevel    int       `json:"support_level"`
	AvatarURL       string    `json:"avatar_url"`
	DefaultGroup    string    `json:"default_group"`
}

type UserShortField UserShort

func (c UserShortField) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *UserShortField) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type UserAuthBase struct {
	UserShort

	Email        string `json:"-"`
	PasswordHash string `json:"-"`
}
