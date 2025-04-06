package anagram

import (
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var res []string
	for _, candidate := range candidates {
		lower := strings.ToLower(candidate)
		if lower == subject {
			continue
		}
		if isAnagram(subject, lower) {
			res = append(res, candidate)
		}
	}
	return res
}

func isAnagram(s1, s2 string) bool {
	f1 := freq(s1)
	f2 := freq(s2)
	if len(f1) != len(f2) {
		return false
	}

	for k, v1 := range f1 {
		if v2, ok := f2[k]; !ok || v2 != v1 {
			return false
		}
	}

	return true
}

func freq(s string) map[rune]int {
	f := make(map[rune]int)
	for _, r := range s {
		f[r]++
	}
	return f
}
