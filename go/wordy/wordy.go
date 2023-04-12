package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

const (
	regexQuestion string = `^What is ((-?\d+)( (plus|minus|multiplied by|divided by) (-?\d+))*)\?$`
)

var (
	reQuestionStructure *regexp.Regexp = regexp.MustCompile(regexQuestion)

	operations map[string]func(int, int)int = map[string]func(int, int)int{
		"plus": func(a, b int) int {return a + b},
		"minus": func(a, b int) int {return a - b},
		"multiplied": func(a, b int) int {return a * b},
		"divided": func(a, b int) int {return a / b},
	}
)

func Answer(question string) (int, bool) {
	if !reQuestionStructure.MatchString(question) {
		return 0, false
	}

	var counter int
	var number int
	var operation string
	var toOperate []string = strings.Fields(reQuestionStructure.FindStringSubmatch(question)[1])

	base, _ := strconv.Atoi(toOperate[0])

	for _, element := range toOperate[1:] {
		if element == "by" {
			continue
		}

		switch counter {
		case 0:
			operation = element
		case 1:
			number, _ = strconv.Atoi(element)
		}

		counter++

		if counter == 2 {
			base = operations[operation](base, number)
			counter = 0
		}
	}

	return base, true

}
