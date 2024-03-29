// Package line implements a 1D line in 2D ambient space.
package line

import (
	"math"

	"github.com/downflux/go-geometry/2d/hypersphere"
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/line"
	"github.com/downflux/go-geometry/nd/vector"

	v2d "github.com/downflux/go-geometry/2d/vector"
)

type L line.L

func New(p v2d.V, d v2d.V) *L {
	l := L(*line.New(vector.V(p), vector.V(d)))
	return &l
}

func (l L) P() v2d.V          { return v2d.V(line.L(l).P()) }
func (l L) D() v2d.V          { return v2d.V(line.L(l).D()) }
func (l L) L(t float64) v2d.V { return v2d.V(line.L(l).L(t)) }
func (l L) T(v v2d.V) float64 { return line.L(l).T(vector.V(v)) }
func (l L) Parallel(m L) bool { return line.L(l).Parallel(line.L(m)) }

// We are defining a line normal which is consistent with our 2D hyperplane
// definition -- that is, if this line is a 2D hyperplane, the normal points
// into the feasible region.
//
// We define the normal as a clockwise π / 2 rotation of the directional
// vector.
func (l L) N() v2d.V { return *v2d.New(l.D().Y(), -l.D().X()) }

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
//	L = P + tD
//	M = Q + uE
//
// At their intersection, we know that L meets M:
//
//	L = M
//	=> P + tD = Q + uE
//
// We want to find the projection onto L, which means we need to find a concrete
// value for t. the other parameter u doesn't matter so much -- let's try to get
// rid of it.
//
//	uE = P - Q + tD
//
// Here, we know P, D, Q, and E are vectors, and we can decompose these into a
// system of equations by isolating their orthogonal (e.g. horizontal and
// vertical) components.
//
//	uEx = Px - Qx + tDx
//	uEy = Py - Qy + tDy
//
// Solving for u, we get
//
//	(Px - Qx + tDx) / Ex = (Py - Qy + tDy) / Ey
//	=> Ey (Px - Qx + tDx) = Ex (Py - Qy + tDy)
//
// We leave the task of simplifying the above terms as an exercise to the
// reader. Isolating t, and noting some common substitutions, we get
//
//	t = || E x (P - Q) || / || D x E ||
//
// See https://gamedev.stackexchange.com/a/44733 for more information.
func (l L) Intersect(m L) (v2d.V, bool) {
	d := v2d.Determinant(l.D(), m.D())
	n := v2d.Determinant(m.D(), v2d.Sub(l.P(), m.P()))

	if epsilon.Within(d, 0) {
		return v2d.V{}, false
	}

	return l.L(n / d), true
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
func (l L) IntersectCircle(c hypersphere.C) (v2d.V, v2d.V, bool) {
	d := l.Distance(c.P())
	if d > c.R() {
		return v2d.V{}, v2d.V{}, false
	}

	t := l.T(c.P())

	discrimanant := (c.R()*c.R() - d*d) / v2d.SquaredMagnitude(l.D())
	return l.L(t - math.Sqrt(discrimanant)), l.L(t + math.Sqrt(discrimanant)), true
}

// Distance finds the distance between the line l and a point p.
//
// The distance from a line L to a point Q is given by
//
//	d := || D x (Q - P) || / || D ||
//
// See
// https://en.wikipedia.org/wiki/Distance_from_a_point_to_a_line#Another_vector_formulation
// for more information.
func (l L) Distance(p v2d.V) float64 {
	v := v2d.Sub(p, l.P())
	return math.Abs(v2d.Determinant(l.D(), v) / v2d.Magnitude(l.D()))
}

func WithinEpsilon(l L, m L, e epsilon.E) bool { return line.WithinEpsilon(line.L(l), line.L(m), e) }
func Within(l L, m L) bool                     { return line.Within(line.L(l), line.L(m)) }
