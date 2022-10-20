package hyperrectangle

import (
	"math"

	"github.com/downflux/go-geometry/nd/vector"
)

type M R

func (r M) R() R          { return R(r) }
func (r M) Min() vector.V { return R(r).Min() }
func (r M) Max() vector.V { return R(r).Max() }

func (r M) Copy(s R) {
	r.Min().M().Copy(s.Min())
	r.Max().M().Copy(s.Max())
}

func (r M) Zero() {
	r.Min().M().Zero()
	r.Max().M().Zero()
}

func (r M) Intersect(s R) bool {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		r.Min().M().SetX(i, math.Max(r.Min()[i], s.Min()[i]))
		r.Max().M().SetX(i, math.Min(r.Max()[i], s.Max()[i]))

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
		r.Min().M().SetX(i, math.Min(r.Min()[i], s.Min()[i]))
		r.Max().M().SetX(i, math.Max(r.Max()[i], s.Max()[i]))
	}
}
