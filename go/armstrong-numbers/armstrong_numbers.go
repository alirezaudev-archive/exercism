package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	n2 := float64(n)
	var arm float64
	digits := strconv.Itoa(n)
	length := float64(len(digits))
	for _, d := range digits {
		arm += math.Pow(float64(d-'0'), length)
		if arm > n2 {
			return false
		}
	}
	return arm == n2
}
