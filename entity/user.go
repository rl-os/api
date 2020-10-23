package entity

import (
	"encoding/json"
	"gorm.io/gorm"

	"database/sql/driver"
	"github.com/deissh/go-utils"
	"time"
)

type User struct {
	gorm.Model

	// internal field
	Mode string `json:"-" db:"-"`

	ID              uint          `json:"id"`
	IsActive        bool          `json:"is_active"`
	IsBot           bool          `json:"is_bot"`
	IsOnline        bool          `json:"is_online"`
	IsSupporter     bool          `json:"is_supporter"`
	LastVisit       time.Time     `json:"last_visit"`
	PmFriendsOnly   bool          `json:"pm_friends_only"`
	ProfileColour   interface{}   `json:"profile_colour"`
	Username        string        `json:"username"`
	Country         Country       `json:"country"`
	Cover           Cover         `json:"cover"`
	CurrentModeRank int           `json:"current_mode_rank"`
	Groups          []interface{} `json:"groups"`
	SupportLevel    int           `json:"support_level"`
	AvatarURL       string        `json:"avatar_url"`
	CountryCode     string        `json:"country_code"`
	DefaultGroup    string        `json:"default_group"`

	CoverURL string    `json:"cover_url"`
	JoinDate time.Time `json:"join_date"`

	HasSupported bool `json:"has_supported"`

	Discord                          utils.NullString   `json:"discord"`
	Skype                            utils.NullString   `json:"skype"`
	Title                            utils.NullString   `json:"title"`
	TitleURL                         utils.NullString   `json:"title_url"`
	Twitter                          utils.NullString   `json:"twitter"`
	Website                          utils.NullString   `json:"website"`
	Interests                        utils.NullString   `json:"interests"`
	Kudosu                           Kudosu             `json:"kudosu"`
	Location                         utils.NullString   `json:"location"`
	MaxBlocks                        int                `json:"max_blocks"`
	MaxFriends                       int                `json:"max_friends"`
	Occupation                       string             `json:"occupation"`
	Playmode                         string             `json:"playmode"`
	Playstyle                        []string           `json:"playstyle"`
	PostCount                        int                `json:"post_count"`
	ProfileOrder                     []string           `json:"profile_order"`
	AccountHistory                   []interface{}      `json:"account_history"`
	ActiveTournamentBanner           []interface{}      `json:"active_tournament_banner"`
	Badges                           []interface{}      `json:"badges"`
	BeatmapPlaycountsCount           int                `json:"beatmap_playcounts_count"`
	FavouriteBeatmapsetCount         int                `json:"favourite_beatmapset_count"`
	FollowerCount                    int                `json:"follower_count"`
	GraveyardBeatmapsetCount         int                `json:"graveyard_beatmapset_count"`
	LovedBeatmapsetCount             int                `json:"loved_beatmapset_count"`
	Page                             Page               `json:"page"`
	PreviousUsernames                []string           `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int                `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCounts             []interface{}      `json:"replays_watched_counts"`
	ScoresBestCount                  int                `json:"scores_best_count"`
	ScoresFirstCount                 int                `json:"scores_first_count"`
	ScoresRecentCount                int                `json:"scores_recent_count"`
	Statistics                       Statistics         `json:"statistics"`
	UnrankedBeatmapsetCount          int                `json:"unranked_beatmapset_count"`
	UserAchievements                 []UserAchievements `json:"user_achievements"`
	RankHistory                      RankHistory        `json:"rank_history"`

	MonthlyPlaycounts []MonthlyPlaycountsField `json:"monthly_playcounts" gorm:"foreignkey:UserName;references:name"`
}

// Country with code
type Country struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

// Cover file url
type Cover struct {
	CustomURL interface{} `json:"custom_url"`
	URL       string      `json:"url"`
	ID        string      `json:"id"`
}

func (c Cover) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Cover) Scan(value interface{}) error {
	result := Cover{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// Kudosu value in user profile
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}

func (c Kudosu) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Kudosu) Scan(value interface{}) error {
	result := Kudosu{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// Page customization
type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

func (c Page) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Page) Scan(value interface{}) error {
	result := Page{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// UserAchievements with datetime
type UserAchievements struct {
	AchievedAt    time.Time `json:"achieved_at" db:"created_at"`
	AchievementID int       `json:"achievement_id" db:"achievement_id"`
}

func (c UserAchievements) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *UserAchievements) Scan(value interface{}) error {
	result := UserAchievements{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// RankHistory recor
type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

func (c RankHistory) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *RankHistory) Scan(value interface{}) error {
	result := RankHistory{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// GetShort version of user
func (u *User) GetShort() *UserShort {
	return &UserShort{
		ID:            u.ID,
		Username:      u.Username,
		IsBot:         u.IsBot,
		IsActive:      u.IsActive,
		IsSupporter:   u.IsSupporter,
		SupportLevel:  u.SupportLevel,
		PmFriendsOnly: u.PmFriendsOnly,
		AvatarURL:     u.AvatarURL,
		CountryCode:   u.CountryCode,
		DefaultGroup:  u.DefaultGroup,
		LastVisit:     u.LastVisit,
		IsOnline:      u.IsOnline,
	}
}
