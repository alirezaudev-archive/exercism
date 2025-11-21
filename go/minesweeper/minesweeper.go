package minesweeper

import (
	"strconv"
	"strings"
)

// Annotate returns an annotated board
func Annotate(board []string) []string {
	dirs := [8][2]int{{1, 1}, {1, 0}, {0, 1}, {-1, -1}, {-1, 0}, {0, -1}, {1, -1}, {-1, 1}}
	rows := len(board)

	for r := range board {
		cols := len(board[r])
		var sb strings.Builder
		sb.Grow(cols)

		for c := 0; c < cols; c++ {
			if board[r][c] == '*' {
				sb.WriteByte('*')
				continue
			}

			count := 0
			for _, dir := range dirs {
				rn, cn := r+dir[0], c+dir[1]
				if rn >= 0 && rn < rows && cn >= 0 && cn < cols && board[rn][cn] == '*' {
					count++
				}
			}

			if count != 0 {
				sb.WriteString(strconv.Itoa(count))
			} else {
				sb.WriteByte(' ')
			}
		}

		board[r] = sb.String()
	}

	return board
}
