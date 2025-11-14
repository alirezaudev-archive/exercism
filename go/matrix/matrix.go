package matrix

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, len(rows))

	for i, row := range rows {
		re := regexp.MustCompile(`\S+`)
		cols := re.FindAllString(row, -1)
		colsN := len(cols)

		if i > 0 && colsN != len(m[i-1]) {
			return m, errors.New("uneven rows")
		}

		m[i] = make([]int, colsN)
		for j := range cols {
			value, err := strconv.Atoi(cols[j])
			if err != nil {
				return nil, err
			}

			m[i][j] = value
		}
	}

	return m, nil
}

func (m Matrix) Cols() [][]int {
	rows := len(m)
	if rows == 0 || len(m[0]) == 0 {
		return [][]int{}
	}

	cols := len(m[0])
	result := make([][]int, cols)
	for c := 0; c < cols; c++ {
		result[c] = make([]int, rows)
		for r := 0; r < rows; r++ {
			result[c][r] = m[r][c]
		}
	}

	return result
}

func (m Matrix) Rows() [][]int {
	out := make([][]int, len(m))
	for i := range m {
		out[i] = append([]int(nil), m[i]...)
	}
	return out
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}

	m[row][col] = val
	return true
}