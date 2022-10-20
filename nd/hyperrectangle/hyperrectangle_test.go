package hyperrectangle

import (
	"testing"

	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
	"github.com/google/go-cmp/cmp"
)

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
			got := Union(c.r, c.s)
			if diff := cmp.Diff(
				c.want,
				got,
				cmp.AllowUnexported(
					R{},
				),
			); diff != "" {
				t.Errorf("Union() mismatch (-want +got):\n%v", diff)
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
			if diff := cmp.Diff(
				c.want,
				got,
				cmp.AllowUnexported(
					R{},
				),
			); diff != "" {
				t.Errorf("Intersect() mismatch (-want +got):\n%v", diff)
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
