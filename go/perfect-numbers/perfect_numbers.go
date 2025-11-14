package perfect

import (
	"errors"
)

var ErrOnlyPositive = errors.New("negative")

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	sum := int64(0)

	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
