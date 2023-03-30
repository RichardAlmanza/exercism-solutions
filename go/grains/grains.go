package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 {
		return 0, errors.New("why would you ask for 0 squares or less")
	}

	if number > 64 {
		return 0, errors.New("there are no more than 64 squares")
	}

	return uint64(math.Pow(2, float64(number - 1))), nil
}

func Total() uint64 {
	totalGrains, _ := Square(64)
	return totalGrains * 2 - 1
}
