package protein

import (
	"errors"
)


var (
	ErrInvalidBase error = errors.New("the rna is broken")
	ErrStop error = errors.New("stop codon")

	codonsToStop = map[string]bool{
		"UAA": true,
		"UAG": true,
		"UGA": true,
	}
	codonsProteins = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
	}
)

func FromRNA(rna string) ([]string, error) {
	var codons []string
	var rnaCode []rune = []rune(rna)

	if len(rnaCode) % 3 != 0 {
		return nil, ErrInvalidBase
	}

	codons = make([]string, 0, len(rnaCode) / 3)

rnaLoop:
	for i := 0; i < len(rnaCode); i += 3 {
		codon := string(rnaCode[i : i+3])
		protein, err := FromCodon(codon)

		switch err {
		case ErrStop:
			break rnaLoop
		case ErrInvalidBase:
			return nil, ErrInvalidBase
		}

		codons = append(codons, protein)
	}

	return codons, nil
}

func FromCodon(codon string) (string, error) {
	if codonsToStop[codon] {
		return "", ErrStop
	}

	protein, exists := codonsProteins[codon]

	if !exists {
		return "", ErrInvalidBase
	}

	return protein, nil
}
