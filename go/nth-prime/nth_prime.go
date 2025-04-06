package prime

import (
	"errors"
	"math"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("")
	}

	if n == 1 {
		return 2, nil
	}

	count := 1
	i := 3
	for ; count < n; i += 2 {
		if isPrime(i) {
			count++
		}
	}

	return i - 2, nil
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
