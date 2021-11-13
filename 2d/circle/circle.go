// Package circle is a 2D circle embedded in 2D ambient space.
package circle

import (
	"github.com/downflux/go-geometry/2d/vector"
)

type C struct {
	r float64
	p vector.V
}

func New(p vector.V, r float64) *C {
	return &C{r: r, p: p}
}

func (c C) R() float64  { return c.r }
func (c C) P() vector.V { return c.p }

func (c C) In(p vector.V) bool {
	return vector.SquaredMagnitude(vector.Sub(p, c.P())) <= c.r*c.r
}
