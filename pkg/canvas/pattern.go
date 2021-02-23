package canvas

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
)

//Pattern represents a pattern of colors
type Pattern struct{
	a *Color
	b *Color
	getPattern func(vector *algebra.Vector, colorA *Color, colorB *Color) *Color
	Transform *algebra.Matrix
}

//GetColor returns the color of the pattern at the given point
func (p *Pattern) GetColor(point *algebra.Vector) *Color{
	return p.getPattern(point, p.a, p.b)
}

//StripePattern Creates a new Stripe Pattern in the x coordinate with default identity transformation
func StripePattern(a *Color, b*Color) *Pattern{
	return &Pattern{a :a, b:b, getPattern: func(p *algebra.Vector, a *Color, b *Color) *Color{
		if int(math.Floor(p.Get()[0])) % 2 == 0{
			return a
		} else {
			return b
		}
	}, Transform: algebra.IdentityMatrix(4)}
}

//GradientPattern creates a new Gradient Pattern for two colors using linear interpolation in the x coordinate by default
func GradientPattern(a *Color, b*Color) *Pattern{
	return &Pattern{a :a, b:b, getPattern: func(p *algebra.Vector, a *Color, b *Color) *Color{
		return a.Add(b.Subtract(a).ScalarMult(p.Get()[0] -math.Floor(p.Get()[0])))
	}, Transform: algebra.IdentityMatrix(4)}
}

//RingPattern Creates a new Ring Pattern as rings that expand in the x-z directions by default
func RingPattern(a *Color, b*Color) *Pattern{
	return &Pattern{a :a, b:b, getPattern: func(p *algebra.Vector, a *Color, b *Color) *Color{
		if int(math.Floor(math.Sqrt(math.Pow(p.Get()[0],2) + math.Pow(p.Get()[2],2)))) % 2 == 0{
			return a
		} else {
			return b
		}
	}, Transform: algebra.IdentityMatrix(4)}
}

//CheckerPattern Creates a new Checker Pattern as cubes that expand in every direction
func CheckerPattern(a *Color, b *Color) *Pattern{
	return &Pattern{a :a, b:b, getPattern: func(p *algebra.Vector, a *Color, b *Color) *Color{
		if int(math.Floor(p.Get()[0]) + math.Floor(p.Get()[1]) + math.Floor(p.Get()[2]))%2 == 0{
			return a
		} else {
			return b
		}
	}, Transform: algebra.IdentityMatrix(4)}
}

//SolidPattern Creates a new Solid Pattern that returns a single constant color
func SolidPattern(a *Color) *Pattern{
	return &Pattern{a : a, b: nil, getPattern: func (p *algebra.Vector, a *Color, b *Color) *Color{
		return a
	}, Transform: algebra.IdentityMatrix(4)}
}

//NestedPattern Creates a new nested Pattern from a new pattern that combines two other patterns
func NestedPattern(pattern *Pattern, patternA *Pattern, patternB *Pattern) *Pattern{
	return &Pattern{a : nil, b : nil, getPattern: func(p *algebra.Vector, a *Color, b *Color) *Color{
		colorA := patternA.GetColor(p)
		colorB := patternB.GetColor(p)
		return pattern.getPattern(p, colorA, colorB)
	}, Transform: algebra.IdentityMatrix(4)}
}

//BlendedPattern Creates a new blended Pattern from two patterns and a heuristic function blend
//that should take two colors and return a new color, if it is provided nil as a heuristic it takes the average of two
//colors.
func BlendedPattern(patternA *Pattern, patternB *Pattern, blend func(colorA, colorB *Color) *Color) *Pattern{
	if blend == nil{
		return &Pattern{a : nil, b: nil, getPattern: func(p *algebra.Vector, colorA *Color, colorB *Color) *Color {
			colorFromA := patternA.GetColor(p)
			colorFromB := patternB.GetColor(p)
			return (colorFromA.Add(colorFromB)).ScalarMult(1/2)
		}, Transform: algebra.IdentityMatrix(4)}
	} else{
		return &Pattern{a : nil, b : nil, getPattern: func(p *algebra.Vector, colorA *Color, colorB *Color) *Color {
			colorFromA := patternA.GetColor(p)
			colorFromB := patternB.GetColor(p)
			return blend(colorFromA, colorFromB)
		}, Transform: algebra.IdentityMatrix(4)}
	}
}

//SetTransform sets the transform of the pattern
func (p *Pattern) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	} else{
		p.Transform = m
	}
}

//Random Noise Patterns

func PerlinNoisePattern(pattern *Pattern) *Pattern{
	return &Pattern{}
}


