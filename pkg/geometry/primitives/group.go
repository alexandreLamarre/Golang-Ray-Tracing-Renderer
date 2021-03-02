package primitives

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

type GroupNormalError Group

func (e GroupNormalError) Error() string{
	return fmt.Sprintf("Invalid call to LocalNormalAt on Abstract Group Shape")
}

//Group represents a collection of shapes, a container for related shapes
type Group struct{
	parent    Shape
	shapes    []Shape
	transform *algebra.Matrix
	bounds [2]*algebra.Vector
}

//NewGroup initializer for a group struct with a given 4x4 transformation matrix, if the matrix
// is not 4x4 or nil, set the identity transform
func NewGroup(m *algebra.Matrix) *Group {
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	emptyShapes := make([]Shape, 0, 0)
	return &Group{transform: mat, parent: nil, shapes: emptyShapes, bounds: [2]*algebra.Vector{}}
}

//GetShapes Getter for shapes array field of a group
func (g *Group)GetShapes() []Shape{
	return g.shapes
}

//AddChild adds a new shape to the Group's container
func (g *Group) AddChild(s Shape){
	s.SetParent(g)
	g.shapes = append(g.shapes, s)
	min, max := g.getBounds()
	g.bounds = [2]*algebra.Vector{min, max}
	fmt.Println(g.bounds)
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
func (g *Group) GetParent() Shape {
	return g.parent
}

//SetParent Setter for Parent of Shape
func (g *Group) SetParent(shape Shape){
	g.parent = shape
}

func (g *Group) GetBounds()(*algebra.Vector, *algebra.Vector){
	return g.bounds[0], g.bounds[1]
}

//getBounds Calculates bounds of this Shape
func (g *Group) getBounds() (*algebra.Vector, *algebra.Vector){
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
func (g *Group) LocalIntersect(r *algebra.Ray) ([]*Intersection, bool){
	xs := make([]*Intersection, 0, 0)
	var hit bool = false

	// Get the AABB of the group
	min, max := g.GetBounds()
	fmt.Println(min, max)

	if min == nil{
		return xs, false
	}
	if GetBoundsTransform(min, max, g.GetTransform()).Intersect(r.Transform(g.GetTransform())) == false{
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
func (g *Group) LocalNormalAt(p *algebra.Vector, hit *Intersection) (*algebra.Vector, error){
	return nil, GroupNormalError(*g)
}
