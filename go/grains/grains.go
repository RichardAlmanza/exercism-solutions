package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("why would you ask for 0 squares or less")
	}

	if number > 64 {
		return 0, errors.New("there are no more than 64 squares")
	}

	// I saw the solution from paulwilljones. the solution reminded me
	// that Go can interact directly with bits and bytes using bitwise operations.
	// The magic of a better performance based on most low basic structure
	return 1 << (number - 1), nil
}

func Total() uint64 {
	totalGrains, _ := Square(64)
	return totalGrains << 1 - 1
}
