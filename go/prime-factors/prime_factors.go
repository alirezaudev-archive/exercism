package prime

func Factors(n int64) []int64 {
	var factors []int64
	if n < 2 {
		return factors
	}
	for i := int64(2); i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}
