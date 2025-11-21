package queenattack

import "fmt"

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) != 2 || len(blackPosition) != 2 || blackPosition == whitePosition{
		return false, fmt.Errorf("invalid position")
	}

	a := int(whitePosition[0] - 'a')
	b := int(whitePosition[1] - '1')
	x := int(blackPosition[0] - 'a')
	y := int(blackPosition[1] - '1')

	if a < 0 || a > 7 || b < 0 || b > 7 || x < 0 || x > 7 || y < 0 || y > 7 {
		return false, fmt.Errorf("position out of range")
	}

	return a == x || b == y || abs(a-x) == abs(b-y), nil
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
