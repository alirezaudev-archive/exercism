package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("")
	}
	if n == 1 {
		return 0, nil
	}

	var next int
	if n%2 == 0 {
		next = n / 2
	} else {
		next = 3*n + 1
	}

	steps, err := CollatzConjecture(next)
	return steps + 1, err
}
