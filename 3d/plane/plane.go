// Package plane is a 2D plane embedded in 3D ambient space.
package plane

import (
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/plane"
	"github.com/downflux/go-geometry/nd/vector"

	l2d "github.com/downflux/go-geometry/2d/line"
	v2d "github.com/downflux/go-geometry/2d/vector"
	v3d "github.com/downflux/go-geometry/3d/vector"
	l3d "github.com/downflux/go-geometry/nd/line"
)

type P plane.P

func New(p v3d.V, n v3d.V) *P {
	pl := P(*plane.New(vector.V(p), vector.V(n)))
	return &pl
}

func (p P) N() v3d.V                 { return v3d.V(p.N()) }
func (p P) P() v3d.V                 { return v3d.V(p.P()) }
func (p P) Distance(v v3d.V) float64 { return plane.P(p).Distance(vector.V(v)) }
func (p P) R() float64               { return plane.P(p).R() }

func (p P) Intersect(q P) (l3d.L, bool) {
	d := vector.V(v3d.Cross(p.N(), q.N()))
	if epsilon.Within(vector.SquaredMagnitude(d), 0) {
		return l3d.L{}, false
	}

	d = vector.Unit(d)

	// A line is defined by both the direction and the offset; in order to
	// find a point X_0 on this line, we set a non-zero component of the
	// normal vector to 0, which lets us solve a two-constraint system of
	// equations
	//
	//   X = X_0 + tD
	//
	// for the remaining components u and w -- that is, we can model the
	// problem as
	//
	//   u = u_0 + td_u
	//   w = w_0 + td_w
	//
	// Note that these lines are still in 2D space, i.e.
	//
	//   u := (u, 0)
	//   w := (0, w)
	//
	// but just forcing a specific dimension to be 0; thus, we can solve the
	// intersection of these 1D lines in 2D space to find a specific
	// solution.
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
			axes[i] = axes[len(axes)-1]
			axes = axes[:len(axes)-1]
		}
	}

	// TODO(minkezhang): Solve for the point-intersect r.
	r := vector.V{}

	u := *l2d.New(*v2d.New(r.X(axes[0]), 0), *v2d.New(d.X(axes[0]), 0))
	w := *l2d.New(*v2d.New(0, r.X(axes[0])), *v2d.New(0, d.X(axes[1])))

	x, ok := u.Intersect(w)
	if !ok {
		return l3d.L{}, false
	}

	xs := make([]float64, d.Dimension())
	for _, a := range axes {
		xs[a] = vector.V(x).X(a)
	}
	xs[zero] = 0

	return *l3d.New(
		*vector.New(xs...),
		vector.V(d),
	), true
}
