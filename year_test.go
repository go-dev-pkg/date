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
	for _, w := range week {
		_ = w
	}
}
