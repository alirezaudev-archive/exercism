package complexnumbers

import (
	"math"
)

type Number struct {
	real float64
	imag float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imag
}

func (n Number) Add(n2 Number) Number {
	return Number{
		real: n.real + n2.real,
		imag: n.imag + n2.imag,
	}
}

func (n Number) Subtract(n2 Number) Number {
	return Number{
		real: n.real - n2.real,
		imag: n.imag - n2.imag,
	}
}

func (n Number) Multiply(n2 Number) Number {
	return Number{
		real: n.real*n2.real - n.imag*n2.imag,
		imag: n.imag*n2.real + n.real*n2.imag,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real: n.real * factor,
		imag: n.imag * factor,
	}
}

func (n Number) Divide(n2 Number) Number {
	denominator := n2.real*n2.real + n2.imag*n2.imag
	return Number{
		real: (n.real*n2.real + n.imag*n2.imag) / denominator,
		imag: (n.imag*n2.real - n.real*n2.imag) / denominator,
	}
}

func (n Number) Conjugate() Number {
	return Number{
		real: n.real,
		imag: -n.imag,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imag*n.imag)
}

func (n Number) Exp() Number {
	expReal := math.Exp(n.real)
	return Number{
		real: expReal * math.Cos(n.imag),
		imag: expReal * math.Sin(n.imag),
	}
}
