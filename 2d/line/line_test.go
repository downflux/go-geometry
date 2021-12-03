package line

import (
	"testing"

	"github.com/downflux/go-geometry/2d/hypersphere"
	"github.com/downflux/go-geometry/2d/vector"
	"github.com/downflux/go-geometry/epsilon"
)

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
				*vector.New(0, 1),
				*vector.New(1, 0),
			),
			m: *New(
				*vector.New(0, 2),
				*vector.New(2, 0),
			),
			want: true,
		},

		{
			name: "AntiParallel",
			l: *New(
				*vector.New(0, 1),
				*vector.New(1, 0),
			),
			m: *New(
				*vector.New(0, 2),
				*vector.New(-2, 0),
			),
			want: false,
		},

		{
			name: "Intersecting",
			l: *New(
				*vector.New(0, 0),
				*vector.New(1, 0),
			),
			m: *New(
				*vector.New(0, 0),
				*vector.New(0, 1),
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

func TestDistance(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		p    vector.V
		want float64
	}{
		{
			name: "Trivial",
			l:    *New(*vector.New(0, 0), *vector.New(0, 1)),
			p:    *vector.New(0, 0),
			want: 0,
		},
		{
			name: "SimpleUnitDirection",
			l:    *New(*vector.New(0, 0), *vector.New(0, 1)),
			p:    *vector.New(1, 1),
			want: 1,
		},
		{
			name: "SimpleLargeDirection",
			l:    *New(*vector.New(0, 0), *vector.New(0, 100)),
			p:    *vector.New(1, 1),
			want: 1,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.Distance(c.p); !epsilon.Within(c.want, got) {
				t.Errorf("Distance() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestL(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		t    float64
		want vector.V
	}{
		{
			name: "SimpleHorizontal",
			l:    *New(*vector.New(0, 1), *vector.New(1, 0)),
			t:    1,
			want: *vector.New(1, 1),
		},
		{
			name: "SimpleVertical",
			l:    *New(*vector.New(1, 0), *vector.New(0, 1)),
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
			l:    *New(*vector.New(0, 1), *vector.New(1, 0)),
			v:    *vector.New(1, 1),
			want: 1,
		},
		{
			name: "OnLine/SimpleVertical",
			l:    *New(*vector.New(1, 0), *vector.New(0, 1)),
			v:    *vector.New(1, 1),
			want: 1,
		},

		{
			name: "OffLine/SimpleHorizontal",
			l:    *New(*vector.New(0, 1), *vector.New(1, 0)),
			v:    *vector.New(0, 0),
			want: 0,
		},

		// Test based on specific potential bug.
		//
		// The intersection between the two contraints
		//
		//   x <= 5 and
		//   x + 4y <= 16
		//
		// Should occur at (5, 2.75). This corresponds to
		//
		//   t = -1.25
		//
		// for the line
		//
		//   L = (0, 4) + t(-4, -1)
		{
			name: "LPConstraint",
			l:    *New(*vector.New(0, 4), *vector.New(-4, 1)),
			v:    *vector.New(5, 2.75),
			want: -1.25,
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.T(c.v); !epsilon.Within(got, c.want) {
				t.Errorf("T() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestN(t *testing.T) {
	testConfigs := []struct {
		name string
		l    L
		want vector.V
	}{
		{
			name: "Vertical/Positive",
			l: *New(
				*vector.New(1, 2),
				*vector.New(0, 1),
			),
			want: *vector.New(1, 0),
		},
		{
			name: "Vertical/Negative",
			l: *New(
				*vector.New(1, 2),
				*vector.New(0, -1),
			),
			want: *vector.New(-1, 0),
		},
		{
			name: "Horizontal/Positive",
			l: *New(
				*vector.New(1, 2),
				*vector.New(1, 0),
			),
			want: *vector.New(0, -1),
		},
		{
			name: "Horizontal/Negative",
			l: *New(
				*vector.New(1, 2),
				*vector.New(-1, 0),
			),
			want: *vector.New(0, 1),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.N(); !vector.Within(got, c.want) {
				t.Errorf("N() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	testConfigs := []struct {
		name    string
		l       L
		m       L
		success bool
		want    vector.V
	}{
		{
			name:    "SimpleOrigin",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			m:       *New(*vector.New(0, 0), *vector.New(0, 1)),
			success: true,
			want:    *vector.New(0, 0),
		},
		{
			name:    "SimpleYIntercept",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			m:       *New(*vector.New(1, 0), *vector.New(0, 1)),
			success: true,
			want:    *vector.New(1, 0),
		},

		// Test based on specific potential bug.
		//
		// The intersection between the two contraints
		//
		//   x <= 5 and
		//   x + 4y <= 16
		//
		// Should occur at (5, 2.75).
		//
		// N.B.: Lines are defined by
		//
		//   L = P + tD
		//
		// notation, not by the Hesse-normal planar form.
		{
			name:    "LPConstraint",
			l:       *New(*vector.New(0, 4), *vector.New(-4, 1)),
			m:       *New(*vector.New(5, 0), *vector.New(0, 5)),
			success: true,
			want:    *vector.New(5, 2.75),
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got, ok := c.l.Intersect(c.m); ok != c.success || !vector.Within(got, c.want) {
				t.Errorf("Intersect() = %v, %v, want = %v, %v", got, ok, c.want, c.success)
			}
		})
	}
}

func TestIntersectCircle(t *testing.T) {
	testConfigs := []struct {
		name    string
		l       L
		c       hypersphere.C
		success bool
		want    []vector.V
	}{
		{
			name:    "SimpleOriginIntersection",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *hypersphere.New(*vector.New(0, 0), 1),
			success: true,
			want: []vector.V{
				*vector.New(-1, 0),
				*vector.New(1, 0),
			},
		},
		{
			name:    "SimpleNoIntersection",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *hypersphere.New(*vector.New(0, 2), 1),
			success: false,
			want: []vector.V{
				vector.V{},
				vector.V{},
			},
		},
		{
			name:    "SimpleTangent",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *hypersphere.New(*vector.New(0, 1), 1),
			success: true,
			want: []vector.V{
				*vector.New(0, 0),
				*vector.New(0, 0),
			},
		},
		{
			name:    "OffCenterLineIntersect",
			l:       *New(*vector.New(0, 1), *vector.New(1, 0)),
			c:       *hypersphere.New(*vector.New(0, 0), 1),
			success: true,
			want: []vector.V{
				*vector.New(0, 1),
				*vector.New(0, 1),
			},
		},
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			vl, vr, ok := c.l.IntersectCircle(c.c)
			if ok != c.success || !vector.Within(vl, c.want[0]) || !vector.Within(vr, c.want[1]) {
				t.Fatalf("IntersectCircle() = %v, %v, %v, want = %v, %v, %v", vl, vr, ok, c.want[0], c.want[1], c.success)
			}
		})
	}
}
