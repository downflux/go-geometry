// Package hyperplane defines an (N - 1)-dimensional geometric hyperplane which is
// embedded in the N-dimensional ambient space.
//
// In 2D ambient space, the hyperplane is a 1D linear constraint of two
// variables (i.e. a line that lies on the XY-plane).
package hyperplane

import (
	"fmt"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

// HP defines an (N - 1)-dimensional hyperplane geometrically consisting of an
// N-dimensional normal vector n to the plane, and an N-dimensional offset p
// which relates the hyperplane to the ambient space origen.
//
// N.B.: By arbitrary convention, vectors pointing away from N are not
// permissible within the half-plane.
type HP struct {
	n vector.V
	p vector.V
}

// New constructs a half-plane passing through a point P and with normal N.
func New(p vector.V, n vector.V) *HP {
	if p.Dimension() != n.Dimension() {
		panic(
			fmt.Sprintf(
				"cannot construct a hyperplane with mismatching %v-dimensional offset and %v-dimensional normal vectors",
				p.Dimension(),
				n.Dimension(),
			),
		)
	}
	return &HP{
		n: n,
		p: p,
	}
}

func (hp HP) P() vector.V { return hp.p }

// N returns the normal vector of the plane, pointing away from the invalid
// region.
func (hp HP) N() vector.V { return hp.n }

// In checks if a given point in vector space is in valid region of the
// half-plane.
func (hp HP) In(p vector.V) bool {
	// Generate a vector with tail on D and pointing towards the input.
	v := vector.Sub(p, hp.P())

	// Check relative orientation between v and D.
	//
	// Remember that by the right hand rule, if v is on the "left" of the
	// hyperplane if for all basis vectors b of the hyperplane,
	//
	//   b x v > 0, and
	//   N • v < 0
	//
	// As the left half of the plane is considered invalid, we are looking
	// instead for the complementary result.
	return vector.Dot(hp.N(), v) >= 0
}

// Disjoint returns if the region of interection between two planes is empty.
//
// Disjoint checks if the characteristic lines of the two planes are parallel,
// and if the a line drawn between two points on the planes, away from the first
// plane, lie in the feasible region of the first plane.
func Disjoint(a HP, b HP) bool {
	return vector.Within(a.N(), vector.Scale(-1, b.N())) && !a.In(vector.Sub(b.P(), a.P()))
}

func WithinEpsilon(a HP, b HP, e epsilon.E) bool {
	return vector.WithinEpsilon(a.N(), b.N(), e) && vector.WithinEpsilon(a.P(), b.P(), e)
}

func Within(a HP, b HP) bool { return WithinEpsilon(a, b, epsilon.DefaultE) }
