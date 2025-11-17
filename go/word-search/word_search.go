package wordsearch

import "errors"

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	for _, word := range words {
		found, ok := find(word, puzzle)
		if !ok {
			return nil, errors.New("word not found: " + word)
		}
		result[word] = found
	}
	return result, nil
}

func find(word string, puzzle []string) ([2][2]int, bool) {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for row := 0; row < len(puzzle); row++ {
		for col := 0; col < len(puzzle[row]); col++ {
			for _, dir := range directions {
				if found, ok := search(word, puzzle, row, col, dir[0], dir[1]); ok {
					return found, true
				}
			}
		}
	}

	return [2][2]int{{-1, -1}, {-1, -1}}, false
}

func search(word string, puzzle []string, startRow, startCol, rowDir, colDir int) ([2][2]int, bool) {
	if len(word) == 0 {
		return [2][2]int{}, false
	}

	rows := len(puzzle)
	if rows == 0 {
		return [2][2]int{}, false
	}
	cols := len(puzzle[0])

	endRow := startRow + rowDir*(len(word)-1)
	endCol := startCol + colDir*(len(word)-1)

	if endRow < 0 || endRow >= rows || endCol < 0 || endCol >= cols {
		return [2][2]int{}, false
	}

	row, col := startRow, startCol
	for i := 0; i < len(word); i++ {
		if puzzle[row][col] != word[i] {
			return [2][2]int{}, false
		}
		row += rowDir
		col += colDir
	}

	return [2][2]int{{startCol, startRow}, {endCol, endRow}}, true
}
