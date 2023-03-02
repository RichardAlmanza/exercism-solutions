package parsinglogfiles

import (
	"fmt"
	"regexp"
)

// IsValidLine function checks is the string argument is a valid log line
func IsValidLine(text string) bool {
	regex := regexp.MustCompile(`^\[((TRC)|(DBG)|(INF)|(WRN)|(ERR)|(FTL))\]`)
	return regex.MatchString(text)
}

// SplitLogLine function splits a string by this pattern `<[-=~\*]*>`.
// The function returns a slice of string
func SplitLogLine(text string) []string {
	regex := regexp.MustCompile(`<[-=~\*]*>`)
	return regex.Split(text, -1)
}

// CountQuotedPasswords function counts every time the word "password"
// is found between double quotes in a line, it counts it even if
// there are more words quoted and it is case insensitive.
// returns an int with the total count.
func CountQuotedPasswords(lines []string) int {
	var counter int
	regex := regexp.MustCompile(`(?i)"(\w*\s*)*password(\w*\s*)*"`)

	for _, line := range lines {
		counter += len(regex.FindAllString(line, -1))
	}

	return counter
}

// RemoveEndOfLineText removes the "end-of-line" text with the
// following number from the string.
func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(`end-of-line\d*`)
	return regex.ReplaceAllString(text, "")
}

// TagWithUserName tags lines with a user in it
func TagWithUserName(lines []string) []string {
	var user string
	var newLines []string = make([]string, len(lines))
	regex := regexp.MustCompile(`User\s+(\w+)`)

	for index, line := range lines {
		if regex.MatchString(line) {
			user = regex.FindStringSubmatch(line)[1]
			newLines[index] = fmt.Sprintf("[USR] %s %s", user, line)
		} else {
			newLines[index] = line
		}
	}

	return newLines
}
