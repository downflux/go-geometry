package hyperplane

import (
	"testing"

	"github.com/downflux/go-geometry/nd/vector"
)

func TestIn(t *testing.T) {
	type check struct {
		v    vector.V
		want bool
	}
	testConfigs := []struct {
		name  string
		hp    HP
		tests []check
	}{
		{
			name: "Vertical",
			hp: *New(
				*vector.New(0, 0),
				*vector.New(1, 0),
			),
			tests: []check{
				{v: *vector.New(0, 0), want: true},
				{v: *vector.New(1, 0), want: true},
				{v: *vector.New(-1, 0), want: false},
			},
		},
		{
			name: "Horizontal",
			hp: *New(
				*vector.New(0, 0),
				*vector.New(0, 1),
			),
			tests: []check{
				{v: *vector.New(0, 0), want: true},
				{v: *vector.New(0, 1), want: true},
				{v: *vector.New(0, -1), want: false},
			},
		},
		{
			name: "Sloped",
			hp: *New(
				*vector.New(0, 0),
				*vector.New(1, 1),
			),
			tests: []check{
				{v: *vector.New(0, 0), want: true},
				{v: *vector.New(1, 1), want: true},
				{v: *vector.New(-1, -1), want: false},
			},
		},
		{
			name: "SlopedOffset",
			hp: *New(
				*vector.New(0, 1),
				*vector.New(1, 1),
			),
			tests: []check{
				{v: *vector.New(0, 0), want: false},
				{v: *vector.New(1, 1), want: true},
				{v: *vector.New(2, 2), want: true},
			},
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			for _, test := range c.tests {
				if got := c.hp.In(test.v); got != test.want {
					t.Errorf("In(%v) = %v, want = %v", test.v, got, test.want)
				}
			}
		})
	}
}

func TestDisjoint(t *testing.T) {
	testConfigs := []struct {
		name string
		a    HP
		b    HP
		want bool
	}{
		{
			name: "Parallel",
			a: *New(
				*vector.New(0, 1),
				*vector.New(0, 1),
			),
			b: *New(
				*vector.New(0, -1),
				*vector.New(0, 1),
			),
			want: false,
		},
		{
			name: "Parallel/OrderInvariance",
			a: *New(
				*vector.New(0, -1),
				*vector.New(0, 1),
			),
			b: *New(
				*vector.New(0, 1),
				*vector.New(0, 1),
			),
			want: false,
		},

		{
			name: "AntiParallel/Facing",
			a: *New(
				*vector.New(0, 1),
				*vector.New(0, -1),
			),
			b: *New(
				*vector.New(0, -1),
				*vector.New(0, 1),
			),
			want: false,
		},
		{
			name: "AntiParallel/Facing/OrderInvariance",
			a: *New(
				*vector.New(0, -1),
				*vector.New(0, 1),
			),
			b: *New(
				*vector.New(0, 1),
				*vector.New(0, -1),
			),
			want: false,
		},

		{
			name: "AntiParallel/Away/OrderInvariance",
			a: *New(
				*vector.New(0, 1),
				*vector.New(0, 1),
			),
			b: *New(
				*vector.New(0, -1),
				*vector.New(0, -1),
			),
			want: true,
		},
		{
			name: "AntiParallel/Away/OrderInvariance",
			a: *New(
				*vector.New(0, -1),
				*vector.New(0, -1),
			),
			b: *New(
				*vector.New(0, 1),
				*vector.New(0, 1),
			),
			want: true,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Disjoint(c.a, c.b); got != c.want {
				t.Errorf("Disjoint() = %v, want = %v", got, c.want)
			}
		})
	}
}
