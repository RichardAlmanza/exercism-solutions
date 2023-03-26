package strand

import "strings"

type transcription map[rune]byte

var dnaToRna = transcription{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA transcripts DNA to RNA
func ToRNA(dna string) string {
	var newRna strings.Builder

	for _, nucleotide := range dna {
		newRna.WriteByte(dnaToRna[nucleotide])
	}

	return newRna.String()
}
