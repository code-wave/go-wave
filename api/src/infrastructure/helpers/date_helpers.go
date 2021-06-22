package helpers

import (
	"time"
)

func GetCurrentTimeForDB() string {
	now := time.Now()
	res := now.Format("2006-01-02 3:4:5 pm")
	return res
}
