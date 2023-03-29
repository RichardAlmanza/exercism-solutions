package pangram

import (
	"strings"
	"unicode/utf8"
)

type alphabet map[rune]uint8

const EnglishAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NewAlphabet returns an alphabet (map[rune]uint8) based on sAlphabet (string)
func NewAlphabet(sAlphabet string) alphabet {
	var newAlpha alphabet = make(alphabet, utf8.RuneCountInString(sAlphabet))

	for _, runeKey := range sAlphabet {
		newAlpha[runeKey] = 0
	}

	return newAlpha
}

// IsPangram verifies if input is a pangram
func IsPangram(input string) bool {
	var counter alphabet = NewAlphabet(EnglishAlphabet)

	input = strings.ToUpper(input)

	for _, runeKey := range input {
		_, exists :=  counter[runeKey]

		if exists {
			counter[runeKey]++
		}
	}

	for _, value := range counter {
		if value == 0 {
			return false
		}
	}

	return true
}
