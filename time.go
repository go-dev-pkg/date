package date

import (
	"time"
)

// DayStartAndEnd 当天的开始和结束时间
func DayStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), nt.Month(), nt.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), nt.Month(), nt.Day(), 23, 59, 59, 0, time.Local)
	return
}

// WeekStartAndEnd 当周的开始和结束时间
func WeekStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	weekday := int(nt.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	s := nt.AddDate(0, 0, -weekday+1)
	e := nt.AddDate(0, 0, 7-weekday)
	start = time.Date(s.Year(), s.Month(), s.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(e.Year(), e.Month(), e.Day(), 23, 59, 59, 0, time.Local)
	return
}

// MonthStartAndEnd 当月的开始和结束时间
func MonthStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), nt.Month(), 1, 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), nt.Month(), start.AddDate(0, 1, -1).Day(), 23, 59, 59, 0, time.Local)
	return
}

// QuarterStartAndEnd 当季度的开始和结束时间
func QuarterStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	month := int(nt.Month())
	if month == 1 || month == 2 || month == 3 {
		start = time.Date(nt.Year(), 1, 1, 0, 0, 0, 0, time.Local)
		end = time.Date(nt.Year(), 3, 31, 23, 59, 59, 0, time.Local)
		return
	}
	if month == 4 || month == 5 || month == 6 {
		start = time.Date(nt.Year(), 4, 1, 0, 0, 0, 0, time.Local)
		end = time.Date(nt.Year(), 6, 30, 23, 59, 59, 0, time.Local)
		return
	}
	if month == 7 || month == 8 || month == 9 {
		start = time.Date(nt.Year(), 7, 1, 0, 0, 0, 0, time.Local)
		end = time.Date(nt.Year(), 9, 30, 23, 59, 59, 0, time.Local)
		return
	}
	if month == 10 || month == 11 || month == 12 {
		start = time.Date(nt.Year(), 10, 1, 0, 0, 0, 0, time.Local)
		end = time.Date(nt.Year(), 12, 31, 23, 59, 59, 0, time.Local)
		return
	}
	return
}

// YearStartAndEnd 当年的开始和结束时间
func YearStartAndEnd(t ...time.Time) (start, end time.Time) {
	nt := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		nt = t[0]
	}
	start = time.Date(nt.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	end = time.Date(nt.Year(), 12, 31, 23, 59, 59, 0, time.Local)
	return
}
