package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	if distance > 10 {
		return 0
	} else if distance > 5 {
		return 1
	} else if distance > 1 {
		return 5
	} else {
		return 10
	}
}
