package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var result []Triplet

	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c2 := a*a + b*b
			c := intSqrt(c2)
			if c > max {
				continue
			}
			if c*c == c2 && c >= b && c >= min && c <= max {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}

func Sum(p int) []Triplet {
	var result []Triplet

	for a := 1; a < p; a++ {
		for b := a + 1; b < p; b++ {
			c := p - a - b
			if c <= b {
				continue
			}
			if a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
		}
	}
	return result
}

func intSqrt(x int) int {
	r := 0
	for r*r <= x {
		r++
	}
	return r - 1
}
