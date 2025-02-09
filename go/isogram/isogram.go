package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	chars := make(map[rune]int)
	for _, char := range word {
		if char < 97 || char > 122 {
			continue
		}
		if chars[char] == 1 {
			return false
		}

		chars[char] += 1
	}
	return true
}
