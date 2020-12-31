package entity

import "github.com/lib/pq"

type Failtimes struct {
	BeatmapId int64 `json:"-" gorm:"beatmap_id"`

	Fail pq.Int64Array `json:"fail" gorm:"type:int[]"`
	Exit pq.Int64Array `json:"exit" gorm:"type:int[]"`
}

func (Failtimes) TableName() string {
	return "beatmap_failtimes"
}
