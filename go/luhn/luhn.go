package luhn

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Valid return if the given id is valid using the luhn algorithm
func Valid(id string) bool {
	var digit int
	var checkSum int
	var lenId int
	var isIdEven bool

	id = strings.ReplaceAll(id, " ", "")
	lenId = utf8.RuneCountInString(id)

	if lenId < 2 {
		return false
	}

	isIdEven = lenId % 2 == 0

	for i, r := range id {
		if !unicode.IsDigit(r) {
			return false
		}

		digit = int(r - '0')

		if (i % 2 == 0) == isIdEven {
			checkSum += ( 2 * digit) % 10

			if digit > 4 {
				checkSum++
			}
		}else{
			checkSum += digit
		}
	}

	return checkSum % 10 == 0
}
