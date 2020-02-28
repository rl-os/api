package entity

import "time"

type BeatmapSet struct {
	ID                    int64                 `json:"id"`
	Title                 string                `json:"title"`
	Artist                string                `json:"artist"`
	PlayCount             int64                 `json:"play_count"`
	FavouriteCount        int64                 `json:"favourite_count"`
	HasFavourited         bool                  `json:"has_favourited"`
	SubmittedDate         time.Time             `json:"submitted_date"`
	LastUpdated           time.Time             `json:"last_updated"`
	RankedDate            time.Time             `json:"ranked_date"`
	Creator               string                `json:"creator"`
	UserID                int64                 `json:"user_id"`
	BPM                   int64                 `json:"bpm"`
	Source                string                `json:"source"`
	Covers                Covers                `json:"covers"`
	PreviewURL            string                `json:"preview_url"`
	Tags                  string                `json:"tags"`
	Video                 bool                  `json:"video"`
	Storyboard            bool                  `json:"storyboard"`
	Ranked                int64                 `json:"ranked"`
	Status                Status                `json:"status"`
	IsScoreable           bool                  `json:"is_scoreable"`
	DiscussionEnabled     bool                  `json:"discussion_enabled"`
	DiscussionLocked      bool                  `json:"discussion_locked"`
	CanBeHyped            bool                  `json:"can_be_hyped"`
	Availability          Availability          `json:"availability"`
	Hype                  Hype                  `json:"hype"`
	Nominations           Hype                  `json:"nominations"`
	LegacyThreadURL       string                `json:"legacy_thread_url"`
	Beatmaps              []Beatmap             `json:"beatmaps"`
	Converts              []Beatmap             `json:"converts"`
	CurrentUserAttributes CurrentUserAttributes `json:"current_user_attributes"`
	Description           Description           `json:"description"`
	Genre                 Genre                 `json:"genre"`
	Language              Genre                 `json:"language"`
	Ratings               []int64               `json:"ratings"`
	RecentFavourites      []User                `json:"recent_favourites"`
	User                  User                  `json:"user"`
}

type Availability struct {
	DownloadDisabled bool        `json:"download_disabled"`
	MoreInformation  interface{} `json:"more_information"`
}

type Covers struct {
	Cover       string `json:"cover"`
	Cover2X     string `json:"cover@2x"`
	Card        string `json:"card"`
	Card2X      string `json:"card@2x"`
	List        string `json:"list"`
	List2X      string `json:"list@2x"`
	Slimcover   string `json:"slimcover"`
	Slimcover2X string `json:"slimcover@2x"`
}

type CurrentUserAttributes struct {
	CanDelete     bool        `json:"can_delete"`
	CanHype       bool        `json:"can_hype"`
	CanHypeReason interface{} `json:"can_hype_reason"`
	CanLove       bool        `json:"can_love"`
	IsWatching    bool        `json:"is_watching"`
	NewHypeTime   interface{} `json:"new_hype_time"`
	RemainingHype int64       `json:"remaining_hype"`
}

type Description struct {
	Description string `json:"description"`
}

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Hype struct {
	Current  int64 `json:"current"`
	Required int64 `json:"required"`
}

type User struct {
	ID            int64        `json:"id"`
	Username      string       `json:"username"`
	ProfileColour interface{}  `json:"profile_colour"`
	AvatarURL     string       `json:"avatar_url"`
	CountryCode   string       `json:"country_code"`
	DefaultGroup  DefaultGroup `json:"default_group"`
	IsActive      bool         `json:"is_active"`
	IsBot         bool         `json:"is_bot"`
	IsOnline      bool         `json:"is_online"`
	IsSupporter   bool         `json:"is_supporter"`
	LastVisit     *string      `json:"last_visit"`
	PmFriendsOnly bool         `json:"pm_friends_only"`
}

type DefaultGroup string

const (
	Default DefaultGroup = "default"
)
