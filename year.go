package date

type Quarter struct {
	Year    int           `json:"year"`    // 年份
	Quarter int           `json:"quarter"` // 季度
	Weeks   []QuarterWeek `json:"weeks"`   // 季周
}

type Month struct {
	Year    int    `json:"year"`    // 年份
	Quarter int    `json:"quarter"` // 季度
	Month   int    `json:"month"`   // 月份
	Ym      int    `json:"ym"`      // 年月
	Weeks   []Week `json:"weeks"`   // 月周
	Days    []Day  `json:"days"`    // 日期
}

type Week struct {
	Year    int   `json:"year"`    // 年份
	Quarter int   `json:"quarter"` // 季度
	Month   int   `json:"month"`   // 月份
	Week    int   `json:"week"`    // 周数
	Days    []Day `json:"days"`    // 日期
}

type QuarterWeek struct {
	Year    int   `json:"year"`    // 年份
	Quarter int   `json:"quarter"` // 季度
	Week    int   `json:"week"`    // 周数
	Days    []Day `json:"days"`    // 日期
}

type YearWeek struct {
	Year int   `json:"year"` // 年份
	Week int   `json:"week"` // 周数
	Days []Day `json:"days"` // 日期
}

type Day struct {
	Year         int `json:"year"`          // 年份
	Quarter      int `json:"quarter"`       // 季度
	Month        int `json:"month"`         // 月份
	Day          int `json:"day"`           // 日期
	Ymd          int `json:"ymd"`           // 年月日
	Weekday      int `json:"weekday"`       // 星期几
	MonthWeeks   int `json:"month_weeks"`   // 当前月份的周数
	QuarterWeeks int `json:"quarter_weeks"` // 当前季度的周数
	YearWeeks    int `json:"year_weeks"`    // 当前年份的周数
}

func year2Months(year int) []Month {
	if year < 1900 {
		return nil
	}
	return generateMonths(year)
}

func generateMonths(year int) []Month {
	sumDay := 0

	// 1、不累加到当前年份的总天数
	for _year := 1900; _year < year; _year++ {
		if (_year%4 == 0 && _year%100 != 0) || _year%400 == 0 {
			sumDay += 366
			continue
		}
		sumDay += 365
	}

	// 2、累加到当前年份的每月天数
	months := make([]Month, 0, 12)
	yearWeeks := 0
	q1Weeks := 0
	q2Weeks := 0
	q3Weeks := 0
	q4Weeks := 0
	for _month := 1; _month <= 12; _month++ {
		quarter := 4
		if _month >= 1 && _month <= 3 {
			quarter = 1
		}
		if _month >= 4 && _month <= 6 {
			quarter = 2
		}
		if _month >= 7 && _month <= 9 {
			quarter = 3
		}

		days := 0
		if _month == 1 || _month == 3 || _month == 5 || _month == 7 || _month == 8 || _month == 10 || _month == 12 {
			days = 31
		}
		if _month == 4 || _month == 6 || _month == 9 || _month == 11 {
			days = 30
		}
		if _month == 2 {
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				days = 29
			} else {
				days = 28
			}
		}

		month := Month{
			Year:    year,
			Quarter: quarter,
			Month:   _month,
			Ym:      year*100 + _month,
			Days:    make([]Day, 0, days),
		}

		monthWeeks := 0
		for day := 1; day <= days; day++ {
			weekday := (sumDay + day) % 7
			if weekday == 0 {
				weekday = 7
			}

			if _month == 1 && day == 1 {
				yearWeeks = 1
			} else {
				if weekday == 1 {
					yearWeeks++
				}
			}

			if quarter == 1 {
				if _month == 1 && day == 1 {
					q1Weeks = 1
				} else {
					if weekday == 1 {
						q1Weeks++
					}
				}
			}

			if quarter == 2 {
				if _month == 4 && day == 1 {
					q2Weeks = 1
				} else {
					if weekday == 1 {
						q2Weeks++
					}
				}
			}

			if quarter == 3 {
				if _month == 7 && day == 1 {
					q3Weeks = 1
				} else {
					if weekday == 1 {
						q3Weeks++
					}
				}
			}

			if quarter == 4 {
				if _month == 10 && day == 1 {
					q4Weeks = 1
				} else {
					if weekday == 1 {
						q4Weeks++
					}
				}
			}

			quarterWeeks := 0
			if quarter == 1 {
				quarterWeeks = q1Weeks
			} else if quarter == 2 {
				quarterWeeks = q2Weeks
			} else if quarter == 3 {
				quarterWeeks = q3Weeks
			} else {
				quarterWeeks = q4Weeks
			}

			if day == 1 {
				monthWeeks = 1
			} else {
				if weekday == 1 {
					monthWeeks++
				}
			}

			month.Days = append(month.Days, Day{
				Year:         month.Year,
				Quarter:      month.Quarter,
				Month:        month.Month,
				Day:          day,
				Ymd:          month.Year*10000 + month.Month*100 + day,
				Weekday:      weekday,
				MonthWeeks:   monthWeeks,
				YearWeeks:    yearWeeks,
				QuarterWeeks: quarterWeeks,
			})
		}

		months = append(months, month)
		sumDay += days
	}

	return months
}

func Year(year int) ([]Quarter, []Month, []YearWeek) {
	months := year2Months(year)
	if months == nil {
		return nil, nil, nil
	}

	quarter := make([]Quarter, 0, 4)
	lastMonth := months[len(months)-1]
	lastDay := lastMonth.Days[len(lastMonth.Days)-1]
	yearWeeks := make([]YearWeek, lastDay.YearWeeks)

	for i, month := range months {
		if month.Month == 1 {
			quarterLastMonth := months[2]
			quarterLastDay := quarterLastMonth.Days[len(quarterLastMonth.Days)-1]
			quarter = append(quarter, Quarter{
				Year:    month.Year,
				Quarter: month.Quarter,
				Weeks:   make([]QuarterWeek, quarterLastDay.QuarterWeeks),
			})
		}
		if month.Month == 4 {
			quarterLastMonth := months[5]
			quarterLastDay := quarterLastMonth.Days[len(quarterLastMonth.Days)-1]
			quarter = append(quarter, Quarter{
				Year:    month.Year,
				Quarter: month.Quarter,
				Weeks:   make([]QuarterWeek, quarterLastDay.QuarterWeeks),
			})
		}
		if month.Month == 7 {
			quarterLastMonth := months[8]
			quarterLastDay := quarterLastMonth.Days[len(quarterLastMonth.Days)-1]
			quarter = append(quarter, Quarter{
				Year:    month.Year,
				Quarter: month.Quarter,
				Weeks:   make([]QuarterWeek, quarterLastDay.QuarterWeeks),
			})
		}
		if month.Month == 10 {
			quarterLastMonth := months[11]
			quarterLastDay := quarterLastMonth.Days[len(quarterLastMonth.Days)-1]
			quarter = append(quarter, Quarter{
				Year:    month.Year,
				Quarter: month.Quarter,
				Weeks:   make([]QuarterWeek, quarterLastDay.QuarterWeeks),
			})
		}

		monthWeeks := month.Days[len(month.Days)-1].MonthWeeks
		weeks := make([]Week, monthWeeks)
		for j, day := range month.Days {
			if j == 0 || day.Weekday == 1 {
				weeks[day.MonthWeeks-1].Year = day.Year
				weeks[day.MonthWeeks-1].Quarter = day.Quarter
				weeks[day.MonthWeeks-1].Month = day.Month
				weeks[day.MonthWeeks-1].Week = day.MonthWeeks

				yearWeeks[day.YearWeeks-1].Year = day.Year
				yearWeeks[day.YearWeeks-1].Week = day.YearWeeks
			}
			if day.Month == 1 && day.Day == 1 {
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Year = day.Year
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Quarter = day.Quarter
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Week = day.QuarterWeeks
			} else if day.Month == 4 && day.Day == 1 {
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Year = day.Year
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Quarter = day.Quarter
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Week = day.QuarterWeeks
			} else if day.Month == 7 && day.Day == 1 {
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Year = day.Year
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Quarter = day.Quarter
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Week = day.QuarterWeeks
			} else if day.Month == 10 && day.Day == 1 {
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Year = day.Year
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Quarter = day.Quarter
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Week = day.QuarterWeeks
			} else if day.Weekday == 1 {
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Year = day.Year
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Quarter = day.Quarter
				quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Week = day.QuarterWeeks
			}
			weeks[day.MonthWeeks-1].Days = append(weeks[day.MonthWeeks-1].Days, day)
			yearWeeks[day.YearWeeks-1].Days = append(yearWeeks[day.YearWeeks-1].Days, day)
			quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Days = append(quarter[day.Quarter-1].Weeks[day.QuarterWeeks-1].Days, day)
		}
		months[i].Weeks = weeks
	}

	return quarter, months, yearWeeks
}
