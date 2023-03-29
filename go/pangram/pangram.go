package pangram

import "strings"

type alphabet map[rune]uint8

const englishAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func NewAlphabet(sAlphabet string) alphabet {
	var english alphabet = make(alphabet, len(sAlphabet))

	for _, runeKey := range sAlphabet {
		english[runeKey] = 0
	}

	return english
}

func IsPangram(input string) bool {
	var counter alphabet = NewAlphabet(englishAlphabet)

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
