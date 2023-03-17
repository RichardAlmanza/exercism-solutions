// hamming calculates the hamming distances
package hamming

import "errors"

// Distance returns the hamming distance between two DNA strings
func Distance(a, b string) (int, error) {
	var err error
	var distance int
	var lenDifference int = len(a) - len(b)

	if lenDifference != 0 {
		err = errors.New("these DNA sequences don't length match")

		return distance, err
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
