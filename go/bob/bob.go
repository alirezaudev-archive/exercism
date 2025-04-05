package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	switch {
	case isQuestion(remark) && isYelling(remark):
		return "Calm down, I know what I'm doing!"
	case isSilence(remark):
		return "Fine. Be that way!"
	case isQuestion(remark):
		return "Sure."
	case isYelling(remark):
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func isSilence(remark string) bool {
	return strings.TrimSpace(remark) == ""
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}

func isYelling(remark string) bool {
	hasLetters := false
	for _, r := range remark {
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return hasLetters
}
