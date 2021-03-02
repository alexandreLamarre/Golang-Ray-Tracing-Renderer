package primitives

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)
//Plane implements a XZ plane by default
type Plane struct{
	parent    Shape
	origin    *algebra.Vector
	direction *algebra.Vector
	transform *algebra.Matrix
	material  *canvas.Material
}

func NewPlane(m *algebra.Matrix) *Plane {
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		mat = algebra.IdentityMatrix(4)
	}
	return &Plane{ origin: algebra.NewPoint(0,0,0), direction: algebra.NewVector(1,0,1),
		transform: mat, material: canvas.NewDefaultMaterial(), parent: nil}
}

func (p *Plane) GetPosition() (*algebra.Vector, *algebra.Vector){
	return p.origin, p.direction
}

func (p *Plane) SetTransform(m *algebra.Matrix) {
	p.transform = m
}

func (p *Plane) GetTransform() *algebra.Matrix {
	return p.transform
}

func (p *Plane) GetMaterial() *canvas.Material{
	return p.material
}

func (p *Plane) SetMaterial(m *canvas.Material) {
	p.material = m
}

//SetParent Setter for parent shape
func(p *Plane) SetParent(shape Shape){
	p.parent = shape
}

//GetParent Getter for parent shape
func(p *Plane) GetParent() Shape {
	return p.parent
}

//GetBounds Getter for default bounds of this Shape
func (p *Plane) GetBounds() (*algebra.Vector, *algebra.Vector){
	return algebra.NewPoint(math.Inf(-1),0,math.Inf(-1)), algebra.NewPoint(math.Inf(1),0,math.Inf(1))
}

func (p *Plane) LocalIntersect(r *algebra.Ray) ([]*Intersection, bool) {
	EPSILON := 0.00001
	if math.Abs(r.Get()["direction"].Get()[1]) < EPSILON {
		return []*Intersection{}, false // ray direction is parallel
	}

	t := -r.Get()["origin"].Get()[1]/r.Get()["direction"].Get()[1]
	return []*Intersection{NewIntersection(p,t)}, true
}

func (p *Plane) LocalNormalAt(point *algebra.Vector, hit *Intersection) (*algebra.Vector, error){
	return algebra.NewVector(0, 1, 0), nil
}