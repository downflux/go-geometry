package hyperrectangle

import (
	"math"

	"github.com/downflux/go-geometry/nd/vector"
)

type M R

func (r M) R() R          { return R(r) }
func (r M) Min() vector.M { return R(r).Min().M() }
func (r M) Max() vector.M { return R(r).Max().M() }

func (r M) Copy(s R) {
	r.Min().Copy(s.Min())
	r.Max().Copy(s.Max())
}

func (r M) Zero() {
	r.Min().Zero()
	r.Max().Zero()
}

func (r M) Intersect(s R) bool {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		r.Min().SetX(i, math.Max(r.Min()[i], s.Min()[i]))
		r.Max().SetX(i, math.Min(r.Max()[i], s.Max()[i]))

		if r.Min()[i] > r.Max()[i] {
			return false
		}
	}

	return true
}

func (r M) Union(s R) {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		r.Min().SetX(i, math.Min(r.Min()[i], s.Min()[i]))
		r.Max().SetX(i, math.Max(r.Max()[i], s.Max()[i]))
	}
}

// Scale will expand or shrink each dimension of the AABB by the given scalar.
//
// If we want to expand or shrink by the total volume instead, we can use
//
//	b.Scale(math.Pow(c, 1.0 / b.Min().Dimension()))
func (r M) Scale(c float64) {
	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		min := r.Min().X(i)
		max := r.Max().X(i)

		r.Max().M().SetX(i, min+((max-min)*c))
	}
}
