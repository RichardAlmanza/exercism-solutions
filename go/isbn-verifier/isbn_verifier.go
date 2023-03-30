package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	var cleanedISBN []rune
	var sum int
	var value int

	isbn = strings.ToUpper(isbn)
	cleanedISBN = []rune(strings.ReplaceAll(isbn, "-", ""))

	if len(cleanedISBN) != 10 {
		return false
	}

	for index, r := range cleanedISBN {
		runeIsX := r == 'X'

		if r < '0' && '9' > r && !runeIsX{
			return false
		}

		if runeIsX && index < 9 {
			return false
		}

		if runeIsX {
			r = '9' + 1
		}

		value = int(r - '0')

		sum += value * (10 - index)
	}

	return sum % 11 == 0
}
