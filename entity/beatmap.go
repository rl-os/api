package entity

import "time"

type Beatmap struct {
	ID               int64    `json:"id"`
	BeatmapsetID     int64    `json:"beatmapset_id"`
	Mode             PlayMode `json:"mode"`
	ModeInt          uint8    `json:"mode_int"`
	Convert          bool     `json:"convert"`
	DifficultyRating float64  `json:"difficulty_rating"`
	Version          string   `json:"version"`
	TotalLength      int64    `json:"total_length"`
	HitLength        int64    `json:"hit_length"`
	BPM              int64    `json:"bpm"`
	CS               int64    `json:"cs"`
	Drain            int64    `json:"drain"`
	Accuracy         int64    `json:"accuracy"`
	Ar               int64    `json:"ar"`
	Playcount        int64    `json:"playcount"`
	Passcount        int64    `json:"passcount"`
	CountCircles     int64    `json:"count_circles"`
	CountSliders     int64    `json:"count_sliders"`
	CountSpinners    int64    `json:"count_spinners"`
	CountTotal       int64    `json:"count_total"`
	IsScoreable      bool     `json:"is_scoreable"`
	Ranked           int64    `json:"ranked"`
	Status           Status   `json:"status"`
	URL              string   `json:"url"`
	//Failtimes      Failtimes `json:"failtimes"`
	MaxCombo *int64 `json:"max_combo"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"last_updated"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (Beatmap) TableName() string {
	return "beatmaps"
}

type SingleBeatmap struct {
	Beatmap
	Beatmapset BeatmapSet `json:"beatmapset" gorm:"foreignkey:beatmapset_id;references:id"`
}

type Failtimes struct {
	Fail []int64 `json:"fail"`
	Exit []int64 `json:"exit"`
}

type PlayMode string

const (
	Fruits PlayMode = "fruits"
	Mania  PlayMode = "mania"
	Osu    PlayMode = "osu"
	Taiko  PlayMode = "taiko"
)

var Modes = map[PlayMode]uint8{
	Osu:    0,
	Taiko:  1,
	Fruits: 2,
	Mania:  3,
}

type Status string

const (
	BeatmapRanked Status = "ranked"
	BeatmapLoved  Status = "loved"
)
