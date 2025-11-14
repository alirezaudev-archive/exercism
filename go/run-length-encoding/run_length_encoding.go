package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	result := ""

	prev := ""
	counter := 0
	for _, r := range input {
		char := string(r)
		counter++

		if prev != char {
			result += concat(counter, prev)
			prev = char
			counter = 0
		}
	}

	result += concat(counter+1, prev)

	return result
}

func concat(counter int, prev string) string {
	n := ""
	if counter > 1 {
		n = strconv.Itoa(counter)
	}
	return n + prev
}

func RunLengthDecode(input string) string {
	result := ""
	count := ""
	for _, r := range input {
		char := string(r)
		if r >= '0' && r <= '9' {
			count += char
		} else {
			if count == "" {
				count = "1"
			}
			countN, _ := strconv.Atoi(count)
			result += strings.Repeat(char, countN)
			count = ""
		}
	}

	return result
}
