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
type V []float64

func New(xs ...float64) *V {
	v := V(xs)
	return &v
}

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

func Unit(v V) V       { return Scale(1/Magnitude(v), v) }
func UnitBuf(v V, b V) { ScaleBuf(1/Magnitude(v), v, b) }

func Dot(v V, u V) float64 {
	r := 0.0
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}
	for i := D(0); i < v.Dimension(); i++ {
		r += v[i] * u[i]
	}
	return r
}

func Add(v V, u V) V {
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}

	b := V(make([]float64, v.Dimension()))
	AddBuf(v, u, b)
	return b
}

func AddBuf(v V, u V, b V) {
	if (v.Dimension() != u.Dimension()) || (v.Dimension() != b.Dimension()) {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		b[i] = v[i] + u[i]
	}
}

func Sub(v V, u V) V {
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}

	b := V(make([]float64, v.Dimension()))
	SubBuf(v, u, b)
	return b
}

func SubBuf(v V, u V, b V) {
	if (v.Dimension() != u.Dimension()) || (v.Dimension() != b.Dimension()) {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		b[i] = v[i] - u[i]
	}
}

func Scale(c float64, v V) V {
	b := V(make([]float64, v.Dimension()))
	ScaleBuf(c, v, b)
	return b
}

func ScaleBuf(c float64, v V, b V) {
	if v.Dimension() != b.Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		b[i] = c * v[i]
	}
}

func Within(v V, u V) bool {
	for i := D(0); i < v.Dimension(); i++ {
		if !epsilon.Within(u[i], v[i]) {
			return false
		}
	}
	return true
}

func IsOrthogonal(v V, u V) bool { return Dot(v, u) == 0 }
