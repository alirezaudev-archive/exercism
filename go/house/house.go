package house

import "strings"

func Verse(v int) string {
	subjects := []string{
		"the house that Jack built",
		"the malt",
		"the rat",
		"the cat",
		"the dog",
		"the cow with the crumpled horn",
		"the maiden all forlorn",
		"the man all tattered and torn",
		"the priest all shaven and shorn",
		"the rooster that crowed in the morn",
		"the farmer sowing his corn",
		"the horse and the hound and the horn",
	}
	actions := []string{
		"",
		"lay in",
		"ate",
		"killed",
		"worried",
		"tossed",
		"milked",
		"kissed",
		"married",
		"woke",
		"kept",
		"belonged to",
	}

	lines := []string{"This is " + subjects[v-1]}
	for i := v - 1; i > 0; i-- {
		lines = append(lines, "that "+actions[i]+" "+subjects[i-1])
	}
	lines[len(lines)-1] += "."

	return strings.Join(lines, "\n")
}

func Song() string {
	var verses []string
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}
