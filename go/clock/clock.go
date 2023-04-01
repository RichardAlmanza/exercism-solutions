package clock

import "fmt"

// Define the Clock type here.
type Clock int

const maxMinutes int = 1440
const clockFormat string = "%.2d:%.2d"

func (c Clock) mod() Clock {
	c %= Clock(maxMinutes)

	if c < 0 {
		c += Clock(maxMinutes)
	}

	return c
}

func New(h, m int) Clock {
	return Clock(m + h * 60).mod()
}

func (c Clock) Add(m int) Clock {
	return (c + Clock(m)).mod()
}

func (c Clock) Subtract(m int) Clock {
	return (c - Clock(m)).mod()
}

func (c Clock) String() string {
	return fmt.Sprintf(clockFormat, c / 60, c % 60)
}
