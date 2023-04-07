package prime

import "errors"

type naturals [baseSet]uint32
type naturalSets map[uint16]naturals

const baseSet int = 20

var (
	ErrNthPrime error = errors.New("cannot be calculated")
	primeNumbers naturalSets
	lastNth int
)

func init() {
	primeNumbers = make(naturalSets, 20)

	primeNumbers[0] = naturals{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	lastNth = 10
}

func getNthPrime(i int) uint32 {
	i--
	return primeNumbers[uint16(i / baseSet)][i % baseSet]
}

func SetNthPrime(newPrime uint32) {
	var lastSet naturals = primeNumbers[uint16(lastNth / baseSet)]
	lastSet[lastNth % baseSet] = newPrime

	primeNumbers[uint16(lastNth / baseSet)] = lastSet
	lastNth++
}

func isPrime(candidate uint32) bool {
	for i := 1; i < lastNth; i++ {
		prime := getNthPrime(i)

		if candidate % prime == 0 {
			return false
		}
	}

	return true
}

func loadPrime(n int) uint32 {
	var prime uint32 = getNthPrime(lastNth)

	for lastNth < n {
		prime += 2

		if isPrime(prime) {
			SetNthPrime(prime)
		}
	}

	return prime
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, ErrNthPrime
	}

	if n <= lastNth {
		return int(getNthPrime(n)), nil
	}

	return int(loadPrime(n)), nil
}
