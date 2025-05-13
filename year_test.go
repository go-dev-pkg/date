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
		_ = q
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
