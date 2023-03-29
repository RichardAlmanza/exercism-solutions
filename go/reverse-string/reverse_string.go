package reverse

import (
	"strings"
)

// Reverse return the reversed input
func Reverse(input string) string {
	var temp []rune = []rune(input)
	var reversedInput strings.Builder

	for index := len(temp) - 1; index >= 0; index-- {
		reversedInput.WriteRune(temp[index])
	}

	return reversedInput.String()
}
