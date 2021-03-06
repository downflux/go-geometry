// Package hypersphere is a 2D ball embedded in 2D ambient space, i.e. a circle
// on the XY-plane.
package hypersphere

import (
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/hypersphere"
	"github.com/downflux/go-geometry/nd/vector"

	v2d "github.com/downflux/go-geometry/2d/vector"
)

type C hypersphere.C

func New(p v2d.V, r float64) *C {
	c := C(*hypersphere.New(vector.V(p), r))
	return &c
}

func (c C) R() float64      { return hypersphere.C(c).R() }
func (c C) P() v2d.V        { return v2d.V(hypersphere.C(c).P()) }
func (c C) In(p v2d.V) bool { return hypersphere.C(c).In(vector.V(p)) }
func WithinEpsilon(c C, d C, e epsilon.E) bool {
	return hypersphere.WithinEpsilon(hypersphere.C(c), hypersphere.C(d), e)
}
func Within(c C, d C) bool { return hypersphere.Within(hypersphere.C(c), hypersphere.C(d)) }
