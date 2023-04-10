package bottlesong

import (
	"fmt"
	"strings"
)

var (
	lines = [4]string{
		"%s green bottle%s hanging on the wall,",
		"%s green bottle%s hanging on the wall,",
		"And if one green bottle should accidentally fall,",
		"There'll be %s green bottle%s hanging on the wall.",
	}

	numbers = map[int]string{
		0: "No",
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
		10: "Ten",
	}
)

func getPlurals(quantity int) string {
	var result string

	if quantity != 1 {
		result = "s"
	}

	return result
}

func verse(number int) []string {
	var newVerse []string = make([]string, 0, 4)
	var sNumber string = numbers[number]
	var quantity string = getPlurals(number)

	for i, line := range lines {
		switch i {
		case 2:
			newVerse = append(newVerse, line)
		case 3:
			sNumber = strings.ToLower(numbers[number - 1])
			quantity = getPlurals(number - 1)
			fallthrough
		default:
			newVerse = append(newVerse, fmt.Sprintf(line, sNumber, quantity))
		}
	}

	return newVerse
}

func Recite(startBottles, takeDown int) []string {
	var maxLines int = 5 * takeDown // 4 lines per stanza (takedown) plus an empty line
	var result []string

	result = make([]string, 0, maxLines)

	for i := startBottles; i > startBottles - takeDown; i-- {
		result = append(result, verse(i)...)
		result = append(result, "")
	}

	return result[:maxLines - 1] // removes the last empty line
}
