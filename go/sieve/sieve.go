package sieve

func Sieve(limit int) []int {
	var numbers []bool
	var primes []int

	limit++

	numbers = make([]bool, limit)
	primes = make([]int, 0, limit)

	for i := 2; i < limit; i++ {
		if numbers[i] { // is Composite?
			continue
		}

		for ii := 2; ii * i < limit; ii++ {
			numbers[i * ii] = true // all theses numbers are composite
		}

		primes = append(primes, i)
	}

	return primes
}
