package lsproduct

import (
	"errors"
	"regexp"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 0 || len(digits) < span || !regexp.MustCompile(`^[0-9]+$`).MatchString(digits) {
		return 0, errors.New("")
	}

	var m int64
	for i := 0; i <= len(digits)-span; i++ {
		tmp := int64(digits[i] - '0')
		for _, digit := range digits[i+1 : i+span] {
			tmp *= int64(digit - '0')
		}
		if tmp > m {
			m = tmp
		}
	}

	return m, nil
}
