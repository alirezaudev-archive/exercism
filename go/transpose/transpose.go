package transpose

const testVersion = 1

func Transpose(m []string) []string {
	if len(m) == 0 {
		return nil
	}

	suffixMax := make([]int, len(m))
	currentMax := 0
	for i := len(m) - 1; i >= 0; i-- {
		rowLen := len(m[i])
		if rowLen > currentMax {
			currentMax = rowLen
		}
		suffixMax[i] = currentMax
	}

	maxCols := suffixMax[0]
	result := make([]string, maxCols)

	for i := 0; i < len(m); i++ {
		row := m[i]
		rowLen := len(row)

		for j := 0; j < rowLen; j++ {
			result[j] += string(row[j])
		}

		for j := rowLen; j < suffixMax[i]; j++ {
			result[j] += " "
		}
	}

	return result
}
