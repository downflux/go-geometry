package constraint

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/downflux/go-geometry/nd/hyperplane"
	"github.com/downflux/go-geometry/nd/vector"
)

// rn returns a random int between [-100, 100).
func rn() float64 { return rand.Float64()*200 - 100 }

// ra returns a vector with randomized coordinates. The constructed vector must
// have a non-zero length.
func rv() vector.V {
	var v vector.V
	for {
		v = *vector.New(rn(), rn())
		if !vector.Within(v, *vector.New(0, 0)) {
			break
		}
	}
	return v
}

func TestIn(t *testing.T) {
	testConfigs := []struct {
		name string
		p    vector.V
		n    vector.V
		v    vector.V
		want bool
	}{
		{
			name: "PositiveX/Edge",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 0),
			v:    *vector.New(0, 0),
			want: true,
		},
		{
			name: "PositiveX/Feasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 0),
			v:    *vector.New(1, 0),
			want: true,
		},
		{
			name: "PositiveX/Infeasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 0),
			v:    *vector.New(-1, 0),
			want: false,
		},

		{
			name: "PositiveY/Edge",
			p:    *vector.New(0, 0),
			n:    *vector.New(0, 1),
			v:    *vector.New(0, 0),
			want: true,
		},
		{
			name: "PositiveY/Feasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(0, 1),
			v:    *vector.New(0, 1),
			want: true,
		},
		{
			name: "PositiveY/Infeasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(0, 1),
			v:    *vector.New(0, -1),
			want: false,
		},

		{
			name: "Sloped/Edge",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 1),
			v:    *vector.New(0, 0),
			want: true,
		},
		{
			name: "Sloped/Feasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 1),
			v:    *vector.New(1, 1),
			want: true,
		},
		{
			name: "Sloped/Infeasible",
			p:    *vector.New(0, 0),
			n:    *vector.New(1, 1),
			v:    *vector.New(-1, -1),
			want: false,
		},

		{
			name: "SlopedOffset/Edge",
			p:    *vector.New(0, 1),
			n:    *vector.New(1, 1),
			v:    *vector.New(0, 1),
			want: true,
		},
		{
			name: "SlopedOffset/Feasible",
			p:    *vector.New(0, 1),
			n:    *vector.New(1, 1),
			v:    *vector.New(2, 2),
			want: true,
		},
		{
			name: "SlopedOffset/Infeasible",
			p:    *vector.New(0, 1),
			n:    *vector.New(1, 1),
			v:    *vector.New(0, 0),
			want: false,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			p := *New(c.p, c.n)
			if got := p.In(c.v); got != c.want {
				t.Errorf("In() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestConformance(t *testing.T) {
	const nTests = 100

	type config struct {
		name string
		p    vector.V
		n    vector.V
		v    vector.V
	}
	var testConfigs []config
	for i := 0; i < nTests; i++ {
		p := rv()
		n := rv()

		for j := 0; j < nTests; j++ {
			testConfigs = append(
				testConfigs,
				config{
					name: fmt.Sprintf("Random-%v/%v", i, j),
					p:    p,
					n:    n,
					v:    rv(),
				},
			)
		}
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			p := *New(c.p, c.n)
			want := hyperplane.HP(p).In(c.v)

			if got := p.In(c.v); got != want {
				t.Errorf("In() = %v, want = %v", got, want)
			}
		})
	}
}
