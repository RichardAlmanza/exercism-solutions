package cryptosquare

import (
	"math"
	"strings"
)

func isCharValid(char rune) bool {
	switch {
	case 'a' <= char && char <= 'z',
		'0' <= char && char <= '9':
		return true
	}

	return false
}

func normalizer(char rune) rune {
	const diffAa rune = 'A' - 'a'

	switch {
	case isCharValid(char):
		break
	case 'A' <= char && char <= 'Z':
		char -= diffAa
	default:
		char = -1
	}

	return char
}

func normalize(input string) string {
	return strings.Map(normalizer, input)
}

func calculatesRowsAndColumns(length int) (int, int) {
	var columns int = int(math.Ceil(math.Sqrt(float64(length))))
	var rows int = columns - 1

	if columns * rows < length {
		rows++
	}

	return rows, columns
}

func squareText(input string, rows, cols int) string {
	var output strings.Builder
	var pad string
	var width int = rows*cols - len(input)

	if width > 0 {
		pad = strings.Repeat(" ", width)
	}

	output.WriteString(input)
	output.WriteString(pad)

	return output.String()
}

func transposeSquaredText(input string, rows, cols int) string {
	var output strings.Builder

	for i := 0; i < cols; i++ {
		if i > 0 {
			output.WriteByte(' ')
		}

		for j := 0; j < rows; j++ {
			output.WriteByte(input[i+j*cols])
		}
	}

	return output.String()
}

func squareAndTransposeText(input string) string {
	var output string

	rows, cols := calculatesRowsAndColumns(len(input))

	output = squareText(input, rows, cols)
	output = transposeSquaredText(output, rows, cols)

	return output
}

func Encode(pt string) string {
	return squareAndTransposeText(normalize(pt))
}
