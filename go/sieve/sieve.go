package sieve

import (
	"math"
)

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	res := []int{2}
	for i := 3; i <= limit; i += 2 {
		if isPrime(i) {
			res = append(res, i)
		}
	}

	return res
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
