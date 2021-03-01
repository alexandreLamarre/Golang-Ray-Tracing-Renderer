package primitives

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

type Triangle struct {
	parent Shape
	transform *algebra.Matrix
	material *canvas.Material
	p1 *algebra.Vector
	p2 *algebra.Vector
	p3 *algebra.Vector
	e1 *algebra.Vector
	e2 *algebra.Vector
	normal *algebra.Vector
}

//NewTriangle Initializer for Triangle Shape
func NewTriangle(p1, p2, p3 *algebra.Vector) *Triangle{
	e1, err := p2.Subtract(p1)
	if err != nil{
		panic(err)
	}
	e2, err := p3.Subtract(p1)
	if err != nil{
		panic(err)
	}
	normal, err := algebra.CrossProduct(e2, e1)

	if err != nil{
		panic(err)
	}
	normal, err = normal.Normalize()
	if err != nil{
		panic(err)
	}

	return &Triangle{
		p1:p1, p2: p2, p3:p3,
		e1: e1, e2: e2,
		material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4),
		parent: nil,
		normal:normal}
}

//Shape interface methods

//GetTransform Getter for Triangle Shape transform algebra.Matrix
func (t *Triangle) GetTransform() *algebra.Matrix{
	return t.transform
}

//GetMaterial Getter for Triangle Shape material canvas.Material
func (t *Triangle) GetMaterial() *canvas.Material{
	return t.material
}

//GetParent Getter for Triangle Shape parent Shape
func (t *Triangle) GetParent() Shape{
	return t.parent
}

//GetBounds return the bounding box for triangle
func (t *Triangle) GetBounds() (*algebra.Vector, *algebra.Vector){
	//TODO: implement
	var xMin = math.Inf(1); var yMin = math.Inf(1); var zMin = math.Inf(1)
	var xMax = math.Inf(-1); var yMax = math.Inf(-1); var zMax = math.Inf(-1)
	points := []*algebra.Vector{t.p1, t.p2, t.p3}
	for _, p := range points{
		xMin = math.Min(p.Get()[0], xMin); yMin = math.Min(p.Get()[1], yMin); zMin = math.Min(p.Get()[2], zMin)
		xMax = math.Max(p.Get()[0], xMax); yMax = math.Max(p.Get()[1], yMax); zMax = math.Max(p.Get()[2], zMax)
	}
	return algebra.NewPoint(xMin, yMin, zMin), algebra.NewPoint(xMax, yMax, zMax)
}

//SetTransform Setter for Triangle Shape transform algebra.Matrix
func (t *Triangle) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	t.transform = m
}

//SetMaterial Setter for Triangle Shape material canvas.Material
func (t *Triangle) SetMaterial(m *canvas.Material){
	t.material = m
}

//SetParent Setter for Triangle Shape parent Shape
func (t *Triangle) SetParent(s Shape){
	t.parent = s
}

//LocalIntersect Intersect implementation for a Triangle Shape
func (t *Triangle) LocalIntersect(r *algebra.Ray) ([]*Intersection, bool){
	xs := make([]*Intersection, 0, 0)
	direction := r.Get()["direction"]
	dirCrossProduct, err := algebra.CrossProduct(direction, t.e2)
	if err != nil {panic(err)}
	det, err := algebra.DotProduct(t.e1, dirCrossProduct)
	if err != nil{panic(err)}
	if math.Abs(det) <= algebra.EPSILON{
		return xs, false
	}

	origin := r.Get()["origin"]
	f := 1.0/det
	p1ToOrigin, err := origin.Subtract(t.p1)
	if err != nil{
		panic(err)
	}
	u, err := algebra.DotProduct(p1ToOrigin, dirCrossProduct)
	if err != nil{panic(err)}
	u *= f
	if u < 0 || u > 1{
		return xs, false
	}

	originCross, err := algebra.CrossProduct(p1ToOrigin, t.e1)
	if err != nil{
		panic(err)
	}
	v, err:= algebra.DotProduct(direction, originCross)
	v *= f
	if err != nil{panic(err)}
	if v < 0 || (u+v) > 1{
		return xs, false
	}

	pos, err := algebra.DotProduct(t.e2, originCross)
	if err != nil{
		panic(err)
	}

	pos *= f
	xs = append(xs, NewIntersection(t, pos))
	return xs, true
}

//LocalNormalAt Normal implementation for a Triangle Shape
func (t *Triangle) LocalNormalAt( p *algebra.Vector) (*algebra.Vector, error){
	return t.normal, nil
}