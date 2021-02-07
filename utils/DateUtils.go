package utils

import "time"

var TimeNow = time.Now

func GetEndOfCurrentWeek() (timestampResult int) {

	now := TimeNow().UTC()

	offset := int(time.Saturday - now.Weekday())

	weekDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC).AddDate(0, 0, offset)
	timestampResult = int(weekDate.UTC().Unix())
	return
}
