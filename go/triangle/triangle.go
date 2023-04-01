package triangle

type Kind int

const (
    NaT = iota  // not a triangle
    Equ         // equilateral
    Iso         // isosceles
    Sca         // scalene
)

func checkTriangle(sides [3]float64) bool {
	var maxSideIndex int
	var sum float64

	for index, side := range sides {
		if side <= 0 {
			return false
		}

		if side > sides[maxSideIndex] {
			maxSideIndex = index
		}

		sum += side
	}

	sum -= 2 * sides[maxSideIndex]

	return 0 <= sum
}

func countEqualSides(sides [3]float64) uint8 {
	var counter uint8

	for index, sideA := range sides {
		sideB := sides[(index + 1) % 3]

		if sideA == sideB {
			counter++
		}
	}

	return counter
}

func newKind(sides [3]float64) Kind {
	var triangle Kind

	if !checkTriangle(sides) {
		return NaT
	}

	switch countEqualSides(sides) {
	case 0:
		triangle = Sca
	case 1:
		triangle = Iso
	case 2:
		panic("Witchcraft!")
	case 3:
		triangle = Equ
	}

	return triangle
}

// KindFromSides returns the triangle kind based on its sides.
func KindFromSides(a, b, c float64) Kind {
	return newKind([3]float64{a, b, c})
}
