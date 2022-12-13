package ray

import (
	"math"

	"github.com/downflux/go-geometry/nd/hyperrectangle"
	"github.com/downflux/go-geometry/nd/vector"
)

type R struct {
	p vector.V
	d vector.V
}

func New(p vector.V, d vector.V) *R {
	u := vector.Unit(d)
	return &R{
		p: p,
		d: u,
	}
}

func (r R) P() vector.V { return r.p }
func (r R) D() vector.V { return r.d }

// IntersectHyperrectangle checks if the input ray collides with the
// hyperrectangle.
//
// See https://gamedev.stackexchange.com/a/18459.
func IntersectHyperrectangle(r R, s hyperrectangle.R) bool {
	k := s.Min().Dimension()

	if r.P().Dimension() != k {
		panic("mismatching vector dimensions")
	}

	u := vector.V(make([]float64, k)).M()
	u.Copy(r.D())

	for i := vector.D(0); i < k; i++ {
		u[i] = 1 / u[i]
	}

	tl := make([]float64, k)
	tr := make([]float64, k)

	smin, smax := s.Min(), s.Max()
	p := r.P()

	for i := vector.D(0); i < k; i++ {
		tl[i] = (smin[i] - p[i]) * u[i]
		tr[i] = (smax[i] - p[i]) * u[i]
	}

	// WLOG tmin is the minimal parametric value which puts s inside r.
	tmin := math.Inf(-1)
	tmax := math.Inf(1)
	for i := vector.D(0); i < k; i++ {
		min := tl[i]
		if tr[i] < tl[i] {
			tmin = tr[i]
		}
		if min > tmin {
			tmin = min
		}
	}
	for i := vector.D(0); i < k; i++ {
		max := tr[i]
		if tl[i] > tr[i] {
			tmax = tl[i]
		}
		if max < tmax {
			tmax = max
		}
	}

	return tmin <= tmax && tmax >= 0
}
