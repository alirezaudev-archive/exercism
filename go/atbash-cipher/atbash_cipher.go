package atbash

import (
	"strings"
	"unicode"
)

func Atbash(input string) string {
	var result []rune
	count := 0

	for _, r := range strings.ToLower(input) {
		switch {
		case unicode.IsLetter(r):
			result = append(result, 'a'+('z'-r))
			count++
		case unicode.IsDigit(r):
			result = append(result, r)
			count++
		default:
			continue
		}

		if count%5 == 0 {
			result = append(result, ' ')
		}
	}

	res := strings.TrimSpace(string(result))
	return res
}
