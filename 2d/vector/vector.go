package vector

import (
	"github.com/downflux/go-geometry/epsilon"
	"github.com/downflux/go-geometry/nd/vector"
)

type V vector.V

func New(x float64, y float64) *V {
	v := V(*vector.New(x, y))
	return &v
}

func (v V) X() float64 { return v[vector.AXIS_X] }
func (v V) Y() float64 { return v[vector.AXIS_Y] }

func Determinant(v V, u V) float64 {
	return v[vector.AXIS_X]*u[vector.AXIS_Y] - v[vector.AXIS_Y]*u[vector.AXIS_X]
}

// Rotate rotates the vector counterclockwise by the input angle.
func Rotate(theta float64, v V) V {
	b := M(make([]float64, len(v)))
	b.Copy(v)
	b.Rotate(theta)
	return b.V()
}

func Add(v V, u V) V         { return V(vector.Add(vector.V(v), vector.V(u))) }
func Sub(v V, u V) V         { return V(vector.Sub(vector.V(v), vector.V(u))) }
func Dot(v V, u V) float64   { return vector.Dot(vector.V(v), vector.V(u)) }
func Scale(c float64, v V) V { return V(vector.Scale(c, vector.V(v))) }
func WithinEpsilon(v V, u V, e epsilon.E) bool {
	return vector.WithinEpsilon(vector.V(v), vector.V(u), e)
}
func Within(v V, u V) bool         { return vector.Within(vector.V(v), vector.V(u)) }
func SquaredMagnitude(v V) float64 { return vector.SquaredMagnitude(vector.V(v)) }
func Magnitude(v V) float64        { return vector.Magnitude(vector.V(v)) }
func Unit(v V) V                   { return V(vector.Unit(vector.V(v))) }
