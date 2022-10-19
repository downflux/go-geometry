package vector

import (
	"math"

	"github.com/downflux/go-geometry/nd/vector"
)

type M vector.M

func (v M) Copy(u V)        { vector.M(v).Copy(vector.V(u)) }
func (v M) Zero()           { vector.M(v).Zero() }
func (v M) V() V            { return V(v) }
func (v M) X() float64      { return v[vector.AXIS_X] }
func (v M) SetX(c float64)  { v[vector.AXIS_X] = c }
func (v M) Y() float64      { return v[vector.AXIS_Y] }
func (v M) SetY(c float64)  { v[vector.AXIS_Y] = c }
func (v M) Add(u V)         { vector.M(v).Add(vector.V(u)) }
func (v M) Sub(u V)         { vector.M(v).Sub(vector.V(u)) }
func (v M) Scale(c float64) { vector.M(v).Scale(c) }
func (v M) Unit()           { vector.M(v).Unit() }

func (v M) Rotate(theta float64) {
	x := v[vector.AXIS_X]
	y := v[vector.AXIS_Y]
	v[vector.AXIS_X] = x*math.Cos(theta) - y*math.Sin(theta)
	v[vector.AXIS_Y] = x*math.Sin(theta) + y*math.Cos(theta)
}
