package summultiples

type multiples map[int]bool

func (m multiples) add(divisor, limit int) {
	if divisor == 0 {
		return
	}

	for factor := divisor; factor < limit; factor += divisor {
		m[factor] = true
	}
}

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	var unique multiples = make(multiples)

	for _, divisor := range divisors {
		unique.add(divisor, limit)
	}

	for key := range unique {
		sum += key
	}

	return sum
}
