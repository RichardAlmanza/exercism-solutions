package darts

import "math"

func getScore(distanceFromCenter float64) int {
	var score int

	switch{
	case distanceFromCenter <= 1.0:
		score = 10
	case distanceFromCenter <= 5.0:
		score = 5
	case distanceFromCenter <= 10.0:
		score = 1
	}

	return score
}

func calculatesDistanceFromCenter(x, y float64) float64 {
	return math.Pow(
		math.Pow(x, 2) + math.Pow(y, 2),
		0.5,
	)
}

// Score returns the score of the dart landed at x and y coordinates
func Score(x, y float64) int {
	return getScore(calculatesDistanceFromCenter(x, y))
}
