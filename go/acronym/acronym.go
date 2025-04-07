package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	words := regexp.MustCompile(`[-\s_]+`).Split(s, -1)
	acr := ""
	for _, word := range words {
		acr += strings.ToUpper(string(word[0]))
	}
	return acr
}
