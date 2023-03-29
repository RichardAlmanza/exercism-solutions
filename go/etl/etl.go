package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var out map[string]int = make(map[string]int, 26)

	for score, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = score
		}
	}

	return out
}
