// Package hyperrectangle defines an N-dimensional box embedded in N-dimensional
// ambient space.
package hyperrectangle

import (
	"github.com/downflux/go-geometry/epsilon"
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

func (r R) M() M          { return M(r) }
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

func Intersect(r R, s R) (R, bool) {
	b := M(*New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	))
	b.Copy(r)
	if ok := b.Intersect(s); ok {
		return b.R(), ok
	}
	return R{}, false
}

func Union(r R, s R) R {
	b := M(*New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	))
	b.Copy(r)
	b.Union(s)
	return b.R()
}

func V(r R) float64 {
	v := 1.0
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		v *= r.Max().X(i) - r.Min().X(i)
	}
	return v
}

// SA returns the "surface area" of an N-dimensional interval. For N = 2, this
// is the perimeter, and for N = 3, this is the total surface area of the
// rectangular prism.
func SA(r R) float64 {
	k := r.Min().Dimension()
	if k == 1 {
		return 0
	}

	var sa float64
	min := make([]float64, k-1)
	max := make([]float64, k-1)
	for i := vector.D(0); i < k; i++ {
		copy(min, r.Min()[:i])
		copy(min[i:], r.Min()[i+1:])
		copy(max, r.Max()[:i])
		copy(max[i:], r.Max()[i+1:])

		sa += V(*New(min, max))
	}
	return 2 * sa
}

func WithinEpsilon(g R, h R, e epsilon.E) bool {
	return vector.WithinEpsilon(g.Min(), h.Min(), e) && vector.WithinEpsilon(g.Max(), h.Max(), e)
}

func Within(g R, h R) bool { return WithinEpsilon(g, h, epsilon.DefaultE) }
