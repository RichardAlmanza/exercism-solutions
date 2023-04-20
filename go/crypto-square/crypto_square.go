package cryptosquare

import (
	"fmt"
	"math"
	"strings"
)

func isCharValid(char rune) bool {
	if 'a' <= char && char <= 'z' {
		return true
	}

	if '0' <= char && char <= '9' {
		return true
	}

	return false
}

func normalize(input string) string {
	const diffAa rune = 'A' - 'a'
	var output strings.Builder

	for _, char := range input {
		if 'A' <= char && char <= 'Z' {
			char -= diffAa
		}

		if !isCharValid(char) {
			continue
		}

		output.WriteRune(char)
	}

	return output.String()
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

	for i := 0; i < rows; i++ {
		start := i * cols
		finish := start + cols

		if finish > len(input) {
			pad = strings.Repeat(" ", finish - len(input))
			finish = len(input)
		}

		output.WriteString(fmt.Sprintf("%s%s", input[start: finish], pad))
	}

	return output.String()
}

func transposeSquaredText(input string, rows, cols int) string {
	var output strings.Builder

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			output.WriteByte(input[i + j * cols])
		}

		if i + 1 < cols {
			output.WriteRune(' ')
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
