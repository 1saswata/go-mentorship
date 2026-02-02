package geometry

import "math"

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}
