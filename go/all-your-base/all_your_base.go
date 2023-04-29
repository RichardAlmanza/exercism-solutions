package allyourbase

import (
	"errors"
)

var (
	ErrInputBase  = errors.New("input base must be >= 2")
	ErrInput      = errors.New("all digits must satisfy 0 <= d < input base")
	ErrOutputBase = errors.New("output base must be >= 2")
)

func trimDigits(digits []int) []int {
	for i, v := range digits {
		if v != 0 {
			return digits[i:]
		}
	}

	return digits
}

func defaultZero(digits []int) []int {
	if len(digits) == 0 {
		return []int{0}
	}

	return digits
}

func validDigits(base int, digits []int) bool {
	for _, value := range digits {
		switch {
		case base <= value,
			value < 0:
			return false
		}
	}

	return true
}

func normalizeInput(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
	case inputBase < 2:
		return nil, ErrInputBase
	case outputBase < 2:
		return nil, ErrOutputBase
	}

	var digits = trimDigits(inputDigits)

	if !validDigits(inputBase, digits) {
		return nil, ErrInput
	}

	digits = defaultZero(digits)

	return digits, nil
}

func reverseDigits(digits []int) []int {
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}

func toBase10(base int, digits []int) uint64 {
	var numBase10 uint64

	for _, value := range digits {
		numBase10 *= uint64(base)
		numBase10 += uint64(value)
	}

	return numBase10
}

func toBaseN(number uint64, baseN int) []int {
	var digits []int = make([]int, 0, 8)

	for number > 0 {
		digit := int(number % uint64(baseN))
		number /= uint64(baseN)

		digits = append(digits, digit)
	}

	digits = reverseDigits(digits)
	digits = defaultZero(digits)

	return digits
}

func fromBaseAToB(baseA int, digitsA []int, baseB int) []int {
	return toBaseN(
		toBase10(baseA, digitsA),
		baseB,
	)
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	var output []int

	digits, err := normalizeInput(inputBase, inputDigits, outputBase)

	if err != nil {
		return nil, err
	}

	output = fromBaseAToB(inputBase, digits, outputBase)

	return output, nil
}
