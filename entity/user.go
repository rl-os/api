package entity

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/deissh/go-utils"
	"time"
)

type User struct {
	UserShort

	HasSupported bool `json:"has_supported"`

	Statistics        UserStatistics      `json:"statistics" gorm:"foreignkey:user_id;references:id"`
	MonthlyPlaycounts []MonthlyPlaycounts `json:"monthly_playcounts" gorm:"foreignkey:user_id;references:id"`

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
	Playstyle                        string             `json:"playstyle"`
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
	ReplaysWatchedCounts             int                `json:"replays_watched_counts"`
	ScoresBestCount                  int                `json:"scores_best_count"`
	ScoresFirstCount                 int                `json:"scores_first_count"`
	ScoresRecentCount                int                `json:"scores_recent_count"`
	UnrankedBeatmapsetCount          int                `json:"unranked_beatmapset_count"`
	UserAchievements                 []UserAchievements `json:"user_achievements"`
	RankHistory                      RankHistory        `json:"rank_history"`

	// internal fields
	Mode      string    `json:"-" gorm:"-"`
	CreatedAt time.Time `json:"join_date"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"-"`
}

func (User) TableName() string {
	return "users"
}

// Country with code
type Country struct {
	Id   uint   `json:"-"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func (Country) TableName() string {
	return "countries"
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
