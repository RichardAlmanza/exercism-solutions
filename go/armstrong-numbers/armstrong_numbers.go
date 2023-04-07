package armstrong

func countDigits(n int) int {
	var counter int = 1

	for n > 9 {
		n /= 10
		counter++
	}

	return counter
}

func powInt(base, exponent int) int {
	var pow int = 1

	for i := 0; i < exponent; i++ {
		pow *= base
	}

	return pow
}

func IsNumber(n int) bool {
	var checkSum int
	var nDigits int = countDigits(n)

	for i := n; i > 0; i /= 10 {
		checkSum += powInt(i % 10, nDigits)
	}

	return checkSum == n
}
