package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dateTime, _ := time.Parse("January 2, 2006 15:04:05", date)
	return dateTime.Before(time.Now())
}

// changeHour returns a new Time with a new hour. Minutes, seconds and
// nanoseconds are set to 0
func changeHour(date time.Time, newHour int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(),
		newHour, 0, 0, 0, date.Location())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateTime, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	var isBeforeNoon bool = dateTime.Before(changeHour(dateTime, 12))
	var isBeforeNight bool = dateTime.Before(changeHour(dateTime, 18))

	if !isBeforeNoon && isBeforeNight {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var dateTime time.Time = Schedule(date)
	return dateTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
