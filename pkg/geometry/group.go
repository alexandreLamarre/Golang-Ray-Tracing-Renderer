package geometry

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
)

type GroupNormalError Group

func (e GroupNormalError) Error() string{
	return fmt.Sprintf("Invalid call to LocalNormalAt on Abstract Group Shape")
}

//Group represents a collection of shapes, a container for related shapes
type Group struct{
	parent Shape
	shapes []Shape
	transform *algebra.Matrix
}

//NewGroup initializer for a group struct with a given 4x4 transformation matrix, if the matrix
// is not 4x4 or nil, set the identity transform
func NewGroup(m *algebra.Matrix) *Group{
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	emptyShapes := make([]Shape, 0, 0)
	return &Group{transform: mat, parent: nil, shapes: emptyShapes}
}

//AddChild adds a new shape to the Group's container
func (g *Group) AddChild(s Shape){
	s.SetParent(g)
	g.shapes = append(g.shapes, s)
}

//Shape interface methods

//GetMaterial Getter for Shape material, but abstract Group does not have a material so return nil
func (g *Group) GetMaterial() *canvas.Material{
	return nil
}

//SetMaterial Setter for Shape material, but abstract Group does not have a material so do nothing
func (g *Group) SetMaterial(m *canvas.Material){
	return
}

//GetTransform Getter for Shape transform
func (g *Group) GetTransform() *algebra.Matrix{
	return g.transform
}

//SetTransform Setter for Shape transform
func (g *Group) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	g.transform = m
}

//GetParent Getter for Parent of Shape
func (g *Group) GetParent() Shape{
	return g.parent
}

//SetParent Setter for Parent of Shape
func (g *Group) SetParent(shape Shape){
	g.parent = shape
}

//LocalIntersect Intersect Implementation for Group Shape
func (g *Group) LocalIntersect(r *algebra.Ray) ([]*Intersection, bool){
	xs := make([]*Intersection, 0, 0)
	var hit bool = false
	for _, s := range g.shapes{
		m := s.GetTransform()
		ri :=  r.Transform(m.Inverse())
		shapeXs, shapeHit := s.LocalIntersect(ri)
		hit = hit || shapeHit
		for _, t := range shapeXs{
			xs = append(xs, t)
		}
	}
	return xs, hit
}

//LocalNormalAt Normal implementation for Group Shape: Only concrete child shapes have a local Normal
func (g *Group) LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error){
	return nil, GroupNormalError(*g)
}
