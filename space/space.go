// Package space defines a an parametric N-dimensional Euclidean space.
package space

import (
	"fmt"
	// "math"

	// "github.com/downflux/go-geometry/circle"
	// "github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/vector"
	// v2d "github.com/downflux/go-geometry/vector/v2d"
)

// L represents a parametric 1D line embedded into N-dimensional ambient space, for N > 1.
type L struct {
	p vector.V
	n vector.V
}

func (s S) P() vector.V { return s.p }

// S represents a parametric N-dimensional Euclidean space. We may represent an
// N = 1 space (i.e. a line) by the normal parametric equation
//
// l(t) = b + tn
//
// Where l, b, and n are 1D vectors, and t is the parameter. The line is defined
// over all real values of t.
//
// Similarly, we may define a 2D space as
//
//   L(t) := P + tN
//
// where L, P, and N are 2D vectors.
//
// We have renamed "x" in the 1D case to be more suggestive of the nature of
// this vector -- that is, it represents the normal of the sweeping hyperplane.
type S struct {
	// p represents the N-dimensional offset of the space.
	p vector.V

	// n represents the N-dimensional normal of the sweeping hyperplane.
	n vector.V

	d vector.D
}

func New(p vector.V, n vector.V) *S {
	if p.D() != n.D() {
		panic(
			fmt.Sprintf(
				"input %v-dimensional directional vector does not match the %v-dimensional offset",
				p.D(),
				n.D(),
			),
		)
	}

	return &S{p: p, n: n}
}

func (s S) D() int { return s.d }

func (s S) P() vector.V { return s.p }
func (s S) N() vector.V { return s.n }

// L returns the N-dimensional projected hyperplane at the given parametric
// input, i.e. L(t) = P + tN.
//
// Note that the vector returned represents the normal of a hyperplane.
func (s S) L(t float64) vector.V { return vector.Add(s.P(), vector.Scale(t, s.N())) }

// T takes as input a vector in the N-dimensional Euclidean space and
// returns the projected t-value of v onto L.
func (s S) T(v vector.V) float64 { return vector.Dot(s.N(), vector.Sub(v, s.P())) }

func (s S) Intersect(r S) (S, bool) {
	if s.D() != r.D() {
		panic("cannot calculate the intersection of different dimensional spaces")
	}

	if len(s.D() > 3) {

	}
}

/*


// Intersect returns a list of basis vectors spanning the intersection of two
// hyperplanes; the intersection for an N-dimensional hyperplane in N + 1
// ambient space is an (N - 1)-dimensional affine plane.
func (pl P) Intersect(pm P) ([]vector.V, bool) {
	if len(pl.Basis()) != len(pm.Basis()) {
		panic("cannot calculate the intersection of different dimensional hyperplanes")
	}

	// Intersect returns the intersection point between two lines.
	//
	// Returns error if the lines are parallel.
	//
	// Find the intersection between the two lines as a function of the
	// constraint parameter t.
	//
	// Given two constraints L, M, we need to find their intersection; WLOG,
	// let's project the intersection point onto L.
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
	// We want to find the projection onto L, which means we need to find a
	// concrete value for t. the other parameter u doesn't matter so much --
	// let's try to get rid of it.
	//
	//   uE = P - Q + tD
	//
	// Here, we know P, D, Q, and E are vectors, and we can decompose these
	// into a system of equations by isolating their orthogonal (e.g.
	// horizontal and vertical) components.
	//
	//   uEx = Px - Qx + tDx
	//   uEy = Py - Qy + tDy
	//
	// Solving for u, we get
	//
	//   (Px - Qx + tDx) / Ex = (Py - Qy + tDy) / Ey
	//   => Ey (Px - Qx + tDx) = Ex (Py - Qy + tDy)
	//
	// We leave the task of simplifying the above terms as an exercise to
	// the reader. Isolating t, and noting some common substitutions, we get
	//
	//   t = || E x (P - Q) || / || D x E ||
	//
	// See https://gamedev.stackexchange.com/a/44733 for more information.
	if len(pl.Basis()) == 1 {
		d := v2d.V(pl.Basis()[0])
		e := v2d.V(pm.Basis()[0])

		p := v2d.V(pl.C())
		q := v2d.V(pm.C())

		denominator := v2d.Determinant(d, e)

		// Parallel lines do not have an intersection.
		if epsilon.Within(denominator, 0) {
			return nil, false
		}

		t := v2d.Determinant(e, v2d.Sub(p, q)) / denominator

		return []vector.V{pl.T(*vector.New(t))}, true
	}

	// TODO(minkezhang): Implement n-dimensional hyperplane intersection.
	return nil, false
}

/*
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
