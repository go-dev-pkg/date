package date_test

import (
	"fmt"
	"testing"

	"github.com/go-dev-pkg/date"
)

func TestYear(t *testing.T) {
	// 闰年
	year := 2024
	quarter, month, week := date.Year(year)
	for _, q := range quarter {
		quarterWeek := 0
		for _, w := range q.Weeks {
			quarterWeek++
			if w.Week != quarterWeek {
				t.Fatal("quarter week does not match")
				return
			}
			if w.Year != year {
				t.Fatal("year does not match")
				return
			}
			for _, day := range w.Days {
				if day.QuarterWeeks != quarterWeek {
					t.Fatal("day quarter week does not match")
					return
				}
				if day.Year != year {
					t.Fatal("day year does not match")
					return
				}
			}
		}

		if q.Quarter == 1 {
			firstDay := q.Weeks[0].Days[0]
			if firstDay.Month != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首月不是1月", year, q.Quarter), firstDay.Month, "!=", 1)
				return
			}
			if firstDay.Day != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日不是1号", year, q.Quarter), firstDay.Day, "!=", 1)
				return
			}
			if firstDay.Weekday != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日星期几有误", year, q.Quarter), firstDay.Weekday, "!=", 1)
				return
			}
			if firstDay.QuarterWeeks != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日是季度第几周有误", year, q.Quarter), firstDay.QuarterWeeks, "!=", 1)
				return
			}
			lastWeek := q.Weeks[len(q.Weeks)-1]
			lastDay := lastWeek.Days[len(lastWeek.Days)-1]
			if lastDay.Month != 3 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾月不是3月", year, q.Quarter), lastDay.Month, "!=", 1)
				return
			}
			if lastDay.Day != 31 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日不是31号", year, q.Quarter), lastDay.Day, "!=", 31)
				return
			}
			if lastDay.Weekday != 7 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日星期几有误", year, q.Quarter), lastDay.Weekday, "!=", 7)
				return
			}
			if lastDay.QuarterWeeks != 13 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日是季度第几周有误", year, q.Quarter), lastDay.QuarterWeeks, "!=", 13)
				return
			}
		}
		if q.Quarter == 2 {
			firstDay := q.Weeks[0].Days[0]
			if firstDay.Month != 4 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首月不是4月", year, q.Quarter), firstDay.Month, "!=", 4)
				return
			}
			if firstDay.Day != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日不是1号", year, q.Quarter), firstDay.Day, "!=", 1)
				return
			}
			if firstDay.Weekday != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日星期几有误", year, q.Quarter), firstDay.Weekday, "!=", 1)
				return
			}
			if firstDay.QuarterWeeks != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日是季度第几周有误", year, q.Quarter), firstDay.QuarterWeeks, "!=", 1)
				return
			}
			lastWeek := q.Weeks[len(q.Weeks)-1]
			lastDay := lastWeek.Days[len(lastWeek.Days)-1]
			if lastDay.Month != 6 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾月不是6月", year, q.Quarter), lastDay.Month, "!=", 6)
				return
			}
			if lastDay.Day != 30 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日不是30号", year, q.Quarter), lastDay.Day, "!=", 30)
				return
			}
			if lastDay.Weekday != 7 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日星期几有误", year, q.Quarter), lastDay.Weekday, "!=", 7)
				return
			}
			if lastDay.QuarterWeeks != 13 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日是季度第几周有误", year, q.Quarter), lastDay.QuarterWeeks, "!=", 13)
				return
			}
		}
		if q.Quarter == 3 {
			firstDay := q.Weeks[0].Days[0]
			if firstDay.Month != 7 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首月不是7月", year, q.Quarter), firstDay.Month, "!=", 7)
				return
			}
			if firstDay.Day != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日不是1号", year, q.Quarter), firstDay.Day, "!=", 1)
				return
			}
			if firstDay.Weekday != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日星期几有误", year, q.Quarter), firstDay.Weekday, "!=", 1)
				return
			}
			if firstDay.QuarterWeeks != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日是季度第几周有误", year, q.Quarter), firstDay.QuarterWeeks, "!=", 1)
				return
			}
			lastWeek := q.Weeks[len(q.Weeks)-1]
			lastDay := lastWeek.Days[len(lastWeek.Days)-1]
			if lastDay.Month != 9 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾月不是9月", year, q.Quarter), lastDay.Month, "!=", 9)
				return
			}
			if lastDay.Day != 30 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日不是30号", year, q.Quarter), lastDay.Day, "!=", 30)
				return
			}
			if lastDay.Weekday != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日星期几有误", year, q.Quarter), lastDay.Weekday, "!=", 1)
				return
			}
			if lastDay.QuarterWeeks != 14 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日是季度第几周有误", year, q.Quarter), lastDay.QuarterWeeks, "!=", 14)
				return
			}
		}
		if q.Quarter == 4 {
			firstDay := q.Weeks[0].Days[0]
			if firstDay.Month != 10 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首月不是10月", year, q.Quarter), firstDay.Month, "!=", 10)
				return
			}
			if firstDay.Day != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日不是1号", year, q.Quarter), firstDay.Day, "!=", 1)
				return
			}
			if firstDay.Weekday != 2 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日星期几有误", year, q.Quarter), firstDay.Weekday, "!=", 2)
				return
			}
			if firstDay.QuarterWeeks != 1 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的首日是季度第几周有误", year, q.Quarter), firstDay.QuarterWeeks, "!=", 1)
				return
			}
			lastWeek := q.Weeks[len(q.Weeks)-1]
			lastDay := lastWeek.Days[len(lastWeek.Days)-1]
			if lastDay.Month != 12 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾月不是12月", year, q.Quarter), lastDay.Month, "!=", 12)
				return
			}
			if lastDay.Day != 31 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日不是31号", year, q.Quarter), lastDay.Day, "!=", 31)
				return
			}
			if lastDay.Weekday != 2 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日星期几有误", year, q.Quarter), lastDay.Weekday, "!=", 2)
				return
			}
			if lastDay.QuarterWeeks != 14 {
				t.Fatal(fmt.Sprintf("%d年第%d季度的尾日是季度第几周有误", year, q.Quarter), lastDay.QuarterWeeks, "!=", 13)
				return
			}
		}
	}
	for _, m := range month {
		if m.Month == 1 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 1 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 1)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 3 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 3)
				return
			}
		}
		if m.Month == 2 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 29 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 29)
				return
			}
			if w := m.Days[0].Weekday; w != 4 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 4)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 4 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 4)
				return
			}
		}
		if m.Month == 3 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 5)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 7 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 7)
				return
			}
		}
		if m.Month == 4 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 30 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 30)
				return
			}
			if w := m.Days[0].Weekday; w != 1 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 1)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 2 {
				t.Fatal(fmt.Sprintf("%d年%d月30号的星期几有误", year, m.Month), w, "!=", 2)
				return
			}
		}
		if m.Month == 5 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 3 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 3)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 5)
				return
			}
		}
		if m.Month == 6 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 30 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 30)
				return
			}
			if w := m.Days[0].Weekday; w != 6 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 6)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 7 {
				t.Fatal(fmt.Sprintf("%d年%d月30号的星期几有误", year, m.Month), w, "!=", 7)
				return
			}
		}
		if m.Month == 7 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 1 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 1)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 3 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 3)
				return
			}
		}
		if m.Month == 8 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 4 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 4)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 6 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 6)
				return
			}
		}
		if m.Month == 9 {
			if l := len(m.Weeks); l != 6 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 6)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 30 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 30)
				return
			}
			if w := m.Days[0].Weekday; w != 7 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 7)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 1 {
				t.Fatal(fmt.Sprintf("%d年%d月30号的星期几有误", year, m.Month), w, "!=", 1)
				return
			}
		}
		if m.Month == 10 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 2 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 2)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 4 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 4)
				return
			}
		}
		if m.Month == 11 {
			if l := len(m.Weeks); l != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 5)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 30 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 30)
				return
			}
			if w := m.Days[0].Weekday; w != 5 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 5)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 6 {
				t.Fatal(fmt.Sprintf("%d年%d月30号的星期几有误", year, m.Month), w, "!=", 6)
				return
			}
		}
		if m.Month == 12 {
			if l := len(m.Weeks); l != 6 {
				t.Fatal(fmt.Sprintf("%d年%d月的总周数有误", year, m.Month), l, "!=", 6)
				return
			}
			if d := m.Days[len(m.Days)-1].Day; d != 31 {
				t.Fatal(fmt.Sprintf("%d年%d月的天数有误", year, m.Month), d, "!=", 31)
				return
			}
			if w := m.Days[0].Weekday; w != 7 {
				t.Fatal(fmt.Sprintf("%d年%d月1号的星期几有误", year, m.Month), w, "!=", 7)
				return
			}
			if w := m.Days[len(m.Days)-1].Weekday; w != 2 {
				t.Fatal(fmt.Sprintf("%d年%d月31号的星期几有误", year, m.Month), w, "!=", 2)
				return
			}
		}
	}
	firstWeek := week[0]
	firstWeekFirstDay := firstWeek.Days[0]
	if m := firstWeekFirstDay.Month; m != 1 {
		t.Fatal(fmt.Sprintf("%d年第一周第一天的月分有误", year), m, "!=", 1)
		return
	}
	if d := firstWeekFirstDay.Day; d != 1 {
		t.Fatal(fmt.Sprintf("%d年第一周第一天有误", year), d, "!=", 1)
		return
	}
	if w := firstWeekFirstDay.Weekday; w != 1 {
		t.Fatal(fmt.Sprintf("%d年第一周第一天星期几有误", year), w, "!=", 1)
		return
	}
	firstWeekLastDay := firstWeek.Days[len(firstWeek.Days)-1]
	if m := firstWeekLastDay.Month; m != 1 {
		t.Fatal(fmt.Sprintf("%d年第一周最后一天的月分有误", year), m, "!=", 1)
		return
	}
	if d := firstWeekLastDay.Day; d != 7 {
		t.Fatal(fmt.Sprintf("%d年第一周最后一天有误", year), d, "!=", 7)
		return
	}
	if w := firstWeekLastDay.Weekday; w != 7 {
		t.Fatal(fmt.Sprintf("%d年第一周最后一天星期几有误", year), w, "!=", 7)
		return
	}

	lastWeek := week[len(week)-1]
	lastWeekFirstDay := lastWeek.Days[0]
	if m := lastWeekFirstDay.Month; m != 12 {
		t.Fatal(fmt.Sprintf("%d年最后一周第一天的月分有误", year), m, "!=", 12)
		return
	}
	if d := lastWeekFirstDay.Day; d != 30 {
		t.Fatal(fmt.Sprintf("%d年最后一周第一天有误", year), d, "!=", 30)
		return
	}
	if w := lastWeekFirstDay.Weekday; w != 1 {
		t.Fatal(fmt.Sprintf("%d年最后一周第一天星期几有误", year), w, "!=", 1)
		return
	}
	lastWeekLastDay := lastWeek.Days[len(lastWeek.Days)-1]
	if m := lastWeekLastDay.Month; m != 12 {
		t.Fatal(fmt.Sprintf("%d年最后一周最后一天的月分有误", year), m, "!=", 12)
		return
	}
	if d := lastWeekLastDay.Day; d != 31 {
		t.Fatal(fmt.Sprintf("%d年最后一周最后一天有误", year), d, "!=", 31)
		return
	}
	if w := lastWeekLastDay.Weekday; w != 2 {
		t.Fatal(fmt.Sprintf("%d年最后一周最后一天星期几有误", year), w, "!=", 2)
		return
	}
	if w := lastWeekLastDay.YearWeeks; w != 53 {
		t.Fatal(fmt.Sprintf("%d年的年周有误", year), w, "!=", 53)
		return
	}
}
