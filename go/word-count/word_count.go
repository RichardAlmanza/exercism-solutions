package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func isToWrite(r rune, previousIsToWrite bool) bool {
	var tw bool

	if r == '\'' {
		return previousIsToWrite
	}

	switch {
	case 'a' <= r && r <= 'z':
		tw = true
	case 'A' <= r && r <= 'Z':
		tw = true
	case '0' <= r && r <= '9':
		tw = true
	}

	return tw
}

func WordCount(phrase string) Frequency {
	var writer strings.Builder
	var wordCounter Frequency = make(Frequency)
	var previousIsToWrite bool

	for _, character := range phrase {
		charIsToWrite := isToWrite(character, previousIsToWrite)

		if !charIsToWrite  && !previousIsToWrite{
			continue
		}

		if charIsToWrite {
			writer.WriteRune(unicode.ToLower(character))
		} else if previousIsToWrite {
			wordCounter[strings.TrimRight(writer.String(), "'")]++
			writer.Reset()
		}

		previousIsToWrite = charIsToWrite
	}

	if previousIsToWrite {
		wordCounter[strings.TrimRight(writer.String(), "'")]++
	}

	return wordCounter
}
