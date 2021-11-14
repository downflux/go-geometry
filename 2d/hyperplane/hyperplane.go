// Package hyperplane defines an 1D line bisecting 2D ambient space.
package hyperplane

import (
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
