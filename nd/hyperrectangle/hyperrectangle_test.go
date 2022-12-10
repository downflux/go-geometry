package hyperrectangle

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

const (
	dimension = 100
	min       = -1e10
	max       = 1e10
)

func rn(min float64, max float64) float64 { return rand.Float64()*(max-min) + min }
func rv(min float64, max float64, d vector.D) vector.V {
	v := vector.V(make([]float64, d))
	for i := vector.D(0); i < d; i++ {
		v[i] = rn(min, max)
	}
	return v
}
func rh(min float64, max float64, d vector.D) R {
	rmin := rv(min, max, d)
	rmax := rv(min, max, d)

	for i := vector.D(0); i < d; i++ {
		if rmin[i] > rmax[i] {
			rmin[i], rmax[i] = rmax[i], rmin[i]
		}
	}
	return *New(rmin, rmax)
}

func BenchmarkContains(b *testing.B) {
	r, s := rh(min, max, 100), rh(min, max, 100)
	for i := 0; i < b.N; i++ {
		Contains(r, s)
	}
}

func BenchmarkDisjoint(b *testing.B) {
	for _, k := range []vector.D{1, 2, 3, dimension} {
		b.Run(fmt.Sprintf("K=%v", k), func(b *testing.B) {
			r, s := rh(min, max, k), rh(min, max, k)
			for i := 0; i < b.N; i++ {
				Disjoint(r, s)
			}
		})
	}
}

func BenchmarkIn(b *testing.B) {
	r, v := rh(min, max, dimension), rv(min, max, dimension)
	for i := 0; i < b.N; i++ {
		r.In(v)
	}
}

func BenchmarkV(b *testing.B) {
	r := rh(min, max, dimension)
	for i := 0; i < b.N; i++ {
		V(r)
	}
}

func BenchmarkSA(b *testing.B) {
	for _, k := range []vector.D{2, 3, dimension} {
		b.Run(fmt.Sprintf("K=%v", k), func(b *testing.B) {
			r := rh(min, max, k)
			for i := 0; i < b.N; i++ {
				SA(r)
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	b.Run("Unbuffered", func(b *testing.B) {
		r, s := rh(min, max, dimension), rh(min, max, dimension)
		for i := 0; i < b.N; i++ {
			Intersect(r, s)
		}
	})
	b.Run("Buffered", func(b *testing.B) {
		r, s := rh(min, max, dimension), rh(min, max, dimension)
		for i := 0; i < b.N; i++ {
			r.M().Intersect(s)
		}
	})
}

func BenchmarkUnion(b *testing.B) {
	b.Run("Unbuffered", func(b *testing.B) {
		r, s := rh(min, max, dimension), rh(min, max, dimension)
		for i := 0; i < b.N; i++ {
			Union(r, s)
		}
	})
	b.Run("Buffered", func(b *testing.B) {
		r, s := rh(min, max, dimension), rh(min, max, dimension)
		for i := 0; i < b.N; i++ {
			r.M().Union(s)
		}
	})
}

func TestScale(t *testing.T) {
	type config struct {
		name string
		r    R
		c    float64
		want R
	}

	configs := []config{
		{
			name: "Trivial",
			r: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
			c: 1,
			want: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
		},
		{
			name: "Shrink",
			r: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
			c: 0.5,
			want: *New(
				[]float64{5, 5},
				[]float64{10, 10},
			),
		},
		{
			name: "Shrink/Volume",
			r: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
			c: math.Pow(0.25, 0.5),
			want: *New(
				[]float64{5, 5},
				[]float64{10, 10},
			),
		},
		{
			name: "Expand",
			r: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
			c: 2,
			want: *New(
				[]float64{5, 5},
				[]float64{25, 25},
			),
		},
		{
			name: "Expand/Volume",
			r: *New(
				[]float64{5, 5},
				[]float64{15, 15},
			),
			c: math.Pow(4.0, 0.5),
			want: *New(
				[]float64{5, 5},
				[]float64{25, 25},
			),
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			if got := Scale(c.r, c.c); !Within(got, c.want) {
				t.Errorf("Scale() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestDisjoint(t *testing.T) {
	type config struct {
		name string
		r    R
		s    R
		want bool
	}

	configs := []config{
		{
			name: "Simple/Disjoint",
			r: *New(
				[]float64{0},
				[]float64{10},
			),
			s: *New(
				[]float64{11},
				[]float64{20},
			),
			want: true,
		},
		{
			name: "Simple/Overlap",
			r: *New(
				[]float64{0},
				[]float64{10},
			),
			s: *New(
				[]float64{9},
				[]float64{20},
			),
			want: false,
		},
		{
			name: "Simple/Disjoint/Commutative",
			r: *New(
				[]float64{11},
				[]float64{20},
			),
			s: *New(
				[]float64{0},
				[]float64{10},
			),
			want: true,
		},
		{
			name: "2D/Disjoint",
			r: *New(
				[]float64{0, 0},
				[]float64{10, 10},
			),
			s: *New(
				[]float64{5, 11},
				[]float64{20, 20},
			),
			want: true,
		},
		{
			name: "2D/Overlap",
			r: *New(
				[]float64{0, 0},
				[]float64{10, 10},
			),
			s: *New(
				[]float64{5, 5},
				[]float64{20, 20},
			),
			want: false,
		},
	}

	for _, c := range configs {
		if got := Disjoint(c.r, c.s); got != c.want {
			t.Errorf("Disjoint() = %v, want = %v", got, c.want)
		}
	}
}

func TestUnion(t *testing.T) {
	testConfigs := []struct {
		name string
		r    R
		s    R
		want R
	}{
		{
			name: "NoExpansion",
			r: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
			s: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
			want: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Union(c.r, c.s); !Within(c.want, got) {
				t.Errorf("Union() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	testConfigs := []struct {
		name        string
		r           R
		s           R
		want        R
		wantSuccess bool
	}{
		{
			name: "Trivial",
			r: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
			s: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
			want: *New(
				*vector.New(1, 2),
				*vector.New(2, 3),
			),
			wantSuccess: true,
		},
		{
			name: "Enveloped",
			r: *New(
				*vector.New(1, 2),
				*vector.New(100, 200),
			),
			s: *New(
				*vector.New(2, 3),
				*vector.New(99, 199),
			),
			want: *New(
				*vector.New(2, 3),
				*vector.New(99, 199),
			),
			wantSuccess: true,
		},
		{
			name: "Overlap/Left",
			r: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			s: *New(
				*vector.New(-1, 0),
				*vector.New(1, 5),
			),
			want: *New(
				*vector.New(0, 0),
				*vector.New(1, 5),
			),
			wantSuccess: true,
		},
		{
			name: "Overlap/Right",
			r: *New(
				*vector.New(-1, 0),
				*vector.New(1, 5),
			),
			s: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			want: *New(
				*vector.New(0, 0),
				*vector.New(1, 5),
			),
			wantSuccess: true,
		},
		{
			name: "Overlap/Top",
			r: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			s: *New(
				*vector.New(0, 4),
				*vector.New(5, 6),
			),
			want: *New(
				*vector.New(0, 4),
				*vector.New(5, 5),
			),
			wantSuccess: true,
		},
		{
			name: "Overlap/Bottom",
			r: *New(
				*vector.New(0, 4),
				*vector.New(5, 6),
			),
			s: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			want: *New(
				*vector.New(0, 4),
				*vector.New(5, 5),
			),
			wantSuccess: true,
		},
		{
			name: "NoOverlap/Left",
			r: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			s: *New(
				*vector.New(-5, 0),
				*vector.New(-1, 5),
			),
			wantSuccess: false,
		},
		{
			name: "NoOverlap/Top",
			r: *New(
				*vector.New(0, 0),
				*vector.New(5, 5),
			),
			s: *New(
				*vector.New(0, 6),
				*vector.New(5, 9),
			),
			wantSuccess: false,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			got, ok := Intersect(c.r, c.s)
			if ok != c.wantSuccess {
				t.Errorf("Intersect() = _, %v, want = _, %v", ok, c.wantSuccess)
			}
			if !Within(got, c.want) {
				t.Errorf("Intersect() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestV(t *testing.T) {
	testConfigs := []struct {
		name string
		r    R
		want float64
	}{
		{
			name: "1D",
			r: *New(
				*vector.New(1),
				*vector.New(100),
			),
			want: 99,
		},
		{
			name: "2D",
			r: *New(
				*vector.New(1, 1),
				*vector.New(100, 100),
			),
			want: 99 * 99,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := V(c.r); !epsilon.Within(c.want, got) {
				t.Errorf("V() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestSA(t *testing.T) {
	testConfigs := []struct {
		name string
		r    R
		want float64
	}{
		{
			name: "1D",
			r: *New(
				*vector.New(1),
				*vector.New(100),
			),
			want: 0,
		},
		{
			name: "2D",
			r: *New(
				*vector.New(0, 0),
				*vector.New(99, 101),
			),
			want: 400,
		},
		{
			name: "3D",
			r: *New(
				*vector.New(0, 0, 0),
				*vector.New(9, 10, 11),
			),
			want: 598,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := SA(c.r); !epsilon.Within(c.want, got) {
				t.Errorf("SA() = %v, want = %v", got, c.want)
			}
		})
	}
}
