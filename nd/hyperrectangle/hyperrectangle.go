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
		success = success && (r.Min()[i] <= v[i] && v[i] <= r.Max()[i])
	}
	return success
}

func Intersect(r R, s R) (R, bool) {
	b := New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	).M()
	b.Copy(r)
	if ok := b.Intersect(s); ok {
		return b.R(), ok
	}
	return R{}, false
}

func Union(r R, s R) R {
	b := New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	).M()
	b.Copy(r)
	b.Union(s)
	return b.R()
}

func Scale(r R, c float64) R {
	b := New(
		vector.V(make([]float64, r.Min().Dimension())),
		vector.V(make([]float64, r.Min().Dimension())),
	).M()
	b.Copy(r)
	b.Scale(c)
	return b.R()
}

// Contains checks if the input rectangle r fully encloses s. r is a closed
// interval.
func Contains(r R, s R) bool {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		if s.Min()[i] < r.Min()[i] || s.Max()[i] > r.Max()[i] {
			return false
		}
	}
	return true
}

func Disjoint(r R, s R) bool {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	rmin, rmax := r.Min(), r.Max()
	smin, smax := s.Min(), s.Max()
	k := rmin.Dimension()

	for i := vector.D(0); i < k; i++ {
		if rmax[i] < smin[i] {
			return true
		}
	}

	for i := vector.D(0); i < k; i++ {
		if smax[i] < rmin[i] {
			return true
		}
	}

	return false
}

func V(r R) float64 {
	v := 1.0
	rmin, rmax := r.Min(), r.Max()
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		v *= rmax[i] - rmin[i]
	}
	return v
}

// SA returns the "surface area" of an N-dimensional interval. For N = 2, this
// is the perimeter, and for N = 3, this is the total surface area of the
// rectangular prism.
//
// See https://math.stackexchange.com/a/1898563 for the full method.
func SA(r R) float64 {
	k := r.Min().Dimension()

	rmin, rmax := r.Min(), r.Max()

	switch k {
	case 1:
		return 0
	case 2:
		dx := rmax[vector.AXIS_X] - rmin[vector.AXIS_X]
		dy := rmax[vector.AXIS_Y] - rmin[vector.AXIS_Y]
		return 2*dx + 2*dy
	case 3:
		dx := rmax[vector.AXIS_X] - rmin[vector.AXIS_X]
		dy := rmax[vector.AXIS_Y] - rmin[vector.AXIS_Y]
		dz := rmax[vector.AXIS_Z] - rmin[vector.AXIS_Z]
		return 2*dx*dy + 2*dy*dz + 2*dx*dz
	}

	v := V(r)
	var sa float64

	// Special case for k = 1, where SA will return 1, even though the
	// "true" area is 0. This is handled above.
	//
	// Note that this is an optimization from the StackExchange method -- we
	// note that each dimension's contribution to the overall surface area
	// is approximately the same (i.e. the volume), save that we are
	// excluding the volume contribution from that iterated dimension.
	for i := vector.D(0); i < k; i++ {
		if d := rmax[i] - rmin[i]; d > 0 {
			sa += 2.0 * v / d
		}
	}
	return sa
}

func WithinEpsilon(g R, h R, e epsilon.E) bool {
	return vector.WithinEpsilon(g.Min(), h.Min(), e) && vector.WithinEpsilon(g.Max(), h.Max(), e)
}

func Within(g R, h R) bool { return WithinEpsilon(g, h, epsilon.DefaultE) }
