package meetup

import "time"

type WeekSchedule int

const (
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Teenth
	Last
)

func firstDay(date time.Time) time.Time {
	if date.Day() != 1 {
		date = time.Date(
			date.Year(), date.Month(), 1,
			date.Hour(), date.Minute(), date.Second(), date.Nanosecond(),
			date.Location(),
		)
	}

	return date
}

func moveToFirstWDay(date time.Time, wDay time.Weekday) time.Time {
	var diffDay time.Duration = time.Duration(wDay - date.Weekday())

	if diffDay < 0 {
		diffDay	+= 7
	}

	date = firstDay(date)
	date = date.Add(24 * time.Hour * diffDay)

	return date
}

func nextWeek(date time.Time) time.Time {
	date = date.Add(24 * time.Hour * 7)

	return date
}

func prevWeek(date time.Time) time.Time {
	date = date.Add(-24 * time.Hour * 7)

	return date
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var date time.Time = time.Date(year, month, 1, 12, 0, 0, 0, time.UTC)

	date = moveToFirstWDay(date, wDay)

	switch wSched {
	case Second, Third, Fourth:
		for i := 0; i < int(wSched); i++ {
			date = nextWeek(date)
		}
	case Teenth:
		for date.Day() < 13 {
			date = nextWeek(date)
		}
	case Last:
		for date.Month() == month {
			date = nextWeek(date)
		}

		date = prevWeek(date)
	}

	return date.Day()
}
