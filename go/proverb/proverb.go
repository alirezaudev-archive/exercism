package proverb

import "fmt"

func Proverb(rhyme []string) []string {
	var res []string

	if len(rhyme) < 1 {
		return res
	}

	for i := range rhyme[:len(rhyme)-1] {
		res = append(res, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	res = append(res, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return res
}
