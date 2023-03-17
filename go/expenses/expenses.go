package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filtered []Record = make([]Record, 0, len(in))

	for _, v := range in {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}

		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	var sum float64
	var recordsInRange []Record = Filter(in, ByDaysPeriod(p))

	for _, record := range recordsInRange {
		sum += record.Amount
	}

	return sum
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	var sum float64
	var err error
	recordsByCategory := Filter(in, ByCategory(c))

	sum = TotalByPeriod(recordsByCategory, p)

	if len(Filter(in, ByCategory(c))) == 0 {
		sum = 0.0
		err = fmt.Errorf("unknown category %s", c)
	}

	return sum, err
}
