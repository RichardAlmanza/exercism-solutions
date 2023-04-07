package cipher

import (
	"strings"
	"unicode"
)

type shift int8
type vigenere []rune


func encodeLetter(letter, key rune) rune {
	letter = unicode.ToLower(letter)
	letter = (key + letter - 'a') % 26 + 'a'

	return letter
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 {
		return nil
	}

	if distance < -25 || 25 < distance  {
		return nil
	}

	if distance < 0 {
		distance += 26
	}

	return shift(distance)
}

func (c shift) Encode(input string) string {
	var encoded strings.Builder

	for _, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}

		char = encodeLetter(char, rune(c))

		encoded.WriteRune(char)
	}

	return encoded.String()
}

func (c shift) Decode(input string) string {
	return shift(26 - int(c)).Encode(input)
}

func NewVigenere(key string) Cipher {
	var compiled strings.Builder
	var checkSum rune

	for _, char := range key {
		if char < 'a' || 'z' < char {
			return nil
		}

		char -= 'a'
		checkSum += char

		compiled.WriteRune(char)
	}

	if checkSum == '\x00' {
		return nil
	}

	return vigenere(compiled.String())
}

func (v vigenere) Encode(input string) string {
	var encoded strings.Builder
	var index int

	for _, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}

		index %= len(v)

		char = encodeLetter(char, v[index])

		encoded.WriteRune(char)

		index++
	}

	return encoded.String()
}

func (v vigenere) Decode(input string) string {
	var decoderKey vigenere = make(vigenere, len(v))

	for i, key := range v {
		decoderKey[i] = 26 - key
	}

	return decoderKey.Encode(input)
}
