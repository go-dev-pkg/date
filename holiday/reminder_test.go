package holiday_test

import (
	"testing"
	"time"

	"github.com/go-dev-pkg/date/holiday"
)

func TestHoliday(t *testing.T) {
	res := holiday.Reminder(5, time.Date(2025, time.December, 19, 23, 59, 59, 0, time.Local))
	if res != nil {
		t.Log(res.Name, res.Date, res.Countdown)
		return
	}
	t.Log("No")
}
