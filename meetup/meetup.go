package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First = WeekSchedule(iota)
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(ws WeekSchedule, w time.Weekday, m time.Month, year int) int {
	t := time.Date(year, m, 1, 0, 0, 0, 0, time.UTC)
	weekday := t.Weekday()
	if w >= weekday {
		t = t.AddDate(0, 0, int(w-weekday))
	} else {
		t = t.AddDate(0, 0, int(7+w-weekday))
	}

	switch ws {
	case First, Second, Third, Fourth:
		t = t.AddDate(0, 0, int(ws*7))
	case Last:
		t = t.AddDate(0, 0, 28)
		if t.Month() != m {
			t = t.AddDate(0, 0, -7)
		}
	case Teenth:
		for {
			t = t.AddDate(0, 0, 7)
			if t.Day() >= 13 {
				break
			}
		}

	}

	return t.Day()
}
