package twelve

import (
	"fmt"
	"strings"
)

var (
	gifts = []string{
		"",
		"Partridge in a Pear Tree",
		"Turtle Doves",
		"French Hens",
		"Calling Birds",
		"Gold Rings",
		"Geese-a-Laying",
		"Swans-a-Swimming",
		"Maids-a-Milking",
		"Ladies Dancing",
		"Lords-a-Leaping",
		"Pipers Piping",
		"Drummers Drumming",
	}

	numbers = map[int][2]string{
		1: {"first", "a"},
		2: {"second", "two"},
		3: {"third", "three"},
		4: {"fourth", "four"},
		5: {"fifth", "five"},
		6: {"sixth", "six"},
		7: {"seventh", "seven"},
		8: {"eighth", "eight"},
		9: {"ninth", "nine"},
		10: {"tenth", "ten"},
		11: {"eleventh", "eleven"},
		12: {"twelfth", "twelve"},
	}

	firstLine string = "On the %s day of Christmas my true love gave to me: %s"
)


func Verse(i int) string {
	var newVerse strings.Builder

	for index := i; index > 0; index-- {
		newVerse.WriteString(numbers[index][1])
		newVerse.WriteRune(' ')
		newVerse.WriteString(gifts[index])

		switch index {
		case 1:
			newVerse.WriteRune('.')
		case 2:
			newVerse.WriteString(", and ")
		default:
			newVerse.WriteString(", ")
		}
	}

	return fmt.Sprintf(firstLine, numbers[i][0], newVerse.String())
}

func Song() string {
	var fullSong strings.Builder

	for i := 1; i < 12; i++ {
		fullSong.WriteString(Verse(i))
		fullSong.WriteRune('\n')
	}

	fullSong.WriteString(Verse(12))

	return fullSong.String()
}
