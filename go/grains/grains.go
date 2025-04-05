package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("")
	}

	return uint64(1 << (number - 1)), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		grains, _ := Square(i)
		total += grains
	}
	return total
}
