package line

import (
	"testing"

	"github.com/downflux/go-geometry/circle"
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/vector/v2d"
)

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
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got := c.l.T(c.v); !epsilon.Within(got, c.want) {
				t.Errorf("T() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestIntersection(t *testing.T) {
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
	}

	for _, c := range testConfigs {
		t.Run(c.name, func(t *testing.T) {
			if got, ok := c.l.Intersect(c.m); ok != c.success || !vector.Within(got, c.want) {
				t.Errorf("Intersect() = %v, %v, want = %v, %v", got, ok, c.want, c.success)
			}
		})
	}
}

func TestIntersectionCircle(t *testing.T) {
	testConfigs := []struct {
		name    string
		l       L
		c       circle.C
		success bool
		want    []vector.V
	}{
		{
			name:    "SimpleOriginIntersection",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *circle.New(*vector.New(0, 0), 1),
			success: true,
			want: []vector.V{
				*vector.New(-1, 0),
				*vector.New(1, 0),
			},
		},
		{
			name:    "SimpleNoIntersection",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *circle.New(*vector.New(0, 2), 1),
			success: false,
			want: []vector.V{
				*vector.New(0, 0),
				*vector.New(0, 0),
			},
		},
		{
			name:    "SimpleTangent",
			l:       *New(*vector.New(0, 0), *vector.New(1, 0)),
			c:       *circle.New(*vector.New(0, 1), 1),
			success: true,
			want: []vector.V{
				*vector.New(0, 0),
				*vector.New(0, 0),
			},
		},
		{
			name:    "OffCenterLineIntersect",
			l:       *New(*vector.New(0, 1), *vector.New(1, 0)),
			c:       *circle.New(*vector.New(0, 0), 1),
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
