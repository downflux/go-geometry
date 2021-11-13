package epsilon

import (
	"math"
)

const (
	epsilon = 1e-5
)

func Within(a float64, b float64) bool {
	return math.Abs(a-b) < epsilon
}
