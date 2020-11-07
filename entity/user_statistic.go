package entity

// UserStatistics in profile
type UserStatistics struct {
	Id     uint   `json:"-" gorm:"id"`
	UserId uint   `json:"-" gorm:"user_id"`
	Mode   string `json:"-"`

	Level struct {
		Current  int `json:"current"`
		Progress int `json:"progress"`
	} `json:"level" gorm:"embedded;embeddedPrefix:level_"`
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
	} `json:"grade_counts" gorm:"embedded;embeddedPrefix:grade_"`
	Rank struct {
		Global  int `json:"global"`
		Country int `json:"country"`
	} `json:"rank" gorm:"-"`
}

func (UserStatistics) TableName() string {
	return "user_statistics"
}
