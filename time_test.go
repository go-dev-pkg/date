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
}
