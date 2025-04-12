package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	MP = iota
	W
	D
	L
	StatsCount
)

func Tally(reader io.Reader, writer io.Writer) error {
	stats := map[string][]int{}
	update := buildUpdateFunc(stats)

	if err := parseInput(reader, update); err != nil {
		return err
	}

	return writeOutput(writer, stats)
}

func buildUpdateFunc(stats map[string][]int) map[string]func(t1 string, t2 string) {
	ensure := func(name string) {
		if _, ok := stats[name]; !ok {
			stats[name] = make([]int, StatsCount)
		}
	}

	return map[string]func(t1, t2 string){
		"win": func(t1, t2 string) {
			ensure(t1)
			ensure(t2)
			stats[t1][MP]++
			stats[t1][W]++
			stats[t2][MP]++
			stats[t2][L]++
		},
		"loss": func(t1, t2 string) {
			ensure(t1)
			ensure(t2)
			stats[t1][MP]++
			stats[t1][L]++
			stats[t2][MP]++
			stats[t2][W]++
		},
		"draw": func(t1, t2 string) {
			ensure(t1)
			ensure(t2)
			stats[t1][MP]++
			stats[t1][D]++
			stats[t2][MP]++
			stats[t2][D]++
		},
	}
}

func parseInput(r io.Reader, actions map[string]func(string, string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			return fmt.Errorf("invalid line: %q", line)
		}

		t1, t2, result := parts[0], parts[1], parts[2]
		action, ok := actions[result]
		if !ok {
			return fmt.Errorf("invalid result: %q", result)
		}
		action(t1, t2)
	}
	return scanner.Err()
}

func writeOutput(w io.Writer, stats map[string][]int) error {
	type teamResult struct {
		name   string
		values []int
	}

	var list []teamResult
	for name, v := range stats {
		list = append(list, teamResult{name, v})
	}

	sort.Slice(list, func(i, j int) bool {
		p1 := points(list[i].values)
		p2 := points(list[j].values)
		if p1 == p2 {
			return list[i].name < list[j].name
		}
		return p1 > p2
	})

	_, err := fmt.Fprintf(w, "%-30s | MP |  W |  D |  L |  P\n", "Team")
	if err != nil {
		return err
	}

	for _, t := range list {
		s := t.values
		_, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name, s[MP], s[W], s[D], s[L], points(s))
		if err != nil {
			return err
		}
	}
	return nil
}

func points(stat []int) int {
	return 3*stat[W] + stat[D]
}
