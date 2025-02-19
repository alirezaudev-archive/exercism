package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	apps := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, char := range log {
		if v, ok := apps[char]; ok {
			return v
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var result string
	for _, char := range log {
		if char == oldRune {
			char = newRune
		}

		result += string(char)
	}

	return result
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
