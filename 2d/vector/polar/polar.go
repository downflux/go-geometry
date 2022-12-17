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

func Polar(v vector.V) V {
	x, y := v[AXIS_R], v[AXIS_THETA]
	return V(*vector.New(math.Sqrt(x*x+y*y), math.Atan2(y, x)))
}

// Normalize returns a vector whose anglular component is bound between 0 and
// 2π.
func Normalize(v V) V {
	theta := math.Mod(v[AXIS_THETA], 2*math.Pi)
	// theta may be negative in the case the original polar coordinate is
	// negative. Since we want to ensure the angle is positive, we have to
	// take this into consideration.
	if theta < 0 {
		theta += 2 * math.Pi
	}
	return *New(v[AXIS_R], theta)
}

func Cartesian(v V) vector.V {
	return *vector.New(
		v[AXIS_R]*math.Cos(v[AXIS_THETA]),
		v[AXIS_R]*math.Sin(v[AXIS_THETA]),
	)
}

func WithinEpsilon(v V, u V, e epsilon.E) bool {
	return (e.Within(v[AXIS_R], 0) && e.Within(u[AXIS_R], 0)) || vector.WithinEpsilon(vector.V(v), vector.V(u), e)
}
func Within(v, u V) bool { return WithinEpsilon(v, u, epsilon.DefaultE) }
