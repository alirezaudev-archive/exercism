package pascal

func Triangle(numRows int) [][]int {
	var triangle [][]int

	for i := 1; i <= numRows; i++ {
		row := make([]int, i)
		for k := 0; k < i; k++ {
			row[k] = 1
		}

		for j := 1; j < i/2+i%2; j++ {
			value := triangle[len(triangle)-1][j-1] + triangle[len(triangle)-1][j]
			row[j] = value
			row[len(row)-j-1] = value
		}

		triangle = append(triangle, row)
	}

	return triangle
}
