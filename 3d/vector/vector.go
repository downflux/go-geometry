package vector

import (
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

type V vector.V

func New(x float64, y float64, z float64) *V {
	v := V(*vector.New(x, y, z))
	return &v
}

func (v V) X() float64 { return v[vector.AXIS_X] }
func (v V) Y() float64 { return v[vector.AXIS_Y] }
func (v V) Z() float64 { return v[vector.AXIS_Z] }

func Cross(v V, u V) V {
	b := make([]float64, len(v))
	CrossBuf(v, u, b)
	return b
}

func CrossBuf(v V, u V, b V) {
	b[vector.AXIS_X] = v[vector.AXIS_Y]*u[vector.AXIS_Z] - v[vector.AXIS_Z]*u[vector.AXIS_Y]
	b[vector.AXIS_Y] = v[vector.AXIS_Z]*u[vector.AXIS_X] - v[vector.AXIS_X]*u[vector.AXIS_Z]
	b[vector.AXIS_Z] = v[vector.AXIS_X]*u[vector.AXIS_Y] - v[vector.AXIS_Y]*u[vector.AXIS_X]
}

func Add(v V, u V) V               { return V(vector.Add(vector.V(v), vector.V(u))) }
func AddBuf(v V, u V, b V)         { vector.AddBuf(vector.V(v), vector.V(u), vector.V(b)) }
func Sub(v V, u V) V               { return V(vector.Sub(vector.V(v), vector.V(u))) }
func SubBuf(v V, u V, b V)         { vector.SubBuf(vector.V(v), vector.V(u), vector.V(b)) }
func Dot(v V, u V) float64         { return vector.Dot(vector.V(v), vector.V(u)) }
func Scale(c float64, v V) V       { return V(vector.Scale(c, vector.V(v))) }
func ScaleBuf(c float64, v V, b V) { vector.ScaleBuf(c, vector.V(v), vector.V(b)) }
func WithinEpsilon(v V, u V, e epsilon.E) bool {
	return vector.WithinEpsilon(vector.V(v), vector.V(u), e)
}
func Within(v V, u V) bool         { return vector.Within(vector.V(v), vector.V(u)) }
func SquaredMagnitude(v V) float64 { return vector.SquaredMagnitude(vector.V(v)) }
func Magnitude(v V) float64        { return vector.Magnitude(vector.V(v)) }
func Unit(v V) V                   { return V(vector.Unit(vector.V(v))) }
func UnitBuf(v V, b V)             { vector.UnitBuf(vector.V(v), vector.V(b)) }
