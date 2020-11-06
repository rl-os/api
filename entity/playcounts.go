package entity

// MonthlyPlaycounts record
type MonthlyPlaycounts struct {
	Id        uint   `json:"-"`
	UserId    uint   `json:"-"`
	StartDate string `json:"start_date" gorm:"column:year_month"`
	Count     int    `json:"count" gorm:"column:playcount"`
}

func (MonthlyPlaycounts) TableName() string {
	return "user_month_playcount"
}
