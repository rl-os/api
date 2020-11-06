package entity

// MonthlyPlaycounts record
type MonthlyPlaycounts struct {
	Id        uint   `json:"-" gorm:"id"`
	UserId    uint   `json:"-" gorm:"user_id"`
	StartDate string `json:"start_date"`
	Count     int    `json:"count"`
}

func (MonthlyPlaycounts) TableName() string {
	return "user_month_playcount"
}
