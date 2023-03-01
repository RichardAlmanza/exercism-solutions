package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	applications := map[rune]string{
		'\u2757': "recommendation",
		'\U0001f50d': "search",
		'\u2600': "weather",
	}

	for _, runeInLog := range log {
		if applications[runeInLog] != "" {
			return applications[runeInLog]
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	var runeToWrite rune

	for _, runeInLog := range log {
		runeToWrite = runeInLog

		if runeToWrite == oldRune {
			runeToWrite = newRune
		}

		newLog = fmt.Sprintf("%s%c", newLog, runeToWrite)
	}

	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
