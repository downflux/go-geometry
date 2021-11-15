// Package constraint implements an N-dimensional linear constraint embedded in
// N-dimensional ambient space.
package constraint

import (
	"github.com/downflux/go-geometry/nd/hyperplane"
	"github.com/downflux/go-geometry/nd/vector"
)

// C defines an N-dimensional linear constraint of the (standard) form embedded
// in N-dimensional ambient space.
//
//   A • X <= B
//
// For two dimensions, this is
//
//   a_x * x + a_y * y <= b
type C hyperplane.HP

func New(p vector.V, n vector.V) *C {
	c := C(*hyperplane.New(p, n))
	return &c
}

// A returns the A vector of the contraint; returns [a, b] in the 2D case.
func (c C) A() []float64 {
	a := vector.Scale(-1, hyperplane.HP(c).N())

	xs := make([]float64, a.Dimension())
	for i := vector.D(0); i < a.Dimension(); i++ {
		xs[i] = a.X(i)
	}

	return xs
}

// B returns the bound on the constraint.
func (c C) B() float64 {
	a := vector.Scale(-1, hyperplane.HP(c).N())
	return vector.Dot(a, hyperplane.HP(c).P())
}

func (c C) In(v vector.V) bool {
	return vector.Dot(vector.V(c.A()), v) <= c.B()
}
