package rotationalcipher

import "strings"

func rotate(character rune, shiftKey int) rune {
	var base rune

	switch {
	case 'A' <= character && character <= 'Z':
		base = 'A'
	case 'a' <= character && character <= 'z':
		base = 'a'
	default:
		return character
	}

	character -= base
	character += rune(shiftKey)
	character %= 26
	character += base

	return character
}

func RotationalCipher(plain string, shiftKey int) string {
	var ciphered strings.Builder

	for _, character := range plain {
		ciphered.WriteRune(rotate(character, shiftKey))
	}

	return ciphered.String()
}
