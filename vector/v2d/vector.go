package vector

import (
	"math"

	"github.com/downflux/go-geometry/vector"
)

type V vector.V

func New(x float64, y float64) *V {
	v := V(*vector.New(x, y))
	return &v
}

func (v V) X() float64 { return vector.V(v).X(vector.AXIS_X) }
func (v V) Y() float64 { return vector.V(v).X(vector.AXIS_Y) }

func Determinant(v V, u V) float64 { return v.X()*u.Y() - v.Y()*u.X() }

// Rotate rotates the vector counterclockwise by the input angle.
func Rotate(theta float64, v V) V {
	return V(
		*vector.New(
			v.X()*math.Cos(theta)-v.Y()*math.Sin(theta),
			v.X()*math.Sin(theta)+v.Y()*math.Cos(theta),
		),
	)
}

func Add(v V, u V) V               { return V(vector.Add(vector.V(v), vector.V(u))) }
func Sub(v V, u V) V               { return V(vector.Sub(vector.V(v), vector.V(u))) }
func Dot(v V, u V) float64         { return vector.Dot(vector.V(v), vector.V(u)) }
func Scale(c float64, v V) V       { return V(vector.Scale(c, vector.V(v))) }
func Within(v V, u V) bool         { return vector.Within(vector.V(v), vector.V(u)) }
func SquaredMagnitude(v V) float64 { return vector.SquaredMagnitude(vector.V(v)) }
func Magnitude(v V) float64        { return vector.Magnitude(vector.V(v)) }
