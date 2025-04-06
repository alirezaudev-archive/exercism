package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	words := regexp.MustCompile(`\b[a-z0-9]+(?:'[a-z0-9]+)?\b`).FindAllString(phrase, -1)

	freq := Frequency{}
	for _, word := range words {
		freq[word]++
	}
	return freq
}
