package summultiples

func SumMultiples(limit int, divisors ...int) int {
	res := 0
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor != 0 && i%divisor == 0 {
				res += i
				break
			}
		}
	}
	return res
}
