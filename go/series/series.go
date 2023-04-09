package series

import "strings"

func All(n int, s string) []string {
	var results []string
	var digits []rune = []rune(s)

	if n > len(digits) {
		return nil
	}

	for i := n; i <= len(digits); i++ {
		results = append(results, string(digits[i - n: i]))
	}

	return results
}

func First(n int, s string) (string, bool) {
	var result strings.Builder

	for _, digit := range s {
		if n == 0 {
			break
		}

		result.WriteRune(digit)
		n--
	}

	return result.String(), n == 0
}

func UnsafeFirst(n int, s string) string {
	result, _ := First(n, s)

	return result
}
