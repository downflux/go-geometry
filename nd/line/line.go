// Package line implements a 1D line in N-dimensional ambient space.
package line

import (
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

// L defines a parametric line of the form
//
//	L := P + tD
type L struct {
	p vector.V
	d vector.V
}

func New(p vector.V, d vector.V) *L {
	return &L{p: p, d: d}
}

func (l L) P() vector.V { return l.p }
func (l L) D() vector.V { return l.d }

// L calculates the vector value on the line which corresponds to the input
// parametric t-value.
func (l L) L(t float64) vector.V { return vector.Add(l.p, vector.Scale(t, l.d)) }

// T calculates the projected t-value of v onto l by finding the point on L
// closest to v.
//
// Intuitively, we want to find the magnitude of the projected leg onto L; that
// is, we care about
//
//	||v - P||
//
// Therefore, we need a factor of ||D|| in the denominator.
//
// Furthermore, we know that t grows linearly with ||D||, so we need to add
// another ||D|| factor in the denominator as a normalization factor.
func (l L) T(v vector.V) float64 {
	return vector.Dot(l.D(), vector.Sub(v, l.P())) / vector.SquaredMagnitude(l.D())
}

// Parallel checks if two lines are parallel. A return value of false may
// indicate the lines intersect, are skew lines, or are anti-parallel.
//
// We check for the parallel property by examining the ratio of the vector
// directions.
//
// See https://stackoverflow.com/a/45181059/873865 for more details.
func (l L) Parallel(m L) bool {
	return epsilon.Within(
		vector.Dot(l.D(), m.D()),
		vector.Magnitude(l.D())*vector.Magnitude(m.D()),
	)
}

func WithinEpsilon(l L, m L, e epsilon.E) bool {
	return vector.WithinEpsilon(l.D(), m.D(), e) && vector.WithinEpsilon(l.P(), m.P(), e)
}

func Within(l L, m L) bool { return WithinEpsilon(l, m, epsilon.DefaultE) }
