package entity

import (
	"database/sql/driver"
	"encoding/json"
)

// Statistics in profile
type Statistics struct {
	Id     uint `json:"-" gorm:"id"`
	UserId uint `json:"-" gorm:"user_id"`

	Level                  Level   `json:"level"`
	Pp                     float64 `json:"pp"`
	PpRank                 int     `json:"pp_rank"`
	RankedScore            int     `json:"ranked_score"`
	HitAccuracy            float64 `json:"hit_accuracy"`
	PlayCount              int     `json:"play_count"`
	PlayTime               int     `json:"play_time"`
	TotalScore             int64   `json:"total_score"`
	TotalHits              int     `json:"total_hits"`
	MaximumCombo           int     `json:"maximum_combo"`
	ReplaysWatchedByOthers int     `json:"replays_watched_by_others"`
	IsRanked               bool    `json:"is_ranked"`
	GradeCounts            struct {
		Ss  int `json:"ss"`
		Ssh int `json:"ssh"`
		S   int `json:"s"`
		Sh  int `json:"sh"`
		A   int `json:"a"`
	} `json:"grade_counts"`
	Rank struct {
		Global  int `json:"global" db:"global"`
		Country int `json:"country" db:"country"`
	} `json:"rank"`
}

type StatisticField Statistics

func (c StatisticField) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *StatisticField) Scan(value interface{}) error {
	result := StatisticField{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}

// Level progress and current value in user profile
type Level struct {
	Current  int `json:"current"`
	Progress int `json:"progress"`
}

func (c Level) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *Level) Scan(value interface{}) error {
	result := Level{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}
