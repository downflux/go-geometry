// Package hyperplane defines a 1D geometric hyperplane (i.e. a linear
// constraint of two variables) embedded in 2D ambient space.
package hyperplane

import (
	"github.com/downflux/go-geometry/2d/line"
	"github.com/downflux/go-geometry/2d/vector"
)

// HP defines a half-plane, geometrically consisting of a normal vector n to the
// plane, and a point p which defines the origin of n.
//
// N.B.: By arbitrary convention, vectors pointing away from N are not
// permissible within the half-plane.
type HP struct {
	l line.L
}

// New constructs a half-plane passing through a point P and with normal N.
//
// TODO(minkezhang): Refactor to use L(P, N) instead; define D as clockwise
// rotation from N.
func New(p vector.V, n vector.V) *HP {
	return &HP{
		l: *line.New(p, n),
	}
}

// Basis returns a set of N-dimensional vectors whose span is the hyperplane. In
// 2D ambient space, a hyperspace is a line, and its basis is a singular 2D
// vector pointing away from the hyperplane normal. Note that the span of the
// basis is not that of the ambient space -- the cardinality of the returned
// value must match the the dimension of the underlying hyperplane, i.e. N - 1.
//
// A vector v relative to the hyperplane is considered to be in the feasible
// region of the hyperplane if for all basis vectors b,
//
//   v x b <= 0
//
// In 2D ambient space, "left" of the line are not permissible.
//
// N.B.: RVO2 defines the "right" side of the line as non-permissible, but we
// have considered an anti-clockwise rotation of N(), e.g. +X to +Y, to be more
// natural. See
// https://github.com/snape/RVO2/blob/57098835aa27dda6d00c43fc0800f621724884cc/src/Agent.cpp#L314
// for evidence of this distinction.
func (hp HP) Basis() []vector.V {
	return []vector.V{*vector.New(-hp.N().Y(), hp.N().X())}
}

func (hp HP) P() vector.V { return hp.l.P() }

// N returns the normal vector of the plane, pointing away from the invalid
// region.
func (hp HP) N() vector.V { return hp.l.D() }

// In checks if a given point in vector space is in valid region of the
// half-plane.
func (hp HP) In(p vector.V) bool {
	// Generate a vector with tail on D and pointing towards the input.
	v := vector.Sub(p, hp.P())

	// Check relative orientation between v and D.
	//
	// Remember that by the right hand rule, if v is on the "left" of the
	// hyperplane,
	//
	//   D x v > 0, and
	//   N • v > 0
	//
	// As the left half of the plane is considered invalid, we are looking
	// instead for the complementary result.
	in := true
	for _, b := range hp.Basis() {
		in = in && vector.Determinant(b, v) <= 0
	}
	return in
}

// Disjoint returns if the region of interection between two planes is empty.
//
// Disjoint checks if the characteristic lines of the two planes are parallel,
// and if the a line drawn between two points on the planes, away from the first
// plane, lie in the feasible region of the first plane.
func Disjoint(a HP, b HP) bool {
	return vector.Within(a.N(), vector.Scale(-1, b.N())) && !a.In(vector.Sub(b.P(), a.P()))
}

func Within(a HP, b HP) bool {
	return vector.Within(a.N(), b.N()) && vector.Within(a.P(), b.P())
}
