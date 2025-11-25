package twobucket

import (
	"errors"
)

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if sizeBucketOne <= 0 ||
		sizeBucketTwo <= 0 ||
		goalAmount <= 0 ||
		(startBucket != "one" && startBucket != "two") ||
		(goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo) ||
		(goalAmount%gcd(sizeBucketOne, sizeBucketTwo) != 0) {
		return "", 0, 0, errors.New("error")
	}

	bucket1, bucket2 := 0, 0
	size1, size2 := sizeBucketOne, sizeBucketTwo

	if startBucket == "two" {
		size1, size2 = size2, size1
		bucket1, bucket2 = bucket2, bucket1
	}

	moves := 0

	for bucket1 != goalAmount && bucket2 != goalAmount {
		if bucket1 == 0 {
			bucket1 = size1
		} else if size2 == goalAmount {
			bucket2 = goalAmount
		} else if bucket2 == size2 {
			bucket2 = 0
		} else {
			pour := size2 - bucket2
			if bucket1 < pour {
				pour = bucket1
			}
			bucket1 -= pour
			bucket2 += pour
		}
		moves++
	}

	goalBucket := "one"
	otherLevel := bucket2

	if bucket2 == goalAmount {
		goalBucket = "two"
		otherLevel = bucket1
	}

	if startBucket == "two" {
		if goalBucket == "one" {
			goalBucket = "two"
		} else {
			goalBucket = "one"
		}
	}

	return goalBucket, moves, otherLevel, nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
