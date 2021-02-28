package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

//Cylinder defines a default cylinder Shape
type Cylinder struct{
	parent Shape
	closed bool //determines if the cylinder is hollow or has caps on the ends
	transform *algebra.Matrix
	material *canvas.Material
	maximum float64 //maximum y-value by default without transformations
	minimum float64 //minimum y-value by default without transformations
}

//NewCylinder returns a new Cylinder Shape
func NewCylinder(m * algebra.Matrix) *Cylinder{
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	return &Cylinder{transform: mat, material: canvas.NewDefaultMaterial(),
		maximum: math.Inf(1), minimum: math.Inf(-1), closed: false, parent: nil}
}

//SetMinimum Setter for cylinder minimum y-truncation
func (cyl *Cylinder) SetMinimum(min float64){
	cyl.minimum = min
}

//SetMaximum Setter for cylinder maximum y-truncation
func (cyl *Cylinder) SetMaximum(max float64){
	cyl.maximum = max
}

//SetClosed Setter for cylinder closededness: whether or not it has caps on the cylinder
func (cyl *Cylinder) SetClosed(closed bool){
	cyl.closed = closed
}

//Shape interface methods

//GetTransform Getter for Cylinder Shape transform
func (cyl *Cylinder) GetTransform() *algebra.Matrix{
	return cyl.transform
}

//GetMaterial Getter for Cylinder Shape material
func (cyl *Cylinder) GetMaterial() *canvas.Material{
	return cyl.material
}

//SetTransform Setter for Cylinder Shape transform
func (cyl *Cylinder) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	cyl.transform = m
}

//SetMaterial Setter for Cylinder Shape material
func (cyl *Cylinder) SetMaterial(m *canvas.Material){
	cyl.material = m
}

//SetParent Setter for parent shape
func(cyl *Cylinder) SetParent(shape Shape){
	cyl.parent = shape
}

//GetParent Getter for parent shape
func(cyl *Cylinder) GetParent() Shape{
	return cyl.parent
}

//GetBounds Getter for default bounds of this Shape
func (cyl *Cylinder) GetBounds() (*algebra.Vector, *algebra.Vector){
	return algebra.NewPoint(-1,cyl.minimum,-1), algebra.NewPoint(1,cyl.maximum,1)
}

//LocalIntersect returns the itersection values for a Ray with a Cylinder
func (cyl *Cylinder) LocalIntersect(ray *algebra.Ray) ([]*Intersection, bool){
	direction := ray.Get()["direction"]
	origin := ray.Get()["origin"]
	dx := direction.Get()[0]; dy := direction.Get()[1]; dz:= direction.Get()[2]
	a := dx * dx + dz * dz
	EPSILON := 0.00001


	xs := make([]*Intersection, 0, 0)

	xs = cyl.intersectCaps(ray, xs)


	if a <= EPSILON && a >= -EPSILON{
		var hit bool
		if len(xs) == 0{
			hit = false
		} else{
			hit = true
		}
		return xs, hit
	}

	ox := origin.Get()[0]; oy := origin.Get()[1]; oz := origin.Get()[2]
	b := 2 * ox * dx + 2 * oz * dz

	c := ox*ox + oz * oz -1
	disc := b*b - 4 * a *c
	if disc < 0 {
		var hit bool
		if len(xs) == 0{
			hit = false
		} else{
			hit = true
		}
		return xs, hit
	}

	t0 := (-b- math.Sqrt(disc))/(2*a)
	t1 := (-b + math.Sqrt(disc))/(2*a)


	y0 := oy + t0 * dy
	if cyl.minimum < y0 && cyl.maximum > y0{
		xs = append(xs, NewIntersection(cyl,t0))
	}

	y1 := oy + t1* dy
	if cyl.minimum < y1 && cyl.maximum > y1{
		xs = append(xs, NewIntersection(cyl,t1))
	}

	var hit bool
	if len(xs) == 0{
		hit = false
	} else{
		hit = true
	}

	return xs, hit
}

//LocalNormalAt returns the normal at an intersection point of a Cylinder Shape, shape interface method
func (cyl *Cylinder) LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error){
	x := p.Get()[0]; y := p.Get()[1]; z := p.Get()[2]
	dist := x*x + z * z
	EPSILON:= 0.001
	if dist < 1 && y >= cyl.maximum - EPSILON{
		return algebra.NewVector(0, 1, 0), nil
	} else if dist < 1 && y <= cyl.minimum + EPSILON{
		return algebra.NewVector(0, -1, 0), nil
	}
	return algebra.NewVector(p.Get()[0], 0, p.Get()[2]), nil
}

//cylinder helpers

func checkCap(ray *algebra.Ray, t float64) bool{
	origin := ray.Get()["origin"]
	direction := ray.Get()["direction"]
	x := origin.Get()[0] + t * direction.Get()[0]
	z := origin.Get()[2] + t * direction.Get()[2]
	return (x*x + z*z) <= 1
}

func (cyl *Cylinder) intersectCaps(ray *algebra.Ray, xs []*Intersection) []*Intersection{
	origin := ray.Get()["origin"]
	direction := ray.Get()["direction"]
	oy := origin.Get()[1]; dy := direction.Get()[1]
	EPSILON := 0.0001
	if !cyl.closed || (dy < EPSILON && dy > -EPSILON){
		return xs
	}

	t := (cyl.minimum -oy)/dy
	if checkCap(ray , t){
		xs = append(xs , NewIntersection(cyl,t))
	}

	t = (cyl.maximum - oy)/dy
	if checkCap(ray, t){
		xs = append(xs, NewIntersection(cyl,t))
	}
	return xs
}





