package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"math"
	"testing"
)

func TestWorldObjectConversion(t *testing.T) {
	g1 := primitives.NewGroup(algebra.RotationY(math.Pi / 2))
	g2 := primitives.NewGroup(algebra.ScalingMatrix(2, 2, 2))
	g1.AddChild(g2)
	s := primitives.NewSphere(algebra.TranslationMatrix(5, 0, 0))
	g2.AddChild(s)
	p := primitives.WorldToObject(s, algebra.NewPoint(-2, 0, -10))
	testVectorEquals(t, p.Get(), algebra.NewPoint(0, 0, -1).Get())

	g1 = primitives.NewGroup(algebra.RotationY(math.Pi / 2))
	g2 = primitives.NewGroup(algebra.ScalingMatrix(1, 2, 3))
	g1.AddChild(g2)
	s = primitives.NewSphere(algebra.TranslationMatrix(5, 0, 0))
	g2.AddChild(s)
	n := primitives.ObjectToWorld(s, algebra.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	testVectorEquals(t, n.Get(), algebra.NewVector(0.2857, 0.4286, -0.8571).Get())
}

func TestNormalAt(t *testing.T) {
	//Testing normal on a group
	g1 := primitives.NewGroup(algebra.RotationY(math.Pi / 2))
	g2 := primitives.NewGroup(algebra.ScalingMatrix(1, 2, 3))
	g1.AddChild(g2)
	s := primitives.NewSphere(algebra.TranslationMatrix(5, 0, 0))
	g2.AddChild(s)
	n := primitives.NormalAt(s, algebra.NewPoint(1.7321, 1.1547, -5.5774), nil)
	testVectorEquals(t, n.Get(), algebra.NewVector(0.2857, 0.4286, -0.8571).Get())
}
