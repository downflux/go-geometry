// Package hypersphere is an N-dimensional ball embedded into an N-dimensional ambient space.
package hypersphere

import (
	"math"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

type C struct {
	r float64
	p vector.V
}

func New(p vector.V, r float64) *C {
	return &C{r: r, p: p}
}

func (c C) R() float64  { return math.Abs(c.r) }
func (c C) P() vector.V { return c.p }

func (c C) In(p vector.V) bool {
	m := vector.SquaredMagnitude(vector.Sub(p, c.P()))
	r := c.R() * c.R()
	// Rounding errors could result in the vector difference to lie slightly
	// outside the circle.
	return m < r || epsilon.Within(m, r)
}

func WithinEpsilon(c C, d C, e epsilon.E) bool {
	return vector.WithinEpsilon(c.P(), d.P(), e) && e.Within(c.R(), d.R())
}

func Within(c C, d C) bool { return WithinEpsilon(c, d, epsilon.DefaultE) }
