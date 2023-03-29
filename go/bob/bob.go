package bob

import (
	"regexp"
	"strings"
)

const (
	Asking uint8 = 1 << iota // 001
	Yelling                  // 010
	Silence                  // 100
)

var (
	isAsking *regexp.Regexp = regexp.MustCompile(`\?$`)
	hasWords *regexp.Regexp = regexp.MustCompile(`[A-Za-z]+`)
)

func getRemarkType(remark string) uint8 {
	var remarkType uint8
	remark = strings.TrimSpace(remark)

	if remark == "" {
		remarkType = remarkType | Silence
	}

	if isAsking.MatchString(remark) {
		remarkType = remarkType | Asking
	}

	if hasWords.MatchString(remark) && remark == strings.ToUpper(remark) {
		remarkType = remarkType | Yelling
	}

	return remarkType
}

func reply(remarkType uint8) string {
	var response string

	switch remarkType {
	case Silence:
		response = "Fine. Be that way!"
	case Yelling:
		response = "Whoa, chill out!"
	case Asking:
		response = "Sure."
	case Yelling | Asking:
		response = "Calm down, I know what I'm doing!"
	default:
		response = "Whatever."
	}

	return response
}

// Hey returns the bob's reply for remark
func Hey(remark string) string {
	return reply(getRemarkType(remark))
}
