package objects

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"math"
)

//NewHexagon Initializer for a Hexagon Shape
func NewHexagon(m *algebra.Matrix) *primitives.Group{
	hex := primitives.NewGroup(m)

	for n := 0.0; n <= 5; n++{
		side := hexagonSide(algebra.RotationY(n*math.Pi/3))
		hex.AddChild(side)
	}
	return hex
}


func hexagonCorner() primitives.Shape{
	corner := primitives.NewSphere(algebra.TranslationMatrix(0, 0, -1).Multiply(algebra.ScalingMatrix(0.25,0.25,0.25)))
	//corner := primitives.NewSphere(algebra.ScalingMatrix(0.25,0.25,0.25).Multiply(algebra.TranslationMatrix(0, 0, -1)))
	return corner
}

func hexagonEdge() primitives.Shape{
	edge := primitives.NewCylinder(algebra.TranslationMatrix(0, 0, -1).Multiply(
		algebra.RotationY(-math.Pi/6).Multiply(
			algebra.RotationZ(-math.Pi/2).Multiply(
				algebra.ScalingMatrix(0.25, 1, 0.25)))))
	//edge := primitives.NewCylinder(algebra.ScalingMatrix(0.25, 1, 0.25).Multiply(algebra.RotationZ(-math.Pi/2)).Multiply(algebra.RotationY(-math.Pi/6)).Multiply(algebra.TranslationMatrix(0, 0, -1)))
	edge.SetMaximum(1)
	edge.SetMinimum(0)
	return edge
}

func hexagonSide(m *algebra.Matrix) primitives.Shape{
	side := primitives.NewGroup(m)
	side.AddChild(hexagonCorner())
	side.AddChild(hexagonEdge())
	return side
}
