// Package plane implements an 2D plane in N-dimensional ambient space.
package plane

import (
	"fmt"
	"math"

	"github.com/downflux/go-geometry/nd/vector"
)

type P struct {
	n vector.V

	p vector.V
}

func New(p vector.V, n vector.V) *P {
	if n.Dimension() != p.Dimension() {
		panic(fmt.Sprintf("cannot construct a plane with mismatching %v-dimensional offset and %v-dimensional normal vectors", p.Dimension(), n.Dimension()))
	}
	return &P{
		n: n,
		p: p,
	}
}

func (p P) N() vector.V { return p.n }
func (p P) P() vector.V { return p.p }

func (p P) Distance(v vector.V) float64 {
	return math.Abs(
		vector.Dot(
			vector.Unit(p.N()),
			vector.Sub(v, p.P()),
		),
	)
}
