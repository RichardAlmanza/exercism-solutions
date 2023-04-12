package wordy

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	reOperationStructure *regexp.Regexp
	reQuestionStructure *regexp.Regexp

	operations map[string]func(int, int)int
)

func init() {
	const regexNumber string = `(-?\d+)`
	const regexOperation string = `(plus|minus|multiplied by|divided by) %s`
	const regexQuestion string = `^What is %s( %s)*\?$`

	var reOperation string = fmt.Sprintf(regexOperation, regexNumber)
	var reQuestion string = fmt.Sprintf(regexQuestion, regexNumber, reOperation)

	reOperationStructure = regexp.MustCompile(reOperation)
	reQuestionStructure = regexp.MustCompile(reQuestion)

	operations = map[string]func(int, int)int{
		"plus": func(a, b int) int {return a + b},
		"minus": func(a, b int) int {return a - b},
		"multiplied by": func(a, b int) int {return a * b},
		"divided by": func(a, b int) int {return a / b},
	}
}

func Answer(question string) (int, bool) {
	if !reQuestionStructure.MatchString(question) {
		return 0, false
	}

	base, _ := strconv.Atoi(reQuestionStructure.FindStringSubmatch(question)[1])

	for _, re := range reOperationStructure.FindAllStringSubmatch(question, -1) {
		operation := re[1]
		number, _ := strconv.Atoi(re[2])

		base = operations[operation](base, number)
	}

	return base, true

}
