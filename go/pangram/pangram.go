package pangram

import "strings"

func IsPangram(input string) bool {
	expected := "abcdefghijklmnopqrstuvwxyz"

	input = strings.ToLower(input)
	var actual [26]string
	for _, c := range input {
		if c >= 'a' && c <= 'z' {
			actual[c-'a'] = string(c)
		}
	}

	return strings.Join(actual[:], "") == expected
}
