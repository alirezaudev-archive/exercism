package isbn

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		ch := rune(isbn[i])
		var val int

		if ch == 'X' {
			if i != 9 {
				return false
			}
			val = 10
		} else if unicode.IsDigit(ch) {
			val = int(ch - '0')
		} else {
			return false
		}

		sum += val * (10 - i)
	}

	return sum%11 == 0
}
