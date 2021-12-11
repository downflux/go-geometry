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

func New(v vector.V, u vector.V) *R {
	if v.Dimension() != u.Dimension() {
		panic("cannot construct a hyperrectangle with mismatching input vector dimensions")
	}
	min := make([]float64, v.Dimension())
	max := make([]float64, v.Dimension())

	for i := vector.D(0); i < v.Dimension(); i++ {
		min[i] = math.Min(v.X(i), u.X(i))
		max[i] = math.Max(v.X(i), u.X(i))
	}

	return &R{
		min: *vector.New(min...),
		max: *vector.New(max...),
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
