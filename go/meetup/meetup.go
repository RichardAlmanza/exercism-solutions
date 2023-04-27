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

func moveToWDay(date time.Time, wDay time.Weekday) time.Time {
	var diffDay time.Duration = time.Duration(wDay - date.Weekday())

	if diffDay < 0 {
		diffDay	+= 7
	}

	date = date.Add(24 * time.Hour * diffDay)

	return date
}

func moveWeek(date time.Time, nWeeks time.Duration) time.Time {
	date = date.Add(24 * time.Hour * 7 * nWeeks)

	return date
}

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var date time.Time
	var day int = 1

	switch wSched {
	case Last:
		month++
	case Teenth:
		day = 13
	}

	date = time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
	date = moveToWDay(date, wDay)

	switch wSched {
	case Second, Third, Fourth:
		date = moveWeek(date, time.Duration(wSched))
	case Last:
		date = moveWeek(date, -1)
	}

	return date.Day()
}
