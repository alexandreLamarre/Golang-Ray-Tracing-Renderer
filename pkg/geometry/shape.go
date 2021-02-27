package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
)

//Shape represents a shape interface in 3D space
type Shape interface {
	SetTransform(m *algebra.Matrix)
	GetTransform() *algebra.Matrix
	SetMaterial(m *canvas.Material)
	GetMaterial() *canvas.Material
	SetParent(s Shape)
	GetParent() Shape
	LocalIntersect(r *algebra.Ray) ([]*Intersection, bool)
	LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error)
}

//NormalAt is the super class method to get the normal of an object, LocalNormalAt implements the specifics of
// the shape subclasses
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

//PatternAtObject takes a shape and pattern and the point the ray intersects in the world and returns what color should
// be there given these parameters
func PatternAtObject(s Shape, pattern *canvas.Pattern, worldPoint *algebra.Vector) *canvas.Color{
	objectPoint := s.GetTransform().Inverse().MultiplyByVec(worldPoint)
	patternPoint := pattern.Transform.Inverse().MultiplyByVec(objectPoint)
	return pattern.GetColor(patternPoint)
}