package cryptosquare

import (
	"strings"
	"unicode"
)

func Encode(pt string) string {
	text := normalize(pt)
	if len(text) == 0 {
		return ""
	}

	rows, cols := rectangle(len(text))
	grid := buildGrid(text, rows, cols)

	var b strings.Builder
	for c := 0; c < cols; c++ {
		if c > 0 {
			b.WriteByte(' ')
		}
		for r := 0; r < rows; r++ {
			b.WriteByte(grid[r][c])
		}
	}

	return b.String()
}

func normalize(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func rectangle(n int) (int, int) {
	for c := 1; ; c++ {
		r := (n + c - 1) / c
		if c >= r && c-r <= 1 {
			return r, c
		}
	}
}

func buildGrid(s string, rows, cols int) [][]byte {
	grid := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]byte, cols)
		for j := 0; j < cols; j++ {
			idx := i*cols + j
			if idx < len(s) {
				grid[i][j] = s[idx]
			} else {
				grid[i][j] = ' '
			}
		}
	}
	return grid
}
