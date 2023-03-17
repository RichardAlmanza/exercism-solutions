package scrabble

import "strings"

type scoreMap map[rune]uint8

var (
	singleton scoreMap
)

func newScoreMap() scoreMap {
	if singleton == nil {
		singleton = make(scoreMap, 28)

		singleton.addLetters("AEIOULNRST", 1)
		singleton.addLetters("DG", 2)
		singleton.addLetters("BCMP", 3)
		singleton.addLetters("FHVWY", 4)
		singleton.addLetters("K", 5)
		singleton.addLetters("JX", 8)
		singleton.addLetters("QZ", 10)
	}

	return singleton
}

func (sm scoreMap) addLetters(letters string, score uint8) {
	for _, r := range letters {
		sm[r] = score
	}
}

func (sm scoreMap) getScore(word string) int  {
	var finalScore int

	word = strings.ToUpper(word)

	for _, r := range word {
		finalScore += int(sm[r])
	}

	return finalScore
}

// Score calculates the score for the scrabble word
func Score(word string) int {
	return newScoreMap().getScore(word)
}