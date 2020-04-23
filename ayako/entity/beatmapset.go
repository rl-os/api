package entity

import (
	"database/sql/driver"
	"github.com/deissh/go-utils"
	"time"
)

type BeatmapSet struct {
	ID                int64        `json:"id" db:"id"`
	LastChecked       time.Time    `json:"last_checked" db:"last_checked"`
	Title             string       `json:"title" db:"title"`
	Artist            string       `json:"artist" db:"artist"`
	PlayCount         int64        `json:"play_count" db:"play_count"`
	FavouriteCount    int64        `json:"favourite_count" db:"favourite_count"`
	HasFavourited     bool         `json:"has_favourited" db:"has_favourited"`
	SubmittedDate     time.Time    `json:"submitted_date" db:"submitted_date"`
	LastUpdated       time.Time    `json:"last_updated" db:"last_updated"`
	RankedDate        time.Time    `json:"ranked_date" db:"ranked_date"`
	Creator           string       `json:"creator" db:"creator"`
	UserID            int64        `json:"user_id" db:"user_id"`
	BPM               int64        `json:"bpm" db:"bpm"`
	Source            string       `json:"source" db:"source"`
	Covers            Covers       `json:"covers" db:"covers"`
	PreviewURL        string       `json:"preview_url" db:"preview_url"`
	Tags              string       `json:"tags" db:"tags"`
	Video             bool         `json:"video" db:"video"`
	Storyboard        bool         `json:"storyboard" db:"storyboard"`
	Ranked            int64        `json:"ranked" db:"ranked"`
	Status            Status       `json:"status" db:"status"`
	IsScoreable       bool         `json:"is_scoreable" db:"is_scoreable"`
	DiscussionEnabled bool         `json:"discussion_enabled" db:"discussion_enabled"`
	DiscussionLocked  bool         `json:"discussion_locked" db:"discussion_locked"`
	CanBeHyped        bool         `json:"can_be_hyped" db:"can_be_hyped"`
	Availability      Availability `json:"availability" db:"availability"`
	Hype              Hype         `json:"hype" db:"hype"`
	Nominations       Hype         `json:"nominations" db:"nominations"`
	LegacyThreadURL   string       `json:"legacy_thread_url" db:"legacy_thread_url"`
	Description       Description  `json:"description" db:"description"`
	Genre             Genre        `json:"genre" db:"genre"`
	Language          Genre        `json:"language" db:"language"`
	User              User         `json:"user" db:"user"`

	Beatmaps []Beatmap `json:"beatmaps" db:"-"`
}

type BeatmapSetFull struct {
	BeatmapSet

	RecentFavourites      []User                `json:"recent_favourites"`
	CurrentUserAttributes CurrentUserAttributes `json:"current_user_attributes"`
	Ratings               []int64               `json:"ratings"`
	Converts              []Beatmap             `json:"converts"`
}

type Availability struct {
	DownloadDisabled bool        `json:"download_disabled"`
	MoreInformation  interface{} `json:"more_information"`
}

func (c Availability) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *Availability) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

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

func (c Covers) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *Covers) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type CurrentUserAttributes struct {
	CanDelete     bool        `json:"can_delete"`
	CanHype       bool        `json:"can_hype"`
	CanHypeReason interface{} `json:"can_hype_reason"`
	CanLove       bool        `json:"can_love"`
	IsWatching    bool        `json:"is_watching"`
	NewHypeTime   interface{} `json:"new_hype_time"`
	RemainingHype int64       `json:"remaining_hype"`
}

func (c CurrentUserAttributes) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *CurrentUserAttributes) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type Description struct {
	Description string `json:"description"`
}

func (c Description) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *Description) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type Genre struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (c Genre) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *Genre) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type Hype struct {
	Current  int64 `json:"current"`
	Required int64 `json:"required"`
}

func (c Hype) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *Hype) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

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

func (c User) Value() (driver.Value, error)  { return utils.ValueOfStruct(c) }
func (c *User) Scan(value interface{}) error { return utils.ScanToStruct(c, value) }

type DefaultGroup string

const (
	Default DefaultGroup = "default"
)
