package vector

import (
	"math"
	"testing"
)

func TestDeterminant(t *testing.T) {
	v := *New(1, 1)
	u := *New(2, 3)
	want := float64(1)
	if got := Determinant(v, u); got != want {
		t.Errorf("Determinant() = %v, want = %v", got, want)
	}
}

func TestRotate(t *testing.T) {
	const tolerance = 1e-10
	testConfigs := []struct {
		name  string
		theta float64
		v     V
		want  V
	}{
		{name: "0Degree", theta: 0, v: *New(1, 0), want: *New(1, 0)},
		{name: "90Degree", theta: .5 * math.Pi, v: *New(1, 0), want: *New(0, 1)},
		{name: "180Degree", theta: math.Pi, v: *New(1, 0), want: *New(-1, 0)},
		{name: "270Degree", theta: 1.5 * math.Pi, v: *New(1, 0), want: *New(0, -1)},
		{name: "360Degree", theta: 2 * math.Pi, v: *New(1, 0), want: *New(1, 0)},
		{name: "InverseRotate", theta: .1, v: Rotate(-.1, *New(1, 0)), want: *New(1, 0)},
		{
			name:  "FlipYCoordinate",
			theta: .2,
			v:     Rotate(-.1, *New(1, 0)),
			want: *New(
				Rotate(-.1, *New(1, 0)).X(),
				-Rotate(-.1, *New(1, 0)).Y(),
			),
		},
		{
			name:  "FlipXCoordinate",
			theta: math.Pi + .2,
			v:     Rotate(-.1, *New(1, 0)),
			want: *New(
				-Rotate(-.1, *New(1, 0)).X(),
				Rotate(-.1, *New(1, 0)).Y(),
			),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Rotate(c.theta, c.v); !Within(got, c.want) {
				t.Errorf("Rotate() = %v, want = %v", got, c.want)
			}
		})
	}
}
