package segment

import (
	"github.com/downflux/go-geometry/line"
	"github.com/downflux/go-geometry/vector"
)

type S struct {
	l   line.L
	min float64
	max float64
}

func New(l line.L, min float64, max float64) *S {
	return &S{
		l:   l,
		min: min,
		max: max,
	}
}

func (s S) L() line.L     { return s.l }
func (s S) TMin() float64 { return s.min }
func (s S) TMax() float64 { return s.max }

func (s S) Project(v vector.V) float64 {
	t := s.l.T(v)
	if t < s.TMin() {
		return s.TMin()
	} else if t > s.TMax() {
		return s.TMax()
	}
	return t
}

func (s S) Feasible() bool { return s.min <= s.max }
