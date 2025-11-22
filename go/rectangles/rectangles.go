package rectangles

func Count(diagram []string) int {
	if len(diagram) == 0 {
		return 0
	}

	count := 0
	rows := len(diagram)
	cols := len(diagram[0])

	for r1 := 0; r1 < rows; r1++ {
		for c1 := 0; c1 < cols; c1++ {
			if diagram[r1][c1] != '+' {
				continue
			}

			for r2 := r1 + 1; r2 < rows; r2++ {
				for c2 := c1 + 1; c2 < cols; c2++ {
					if diagram[r2][c2] != '+' {
						continue
					}

					if isRectangle(diagram, r1, c1, r2, c2) {
						count++
					}
				}
			}
		}
	}

	return count
}

func isRectangle(diagram []string, r1, c1, r2, c2 int) bool {
	if diagram[r1][c2] != '+' || diagram[r2][c1] != '+' {
		return false
	}

	for c := c1 + 1; c < c2; c++ {
		if diagram[r1][c] != '-' && diagram[r1][c] != '+' {
			return false
		}
	}

	for c := c1 + 1; c < c2; c++ {
		if diagram[r2][c] != '-' && diagram[r2][c] != '+' {
			return false
		}
	}

	for r := r1 + 1; r < r2; r++ {
		if diagram[r][c1] != '|' && diagram[r][c1] != '+' {
			return false
		}
	}

	for r := r1 + 1; r < r2; r++ {
		if diagram[r][c2] != '|' && diagram[r][c2] != '+' {
			return false
		}
	}

	return true
}