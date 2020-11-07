package entity

import (
	"github.com/lib/pq"
)

// RankHistory record
type RankHistory struct {
	Id     uint `json:"-" gorm:"id"`
	UserId uint `json:"-" gorm:"user_id"`

	Mode string        `json:"mode"`
	Data pq.Int64Array `json:"data" gorm:"type:int[]"`
}

// TableName of RankHistory
func (r RankHistory) TableName() string {
	return "user_performance_ranks"
}
