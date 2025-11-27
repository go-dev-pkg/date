package holiday

import (
	"time"

	"github.com/shopspring/decimal"
)

type Holiday struct {
	Name string
	Date time.Time
}

var Holidays = []Holiday{
	{
		Name: "感恩节",
		Date: time.Date(2025, time.November, 27, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "圣诞节",
		Date: time.Date(2025, time.December, 25, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "元旦",
		Date: time.Date(2026, time.January, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "除夕",
		Date: time.Date(2026, time.February, 16, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "春节",
		Date: time.Date(2026, time.February, 17, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "元宵节",
		Date: time.Date(2026, time.March, 3, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "清明节",
		Date: time.Date(2026, time.April, 5, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "劳动节",
		Date: time.Date(2026, time.May, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "端午节",
		Date: time.Date(2026, time.June, 19, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "中秋节",
		Date: time.Date(2026, time.September, 25, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "国庆节",
		Date: time.Date(2026, time.October, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "重阳节",
		Date: time.Date(2026, time.October, 18, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "万圣节",
		Date: time.Date(2026, time.November, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "感恩节",
		Date: time.Date(2026, time.November, 26, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "圣诞节",
		Date: time.Date(2026, time.December, 25, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "元旦",
		Date: time.Date(2027, time.January, 1, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "除夕",
		Date: time.Date(2027, time.February, 5, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "春节",
		Date: time.Date(2027, time.February, 6, 0, 0, 0, 0, time.Local),
	},
	{
		Name: "元宵节",
		Date: time.Date(2027, time.February, 20, 0, 0, 0, 0, time.Local),
	},
}

type Data struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
	Countdown int    `json:"countdown"`
}

// Reminder 节日提醒
func Reminder(beforeDay int, t ...time.Time) (data *Data) {
	_t := time.Now()
	if len(t) > 0 && !t[0].IsZero() {
		_t = t[0]
	}
	year := _t.Year()
	month := _t.Month()
	day := _t.Day()
	for _, v := range Holidays {
		if v.Date.Year() == year && v.Date.Month() == month && v.Date.Day() == day {
			return &Data{
				Name:      v.Name,
				Date:      v.Date.Format(time.DateOnly),
				Timestamp: v.Date.Unix(),
				Countdown: 0,
			}
		}
		if v.Date.Before(_t) {
			continue
		}
		countdown := int(((decimal.NewFromInt(v.Date.Unix() - _t.Unix()).Div(decimal.NewFromInt(86400))).Ceil()).BigInt().Int64())
		if countdown <= beforeDay {
			return &Data{
				Name:      v.Name,
				Date:      v.Date.Format(time.DateOnly),
				Timestamp: v.Date.Unix(),
				Countdown: countdown,
			}
		}
		break
	}
	return
}
