package date_test

import (
	"testing"
	"time"

	"github.com/go-dev-pkg/date"
)

func TestWeekStartAndEnd(t *testing.T) {
	// 跨年的周
	times, err := time.ParseInLocation(time.DateTime, "2025-12-31 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end := date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-12-29 00:00:00" {
		t.Fatal(s, "!=", "2025-12-29 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2026-01-04 23:59:59" {
		t.Fatal(e, "!=", "2026-01-04 23:59:59")
		return
	}

	// 跨年的周
	times, err = time.ParseInLocation(time.DateTime, "2026-01-01 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-12-29 00:00:00" {
		t.Fatal(s, "!=", "2025-12-29 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2026-01-04 23:59:59" {
		t.Fatal(e, "!=", "2026-01-04 23:59:59")
		return
	}

	// 周一
	times, err = time.ParseInLocation(time.DateTime, "2025-12-01 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-12-01 00:00:00" {
		t.Fatal(s, "!=", "2025-12-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-12-07 23:59:59" {
		t.Fatal(e, "!=", "2025-12-07 23:59:59")
		return
	}

	// 周日
	times, err = time.ParseInLocation(time.DateTime, "2025-12-07 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-12-01 00:00:00" {
		t.Fatal(s, "!=", "2025-12-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-12-07 23:59:59" {
		t.Fatal(e, "!=", "2025-12-07 23:59:59")
		return
	}

	// 平年跨月的周
	times, err = time.ParseInLocation(time.DateTime, "2025-02-27 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-02-24 00:00:00" {
		t.Fatal(s, "!=", "2025-02-24 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-03-02 23:59:59" {
		t.Fatal(e, "!=", "2025-03-02 23:59:59")
		return
	}

	// 平年跨月的周
	times, err = time.ParseInLocation(time.DateTime, "2025-03-02 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-02-24 00:00:00" {
		t.Fatal(s, "!=", "2025-02-24 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-03-02 23:59:59" {
		t.Fatal(e, "!=", "2025-03-02 23:59:59")
		return
	}

	// 闰年跨月的周
	times, err = time.ParseInLocation(time.DateTime, "2024-02-27 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2024-02-26 00:00:00" {
		t.Fatal(s, "!=", "2024-02-26 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2024-03-03 23:59:59" {
		t.Fatal(e, "!=", "2024-03-03 23:59:59")
		return
	}

	// 闰年跨月的周
	times, err = time.ParseInLocation(time.DateTime, "2024-03-02 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.WeekStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2024-02-26 00:00:00" {
		t.Fatal(s, "!=", "2024-02-26 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2024-03-03 23:59:59" {
		t.Fatal(e, "!=", "2024-03-03 23:59:59")
		return
	}
}

func TestMonthStartAndEnd(t *testing.T) {
	// 闰年
	times, err := time.ParseInLocation(time.DateTime, "2024-02-12 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end := date.MonthStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2024-02-01 00:00:00" {
		t.Fatal(s, "!=", "2024-02-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2024-02-29 23:59:59" {
		t.Fatal(e, "!=", "2024-02-29 23:59:59")
		return
	}

	// 平年
	times, err = time.ParseInLocation(time.DateTime, "2025-02-12 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.MonthStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-02-01 00:00:00" {
		t.Fatal(s, "!=", "2025-02-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-02-28 23:59:59" {
		t.Fatal(e, "!=", "2025-02-28 23:59:59")
		return
	}

	// 31天
	times, err = time.ParseInLocation(time.DateTime, "2025-01-12 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.MonthStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-01-01 00:00:00" {
		t.Fatal(s, "!=", "2025-01-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-01-31 23:59:59" {
		t.Fatal(e, "!=", "2025-01-31 23:59:59")
		return
	}

	// 30天
	times, err = time.ParseInLocation(time.DateTime, "2025-04-12 08:08:08", time.Local)
	if err != nil {
		t.Fatal(err)
		return
	}
	start, end = date.MonthStartAndEnd(times)
	if s := start.Format(time.DateTime); s != "2025-04-01 00:00:00" {
		t.Fatal(s, "!=", "2025-04-01 00:00:00")
		return
	}
	if e := end.Format(time.DateTime); e != "2025-04-30 23:59:59" {
		t.Fatal(e, "!=", "2025-04-30 23:59:59")
		return
	}
}
