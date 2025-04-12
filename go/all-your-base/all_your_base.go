package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	num := 0
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num = num*inputBase + d
	}

	if num == 0 {
		return []int{0}, nil
	}

	var result []int
	for num > 0 {
		result = append([]int{num % outputBase}, result...)
		num /= outputBase
	}

	return result, nil
}
