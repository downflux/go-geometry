package vector

import (
	"fmt"
)

// M is a mutable n-dimensional vector.
type M []float64

func (v M) Dimension() D { return D(len(v)) }
func (v M) V() V         { return V(v) }

func (v M) X(i D) float64 {
	if i >= v.Dimension() {
		panic(fmt.Sprintf("cannot access %v-dimensional data in a %v dimensional vector", i+1, v.Dimension()))
	}
	return v[i]
}

func (v M) SetX(i D, c float64) {
	if i >= v.Dimension() {
		panic(fmt.Sprintf("cannot access %v-dimensional data in a %v dimensional vector", i+1, v.Dimension()))
	}
	v[i] = c
}

func (v M) Zero() {
	for i := D(0); i < v.Dimension(); i++ {
		v[i] = 0
	}
}

func (v M) Copy(u V) {
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		v[i] = u[i]
	}
}

func (v M) Add(u V) {
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		v[i] += u[i]
	}
}

func (v M) Sub(u V) {
	if v.Dimension() != u.Dimension() {
		panic("mismatching vector dimensions")
	}

	for i := D(0); i < v.Dimension(); i++ {
		v[i] -= u[i]
	}
}

func (v M) Scale(c float64) {
	for i := D(0); i < v.Dimension(); i++ {
		v[i] *= c
	}
}

func (v M) Unit() { v.Scale(1 / Magnitude(v.V())) }
