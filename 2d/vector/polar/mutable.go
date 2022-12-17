package polar

import (
	"math"

	"github.com/downflux/go-geometry/2d/vector"
)

type M V

func (v M) Copy(u V)           { vector.M(v).Copy(vector.V(u)) }
func (v M) Zero()              { vector.M(v).Zero() }
func (v M) V() V               { return V(v) }
func (v M) R() float64         { return v[AXIS_R] }
func (v M) SetR(c float64)     { v[AXIS_R] = c }
func (v M) Theta() float64     { return v[AXIS_THETA] }
func (v M) SetTheta(c float64) { v[AXIS_THETA] = c }
func (v M) Add(u V)            { vector.M(v).Add(vector.V(u)) }
func (v M) Sub(u V)            { vector.M(v).Sub(vector.V(u)) }

func (v M) Normalize() {
	theta := math.Mod(v[AXIS_THETA], 2*math.Pi)
	// theta may be negative in the case the original polar coordinate is
	// negative. Since we want to ensure the angle is positive, we have to
	// take this into consideration.
	if theta < 0 {
		theta += 2 * math.Pi
	}
	v[AXIS_THETA] = theta
}

func (v M) Unit() {
	v.Normalize()
	v[AXIS_R] = 1
}
