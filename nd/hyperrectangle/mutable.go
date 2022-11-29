package hyperrectangle

import (
	"github.com/downflux/go-geometry/nd/vector"
)

type M R

func (r M) R() R          { return R(r) }
func (r M) Min() vector.M { return R(r).Min().M() }
func (r M) Max() vector.M { return R(r).Max().M() }

func (r M) Copy(s R) {
	copy(r.Min(), s.Min())
	copy(r.Max(), s.Max())
}

func (r M) Zero() {
	r.Min().Zero()
	r.Max().Zero()
}

func (r M) Intersect(s R) bool {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	rmin, rmax := r.Min(), r.Max()
	smin, smax := s.Min(), s.Max()

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		if rmin[i] < smin[i] {
			rmin[i] = smin[i]
		}
		if rmax[i] > smax[i] {
			rmax[i] = smax[i]
		}
		if rmin[i] > rmax[i] {
			return false
		}
	}

	return true
}

func (r M) Union(s R) {
	if r.Min().Dimension() != s.Min().Dimension() {
		panic("mismatching vector dimensions")
	}

	rmin, rmax := r.Min(), r.Max()
	smin, smax := s.Min(), s.Max()

	k := r.Min().Dimension()
	for i := vector.D(0); i < k; i++ {
		if smin[i] < rmin[i] {
			rmin[i] = smin[i]
		}
	}
	for i := vector.D(0); i < k; i++ {
		if smax[i] > rmax[i] {
			rmax[i] = smax[i]
		}
	}
}

// Scale will expand or shrink each dimension of the AABB by the given scalar.
//
// If we want to expand or shrink by the total volume instead, we can use
//
//	b.Scale(math.Pow(c, 1.0 / b.Min().Dimension()))
func (r M) Scale(c float64) {
	rmin, rmax := r.Min(), r.Max()

	for i := vector.D(0); i < r.Min().Dimension(); i++ {
		min := rmin[i]
		max := rmax[i]

		r.Max()[i] = min + ((max - min) * c)
	}
}
