package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Value         int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	if fmin > fmax {
		return Product{}, Product{}, errors.New("fmin > fmax")
	}

	var pmin, pmax Product
	found := false
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			v := i * j
			if !isPalindrome(v) {
				continue
			}

			if !found {
				found = true
				pmin = Product{Value: v, Factorizations: [][2]int{{i, j}}}
				pmax = pmin
				continue
			}

			if v < pmin.Value {
				pmin = Product{Value: v, Factorizations: [][2]int{{i, j}}}
			} else if v == pmin.Value {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}

			if v > pmax.Value {
				pmax = Product{Value: v, Factorizations: [][2]int{{i, j}}}
			} else if v == pmax.Value {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}
		}
	}

	if !found {
		return Product{}, Product{}, errors.New("no palindromes")
	}

	return pmin, pmax, nil
}

func isPalindrome(n int) bool {
	if n < 0 {
		n = -n
	}

	s := strconv.Itoa(n)
	l := len(s)
	half := l / 2
	for i := 0; i < half; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
