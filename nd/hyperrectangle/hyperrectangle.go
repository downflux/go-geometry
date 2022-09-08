// Package hyperrectangle defines an N-dimensional box embedded in N-dimensional
// ambient space.
package hyperrectangle

import (
	"math"

	"github.com/downflux/go-geometry/nd/vector"
)

type R struct {
	min vector.V
	max vector.V
}

func New(min vector.V, max vector.V) *R {
	if min.Dimension() != max.Dimension() {
		panic("cannot construct a hyperrectangle with mismatching input vector dimensions")
	}

	for i := vector.D(0); i < min.Dimension(); i++ {
		if min[i] > max[i] {
			panic("cannot construct a hyperrectangle with invalid min and max vectors")
		}
	}

	return &R{
		min: min,
		max: max,
	}
}

func (r R) Min() vector.V { return r.min }
func (r R) Max() vector.V { return r.max }
func (r R) D() vector.V   { return vector.Sub(r.Max(), r.Min()) }

func (r R) In(v vector.V) bool {
	success := true
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		success = success && (r.Min().X(i) <= v.X(i) && v.X(i) <= r.Max().X(i))
	}
	return success
}

func (r R) Intersect(s R) (R, bool) {
	b := New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	)

	if ok := r.IntersectBuf(s, b); ok {
		return *b, ok
	}
	return R{}, false
}

// IntersectBuf calculates the rectangle of intersection between two rectangles
// and writes to a rectangle buffer.
//
// N.B.: The buffer min and max properties must match the dimensions of the
// input rectangles.
func (r R) IntersectBuf(s R, b *R) bool {
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		b.Min()[i] = math.Max(r.Min().X(i), s.Min().X(i))
		b.Max()[i] = math.Min(r.Max().X(i), s.Max().X(i))

		if b.Min()[i] > b.Max()[i] {
			return false
		}
	}

	return true
}

func V(r R) float64 {
	v := 1.0
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		v *= r.Max().X(i) - r.Min().X(i)
	}
	return v
}
