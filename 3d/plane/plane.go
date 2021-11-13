// Package plane is a 2D plane embedded in 3D ambient space.
package plane

import (
	"github.com/downflux/go-geometry/nd/plane"
	"github.com/downflux/go-geometry/nd/vector"

	v3d "github.com/downflux/go-geometry/3d/vector"
)

type P plane.P

func New(p v3d.V, n v3d.V) *P {
	pl := P(*plane.New(vector.V(p), vector.V(n)))
	return &pl
}

func (p P) N() v3d.V                 { return v3d.V(p.N()) }
func (p P) P() v3d.V                 { return v3d.V(p.P()) }
func (p P) Distance(v v3d.V) float64 { return plane.P(p).Distance(vector.V(v)) }
