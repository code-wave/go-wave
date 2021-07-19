package helpers

import (
	"time"
)

const (
	// dateFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02T15:04:05Z07:00"
)

func GetCurrentTimeForDB() string {
	now := time.Now()
	res := now.Format("2006-01-02 3:4:5 pm")
	return res
}

func GetDateString(date time.Time) string {
	return date.Format(dateFormat)
}
