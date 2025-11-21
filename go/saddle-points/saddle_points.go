package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix struct {
	grid [][]int
	rows int
	cols int
}

type Pair [2]int

func New(s string) (*Matrix, error) {
	lines := strings.Split(strings.TrimSpace(s), "\n")
	grid := make([][]int, len(lines))

	rows := len(lines)
	cols := 0

	for i, line := range lines {
		fields := strings.Fields(line)
		row := make([]int, len(fields))
		if len(row) > cols {
			cols = len(row)
		}

		for j, f := range fields {
			num, err := strconv.Atoi(f)
			if err != nil {
				return nil, fmt.Errorf("invalid number %q: %v", f, err)
			}
			row[j] = num
		}

		grid[i] = row
	}

	return &Matrix{grid: grid, rows: rows, cols: cols}, nil
}

func (m *Matrix) Saddle() []Pair {
	if m.rows == 0 || m.cols == 0 {
		return nil
	}

	rowMax := make([]int, m.rows)
	colMin := make([]int, m.cols)
	colInit := make([]bool, m.cols)
	for r := 0; r < m.rows; r++ {
		row := m.grid[r]
		if len(row) == 0 {
			continue
		}

		rowMax[r] = row[0]
		for c, v := range row {
			if v > rowMax[r] {
				rowMax[r] = v
			}

			if !colInit[c] || v < colMin[c] {
				colMin[c] = v
				colInit[c] = true
			}
		}
	}

	var result []Pair
	for r, row := range m.grid {
		for c, v := range row {
			if v == rowMax[r] && v == colMin[c] {
				result = append(result, Pair{r + 1, c + 1})
			}
		}
	}

	return result
}
