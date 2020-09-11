package entity

import "time"

type BeatmapsetSearch struct {
	ID                int          `json:"id"`
	Title             string       `json:"title"`
	Artist            string       `json:"artist"`
	Covers            Covers       `json:"covers"`
	Creator           string       `json:"creator"`
	FavouriteCount    int          `json:"favourite_count"`
	PlayCount         int          `json:"play_count"`
	PreviewURL        string       `json:"preview_url"`
	Source            string       `json:"source"`
	Status            string       `json:"status"`
	UserID            int          `json:"user_id"`
	Video             bool         `json:"video"`
	Availability      Availability `json:"availability"`
	Bpm               int          `json:"bpm"`
	CanBeHyped        bool         `json:"can_be_hyped"`
	DiscussionEnabled bool         `json:"discussion_enabled"`
	DiscussionLocked  bool         `json:"discussion_locked"`
	Hype              Hype         `json:"hype"`
	IsScoreable       bool         `json:"is_scoreable"`
	LastUpdated       time.Time    `json:"last_updated"`
	LegacyThreadURL   string       `json:"legacy_thread_url"`
	Nominations       Nominations  `json:"nominations"`
	Ranked            int          `json:"ranked"`
	RankedDate        time.Time    `json:"ranked_date"`
	Storyboard        bool         `json:"storyboard"`
	SubmittedDate     time.Time    `json:"submitted_date"`
	Tags              string       `json:"tags"`
	HasFavourited     bool         `json:"has_favourited"`
	Beatmaps          []Beatmap    `json:"beatmaps"`
}

type BeatmapsetSearchResult struct {
	Beatmapsets           *[]BeatmapsetSearch `json:"beatmapsets"`
	RecommendedDifficulty float32             `json:"recommended_difficulty"`
	Error                 error               `json:"error"`
	Total                 uint                `json:"total"`
}
