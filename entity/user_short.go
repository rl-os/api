package entity

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	ID           uint   `json:"id"`
	Email        string `json:"-"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
}

// UserShort selector from database
type UserShort struct {
	UserBasic

	IsActive        bool      `json:"is_active"`
	IsBot           bool      `json:"is_bot"`
	IsOnline        bool      `json:"is_online" gorm:"-"`
	IsSupporter     bool      `json:"is_supporter"`
	LastVisit       time.Time `json:"last_visit"`
	PmFriendsOnly   bool      `json:"pm_friends_only"`
	ProfileColour   string    `json:"profile_colour"`
	CountryCode     string    `json:"country_code"`
	Country         Country   `json:"country" gorm:"foreignkey:code;references:country_code"`
	Cover           Cover     `json:"cover"  gorm:"embedded;embeddedPrefix:cover_"`
	CurrentModeRank int       `json:"current_mode_rank"`
	Groups          string    `json:"groups"`
	SupportLevel    int       `json:"support_level"`
	AvatarURL       string    `json:"avatar_url"`
	DefaultGroup    string    `json:"default_group"`

	// internal fields
	Mode      string    `json:"-" gorm:"-"`
	CreatedAt time.Time `json:"join_date"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

func (UserShort) TableName() string {
	return "users"
}

func (u *UserShort) AfterFind(_ *gorm.DB) (err error) {
	if u.LastVisit.Add(time.Minute * 15).After(time.Now()) {
		u.IsOnline = true
	}
	return
}
