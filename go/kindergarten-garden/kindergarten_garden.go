package kindergarten

import (
	"errors"
	"strings"
)

type Garden map[string][]string

var plants = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	sorted := sortChildren(append([]string{}, children...))
	rows := parseDiagram(diagram)
	if sorted == nil || rows == nil || len(children)*2 != len(rows[0]) {
		return nil, errors.New("")
	}

	g := make(Garden)

	for i, child := range sorted {
		s := i * 2
		e := i*2 + 2
		for _, r := range rows[0][s:e] + rows[1][s:e] {
			plant, ok := plants[r]
			if !ok {
				return nil, errors.New("invalid plant character in diagram")
			}

			g[child] = append(g[child], plant)
		}
	}

	return &g, nil
}

func parseDiagram(diagram string) []string {
	rows := strings.Split(diagram[1:], "\n")
	if len(rows) != 2 || len(rows[0]) != len(rows[1]) {
		return nil
	}

	return rows
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if plant, ok := (*g)[child]; ok {
		return plant, true
	}

	return []string{}, false

}

func sortChildren(arr []string) []string {
	n := len(arr)
	for i := 0; i < n; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] == arr[minIdx] {
				return nil
			}

			if arr[j][0] < arr[minIdx][0] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr
}
