package epsilon

import (
	"math"
)

var (
	// minNormal is the minimum delta between two float64 values. In
	// float64 format, 53 bits are reserved for precision (the significand).
	// See
	// https://en.wikipedia.org/wiki/Double-precision_floating-point_format
	// for more information.
	minNormal = math.Float64frombits(1 << 52) // 0x0010000000000000

	DefaultE = *New(
		func(a float64, b float64) float64 {
			return 128 * math.Abs(a-math.Nextafter(a, b))
		},
	)
)

type f func(a, b float64) float64
type E f

func New(epsilon f) *E {
	e := E(epsilon)
	return &e
}

// Within calculates if two float64 values are very close to one another. This
// is based on the principle that if a ~ b, (a - b) / (a + b) ~ 0.
//
// See https://stackoverflow.com/a/32334103/873865.
func (e E) Within(a float64, b float64) bool {
	if a == b {
		return true
	}

	normal := math.Min(math.Abs(a)+math.Abs(b), math.MaxFloat64)
	return math.Abs(a-b) < math.Max(minNormal, f(e)(a, b)*normal)
}

func Within(a float64, b float64) bool { return DefaultE.Within(a, b) }
