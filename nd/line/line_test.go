package line

import (
	"testing"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

func TestL(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		t    float64
		want vector.V
	}{
		{
			name: "SimpleHorizontal",
			l:    L{p: *vector.New(0, 1), d: *vector.New(1, 0)},
			t:    1,
			want: *vector.New(1, 1),
		},
		{
			name: "SimpleVertical",
			l:    L{p: *vector.New(1, 0), d: *vector.New(0, 1)},
			t:    1,
			want: *vector.New(1, 1),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.L(c.t); !vector.Within(got, c.want) {
				t.Errorf("T() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestT(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		v    vector.V
		want float64
	}{
		{
			name: "OnLine/SimpleHorizontal",
			l:    L{p: *vector.New(0, 1), d: *vector.New(1, 0)},
			v:    *vector.New(1, 1),
			want: 1,
		},
		{
			name: "OnLine/SimpleVertical",
			l:    L{p: *vector.New(1, 0), d: *vector.New(0, 1)},
			v:    *vector.New(1, 1),
			want: 1,
		},

		{
			name: "OffLine/SimpleHorizontal",
			l:    L{p: *vector.New(0, 1), d: *vector.New(1, 0)},
			v:    *vector.New(0, 0),
			want: 0,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.T(c.v); !epsilon.Within(got, c.want) {
				t.Errorf("Project() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestParallel(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		m    L
		want bool
	}{
		{
			name: "Parallel",
			l: *New(
				*vector.New(0, 1, 1),
				*vector.New(1, 0, 0),
			),
			m: *New(
				*vector.New(0, 0, 1),
				*vector.New(5, 0, 0),
			),
			want: true,
		},

		{
			name: "AntiParallel",
			l: *New(
				*vector.New(0, 1, 1),
				*vector.New(1, 0, 0),
			),
			m: *New(
				*vector.New(0, 0, 1),
				*vector.New(-5, 0, 0),
			),
			want: false,
		},

		{
			name: "Intersecting",
			l: *New(
				*vector.New(0, 0, 0),
				*vector.New(1, 0, 0),
			),
			m: *New(
				*vector.New(0, 0, 0),
				*vector.New(0, 1, 0),
			),
			want: false,
		},

		{
			name: "Skew",
			l: *New(
				*vector.New(0, 0, 0),
				*vector.New(1, 0, 0),
			),
			m: *New(
				*vector.New(0, 0, 1),
				*vector.New(0, 1, 0),
			),
			want: false,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.Parallel(c.m); got != c.want {
				t.Errorf("Parallel() = %v, want = %v", got, c.want)
			}
		})
	}
}
