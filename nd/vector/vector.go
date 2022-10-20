// Package vector defines an n-dimensional vector. The mutable / immutable
// syntax is based loosely off of the github.com/kvartborg/vector implementation.
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

// V is an immutable n-length vector.
type V []float64

func New(xs ...float64) *V {
	v := V(xs)
	return &v
}

func (v V) M() M { return M(v) }

// Dimension returns the dimension of the vector.
func (v V) Dimension() D { return D(len(v)) }

func (v V) X(i D) float64 {
	if i >= v.Dimension() {
		panic(fmt.Sprintf("cannot access %v-dimensional data in a %v dimensional vector", i+1, v.Dimension()))
	}
	return v[i]
}

func SquaredMagnitude(v V) float64 { return Dot(v, v) }
func Magnitude(v V) float64        { return math.Sqrt(SquaredMagnitude(v)) }

func Unit(v V) V { return Scale(1/Magnitude(v), v) }

func Dot(v V, u V) float64 {
	r := 0.0
	for i := D(0); i < v.Dimension(); i++ {
		r += v[i] * u[i]
	}
	return r
}

func Add(v V, u V) V {
	b := M(make([]float64, v.Dimension()))
	b.Copy(v)
	b.Add(u)
	return b.V()
}

func Sub(v V, u V) V {
	b := M(make([]float64, v.Dimension()))
	b.Copy(v)
	b.Sub(u)
	return b.V()
}

func Scale(c float64, v V) V {
	b := M(make([]float64, v.Dimension()))
	b.Copy(v)
	b.Scale(c)
	return b.V()
}

func WithinEpsilon(v V, u V, e epsilon.E) bool {
	for i := D(0); i < v.Dimension(); i++ {
		if !e.Within(u[i], v[i]) {
			return false
		}
	}
	return true
}

func Within(v V, u V) bool       { return WithinEpsilon(v, u, epsilon.DefaultE) }
func IsOrthogonal(v V, u V) bool { return Dot(v, u) == 0 }
