package acronym

import (
	"strings"
	"unicode"
)

func isSpace(char rune) bool {
	switch char {
	case ' ', '-', '_':
		return true
	}

	return false
}

// Abbreviate return the acronym for input s
func Abbreviate(s string) string {
	var previousIsSpace bool = true
	var acronym strings.Builder

	for _, char := range s {
		isCharSpace := isSpace(char)

		if previousIsSpace && !isCharSpace {
			acronym.WriteRune(unicode.ToUpper(char))
		}

		previousIsSpace = isCharSpace
	}

	return acronym.String()
}
