package geometry

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"math"
)

type GroupNormalError Group

func (e GroupNormalError) Error() string{
	return fmt.Sprintf("Invalid call to LocalNormalAt on Abstract Group Shape")
}

//Group represents a collection of shapes, a container for related shapes
type Group struct{
	parent    primitives.Shape
	shapes    []primitives.Shape
	transform *algebra.Matrix
}

//NewGroup initializer for a group struct with a given 4x4 transformation matrix, if the matrix
// is not 4x4 or nil, set the identity transform
func NewGroup(m *algebra.Matrix) *Group{
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	emptyShapes := make([]primitives.Shape, 0, 0)
	return &Group{transform: mat, parent: nil, shapes: emptyShapes}
}

//AddChild adds a new shape to the Group's container
func (g *Group) AddChild(s primitives.Shape){
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
func (g *Group) GetParent() primitives.Shape {
	return g.parent
}

//SetParent Setter for Parent of Shape
func (g *Group) SetParent(shape primitives.Shape){
	g.parent = shape
}

//GetBounds Getter for default bounds of this Shape
func (g *Group) GetBounds() (*algebra.Vector, *algebra.Vector){
	var min *algebra.Vector = nil; var max *algebra.Vector = nil
	if len(g.shapes) == 0 {
		return min, max
	}
	minX := math.Inf(1); minY := math.Inf(1); minZ := math.Inf(1)
	maxX := math.Inf(-1); maxY := math.Inf(-1); maxZ := math.Inf(-1)

	for _, shape := range g.shapes{
		tempMin, tempMax := shape.GetBounds()
		if tempMin != nil{
			b := GetBoundsTransform(tempMin, tempMax, shape.GetTransform())
			tempMin = b.minimum
			tempMax = b.maximum
			tempMinX := tempMin.Get()[0]; tempMinY := tempMin.Get()[1]; tempMinZ := tempMin.Get()[2]
			tempMaxX := tempMax.Get()[0]; tempMaxY := tempMax.Get()[1]; tempMaxZ := tempMax.Get()[2]
			minX = math.Min(tempMinX, minX); minY = math.Min(tempMinY, minY); minZ = math.Min(tempMinZ, minZ)
			maxX = math.Max(tempMaxX, maxX); maxY = math.Max(tempMaxY, maxY); maxZ = math.Max(tempMaxZ, maxZ)
		}
	}
	min = algebra.NewPoint(minX, minY, minZ)
	max = algebra.NewPoint(maxX, maxY, maxZ)
	return min,max
}

//LocalIntersect Intersect Implementation for Group Shape
func (g *Group) LocalIntersect(r *algebra.Ray) ([]*primitives.Intersection, bool){
	xs := make([]*primitives.Intersection, 0, 0)
	var hit bool = false

	// Get the AABB of the group
	min, max := g.GetBounds()

	if min == nil{
		return xs, false
	}

	if GetBoundsTransform(min, max, g.GetTransform()).Intersect(r) == false{
		return xs, false
	}

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
