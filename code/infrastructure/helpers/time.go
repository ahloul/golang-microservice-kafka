package helpers

import (
	"time"
)

func AddTime(date time.Time, t time.Time) time.Time {
	return date.Add(time.Hour*time.Duration(t.Hour()) + time.Minute*time.Duration(t.Minute()) + time.Second*time.Duration(t.Second()))
}

func SeparateStringDateTime(dateS string, layout string) (time.Time, time.Time) {
	date, _ := time.Parse(layout, dateS)
	t := AddTime(time.Time{}, date)
	date = time.Time{}.AddDate(date.Year()-1, int(date.Month())-1, date.Day()-1)
	return date, t
}
