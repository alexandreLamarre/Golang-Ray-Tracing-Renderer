package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

//Sphere Data type for a 3D sphere
type Sphere struct {
	parent Shape
	origin    *algebra.Vector
	radius    float64
	transform *algebra.Matrix
	material  *canvas.Material
}

// NewSphere creates a new Sphere datatype at origin 0,0,0 with unit radius and no ray intersections
func NewSphere(m *algebra.Matrix) *Sphere {
	mat := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		mat = algebra.IdentityMatrix(4)
	}
	return &Sphere{origin: algebra.NewPoint(0, 0, 0), radius: 1.0,
		transform: mat, material: canvas.NewDefaultMaterial(), parent: nil}
}

func NewGlassSphere(m *algebra.Matrix, refractiveIndex float64) *Sphere{
	matrix := m
	if m == nil || len(m.Get()) != 4 || len(m.Get()[0]) != 4 {
		matrix = algebra.IdentityMatrix(4)
	}
	material := canvas.NewDefaultMaterial()
	material.Transparency = 1.0
	material.RefractiveIndex = refractiveIndex
	return &Sphere{origin: algebra.NewPoint(0, 0, 0), radius: 1.0,
		transform: matrix, material: material, parent: nil}
}


// Sphere interface Shape Methods

func (s *Sphere) GetPosition() *algebra.Vector {
	return s.origin
}

//SetTransform sets the Sphere's transformation
func (s *Sphere) SetTransform(m *algebra.Matrix) {
	s.transform = m
}

func (s *Sphere) GetTransform() *algebra.Matrix {
	return s.transform
}

//SetMaterial sets the Sphere's material
func (s *Sphere) SetMaterial(m *canvas.Material) {
	s.material = m
}

//GetMaterial returns the spheres Material
func (s *Sphere) GetMaterial() *canvas.Material {
	return s.material
}

//SetParent Setter for parent shape
func(s *Sphere) SetParent(shape Shape){
	s.parent = shape
}

//GetParent Getter for parent shape
func(s *Sphere) GetParent() Shape{
	return s.parent
}


//NormalAt returns the normal to the sphere at the location "point"
func (s *Sphere) LocalNormalAt(point *algebra.Vector) (*algebra.Vector, error) {

	sphereNormal, err := point.Subtract(algebra.NewPoint(0, 0, 0))

	return sphereNormal, err
}

//LocalIntersect returns the intersection of a ray with a sphere
func (s *Sphere) LocalIntersect(r *algebra.Ray) ([]float64, bool){
	got := r.Get()
	origin := got["origin"]
	direction := got["direction"]
	sphereToRay, err := origin.Subtract(s.GetPosition())
	if err != nil {
		panic(err)
	}

	a, err := algebra.DotProduct(direction, direction)
	if err != nil {
		panic(err)
	}
	if a == 0 {
		panic(algebra.ZeroDivide(0))
	}

	b, err := algebra.DotProduct(direction, sphereToRay)
	if err != nil {
		panic(err)
	}
	b = 2 * b

	c, err := algebra.DotProduct(sphereToRay, sphereToRay)
	if err != nil {
		panic(err)
	}
	c = c - 1
	discriminant := math.Pow(b, 2) - (4 * a * c)

	if discriminant < 0 { // No rays intersect the sphere
		 return []float64{}, false
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	return []float64{t1, t2}, true
}