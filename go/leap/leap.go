// leap determines if a given year is leap
package leap

// IsLeapYear returns if the year is leap
func IsLeapYear(year int) bool {
	var isLeap bool

	// By400 || ( !By100 && By4 )
	switch {
	case year % 400 == 0:
		isLeap = true

	case year % 100 == 0:
		isLeap = false

	case year % 4 == 0:
		isLeap = true
	}

	return isLeap
}
