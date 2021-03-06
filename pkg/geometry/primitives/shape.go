package primitives

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
	GetBounds() (*algebra.Vector, *algebra.Vector)
	LocalIntersect(r *algebra.Ray) ([]*Intersection, bool)
	LocalNormalAt(p *algebra.Vector, hit *Intersection) (*algebra.Vector, error)
}

//NormalAt is the super class method to get the normal of an object, LocalNormalAt implements the specifics of
// the shape subclasses
func NormalAt(s Shape, point *algebra.Vector, hit *Intersection) *algebra.Vector {
	localPoint := WorldToObject(s, point)
	localNormal, err := s.LocalNormalAt(localPoint, hit)
	if err != nil {
		panic(err)
	}
	return ObjectToWorld(s, localNormal)
}

//PatternAtObject takes a shape and pattern and the point the ray intersects in the world and returns what color should
// be there given these parameters
func PatternAtObject(s Shape, pattern *canvas.Pattern, worldPoint *algebra.Vector) *canvas.Color {
	objectPoint := s.GetTransform().Inverse().MultiplyByVec(worldPoint)
	patternPoint := pattern.Transform.Inverse().MultiplyByVec(objectPoint)
	return pattern.GetColor(patternPoint)
}

// helpers

func WorldToObject(s Shape, point *algebra.Vector) *algebra.Vector {

	if s.GetParent() != nil {
		point = WorldToObject(s.GetParent(), point)
	}
	return s.GetTransform().Inverse().MultiplyByVec(point)
}

func ObjectToWorld(s Shape, normal *algebra.Vector) *algebra.Vector {
	normal = s.GetTransform().Inverse().Transpose().MultiplyByVec(normal)
	normal.Get()[3] = 0
	normal, err := normal.Normalize()
	if err != nil {
		panic(err)
	}
	if s.GetParent() != nil {
		normal = ObjectToWorld(s.GetParent(), normal)
	}
	return normal
}
