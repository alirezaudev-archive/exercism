package spiralmatrix

func SpiralMatrix(size int) [][]int {
	if size == 0 {
		return [][]int{}
	}

	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dir := 0
	row, col := 0, 0

	for i := 1; i <= size*size; i++ {
		matrix[row][col] = i

		nextRow := row + directions[dir][0]
		nextCol := col + directions[dir][1]

		if nextRow < 0 || nextRow >= size || nextCol < 0 || nextCol >= size || matrix[nextRow][nextCol] != 0 {
			dir = (dir + 1) % 4
			nextRow = row + directions[dir][0]
			nextCol = col + directions[dir][1]
		}

		row = nextRow
		col = nextCol
	}

	return matrix
}