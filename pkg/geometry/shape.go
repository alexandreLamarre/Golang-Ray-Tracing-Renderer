package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
)

//Shape represents a shape interface in 3D space
type Shape interface {
	GetPosition() *algebra.Vector
	SetTransform(m *algebra.Matrix)
	GetTransform() *algebra.Matrix
	SetMaterial(m *canvas.Material)
	GetMaterial() *canvas.Material
	NormalAt(point *algebra.Vector) *algebra.Vector //returns the normal at the location "point" on the shape
}
