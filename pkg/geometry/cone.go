package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

//Cone defines a default Cone Shape
type Cone struct{
	parent Shape
	closed bool //determines if the cone is hollow or has caps on the ends
	transform *algebra.Matrix
	material *canvas.Material
	maximum float64 //maximum y-value by default without transformations
	minimum float64 //minimum y-value by default without transformations
}

//NewCone returns a new Cone Shape
func NewCone(m *algebra.Matrix) *Cone{
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	return &Cone{transform: mat, material: canvas.NewDefaultMaterial(),
		maximum: math.Inf(1), minimum: math.Inf(-1), closed: false, parent: nil}
}

//SetMinimum Setter for cylinder minimum y-truncation
func (cone *Cone) SetMinimum(min float64){
	cone.minimum = min
}

//SetMaximum Setter for cylinder maximum y-truncation
func (cone *Cone) SetMaximum(max float64){
	cone.maximum = max
}

//SetClosed Setter for cylinder closededness: whether or not it has caps on the cylinder
func (cone *Cone) SetClosed(closed bool){
	cone.closed = closed
}

//Shape interface methods

//GetTransform Getter for Cylinder Shape transform
func (cone *Cone) GetTransform() *algebra.Matrix{
	return cone.transform
}

//GetMaterial Getter for Cylinder Shape material
func (cone *Cone) GetMaterial() *canvas.Material{
	return cone.material
}

//SetTransform Setter for Cylinder Shape transform
func (cone *Cone) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	cone.transform = m
}

//SetMaterial Setter for Cylinder Shape material
func (cone *Cone) SetMaterial(m *canvas.Material){
	cone.material = m
}

//SetParent Setter for parent shape
func(cone *Cone) SetParent(shape Shape){
	cone.parent = shape
}

//GetParent Getter for parent shape
func(cone *Cone) GetParent() Shape{
	return cone.parent
}

//LocalIntersect returns the itersection values for a Ray with a Cylinder
func (cone *Cone) LocalIntersect(ray *algebra.Ray) ([]float64, bool){
	direction := ray.Get()["direction"]
	origin := ray.Get()["origin"]
	dx := direction.Get()[0]; dy := direction.Get()[1]; dz:= direction.Get()[2]
	ox := origin.Get()[0]; oy := origin.Get()[1]; oz := origin.Get()[2]


	a := (dx * dx) -(dy * dy) + (dz * dz)
	b := (2 * ox * dx) - (2 * oy * dy) + (2 * oz * dz)
	c := (ox*ox) - (oy * oy) + (oz * oz)

	disc := b*b - 4 * a *c

	xs := make([]float64, 0, 0)
	xs = cone.intersectCaps(ray, xs)

	EPSILON := 0.00001
	if a <= EPSILON && a >= -EPSILON{
		if (b >= EPSILON || b <= -EPSILON) && len(xs) == 0{
			xs = append(xs, -c/(2*b))
			return xs, true
		}

		var hit bool
		if len(xs) == 0{
			hit = false
		} else{
			hit = true
		}
		return xs, hit
	}

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
	if cone.minimum < y0 && cone.maximum > y0{
		xs = append(xs, t0)
	}

	y1 := oy + t1* dy
	if cone.minimum < y1 && cone.maximum > y1{
		xs = append(xs, t1)
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
func (cone *Cone) LocalNormalAt(p *algebra.Vector) (*algebra.Vector, error){
	x := p.Get()[0]; y := p.Get()[1]; z := p.Get()[2]
	dist := x*x + z * z
	y = math.Sqrt(dist)
	if p.Get()[1] > 0 {
		y = -y
	}
	EPSILON:= 0.001
	if dist < 1 && y >= cone.maximum - EPSILON{
		return algebra.NewVector(0, 1, 0), nil
	} else if dist < 1 && y <= cone.minimum + EPSILON{
		return algebra.NewVector(0, -1, 0), nil
	}

	return algebra.NewVector(x, y, z), nil
}

//cone helpers

func (cone *Cone) intersectCaps(ray *algebra.Ray, xs []float64) []float64{
	origin := ray.Get()["origin"]
	direction := ray.Get()["direction"]
	oy := origin.Get()[1]; dy := direction.Get()[1]
	EPSILON := 0.0001
	if !cone.closed || (dy < EPSILON && dy > -EPSILON){
		return xs
	}

	t := (cone.minimum -oy)/dy
	if checkCap(ray , t){
		xs = append(xs , t)
	}

	t = (cone.maximum - oy)/dy
	if checkCap(ray, t){
		xs = append(xs, t)
	}
	return xs
}