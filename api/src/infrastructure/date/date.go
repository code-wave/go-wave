package date

import "time"

const (
	// dateFormat = "2006-01-02 15:04:05"
	dateFormat = "2006-01-02T15:04:05Z07:00"
)

func GetDateString(date time.Time) string {
	return date.Format(dateFormat)
}
