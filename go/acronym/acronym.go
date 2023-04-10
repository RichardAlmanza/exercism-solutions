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
		if previousIsSpace && !isSpace(char){
			acronym.WriteRune(unicode.ToUpper(char))
			previousIsSpace = false
		} else {
			previousIsSpace = isSpace(char)
		}
	}

	return acronym.String()
}
