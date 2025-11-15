package piglatin

import "strings"

var vowelsMap = map[byte]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

var vowels = []string{"a", "e", "i", "o", "u"}

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	for i, w := range words {
		words[i] = translate(w)
	}
	return strings.Join(words, " ")
}

func translate(sentence string) string {
	var end string
	var start string

	if strStratsWith(sentence, append(vowels, "xr", "yt")...) {
		return sentence + "ay"
	}

	if hasQU := indexOf(sentence, "qu", 0); hasQU != -1 && isAllConsonants(sentence[:hasQU]) {
		start = sentence[hasQU+2:]
		end = sentence[:hasQU+2] + "ay"
		return start + end
	}

	if yIndex := indexOf(sentence, "y", 0); yIndex > 0 && isAllConsonants(sentence[:yIndex]) {
		start = sentence[yIndex:]
		end = sentence[:yIndex] + "ay"
		return start + end
	}

	constantEnd := 0
	for constantEnd < len(sentence) && !vowelsMap[sentence[constantEnd]] {
		constantEnd++
	}

	start = sentence[constantEnd:]
	end = sentence[:constantEnd] + "ay"

	return start + end
}

func strStratsWith(haystack string, needle ...string) bool {
	l := len(haystack)
	if l == 0 {
		return false
	}
	for _, n := range needle {
		nl := len(n)
		if nl > l {
			continue
		}
		if haystack[0:nl] == n {
			return true
		}
	}
	return false
}

func indexOf(s string, needle string, start int) int {
	nl := len(needle)
	sl := len(s)

	for i := start; i <= sl-nl; i++ {
		if s[i:i+nl] == needle {
			return i
		}
	}

	return -1
}

func isAllConsonants(s string) bool {
	for i := 0; i < len(s); i++ {
		if vowelsMap[s[i]] {
			return false
		}
	}
	return true
}
