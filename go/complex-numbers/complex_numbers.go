package complexnumbers

import "math"

// Define the Number type here.

type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.a + n2.a, n1.b + n2.b}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.a - n2.a, n1.b - n2.b}
}

func (n1 Number) Multiply(n2 Number) Number {
	var newReal float64
	var newImaginary float64

	a, b, c, d := n1.a, n1.b, n2.a, n2.b

	newReal = a * c - b * d
	newImaginary = b * c + a * d

	return Number{newReal, newImaginary}
}

func (n Number) Times(factor float64) Number {
	return Number{n.a * factor, n.b * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	var newReal float64
	var newImaginary float64
	var denominator float64

	a, b, c, d := n1.a, n1.b, n2.a, n2.b

	denominator = c * c + d * d
	newReal = (a * c + b * d) / denominator
	newImaginary = (b * c - a * d) / denominator

	return Number{newReal, newImaginary}
}

func (n Number) Conjugate() Number {
	return Number{n.a, -n.b}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a * n.a + n.b * n.b)
}

func (n Number) Exp() Number {
	return Number{math.Cos(n.b), math.Sin(n.b)}.Times(math.Pow(math.E, n.a))
}
