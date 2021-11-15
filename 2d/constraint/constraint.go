// Package constraint implements a 2D linear constraint embedded in 2D ambient
// space.
package constraint

import (
	"github.com/downflux/go-geometry/nd/constraint"
	"github.com/downflux/go-geometry/nd/vector"

	v2d "github.com/downflux/go-geometry/2d/vector"
)

// C defines an N-dimensional linear constraint of the (standard) form embedded
// in N-dimensional ambient space.
//
//   A • X <= B
type C constraint.C

func New(p v2d.V, n v2d.V) *C {
	c := C(*constraint.New(vector.V(p), vector.V(n)))
	return &c
}

func (c C) A() []float64    { return constraint.C(c).A() }
func (c C) B() float64      { return constraint.C(c).B() }
func (c C) In(v v2d.V) bool { return constraint.C(c).In(vector.V(v)) }
