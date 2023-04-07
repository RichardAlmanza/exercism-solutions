package armstrong

import "math"

func countDigits(n int) int {
	var counter int = 1

	for n > 9 {
		n /= 10
		counter++
	}

	return counter
}

func powInt(base, exponent int) int {
	return int(
		math.Pow(
			float64(base), float64(exponent),
		),
	)
}

func IsNumber(n int) bool {
	var checkSum int
	var tempN int = n
	var nDigits int = countDigits(n)

	for i := 0; i < nDigits; i++ {
		checkSum += powInt(tempN % 10, nDigits)
		tempN /= 10
	}

	return checkSum == n
}
