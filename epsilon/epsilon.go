package epsilon

import (
	"math"
)

const (
	epsilon = 1e-5
)

func Within(a float64, b float64) bool {
	if (a == math.Inf(-1) && b == math.Inf(-1)) || (a == math.Inf(0) && b == math.Inf(0)) {
		return true
	}
	return math.Abs(a-b) < epsilon
}
