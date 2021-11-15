// Package segment is a 1D line segment embedded in 2D ambient space.
package segment

import (
	"github.com/downflux/go-geometry/nd/line"
	"github.com/downflux/go-geometry/nd/segment"
	"github.com/downflux/go-geometry/nd/vector"

	l2d "github.com/downflux/go-geometry/2d/line"
	v2d "github.com/downflux/go-geometry/2d/vector"
)

type S segment.S

func New(l l2d.L, min float64, max float64) *S {
	s := S(*segment.New(line.L(l), min, max))
	return &s
}

func (s S) L() l2d.L          { return l2d.L(segment.S(s).L()) }
func (s S) TMin() float64     { return segment.S(s).TMin() }
func (s S) TMax() float64     { return segment.S(s).TMax() }
func (s S) T(v v2d.V) float64 { return segment.S(s).T(vector.V(v)) }
func (s S) Feasible() bool    { return segment.S(s).Feasible() }
