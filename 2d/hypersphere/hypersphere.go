// Package hypersphere is a 2D ball embedded in 2D ambient space, i.e. a circle
// on the XY-plane.
package hypersphere

import (
	"github.com/downflux/go-geometry/nd/hypersphere"
	"github.com/downflux/go-geometry/nd/vector"

	v2d "github.com/downflux/go-geometry/2d/vector"
)

type C hypersphere.C

func New(p v2d.V, r float64) *C {
	c := C(*hypersphere.New(vector.V(p), r))
	return &c
}

func (c C) R() float64 { return hypersphere.C(c).R() }
func (c C) P() v2d.V   { return v2d.V(hypersphere.C(c).P()) }
func (c C) In(p v2d.V) bool {
	return v2d.SquaredMagnitude(v2d.Sub(p, c.P())) <= c.R()*c.R()
}
