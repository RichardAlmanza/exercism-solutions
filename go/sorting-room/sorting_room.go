package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf(
		"This is a box containing the number %.1f",
		float32(nb.Number()),
	)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, _ := fnb.(FancyNumber)
	var intValue, _ = strconv.Atoi(fn.Value())
	return intValue
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	fn, _ := fnb.(FancyNumber)
	var intValue, _ = strconv.Atoi(fn.Value())
	return fmt.Sprintf("This is a fancy box containing the number %d.0", intValue)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var description string

	switch valueToDescribe := i.(type) {
	case int:
		description = DescribeNumber(float64(valueToDescribe))
	case float64:
		description = DescribeNumber(valueToDescribe)
	case NumberBox:
		description = DescribeNumberBox(valueToDescribe)
	case FancyNumberBox:
		description = DescribeFancyNumberBox(valueToDescribe)
	default:
		description = "Return to sender"
	}

	return description
}
