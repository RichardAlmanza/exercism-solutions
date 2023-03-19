package collatzconjecture

import "errors"

// CollatzConjecture return the
// number of steps to reach 1 given
// n as starting point
func CollatzConjecture(n int) (int, error) {
	var steps int

	if n <= 0 {
		return 0, errors.New("the given number needs to be more than 0")
	}

	for n > 1 {
		steps ++

		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3 * n + 1
		}
	}

	return steps, nil
}
