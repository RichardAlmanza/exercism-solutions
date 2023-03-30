package romannumerals

import (
	"errors"
	"strings"
)

var symbols = []rune("\x00\x00MDCLXVI")

func writeRoman(roman *strings.Builder, unit int, index int) {
	var numberOfOnes int = unit % 5
	var unitHasFive bool = unit >= 5
	var one rune = symbols[index]
	var five rune = symbols[index-1]
	var ten rune = symbols[index-2]

	if numberOfOnes == 4 {
		roman.WriteRune(one)

		if unitHasFive {
			roman.WriteRune(ten)
		} else {
			roman.WriteRune(five)
		}

	} else {
		if unitHasFive {
			roman.WriteRune(five)
		}

		for i := 0; i < numberOfOnes; i++ {
			roman.WriteRune(one)
		}
	}

}

func ToRomanNumeral(input int) (string, error) {
	if input > 3999 || input < 1 {
		return "", errors.New("don't joke around")
	}

	var roman strings.Builder
	var value int = 1000

	for index := 2; index < len(symbols); index += 2 {
		writeRoman(&roman, input/value, index)

		input = input % value
		value = value / 10
	}

	return roman.String(), nil
}
