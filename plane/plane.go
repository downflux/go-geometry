// Plane defines a parametric N-dimensional geometric plane.
package plane

import (
	// "math"

	// "github.com/downflux/go-geometry/circle"
	// "github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/vector"
)

// P represents an N-dimensional plane embedded into a possibly higher
// dimensonal space. When N = 1, the "plane" is a line of the form
//
//   P := C + tD
//
// Where C is some offset vector, and D is the single directional basis of the
// ilne.
type P struct {
	// basis is a list of vectors whose span define the plane. The basis
	// does not have to be orthonormal.
	//
	// The number of elements in the basis represent the dimension of the
	// plane.
	basis []vector.V

	// c represents the offset of the plane in higher-dimensional space.
	c vector.V
}

func New(c vector.V, basis []vector.V) *P { return &P{c: c, basis: basis} }

func (pl P) Basis() []vector.V { return pl.basis }
func (pl P) D() vector.D       { return vector.D(len(pl.Basis())) }
func (pl P) C() vector.V       { return pl.c }

// T returns the N-dimensional projected point of the plane, given a parametric
// vector input. The dimension of the parametric vector must match the dimension
// of the plane.
func (pl P) T(t vector.V) vector.V {
	var r vector.V

	for d, b := range pl.Basis() {
		r = vector.Add(r, vector.Add(pl.C(), vector.Scale(t.X(vector.D(d)), b)))
	}
	return r
}

// Project takes as input a higher-dimensional point, and returns a parametric
// vector representing the point on the plane closest to the input.
func (pl P) Project(v vector.V) vector.V {
	xs := make([]float64, len(pl.Basis()))

	for d, b := range pl.Basis() {
		xs[d] = vector.Dot(b, vector.Sub(v, pl.C()))
	}
	return *vector.New(xs...)
}

/*
// L defines a parametric line of the form
//
//   L := P + tD
type L struct {
	p vector.V
	d vector.V
}

func New(p vector.V, d vector.V) *L { return &L{p: p, d: d} }

func (l L) P() vector.V { return l.p }
func (l L) D() vector.V { return l.d }

// T calculates the vector value on the line which corresponds to the input
// parametric t-value.
func (l L) T(t float64) vector.V { return vector.Add(l.p, vector.Scale(t, l.d)) }

// Project calculates the projected t-value of v onto l by finding the point on
// L closest to v.
func (l L) Project(v vector.V) float64 {
	return vector.Dot(l.D(), vector.Sub(v, l.P()))
}

// Intersect returns the intersection point between two lines.
//
// Returns error if the lines are parallel.
//
// Find the intersection between the two lines as a function of the constraint
// parameter t.
//
// Given two constraints L, M, we need to find their intersection; WLOG, let's
// project the intersection point onto L.
//
// We know the parametric equation form of these lines -- that is,
//
//   L = P + tD
//   M = Q + uE
//
// At their intersection, we know that L meets M:
//
//   L = M
//   => P + tD = Q + uE
//
// We want to find the projection onto L, which means we need to find a concrete
// value for t. the other parameter u doesn't matter so much -- let's try to get
// rid of it.
//
//   uE = P - Q + tD
//
// Here, we know P, D, Q, and E are vectors, and we can decompose these into a
// system of equations by isolating their orthogonal (e.g. horizontal and
// vertical) components.
//
//   uEx = Px - Qx + tDx
//   uEy = Py - Qy + tDy
//
// Solving for u, we get
//
//   (Px - Qx + tDx) / Ex = (Py - Qy + tDy) / Ey
//   => Ey (Px - Qx + tDx) = Ex (Py - Qy + tDy)
//
// We leave the task of simplifying the above terms as an exercise to the
// reader. Isolating t, and noting some common substitutions, we get
//
//   t = || E x (P - Q) || / || D x E ||
//
// See https://gamedev.stackexchange.com/a/44733 for more information.
func (l L) Intersect(m L) (vector.V, bool) {
	d := vector.Determinant(l.D(), m.D())
	n := vector.Determinant(m.D(), vector.Sub(l.P(), m.P()))

	if epsilon.Within(d, 0) {
		return vector.V{}, false
	}

	return l.T(n / d), true
}

// IntersectCircle returns the intersection points between a line and a circle.
// If the line does not intersect a circle, the function will return not
// successful.
//
// As a line stretches to infinity in both directions, it is not possible for a
// line to intersect the circle partway.
//
// If the line lies tangent to the circle, then the returned t-values are the
// same.
//
// See https://stackoverflow.com/a/1084899 for more information.
func (l L) IntersectCircle(c circle.C) (vector.V, vector.V, bool) {
	p := vector.Sub(l.P(), c.P())

	dot := vector.Dot(p, l.D())
	discriminant := dot*dot + c.R()*c.R() - vector.SquaredMagnitude(p)

	// The line does not intersect the circle.
	if discriminant < 0 {
		return vector.V{}, vector.V{}, false
	}

	// Find two intersections between line and circle. This is equivalent to
	// having two additional constraints which lie tangent to the circle at
	// these two points.
	tl := -dot - math.Sqrt(discriminant)
	tr := -dot + math.Sqrt(discriminant)

	return l.T(tl), l.T(tr), true
}

// Distance finds the distance between the line l and a point p.
//
// The distance from a line L to a point Q is given by
//
//   d := || D x (Q - P) || / || D ||
//
// See
// https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line#Another_vector_formulation
// for more information.
func (l L) Distance(p vector.V) float64 {
	v := vector.Sub(p, l.P())
	return math.Abs(vector.Determinant(l.D(), v) / vector.Magnitude(l.D()))
}
*/
