package epsilon

import (
	"math"
	"testing"
)

func TestWithin(t *testing.T) {
	testConfigs := []struct {
		name string
		a    float64
		b    float64
		want bool
	}{
		{
			name: "Infinity/Positive",
			a:    math.Inf(0),
			b:    math.Inf(0),
			want: true,
		},
		{
			name: "Infinity/Negative",
			a:    math.Inf(-1),
			b:    math.Inf(-1),
			want: true,
		},

		{
			name: "Infinity/NotEqual",
			a:    math.Inf(-1),
			b:    math.Inf(0),
			want: false,
		},

		{
			name: "Finite/Equal",
			a:    1.2,
			b:    1.2,
			want: true,
		},
		{
			name: "Finite/Equal/NextAfter",
			a:    1.1,
			b:    math.Nextafter(1.1, 2),
			want: true,
		},
		{
			name: "Finite/Equal/NextAfter/Reversed",
			a:    1.1,
			b:    math.Nextafter(1.1, 0),
			want: true,
		},
		{
			name: "Finite/NotEqual",
			a:    1.2,
			b:    1.1,
			want: false,
		},
		{
			name: "Finite/NotEqual/Reversed",
			a:    1.1,
			b:    1.2,
			want: false,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Within(c.a, c.b); got != c.want {
				t.Errorf("Within() = %v, want = %v", got, c.want)
			}
		})
	}
}
