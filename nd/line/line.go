// Package line implements a 1D line in N-dimensional ambient space.
package line

import (
	"github.com/downflux/go-geometry/nd/vector"
)

// L defines a parametric line of the form
//
//   L := P + tD
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
func (l L) T(v vector.V) float64 {
	return vector.Dot(l.D(), vector.Sub(v, l.P()))
}
