// Package hyperplane defines a 2D geometric half-plane embedded in 3D ambient space.
package hyperplane

import (
	"github.com/downflux/go-geometry/line/l2d"
	"github.com/downflux/go-geometry/vector/v2d"
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
func New(p vector.V, n vector.V) *HP {
	// D returns the characteristic line along the plane is bisected. Points
	// to the "left" of the line are not permissible.
	//
	// N.B.: RVO2 defines the "right" side of the line as non-permissible,
	// but we have considered an anti-clockwise rotation of N(), e.g. +X to
	// +Y, to be more natural. See
	// https://github.com/snape/RVO2/blob/57098835aa27dda6d00c43fc0800f621724884cc/src/Agent.cpp#L314
	// for evidence of this distinction.
	d := *vector.New(-n.Y(), n.X())

	return &HP{
		l: *line.New(p, d),
	}
}

func (hp HP) L() line.L { return hp.l }

func (hp HP) D() vector.V { return hp.l.D() }
func (hp HP) P() vector.V { return hp.l.P() }

// N returns the normal vector of the plane, pointing away from the invalid
// region.
func (hp HP) N() vector.V { return *vector.New(hp.D().Y(), -hp.D().X()) }

// In checks if a given point in vector space is in valid region of the
// half-plane.
func (hp HP) In(p vector.V) bool {
	// Generate a vector with tail on D and pointing towards the input.
	v := vector.Sub(p, hp.P())

	// Check relative orientation between w and D.
	//
	// Remember that by the right hand rule, if v is on the "left" of the
	// plane,
	//
	//   D x v > 0, and
	//   N • v > 0
	//
	// As the left half of the plane is considered invalid, we are looking
	// instead for the complementary result.
	return vector.Determinant(hp.D(), v) <= 0
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
