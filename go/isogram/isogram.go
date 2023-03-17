package isogram

import "strings"

// IsIsogram return if the word is a isogram
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letters :=  make(map[rune]struct{})

	for _, char := range word {
		_, exists := letters[char]

		if exists {
			return false
		}

		switch char {
		case ' ', '-':
			continue
		}

		letters[char] = struct{}{}
	}

	return true
}
