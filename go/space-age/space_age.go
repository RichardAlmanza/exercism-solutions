package space

type Planet string

const earthYearInSeconds float64 = 31557600

func (p Planet) getFactor() float64 {
	// factor is based on a year in Earth
	var factor float64

	switch p {
	case "Mercury":
		factor = 0.2408467
	case "Venus":
		factor = 0.61519726
	case "Earth":
		factor = 1.0
	case "Mars":
		factor = 1.8808158
	case "Jupiter":
		factor = 11.862615
	case "Saturn":
		factor = 29.447498
	case "Uranus":
		factor = 84.016846
	case "Neptune":
		factor = 164.79132
	}

	return factor
}

func (p Planet) getYears(seconds float64) float64 {
	if p.getFactor() == 0.0 {
		return -1
	}

	return seconds / ( earthYearInSeconds * p.getFactor() )
}

// Age returns how many years are on the planet, the years are
// calculated based on the seconds' input.
func Age(seconds float64, planet Planet) float64 {
	return planet.getYears(seconds)
}
