package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

//Cube defines a 3d cube Shape
type Cube struct{
	parent Shape
	transform *algebra.Matrix
	material *canvas.Material
}

//NewCube returns a new Cube Shape with an identity matrix/ default material
func NewCube(m *algebra.Matrix) *Cube{
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	return &Cube{transform: mat, material: canvas.NewDefaultMaterial(), parent: nil}
}

// Shape interface functions

//GetTransform Getter for Cube transform, Shape interface method
func (c *Cube) GetTransform() *algebra.Matrix{
	return c.transform
}

//SetTransform Setter for Cube transform, Shape interface method
func (c *Cube) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	c.transform = m
}

//GetMaterial Getter for Cube material, Shape interface method
func (c *Cube) GetMaterial() *canvas.Material{
	return c.material
}

//SetMaterial Setter for Cube material, Shape interface method
func (c *Cube) SetMaterial(m *canvas.Material){
	c.material = m
}

//SetParent Setter for parent shape
func(c * Cube) SetParent(shape Shape){
	c.parent = shape
}

//GetParent Getter for parent shape
func(c *Cube) GetParent() Shape{
	return c.parent
}

//GetBounds Getter for default bounds of this Shape
func (c *Cube) GetBounds() (*algebra.Vector, *algebra.Vector){
	return algebra.NewPoint(-1,-1,-1), algebra.NewPoint(1,1,1)
}

//LocalIntersect returns the itersection values for a Ray with a Cube
func (c *Cube) LocalIntersect(ray *algebra.Ray) ([]*Intersection, bool){
	origin := ray.Get()["origin"]; direction := ray.Get()["direction"]
	xtmin, xtmax := checkAxis(origin.Get()[0], direction.Get()[0])
	ytmin, ytmax := checkAxis(origin.Get()[1], direction.Get()[1])
	ztmin, ztmax := checkAxis(origin.Get()[2], direction.Get()[2])

	tmin := max(xtmin, ytmin, ztmin)
	tmax := min(xtmax, ytmax, ztmax)

	if tmin > tmax{
		return []*Intersection{}, false
	}

	return []*Intersection{NewIntersection(c,tmin), NewIntersection(c,tmax)}, true
}

//LocalNormalAt returns the normal at a ray intersection point
func (c *Cube) LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error){
	maxc := max(math.Abs(p.Get()[0]), math.Abs(p.Get()[1]), math.Abs(p.Get()[2]))

	if maxc == math.Abs(p.Get()[0]){
		return algebra.NewVector(p.Get()[0], 0, 0), nil
	} else if maxc == math.Abs(p.Get()[1]){
		return algebra.NewVector(0, p.Get()[1], 0), nil
	}
	return algebra.NewVector(0, 0, p.Get()[2]), nil
}

//helpers for cube methods
func checkAxis(origin, direction float64) (float64, float64){
	tminNumerator := -1 -origin
	tmaxNumerator := 1 - origin

	var tmin float64
	var tmax float64
	EPSILON := 0.0001
	if math.Abs(direction) >= EPSILON {
		tmin = tminNumerator/direction
		tmax = tmaxNumerator/direction
	} else {
		tmin = tminNumerator * math.Inf(1)
		tmax = tmaxNumerator * math.Inf(1)
	}

	if tmin > tmax {
		temp := tmin
		tmin = tmax
		tmax = temp
	}

	return tmin, tmax
}

func max(values ...float64) float64{
	if len(values) == 1{
		return values[0]
	}
	maxVal := math.Inf(-1)
	for i:= 0; i < len(values); i++{
		maxVal = math.Max(values[i], maxVal)
	}
	return maxVal
}

func min(values ...float64) float64{
	if len(values) == 1{
		return values[0]
	}
	minVal := math.Inf(1)
	for i:= 0; i < len(values); i++{
		minVal = math.Min(values[i], minVal)
	}
	return minVal
}
