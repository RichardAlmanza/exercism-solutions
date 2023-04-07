package atbash

import (
	"strings"
	"unicode"
)

var groupSize int = 5
var alphabet []rune = []rune("abcdefghijklmnopqrstuvwxyz")

func Atbash(s string) string {
	var ciphered strings.Builder
	var counter int

	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			continue
		}

		if unicode.IsLetter(char) {
			char = unicode.ToLower(char)
			
			index := int(char - 'a' + 1)
			index = len(alphabet) - index

			char = alphabet[index]
		}

		if counter % groupSize == 0 && counter > 0 {
			ciphered.WriteRune(' ')
		}

		ciphered.WriteRune(char)
		counter++
	}

	return ciphered.String()
}
