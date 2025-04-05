package triangle

type Kind string

const (
	NaT = Kind("not a triangle")
	Equ = Kind("equilateral")
	Iso = Kind("isosceles")
	Sca = Kind("scalene")
)

// KindFromSides determines the type of triangle given side lengths.
func KindFromSides(a, b, c float64) Kind {
	anySideIsLessThanZero := a <= 0 || b <= 0 || c <= 0
	isNotValid := a+b < c || b+c < a || a+c < b
	if anySideIsLessThanZero || isNotValid {
		return NaT
	}

	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}
