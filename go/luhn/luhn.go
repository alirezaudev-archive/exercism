package luhn

import "strconv"

func Valid(id string) bool {
	length := len(id)
	if length <= 1 {
		return false
	}

	if id[0] == ' ' || id[length-1] == ' ' {
		return false
	}

	var doublySum int
	intIndex := 0
	previousSpacePosition := length
	for i := length - 1; i > -1; i-- {
		char := id[i]

		if char == ' ' {
			if previousSpacePosition != length && previousSpacePosition-1 == i {
				return false
			}

			previousSpacePosition = i
			continue
		}

		if char < '0' || char > '9' {
			return false
		}

		intChar, _ := strconv.Atoi(string(char))
		double := intChar
		if intIndex%2 != 0 {
			double = double * 2
			if double > 9 {
				double -= 9
			}
		}
		doublySum += double
		intIndex++
	}
	if doublySum%10 != 0 {
		return false
	}

	return true
}
