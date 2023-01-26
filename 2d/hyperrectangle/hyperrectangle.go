package hyperrectangle

import (
	"github.com/downflux/go-geometry/2d/vector"
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/hyperrectangle"

	vnd "github.com/downflux/go-geometry/nd/vector"
)

type R hyperrectangle.R

func New(min vector.V, max vector.V) *R { return (*R)(hyperrectangle.New(vnd.V(min), vnd.V(max))) }

func (r R) M() M          { return M(r) }
func (r R) Min() vector.V { return vector.V(hyperrectangle.R(r).Min()) }
func (r R) Max() vector.V { return vector.V(hyperrectangle.R(r).Max()) }
func (r R) D() vector.V   { return vector.V(hyperrectangle.R(r).D()) }

func (r R) In(v vector.V) bool { return hyperrectangle.R(r).In(vnd.V(v)) }

func Intersect(r R, s R) (R, bool) {
	t, ok := hyperrectangle.Intersect(hyperrectangle.R(r), hyperrectangle.R(s))
	return R(t), ok
}

func Union(r R, s R) R       { return R(hyperrectangle.Union(hyperrectangle.R(r), hyperrectangle.R(s))) }
func Scale(r R, c float64) R { return R(hyperrectangle.Scale(hyperrectangle.R(r), c)) }

func Contains(r R, s R) bool {
	return hyperrectangle.Contains(hyperrectangle.R(r), hyperrectangle.R(s))
}

func Disjoint(r R, s R) bool {
	return hyperrectangle.Disjoint(hyperrectangle.R(r), hyperrectangle.R(s))
}

func V(r R) float64  { return hyperrectangle.V(hyperrectangle.R(r)) }
func SA(r R) float64 { return hyperrectangle.SA(hyperrectangle.R(r)) }

func WithinEpsilon(g R, h R, e epsilon.E) bool {
	return hyperrectangle.WithinEpsilon(hyperrectangle.R(g), hyperrectangle.R(h), e)
}
func Within(g R, h R) bool { return hyperrectangle.Within(hyperrectangle.R(g), hyperrectangle.R(h)) }
