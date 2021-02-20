package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
)

//Sphere Data type for a 3D sphere
type Sphere struct {
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
		transform: mat, material: canvas.NewDefaultMaterial()}
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

//NormalAt returns the normal to the sphere at the location "point"
func (s *Sphere) NormalAt(point *algebra.Vector) *algebra.Vector {
	inverseTransform := s.transform.Inverse()
	sphereBoundary := inverseTransform.MultiplyByVec(point)
	sphereNormal, err := sphereBoundary.Subtract(algebra.NewPoint(0, 0, 0))
	if err != nil {
		panic(err)
		return nil
	}
	worldNormal := inverseTransform.Transpose().MultiplyByVec(sphereNormal)

	res, err := worldNormal.Normalize()
	if err != nil {
		panic(err)
		return nil
	}
	res = algebra.NewVector(res.Get()[:3]...)
	return res
}
