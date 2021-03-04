package objects

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"math"
	"reflect"
)

//ErrorCSGNormal returns an error if csg.LocalNormalAt is ever called, only a child shape should have its local noral
//called
type ErrorCSGNormal int

func (e ErrorCSGNormal) Error() string {
	return fmt.Sprintf("CSG Local normal called : Invalid use of LocalNormal method")
}

//Contains Check if the Shape is the provided Shape/contains the provided Shape
func Contains(s primitives.Shape, testContained primitives.Shape) bool {
	if reflect.TypeOf(s) == reflect.TypeOf(&primitives.Group{}) {
		g := s.(*primitives.Group)
		for _, v := range g.GetShapes() {
			if Contains(v, testContained) {
				return true
			}
		}
	} else if reflect.TypeOf(s) == reflect.TypeOf(&CSGShape{}) {
		csg := s.(*CSGShape)
		if Contains(csg.Left(), testContained) {
			return true
		}
		if Contains(csg.Right(), testContained) {
			return true
		}
	}
	return s == testContained
}

//CSGShape is the datatype that uses the Constructive Geometry
// approach to handling groups of objects
type CSGShape struct {
	Parent    primitives.Shape
	material  *canvas.Material
	transform *algebra.Matrix
	left      primitives.Shape
	right     primitives.Shape
	action    string
}

//UnionCSG initializes a CSG shape based on the union operation of two Shapes: left and right
func UnionCSG(left, right primitives.Shape) *CSGShape {
	if left == nil {
		panic("Undefined left Shape in CSG operation `union`")
	}
	if right == nil {
		panic("Undefined right Shape in CSG operation `union`")
	}

	c := &CSGShape{Parent: nil, material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4), left: left, right: right, action: "union"}
	left.SetParent(c)
	right.SetParent(c)
	return c
}

//IntersectCSG initializes a CSG shape based on the union operation of two Shapes: left and right
func IntersectCSG(left, right primitives.Shape) *CSGShape {
	if left == nil {
		panic("Undefined left Shape in CSG operation `intersect`")
	}
	if right == nil {
		panic("Undefined right Shape in CSG operation `intersect`")
	}
	c := &CSGShape{Parent: nil, material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4), left: left, right: right, action: "intersect"}
	left.SetParent(c)
	right.SetParent(c)
	return c
}

//DifferenceCSG initializes a CSG shape based on the union operation of two Shapes: left and right
func DifferenceCSG(left, right primitives.Shape) *CSGShape {
	if left == nil {
		panic("Undefined left Shape in CSG operation `difference`")
	}
	if right == nil {
		panic("Undefined right Shape in CSG operation `difference`")
	}
	c := &CSGShape{Parent: nil, material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4), left: left, right: right, action: "difference"}
	left.SetParent(c)
	right.SetParent(c)
	return c
}

//Left returns the left shape of the CSG Shape
func (s *CSGShape) Left() primitives.Shape {
	return s.left
}

//Right returns the left shape of the CSG Shape
func (s *CSGShape) Right() primitives.Shape {
	return s.right
}

//primitives.Shape interface methods

//GetTransform Getter for CSG Shape transform, primitives.Shape interface method
func (s *CSGShape) GetTransform() *algebra.Matrix {
	return s.transform
}

//GetMaterial Getter for CSG Shape material, primitives.Shape interface method
func (s *CSGShape) GetMaterial() *canvas.Material {
	return s.material
}

//GetParent Getter for CSG Parent Shape, primitives.Shape interface method
func (s *CSGShape) GetParent() primitives.Shape {
	return s.Parent
}

//GetBounds Getter for CSG Shape bounding box, primitives.Shape interface method
func (s *CSGShape) GetBounds() (*algebra.Vector, *algebra.Vector) {

	minL, maxL := s.left.GetBounds()
	minR, maxR := s.right.GetBounds()
	bL := primitives.GetBoundsTransform(minL, maxL, s.left.GetTransform())
	br := primitives.GetBoundsTransform(minL, maxL, s.right.GetTransform())
	minL, maxL = bL.Get()
	minR, maxR = br.Get()

	minX := math.Inf(1)
	minY := math.Inf(1)
	minZ := math.Inf(1)
	maxX := math.Inf(-1)
	maxY := math.Inf(-1)
	maxZ := math.Inf(-1)

	minX = math.Min(minX, minL.Get()[0])
	minX = math.Min(minX, minR.Get()[0])
	maxX = math.Max(maxX, maxL.Get()[0])
	maxX = math.Max(maxX, maxR.Get()[0])

	minY = math.Min(minY, minL.Get()[1])
	minY = math.Min(minY, minR.Get()[1])
	maxY = math.Max(maxY, maxL.Get()[1])
	maxY = math.Max(maxY, maxR.Get()[1])

	minZ = math.Min(minZ, minL.Get()[2])
	minZ = math.Min(minZ, minR.Get()[2])
	maxZ = math.Max(maxZ, maxL.Get()[2])
	maxZ = math.Max(maxZ, maxR.Get()[2])

	return algebra.NewVector(minX, minY, minZ), algebra.NewVector(maxX, maxY, maxZ)
}

//SetTransform Setter for CSG Shape transform, primitives.Shape interface method
func (s *CSGShape) SetTransform(m *algebra.Matrix) {
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		panic(algebra.ExpectedDimension(4))
	}
	s.transform = m
}

//SetMaterial Setter for CSG Shape material, primitives.Shape interface method
func (s *CSGShape) SetMaterial(m *canvas.Material) {
	s.material = m
}

//SetParent Setter for CSG Shape parent, primitives.Shape interface method
func (s *CSGShape) SetParent(shape primitives.Shape) {
	s.Parent = shape
}

//LocalIntersect Intersect implementation for CSG Shape, Check intersections of left and right
//and apply the appropriate CSG Shape action. primitives.Shape interface method.
func (s *CSGShape) LocalIntersect(r *algebra.Ray) ([]*primitives.Intersection, bool) {
	xs := make([]*primitives.Intersection, 0, 0)
	is := primitives.NewIntersections()
	err := is.Intersect(s.Left(), r)
	if err != nil {
		panic(err)
	}
	err = is.Intersect(s.Right(), r)
	if err != nil {
		panic(err)
	}

	if is.Count() == 0 {
		return xs, false
	}
	hits := is.GetHits()
	refs := is.GetRef()

	k := len(refs.Get())
	for i := 0; i < k; i++ {
		if min := refs.ExtractMin(); min != nil {
			xs = append(xs, refs.ExtractMin())
		}
	}

	k = len(hits.Get())
	for i := 0; i < k; i++ {
		if min := hits.ExtractMin(); min != nil {
			xs = append(xs, min)
		}
	}

	xs = s.filterIntersections(xs)
	return xs, true
}

//LocalNormalAt Local normal implementation for CSG Shape, should never be directly called on. Only
//the children of CSG shape should be intersected. primitives.Shape interface method.
func (s *CSGShape) LocalNormalAt(p *algebra.Vector, hit *primitives.Intersection) (*algebra.Vector, error) {
	return nil, ErrorCSGNormal(0)
}

// CSG Shape helpers

func (s *CSGShape) filterIntersections(xs []*primitives.Intersection) []*primitives.Intersection {
	inl := false
	inr := false

	result := make([]*primitives.Intersection, 0, 0)

	for _, i := range xs {
		lhit := Contains(s.Left(), i.Object)
		if intersectionAllowed(s.action, lhit, inl, inr) {
			result = append(result, i)
		}

		if lhit {
			inl = !(inl)
		} else {
			inr = !(inr)
		}
	}
	return result
}

func intersectionAllowed(action string, lhit, inl, inr bool) bool {
	if action == "union" {
		return (lhit && (!inr)) || ((!lhit) && (!inl))
	}
	if action == "intersect" {
		return (lhit && inr) || ((!lhit) && inl)
	}
	if action == "difference" {
		return (lhit && !inr) || (!lhit && inl)
	}
	return false
}
