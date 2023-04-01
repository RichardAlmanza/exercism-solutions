package proverb

import "fmt"

const proverbFormat = "For want of a %s the %s was lost."

// Proverb return proverb []string with the items in rhyme
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}

	var proverb []string = make([]string, 0, len(rhyme))

	for previousIndex, item := range rhyme[1:] {
		proverb = append(proverb, fmt.Sprintf(proverbFormat, rhyme[previousIndex], item))
	}

	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return proverb
}
