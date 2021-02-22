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
	LocalIntersect(r *algebra.Ray) ([]float64, bool)
	LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error)
}

func NormalAt(s Shape, point *algebra.Vector) *algebra.Vector {
	inverseTransform := s.GetTransform().Inverse()
	localPoint := inverseTransform.MultiplyByVec(point)
	localNormal, err := s.LocalNormalAt(localPoint)

	worldNormal := inverseTransform.Transpose().MultiplyByVec(localNormal)

	res, err := worldNormal.Normalize()
	if err != nil {
		panic(err)
		return nil
	}
	res = algebra.NewVector(res.Get()[:3]...)
	return res
}