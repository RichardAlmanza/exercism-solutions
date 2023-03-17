package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (err *SillyNephewError) Error() string {
	msg := "silly nephew, there cannot be %d cows"
	return fmt.Sprintf(msg, err.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	switch {
	case err != nil && err != ErrScaleMalfunction:
		break

	case fodderAmount < 0:
		err = errors.New("negative fodder")

	case err == ErrScaleMalfunction:
		fodderAmount = fodderAmount * 2
		err = nil

	case cows == 0:
		cows = 1
		err = errors.New("division by zero")

	case cows < 0:
		err = &SillyNephewError{cows: cows}
	}

	if err != nil {
		fodderAmount = 0
	}

	return fodderAmount / float64(cows), err
}
