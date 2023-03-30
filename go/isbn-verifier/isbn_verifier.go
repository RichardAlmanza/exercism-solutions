package isbn

import (
	"regexp"
	"strings"
	"unicode"
)

// using regex, it is just a little less efficient but more readable than my previous solution
var isbnStructure *regexp.Regexp = regexp.MustCompile(`^\d{9}[xX\d]$`)

func IsValidISBN(isbn string) bool {
	var cleanedISBN []rune
	var sum int
	var value int

	isbn = strings.ReplaceAll(isbn, "-", "")

	if !isbnStructure.MatchString(isbn) {
		return false
	}

	cleanedISBN = []rune(isbn)

	if !unicode.IsDigit(cleanedISBN[9]) {
		cleanedISBN[9] = '9' + 1
	}

	for index, character := range cleanedISBN {
		value = int(character - '0')

		sum += value * (10 - index)
	}

	return sum % 11 == 0
}
