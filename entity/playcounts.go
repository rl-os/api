package entity

import (
	"database/sql/driver"
	"encoding/json"
)

// MonthlyPlaycounts record
type MonthlyPlaycounts struct {
	Id        uint   `json:"-" gorm:"id"`
	UserId    uint   `json:"-" gorm:"user_id"`
	StartDate string `json:"start_date" gorm:"year_month"`
	Count     int    `json:"count" gorm:"playcount"`
}

type MonthlyPlaycountsField MonthlyPlaycounts

func (c MonthlyPlaycountsField) Value() (driver.Value, error) { return json.Marshal(c) }
func (c *MonthlyPlaycountsField) Scan(value interface{}) error {
	result := MonthlyPlaycountsField{}
	err := json.Unmarshal(value.([]byte), &result)
	return err
}
