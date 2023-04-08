package prime

import "math"


type uIntStd uint64
type naturals [baseSet]uIntStd
type naturalsSet map[uIntStd]naturals
type primeSet map[uIntStd]bool

type primes struct {
	lastLimit uIntStd
	lastNth uIntStd
	numbers naturalsSet
	primeNumbers primeSet
}

const baseSet uIntStd = 1024

var (
	numbers *primes
)

func init() {
	const startWith uIntStd = 100000

	numbers = &primes{}
	numbers.numbers = make(naturalsSet)
	numbers.primeNumbers = make(primeSet)

	numbers.numbers[0] = naturals{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	numbers.lastNth = 10

	numbers.LoadPrimesLimit(startWith)
}

func NewPrimeNumbers() *primes {
	return numbers
}

func Factors(n int64) []int64 {
	var newFactors []int64
	var primeSet *primes = NewPrimeNumbers()

	for i := 1; i < int(n); {
		if primeSet.GetPrime(uIntStd(n)) {
			newFactors = append(newFactors, n)
			break
		}

		prime := primeSet.GetNth(uIntStd(i))

		if uIntStd(n) % prime == 0 {
			newFactors = append(newFactors, int64(prime))
			n /= int64(prime)

			if n == 1 {
				break
			}

            i = 1
			continue
		}

        i++
	}

	return newFactors
}

func (p *primes) Len() int {
	return int(p.lastNth)
}

func (p *primes) LoadPrimesLimit(limit uIntStd) {
	if limit <= p.lastLimit {
		return
	}

	p.loadPrimeLimit(limit)
}

func (p *primes) GetNth(i uIntStd) uIntStd {
	for p.lastNth < i {
		p.loadNextPrime()
	}

	i--
	return p.numbers[i / baseSet][i % baseSet]
}

func (p *primes) Append(newPrime uIntStd) {
	var lastSet naturals = p.numbers[p.lastNth / baseSet]
	lastSet[p.lastNth % baseSet] = newPrime

	p.numbers[p.lastNth / baseSet] = lastSet
	p.primeNumbers[newPrime] = true
	p.lastNth++
	p.lastLimit = newPrime
}

func (p *primes) GetPrime(candidate uIntStd) bool {
	return p.primeNumbers[candidate]
}

func (p *primes) isPrime(candidate uIntStd) bool {
	var primeLimit uIntStd = uIntStd(math.Sqrt(float64(candidate)))

	for i := uIntStd(1); i < p.lastNth; i++ {
		prime := p.GetNth(i)

		if candidate % prime == 0 {
			return false
		}

		if prime > primeLimit {
			return true
		}
	}

	return true
}

func (p *primes) loadNextPrime() uIntStd {
	var prime uIntStd = p.GetNth(p.lastNth)

	for {
		prime += 2

		if p.isPrime(prime) {
			p.Append(prime)
			break
		}
	}

	return prime
}

func (p *primes) loadPrimeLimit(limit uIntStd) {
	var prime uIntStd = p.GetNth(p.lastNth)

	for p.lastLimit <= limit {
		prime += 2

		if p.isPrime(prime) {
			p.Append(prime)
		}
	}
}