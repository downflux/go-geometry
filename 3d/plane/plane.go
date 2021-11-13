// Package plane is a 2D plane embedded in 3D ambient space.
package plane

import (
	"github.com/downflux/go-geometry/nd/plane"
	"github.com/downflux/go-geometry/nd/vector"

	v3d "github.com/downflux/go-geometry/3d/vector"
	l3d "github.com/downflux/go-geometry/3d/line"
)

type P plane.P

func New(p v3d.V, n v3d.V) *P {
	pl := P(*plane.New(vector.V(p), vector.V(n)))
	return &pl
}

func (p P) N() v3d.V                 { return v3d.V(p.N()) }
func (p P) P() v3d.V                 { return v3d.V(p.P()) }
func (p P) Distance(v v3d.V) float64 { return plane.P(p).Distance(vector.V(v)) }

func (p P) Intersect(q P) (l3d.L, bool) {
	d := v3d.Cross(p.N(), q.N())
	if epsilon(v3d.SquaredMagnitude(d), 0) {
		return l3d.L{}, false
	}

	d = v3d.Unit(d)

	// A line is defined by both the direction and the offset; in order to
	// find a point X_0 on this line, we set a non-zero component of the
	// normal vector to 0, which lets us solve a two-constraint system of
	// equations
	//
	//   X = X_0 + tD
	//
	// for the remaining components -- that is, we can model the problem as
	//
	//   x = x_0 + td_x
	//   y = y_0 + td_y
	//
	// For the other two axis x and y.
	//
	// These lines are still in 2D space -- so we can solve the intersection
	// of these 1D lines in 2D space to find a specific solution.
	//
	// See https://mathworld.wolfram.com/Plane-PlaneIntersection.html for
	// more information.
	axes := []vector.D{
		vector.AXIS_X,
		vector.AXIS_Y,
		vector.AXIS_Z,
	}
	var zero vector.D

	for i := vector.D(0); i < d.Dimension(); i++ {
		if d.X(i) != 0 && len(axes) > 2 {
			zero = i
			delete(axes, i)
		}
	}

	projected := []vector.D{}
	for i := range axes

	return l3d.New(x, d), true
}
