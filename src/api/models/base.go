package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type JSONTime struct {
	Time time.Time
}

const TimeFmt = "2006-01-02T15:04:05Z"

func (t *JSONTime) Scan(src interface{}) error {
	t.Time = src.(time.Time)
	return nil
}

func (t JSONTime) Value() (driver.Value, error) {
	return t.Time, nil
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t.Time).Format(TimeFmt))
	return []byte(stamp), nil
}

type Base struct {
	ID        int64    `json:"id" gorm:"primary_key"`
	CreatedAt JSONTime `json:"created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt JSONTime `json:"updated_at" sql:"DEFAULT:current_timestamp"`
}
