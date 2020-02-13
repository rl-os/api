package entity

import (
	"github.com/deissh/osu-lazer/api/pkg/common/utils"
	"github.com/lib/pq"
	"time"
)

// User data structure
type User struct {
	UserShort

	CanModerate  bool             `json:"can_moderate" db:"can_moderate"`
	Interests    utils.NullString `json:"interests" db:"interests"`
	Occupation   string           `json:"occupation" db:"occupation"`
	Title        utils.NullString `json:"title" db:"title"`
	Location     utils.NullString `json:"location" db:"location"`
	Twitter      utils.NullString `json:"twitter" db:"twitter"`
	Lastfm       utils.NullString `json:"lastfm" db:"lastfm"`
	Skype        utils.NullString `json:"skype" db:"skype"`
	Website      utils.NullString `json:"website" db:"website"`
	Discord      utils.NullString `json:"discord" db:"discord"`
	Playstyle    pq.StringArray   `json:"playstyle" db:"playstyle"`
	Playmode     string           `json:"playmode" db:"playmode"`
	ProfileOrder pq.StringArray   `json:"profile_order" db:"profile_order"`
	CoverURL     string           `json:"cover_url" db:"cover_url"`
	MaxBlocks    int              `json:"max_blocks" db:"max_blocks"`
	MaxFriends   int              `json:"max_friends" db:"max_friends"`

	Cover         Cover       `json:"cover"`
	Kudosu        Kudosu      `json:"kudosu"`
	Page          Page        `json:"page"`
	Statistics    Statistics  `json:"statistics"`
	RankHistory   RankHistory `json:"rankHistory"`
	ProfileColour interface{} `json:"profile_colour"`

	// joins
	AccountHistory         []interface{}       `json:"account_history"`
	ActiveTournamentBanner []interface{}       `json:"active_tournament_banner"`
	Badges                 []interface{}       `json:"badges"`
	MonthlyPlaycounts      []MonthlyPlaycounts `json:"monthly_playcounts"`
	PreviousUsernames      []string            `json:"previous_usernames"`
	ReplaysWatchedCounts   []interface{}       `json:"replays_watched_counts"`
	UserAchievements       []UserAchievements  `json:"user_achievements"`

	// computed
	IsGmt                            bool `json:"is_gmt" db:"is_gmt"`
	IsNat                            bool `json:"is_nat" db:"is_nat"`
	IsBng                            bool `json:"is_bng" db:"is_bng"`
	IsFullBn                         bool `json:"is_full_bn" db:"is_full_bn"`
	IsLimitedBn                      bool `json:"is_limited_bn" db:"is_limited_bn"`
	FavouriteBeatmapsetCount         int  `json:"favourite_beatmapset_count" db:"favourite_beatmapset_count"`
	FollowerCount                    int  `json:"follower_count" db:"follower_count"`
	GraveyardBeatmapsetCount         int  `json:"graveyard_beatmapset_count" db:"graveyard_beatmapset_count"`
	LovedBeatmapsetCount             int  `json:"loved_beatmapset_count" db:"loved_beatmapset_count"`
	RankedAndApprovedBeatmapsetCount int  `json:"ranked_and_approved_beatmapset_count" db:"ranked_and_approved_beatmapset_count"`
	ScoresFirstCount                 int  `json:"scores_first_count" db:"scores_first_count"`
	UnrankedBeatmapsetCount          int  `json:"unranked_beatmapset_count" db:"unranked_beatmapset_count"`
	PostCount                        int  `json:"post_count"`
}

// Country with code
type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Cover file url
type Cover struct {
	CustomURL interface{} `json:"custom_url"`
	URL       string      `json:"url"`
	ID        string      `json:"id"`
}

// Kudosu value in user profile
type Kudosu struct {
	Total     int `json:"total"`
	Available int `json:"available"`
}

// MonthlyPlaycounts record
type MonthlyPlaycounts struct {
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

// Page customization
type Page struct {
	HTML string `json:"html"`
	Raw  string `json:"raw"`
}

// Level progress and current value in user profile
type Level struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

// GradeCounts data
type GradeCounts struct {
	Ss  int `json:"ss"`
	SSH int `json:"ssh"`
	S   int `json:"s"`
	Sh  int `json:"sh"`
	A   int `json:"a"`
}

// Rank in world and in the user country
type Rank struct {
	Global  int `json:"global"`
	Country int `json:"country"`
}

// Statistics in profile
type Statistics struct {
	Level                  Level       `json:"level"`
	Pp                     float64     `json:"pp"`
	PpRank                 int         `json:"pp_rank"`
	RankedScore            int         `json:"ranked_score"`
	HitAccuracy            float64     `json:"hit_accuracy"`
	PlayCount              int         `json:"play_count"`
	PlayTime               int         `json:"play_time"`
	TotalScore             int         `json:"total_score"`
	TotalHits              int         `json:"total_hits"`
	MaximumCombo           int         `json:"maximum_combo"`
	ReplaysWatchedByOthers int         `json:"replays_watched_by_others"`
	IsRanked               bool        `json:"is_ranked"`
	GradeCounts            GradeCounts `json:"grade_counts"`
	Rank                   Rank        `json:"rank"`
}

// UserAchievements with datetime
type UserAchievements struct {
	AchievedAt    time.Time `json:"achieved_at"`
	AchievementID int       `json:"achievement_id"`
}

// RankHistory recor
type RankHistory struct {
	Mode string `json:"mode"`
	Data []int  `json:"data"`
}

// GetShort version of user
func (u *User) GetShort() *UserShort {
	return &UserShort{
		ID:            u.ID,
		Username:      u.Username,
		Email:         u.Email,
		PasswordHash:  u.PasswordHash,
		IsBot:         u.IsBot,
		IsActive:      u.IsActive,
		IsSupporter:   u.IsSupporter,
		HasSupported:  u.HasSupported,
		SupportLevel:  u.SupportLevel,
		PmFriendsOnly: u.PmFriendsOnly,
		AvatarURL:     u.AvatarURL,
		CountryCode:   u.CountryCode,
		DefaultGroup:  u.DefaultGroup,
		LastVisit:     u.LastVisit,
		JoinDate:      u.JoinDate,
		IsOnline:      u.IsOnline,
	}
}
