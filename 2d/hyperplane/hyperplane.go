// Package hyperplane defines an 1D line bisecting 2D ambient space.
package hyperplane

import (
	"github.com/downflux/go-geometry/2d/line"
	"github.com/downflux/go-geometry/nd/hyperplane"
	"github.com/downflux/go-geometry/nd/vector"

	v2d "github.com/downflux/go-geometry/2d/vector"
)

type HP hyperplane.HP

func New(p v2d.V, n v2d.V) *HP {
	hp := HP(*hyperplane.New(vector.V(p), vector.V(n)))
	return &hp
}

func (hp HP) P() v2d.V        { return v2d.V(hyperplane.HP(hp).P()) }
func (hp HP) N() v2d.V        { return v2d.V(hyperplane.HP(hp).N()) }
func (hp HP) In(p v2d.V) bool { return hyperplane.HP(hp).In(vector.V(p)) }

func Disjoint(a HP, b HP) bool { return hyperplane.Disjoint(hyperplane.HP(a), hyperplane.HP(b)) }
func Within(a HP, b HP) bool   { return hyperplane.Within(hyperplane.HP(a), hyperplane.HP(b)) }

// Line returns the line of bisection representing the hyperplane.
//
// By arbitrary convention, points on the "left" of the line are infeasible in
// the hyperplane, so we orient the line direction vector π / 2 clockwise to N.
//
// N.B.: RVO2 defines the "right" side of the line as non-permissible, but we
// have considered an anti-clockwise rotation of D() to N() (e.g. +X to +Y) to
// be more natural. See
// https://github.com/snape/RVO2/blob/57098835aa27dda6d00c43fc0800f621724884cc/src/Agent.cpp#L314
// for evidence of this distinction.
func Line(hp HP) line.L {
	d := *v2d.New(hp.N().Y(), -hp.N().X())

	return *line.New(hp.P(), d)
}
