package polar

import (
	"math"

	"github.com/downflux/go-geometry/2d/vector"
	"github.com/downflux/go-geometry/epsilon"

	vnd "github.com/downflux/go-geometry/nd/vector"
)

type V vector.V

const (
	AXIS_R     = vnd.AXIS_X
	AXIS_THETA = vnd.AXIS_Y
)

func New(r float64, theta float64) *V {
	v := V(*vector.New(r, theta))
	return &v
}

func (v V) M() M       { return M(v) }
func (v V) R() float64 { return vector.V(v).X() }

// Theta returns the angular component of the polar coordinate. Note that theta
// may extend beyond 2π, as polar coordinates may also represent angular
// acceleration and velocity, which are not bound by a single rotation.
func (v V) Theta() float64 { return vector.V(v).Y() }

func Add(v V, u V) V       { return V(vector.Add(vector.V(v), vector.V(u))) }
func Sub(v V, u V) V       { return V(vector.Sub(vector.V(v), vector.V(u))) }
func Dot(v V, u V) float64 { return v[AXIS_R] * u[AXIS_R] * math.Cos(v[AXIS_THETA]-u[AXIS_THETA]) }
func Determinant(v V, u V) float64 {
	return v[AXIS_R] * u[AXIS_R] * math.Sin(v[AXIS_THETA]-u[AXIS_THETA])
}

// Normalize returns a vector whose anglular component is bound between 0 and
// 2π.
func Normalize(v V) V {
	buf := M([]float64{0, 0})
	buf.Copy(v)
	buf.Normalize()
	return buf.V()
}

func Unit(v V) V {
	buf := M([]float64{0, 0})
	buf.Copy(v)
	buf.Unit()
	return buf.V()
}

func Cartesian(v V) vector.V {
	return *vector.New(
		v[AXIS_R]*math.Cos(v[AXIS_THETA]),
		v[AXIS_R]*math.Sin(v[AXIS_THETA]),
	)
}

func Polar(v vector.V) V {
	x, y := v[AXIS_R], v[AXIS_THETA]
	buf := M([]float64{0, 0})
	buf.Copy(V(*vector.New(math.Sqrt(x*x+y*y), math.Atan2(y, x))))
	buf.Normalize()
	return buf.V()
}

func WithinEpsilon(v V, u V, e epsilon.E) bool {
	return (e.Within(v[AXIS_R], 0) && e.Within(u[AXIS_R], 0)) || (e.Within(v[AXIS_R], u[AXIS_R]) && e.Within(math.Mod(v[AXIS_THETA], 2*math.Pi), math.Mod(u[AXIS_THETA], 2*math.Pi)))
}
func Within(v, u V) bool { return WithinEpsilon(v, u, epsilon.DefaultE) }
