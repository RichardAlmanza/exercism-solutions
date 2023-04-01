package anagram

import (
	// "sort"
	"strings"
)

type runeCounter map[rune]uint8
type anagram struct {
	lower string
	mapped runeCounter
}

func newAnagram(word string) anagram {
	word = strings.ToLower(word)

	return anagram{
		lower: word,
		mapped: newRuneCounter(word),
	}
}

func (a *anagram) isEqual(a2 *anagram) bool {
	if len(a.lower) != len(a2.lower) {
		return false
	}

	if a.lower == a2.lower {
		return false
	}

	return a.mapped.isEqual(a2.mapped)
}

func newRuneCounter(word string) runeCounter {
	var rc runeCounter = make(runeCounter, len(word))

	for _, character := range word {
		rc[character]++
	}

	return rc
}

func (rc runeCounter) isEqual(rc2 runeCounter) bool {
	if len(rc) != len(rc2) {
		return false
	}

	for character, count := range rc {
		count2, exists := rc2[character]

		if !exists {
			return false
		}

		if count != count2 {
			return false
		}
	}

	return true
}

func Detect(subject string, candidates []string) []string {
	var subjectCompilation anagram = newAnagram(subject)
	var detected []string = make([]string, 0, len(candidates))

	for _, candidate := range candidates {
		candidateCompilation := newAnagram(candidate)
		isAnagram := subjectCompilation.isEqual(&candidateCompilation)

		if isAnagram {
			detected = append(detected, candidate)
		}
	}

	return detected
}
