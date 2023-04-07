package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

const phoneNumberFormat string = "(%s) %s-%s"

var (
	ErrInvalidNumber error = errors.New("invalid phone number")
	phoneNumberStructure *regexp.Regexp = regexp.MustCompile(`^1?(([2-9]\d\d){2}\d{4})$`)
)

func isValidNumberChar(char rune) bool {
	if '0' <= char && char <= '9' {
		return true
	}

	switch char {
	case ' ', '-', '.', '(', ')', '+':
		return true
	}

	return false
}

func Number(phoneNumber string) (string, error) {
	var cleanedNumber strings.Builder
	var number string

	for _, char := range phoneNumber {
		if !isValidNumberChar(char) {
			return "", ErrInvalidNumber
		}

		if unicode.IsDigit(char) {
			cleanedNumber.WriteRune(char)
		}
	}

	number = cleanedNumber.String()

	if !phoneNumberStructure.MatchString(number) {
		return "", ErrInvalidNumber
	}

	number = phoneNumberStructure.FindAllStringSubmatch(number, 1)[0][1]

	return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err == nil {
		phoneNumber = phoneNumber[:3]
	}

	return phoneNumber, err
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)

	if err == nil {
		phoneNumber = fmt.Sprintf(
			phoneNumberFormat,
			phoneNumber[:3],
			phoneNumber[3:6],
			phoneNumber[6:],
		)
	}

	return phoneNumber, err
}
