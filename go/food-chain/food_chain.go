package foodchain

import "strings"

var animals = []string{"", "fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var reactions = map[int]string{
	2: "It wriggled and jiggled and tickled inside her.",
	3: "How absurd to swallow a bird!",
	4: "Imagine that, to swallow a cat!",
	5: "What a hog, to swallow a dog!",
	6: "Just opened her throat and swallowed a goat!",
	7: "I don't know how she swallowed a cow!",
}
var extras = map[int]string{
	2: " that wriggled and jiggled and tickled inside her",
}
var theEnd = "I don't know why she swallowed the fly. Perhaps she'll die."

func Verse(v int) string {
	if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	lines := []string{"I know an old lady who swallowed a " + animals[v] + "."}

	if reaction, ok := reactions[v]; ok {
		lines = append(lines, reaction)
	}

	for i := v; i > 1; i-- {
		line := "She swallowed the " + animals[i] + " to catch the " + animals[i-1]
		if extra, ok := extras[i-1]; ok {
			line += extra
		}
		line += "."
		lines = append(lines, line)
	}

	lines = append(lines, theEnd)

	return strings.Join(lines, "\n")
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
