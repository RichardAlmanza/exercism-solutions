package lsproduct

import "errors"

var (
	ErrSmallSpan error = errors.New("span must be smaller than string length")
	ErrNegativeSpan error = errors.New("span must not be negative")
	ErrNotDigit error = errors.New("digits input must only contain digits")
)

func isDigit(digit rune) bool {
	return '0' <= digit && digit <= '9'
}

func rDigitToInt64 (digit rune) int64 {
	return int64(digit - '0')
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var factorCounter int
	var maxProduct int64
	var product int64 = 1
	var digitChars []rune = []rune(digits)

	if span < 0 {
		return 0, ErrNegativeSpan
	}

	if len(digitChars) < span {
		return 0, ErrSmallSpan
	}

	if span == 0 {
		return 1, nil
	}

	for i := 0; i < len(digitChars); i++ {
		digitChar := digitChars[i]

		if !isDigit(digitChar) {
			return 0, ErrNotDigit
		}

		if digitChar == '0' {
			product = 1
			factorCounter = 0
			continue
		}

		product *= rDigitToInt64(digitChar)
		factorCounter++

		if factorCounter > span {
			product /= rDigitToInt64(digitChars[i-span])
		}

		if product > maxProduct && factorCounter >= span {
			maxProduct = product
		}
	}

	return maxProduct, nil
}
