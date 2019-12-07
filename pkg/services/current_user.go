package services

import (
	"time"
)

// CurrentUser data structure
type CurrentUser struct {
	ID                               int                 `json:"id"`
	Username                         string              `json:"username"`
	JoinDate                         time.Time           `json:"join_date"`
	Country                          Country             `json:"country"`
	AvatarURL                        string              `json:"avatar_url"`
	IsSupporter                      bool                `json:"is_supporter"`
	HasSupported                     bool                `json:"has_supported"`
	IsGmt                            bool                `json:"is_gmt"`
	IsNat                            bool                `json:"is_nat"`
	IsBng                            bool                `json:"is_bng"`
	IsFullBn                         bool                `json:"is_full_bn"`
	IsLimitedBn                      bool                `json:"is_limited_bn"`
	IsBot                            bool                `json:"is_bot"`
	IsActive                         bool                `json:"is_active"`
	CanModerate                      bool                `json:"can_moderate"`
	Interests                        interface{}         `json:"interests"`
	Occupation                       string              `json:"occupation"`
	Title                            interface{}         `json:"title"`
	Location                         interface{}         `json:"location"`
	LastVisit                        time.Time           `json:"last_visit"`
	IsOnline                         bool                `json:"is_online"`
	Twitter                          string              `json:"twitter"`
	Lastfm                           interface{}         `json:"lastfm"`
	Skype                            interface{}         `json:"skype"`
	Website                          string              `json:"website"`
	Discord                          string              `json:"discord"`
	Playstyle                        []string            `json:"playstyle"`
	Playmode                         string              `json:"playmode"`
	PmFriendsOnly                    bool                `json:"pm_friends_only"`
	PostCount                        int                 `json:"post_count"`
	ProfileColour                    interface{}         `json:"profile_colour"`
	ProfileOrder                     []string            `json:"profile_order"`
	CoverURL                         string              `json:"cover_url"`
	Cover                            Cover               `json:"cover"`
	Kudosu                           Kudosu              `json:"kudosu"`
	MaxBlocks                        int                 `json:"max_blocks"`
	MaxFriends                       int                 `json:"max_friends"`
	AccountHistory                   []interface{}       `json:"account_history"`
	ActiveTournamentBanner           []interface{}       `json:"active_tournament_banner"`
	Badges                           []interface{}       `json:"badges"`
	FavouriteBeatmapsetCount         int                 `json:"favourite_beatmapset_count"`
	FollowerCount                    int                 `json:"follower_count"`
	GraveyardBeatmapsetCount         int                 `json:"graveyard_beatmapset_count"`
	LovedBeatmapsetCount             int                 `json:"loved_beatmapset_count"`
	MonthlyPlaycounts                []MonthlyPlaycounts `json:"monthly_playcounts"`
	Page                             Page                `json:"page"`
	PreviousUsernames                []interface{}       `json:"previous_usernames"`
	RankedAndApprovedBeatmapsetCount int                 `json:"ranked_and_approved_beatmapset_count"`
	ReplaysWatchedCounts             []interface{}       `json:"replays_watched_counts"`
	ScoresFirstCount                 int                 `json:"scores_first_count"`
	Statistics                       Statistics          `json:"statistics"`
	SupportLevel                     int                 `json:"support_level"`
	UnrankedBeatmapsetCount          int                 `json:"unranked_beatmapset_count"`
	UserAchievements                 []UserAchievements  `json:"user_achievements"`
	RankHistory                      RankHistory         `json:"rankHistory"`
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