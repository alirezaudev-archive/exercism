package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := map[string]int{}
	for point, letters := range in {
		for _, letter := range letters {
			res[strings.ToLower(letter)] = point
		}
	}
	return res
}
