package vector

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	v := *New(1, 2)
	u := *New(0, 0)
	want := *New(1, 2)

	if got := Add(v, u); !Within(got, want) {
		t.Errorf("Add() = %v, want = %v", got, want)
	}
}

func TestSub(t *testing.T) {
	v := *New(1, 2)
	u := *New(1, 2)
	want := *New(0, 0)

	if got := Sub(v, u); !Within(got, want) {
		t.Errorf("Sub() = %v, want = %v", got, want)
	}
}

func TestScale(t *testing.T) {
	c := 2.0
	v := *New(1, 2)
	want := *New(2, 4)

	if got := Scale(c, v); !Within(got, want) {
		t.Errorf("Scale() = %v, want = %v", got, want)
	}
}

func TestDot(t *testing.T) {
	v := *New(1, 2)
	u := *New(2, 3)
	want := 8.0

	if got := Dot(v, u); got != want {
		t.Errorf("Dot() = %v, want = %v", got, want)
	}
}

func TestUnit(t *testing.T) {
	v := *New(5, 0)
	want := *New(1, 0)

	if got := Unit(v); !Within(got, want) {
		t.Errorf("Unit() = %v, want = %v", got, want)
	}
}

func TestIsOrthogonal(t *testing.T) {
	v := *New(1, 1)
	u := *New(-1, 1)
	want := true

	if got := IsOrthogonal(v, u); got != want {
		t.Errorf("IsOrthogonal() = %v, want = %v", got, want)
	}
}

func TestWithin(t *testing.T) {
	testConfigs := []struct {
		name string
		v    V
		u    V
		want bool
	}{
		{
			name: "Simple/Equal",
			v:    *New(1, 2),
			u:    *New(1, 2),
			want: true,
		},
		{
			name: "Simple/NotEqual",
			v:    *New(1, 2),
			u:    *New(1, 3),
			want: false,
		},
		{
			name: "Infinity/Equal",
			v:    *New(math.Inf(-1), 2),
			u:    *New(math.Inf(-1), 2),
			want: true,
		},
		{
			name: "Simple/NotEqual",
			v:    *New(math.Inf(-1), 2),
			u:    *New(1, 2),
			want: false,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := Within(c.v, c.u); got != c.want {
				t.Errorf("Within() = %v, want = %v", got, c.want)
			}
		})
	}
}
