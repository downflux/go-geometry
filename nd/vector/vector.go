package vector

import (
	"fmt"
	"math"

	"github.com/downflux/go-geometry/epsilon"
)

type D int

const (
	// AXIS_X is a common alias for the first dimension.
	AXIS_X D = iota

	// AXIS_Y is a common alias for the second dimension.
	AXIS_Y

	// AXIS_Z is a common alias for the third dimension.
	AXIS_Z

	// AXIS_W is a common alias for the fourth dimension.
	AXIS_W
)

// V is an n-length vector.
type V struct {
	xs []float64
}

func New(xs ...float64) *V { return &V{xs: xs} }

// Dimension returns the dimension of the vector.
func (v V) Dimension() D { return D(len(v.xs)) }

func (v V) X(i D) float64 {
	if i >= v.Dimension() {
		panic(fmt.Sprintf("cannot access %v-dimensional data in a %v dimensional vector", i, v.Dimension()))
	}
	return v.xs[i]
}

func SquaredMagnitude(v V) float64 { return Dot(v, v) }
func Magnitude(v V) float64        { return math.Sqrt(SquaredMagnitude(v)) }
func Unit(v V) V                   { return Scale(1/Magnitude(v), v) }

func Dot(v V, u V) float64 {
	r := 0.0
	for i := D(0); i < max(v.Dimension(), u.Dimension()); i++ {
		r += v.X(i) * u.X(i)
	}
	return r
}

func Add(v V, u V) V {
	d := max(v.Dimension(), u.Dimension())
	xs := make([]float64, int(d))
	for i := D(0); i < d; i++ {
		xs[i] = v.X(i) + u.X(i)
	}
	return V{xs: xs}
}

func Sub(v V, u V) V {
	d := max(v.Dimension(), u.Dimension())
	xs := make([]float64, int(d))
	for i := D(0); i < v.Dimension(); i++ {
		xs[i] = v.X(i) - u.X(i)
	}
	return V{xs: xs}
}

func Scale(c float64, v V) V {
	xs := make([]float64, int(v.Dimension()))
	for i := D(0); i < v.Dimension(); i++ {
		xs[i] = c * v.X(i)
	}
	return V{xs: xs}
}

func Within(v V, u V) bool       { return epsilon.Within(SquaredMagnitude(Sub(u, v)), 0) }
func IsOrthogonal(v V, u V) bool { return Dot(v, u) == 0 }

func max(i D, j D) D {
	if int(i) > int(j) {
		return i
	}
	return j
}
