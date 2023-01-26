package hyperrectangle

import (
	"github.com/downflux/go-geometry/2d/vector"
	"github.com/downflux/go-geometry/nd/hyperrectangle"
)

type M hyperrectangle.M

func (r M) Copy(s R)           { hyperrectangle.M(r).Copy(hyperrectangle.R(s)) }
func (r M) Zero()              { hyperrectangle.M(r).Zero() }
func (r M) R() R               { return R(r) }
func (r M) Min() vector.M      { return vector.M(hyperrectangle.R(r).Min()) }
func (r M) Max() vector.M      { return vector.M(hyperrectangle.R(r).Max()) }
func (r M) Intersect(s R) bool { return hyperrectangle.M(r).Intersect(hyperrectangle.R(s)) }
func (r M) Union(s R)          { hyperrectangle.M(r).Union(hyperrectangle.R(s)) }
func (r M) Scale(c float64)    { hyperrectangle.M(r).Scale(c) }
