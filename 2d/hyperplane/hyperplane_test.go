package hyperplane

import (
	"testing"

	"github.com/downflux/go-geometry/2d/vector"
)

// Ensure generated lines are aligned in the expected direction -- that is, D is
// N rotated π / 2 counter-clockwise.
func TestLine(t *testing.T) {
	testConfigs := []struct {
		name string
		p    vector.V
		n    vector.V
		t    float64
		want vector.V
	}{
		{
			name: "VerticalNormal/Positive",
			p:    *vector.New(0, 0),
			n:    *vector.New(0, 1),
			t:    1,
			want: *vector.New(-1, 0),
		},
		{
			name: "VerticalNormal/Negative",
			p:    *vector.New(0, 0),
			n:    *vector.New(0, -1),
			t:    1,
			want: *vector.New(1, 0),
		},
		{
			name: "HorizontalNormal/Positive",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 0),
			t:    1,
			want: *vector.New(0, 1),
		},
		{
			name: "HorizontalNormal/Negative",
			p:    *vector.New(0, 0),
			n:    *vector.New(-1, 0),
			t:    1,
			want: *vector.New(0, -1),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			hp := *New(c.p, c.n)
			l := Line(hp)

			if got := l.L(c.t); !vector.Within(got, c.want) {
				t.Errorf("L() = %v, want = %v", got, c.want)
			}
		})
	}
}
