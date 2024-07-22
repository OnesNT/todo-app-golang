package models

import (
	"encoding/json"
	"time"
)

type DateOnly time.Time

// UnmarshalJSON parses the date-only format
func (d *DateOnly) UnmarshalJSON(b []byte) (err error) {
	s := string(b)
	s = s[1 : len(s)-1] // Remove quotes
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

// MarshalJSON formats the date-only value
func (d DateOnly) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	return json.Marshal(t.Format("2006-01-02"))
}

type Task struct {
	ID          uint     `json:"id" gorm:"primary_key"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     DateOnly `json:"due_date" gorm:"type:date"`
	Status      bool     `json:"status"`
}
