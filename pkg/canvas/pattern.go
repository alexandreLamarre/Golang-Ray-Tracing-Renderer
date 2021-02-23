package canvas

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
)

//Pattern represents a pattern of colors
type Pattern struct{
	a *Color
	b *Color
	GetPattern func(vector *algebra.Vector) *Color
	Transform *algebra.Matrix
}

//Stripe Pattern Creates a new Stripe Pattern in the x coordinate with default identity transformation
func StripePattern(a *Color, b*Color) *Pattern{
	return &Pattern{a :a, b:b, GetPattern: func(p *algebra.Vector) *Color{
		if int(math.Floor(p.Get()[0])) % 2 == 0{
			return a
		} else {
			return b
		}
	}, Transform: algebra.IdentityMatrix(4)}
}

//SetTransform sets the transform of the pattern
func (p *Pattern) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	} else{
		p.Transform = m
	}
}
