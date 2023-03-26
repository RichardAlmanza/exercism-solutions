package strand

import "strings"

type translation map[rune]rune

var dnaToRna = translation{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var newRna strings.Builder

	for _, nucleotide := range dna {
		newRna.WriteRune(dnaToRna[nucleotide])
	}

	return newRna.String()
}
