package entity

import "time"

// UserAchievements with datetime
type UserAchievements struct {
	ID            uint `json:"-"`
	UserId        uint `json:"-"`
	AchievementID int  `json:"achievement_id"`

	CreatedAt time.Time `json:"achieved_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func (UserAchievements) TableName() string {
	return "user_achievements"
}
