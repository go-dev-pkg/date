package date

import (
	"time"
)

func DayStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), nt.Month(), nt.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), nt.Month(), nt.Day(), 23, 59, 59, 0, time.Local)
	return
}

func MonthStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), nt.Month(), 1, 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), nt.Month(), start.AddDate(0, 1, -1).Day(), 23, 59, 59, 0, time.Local)
	return
}

func YearStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), 12, 31, 23, 59, 59, 0, time.Local)
	return
}
