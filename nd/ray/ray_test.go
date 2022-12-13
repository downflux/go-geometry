package ray

import (
	"math/rand"
	"testing"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/hyperrectangle"
	"github.com/downflux/go-geometry/nd/vector"
)

func rn(min, max float64) float64 { return min + rand.Float64()*(max-min) }
func rv(min, max float64, k vector.D) vector.V {
	v := vector.V(make([]float64, k))
	for i := vector.D(0); i < k; i++ {
		v[i] = rn(min, max)
	}
	return v
}
func rr(min, max float64, k vector.D) R {
	p, d := rv(min, max, k), rv(min, max, k)
	if epsilon.Within(vector.Magnitude(d), 0) {
		d[rand.Intn(int(k))] += 1
	}
	return *New(p, d)
}
func rh(min, max float64, k vector.D) hyperrectangle.R {
	vmin, vmax := rv(min, max, k), rv(min, max, k)
	for i := vector.D(0); i < k; i++ {
		if vmin[i] > vmax[i] {
			vmin[i], vmax[i] = vmax[i], vmin[i]
		}
	}
	return *hyperrectangle.New(vmin, vmax)
}

func BenchmarkIntersectHyperrectangle(b *testing.B) {
	r, s := rr(100, 200, 100), rh(-200, 200, 100)
	for i := 0; i < b.N; i++ {
		IntersectHyperrectangle(r, s)
	}
}

func TestIntersectHyperrectangle(t *testing.T) {
	type config struct {
		name string
		r    R
		s    hyperrectangle.R
		want bool
	}

	configs := []config{
		{
			name: "Touch/1D",
			r:    *New(vector.V{0}, vector.V{1}),
			s:    *hyperrectangle.New(vector.V{-1}, vector.V{0}),
			want: true,
		},
		{
			name: "Hit/1D",
			r:    *New(vector.V{0}, vector.V{1}),
			s:    *hyperrectangle.New(vector.V{2}, vector.V{3}),
			want: true,
		},
		{
			name: "Hit/2D",
			r:    *New(vector.V{0, 0}, vector.V{1, 1}),
			s:    *hyperrectangle.New(vector.V{2, 2}, vector.V{3, 3}),
			want: true,
		},
		{
			name: "Hit/1D/Inside",
			r:    *New(vector.V{0}, vector.V{1}),
			s:    *hyperrectangle.New(vector.V{-1}, vector.V{3}),
			want: true,
		},
		{
			name: "Miss/1D",
			r:    *New(vector.V{0}, vector.V{1}),
			s:    *hyperrectangle.New(vector.V{-2}, vector.V{-1}),
			want: false,
		},
		{
			name: "Miss/2D",
			r:    *New(vector.V{0, 0}, vector.V{1, 0}),
			s:    *hyperrectangle.New(vector.V{0, 1}, vector.V{1, 2}),
			want: false,
		},
		{
			name: "Miss/1D/Reverse",
			r:    *New(vector.V{0}, vector.V{-1}),
			s:    *hyperrectangle.New(vector.V{1}, vector.V{2}),
			want: false,
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			if got := IntersectHyperrectangle(c.r, c.s); got != c.want {
				t.Errorf("IntersectHyperrectangle() = %v, want = %v", got, c.want)
			}
		})
	}
}
