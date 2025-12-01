package connect

import "strings"

var directions = [6][2]int{{-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}}

func ResultOf(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", nil
	}

	board := make([][]rune, len(lines))
	for i, line := range lines {
		row := strings.ReplaceAll(line, " ", "")
		board[i] = []rune(row)
	}

	rows, cols := len(board), len(board[0])
	for c := 0; c < cols; c++ {
		if board[0][c] == 'O' && dfs(board, 0, c, 'O') {
			return "O", nil
		}
	}

	for r := 0; r < rows; r++ {
		if board[r][0] == 'X' && dfs(board, r, 0, 'X') {
			return "X", nil
		}
	}

	return "", nil
}

func dfs(board [][]rune, x, y int, player rune) bool {
	rows, cols := len(board), len(board[0])

	if player == 'O' && x == rows-1 {
		return true
	}
	if player == 'X' && y == cols-1 {
		return true
	}

	board[x][y] = '.'
	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
			continue
		}
		if board[nx][ny] != player {
			continue
		}
		if dfs(board, nx, ny, player) {
			return true
		}
	}

	return false
}
