package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/datatypes"
	"math"
)

//Shape represents a shape interface in 3D space
type Shape interface {
	SetTransform(m *algebra.Matrix)
	SetMaterial(m *canvas.Material)
	NormalAt(point *algebra.Vector) *algebra.Vector //returns the normal at the location "point" on the shape
}

//Sphere Data type for a 3D sphere
type Sphere struct {
	origin    *algebra.Vector
	radius    float64
	transform *algebra.Matrix
	material *canvas.Material
}

//Intersections data type keeps track of t values of the intersections of rays with a sphere
type Intersections struct {
	hits map[Shape]map[*algebra.Ray]*datatypes.MinHeap // hits on contours of objects
	ref  map[Shape]map[*algebra.Ray]*datatypes.MinHeap // used in ray reflections/refractions
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

//SetTransform sets the Sphere's transformation
func (s *Sphere) SetTransform(m *algebra.Matrix) {
	s.transform = m
}

//SetMaterial sets the Sphere's material
func (s *Sphere) SetMaterial(m *canvas.Material){
	s.material = m
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

//NewIntersections creates a new intersection data type
func NewIntersections() *Intersections {
	hits := make(map[Shape]map[*algebra.Ray]*datatypes.MinHeap)
	ref := make(map[Shape]map[*algebra.Ray]*datatypes.MinHeap)
	return &Intersections{hits: hits, ref: ref}
}

//Intersect Updates intersections of a Sphere with the given algebra.Ray
func (intersections *Intersections) Intersect(s *Sphere, r *algebra.Ray) error {
	r2 := r.Transform(s.transform.Inverse())
	got := r2.Get()
	origin := got["origin"]
	direction := got["direction"]
	sphereToRay, err := origin.Subtract(s.origin)
	if err != nil {
		return err
	}

	a, err := algebra.DotProduct(direction, direction)
	if err != nil {
		return err
	}
	if a == 0 {
		return algebra.ZeroDivide(0)
	}

	b, err := algebra.DotProduct(direction, sphereToRay)
	if err != nil {
		return err
	}
	b = 2 * b

	c, err := algebra.DotProduct(sphereToRay, sphereToRay)
	if err != nil {
		return err
	}
	c = c - 1
	discriminant := math.Pow(b, 2) - (4 * a * c)

	if discriminant < 0 { // No rays intersect the sphere
		return nil
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)

	if (t1 > 0 || t2 > 0) && intersections.hits[s] == nil {
		res := make(map[*algebra.Ray]*datatypes.MinHeap)
		intersections.hits[s] = res
		newHeap := datatypes.NewMinHeap()
		intersections.hits[s][r] = newHeap
	}

	if (t1 < 0 || t2 < 0) && intersections.ref[s] == nil {
		res := make(map[*algebra.Ray]*datatypes.MinHeap)
		intersections.ref[s] = res
		newHeap := datatypes.NewMinHeap()
		intersections.ref[s][r] = newHeap
	}
	if t1 >= 0 {
		intersections.hits[s][r].Push(t1)
	} else {
		intersections.ref[s][r].Push(t1)
	}
	if t2 >= 0 {
		intersections.hits[s][r].Push(t2)
	} else {
		intersections.ref[s][r].Push(t2)
	}

	return nil
}

//Count returns the number of intersections of a given algebra.Ray with the Sphere
func (intersections *Intersections) Count(s Shape, r *algebra.Ray) int {
	var numHits int
	var numRef int
	if intersections.hits[s] == nil || intersections.hits[s][r] == nil {
		numHits = 0
	} else {
		numHits = len(intersections.hits[s][r].Get())
	}

	if intersections.ref[s] == nil || intersections.ref[s][r] == nil {
		numRef = 0
	} else {
		numRef = len(intersections.ref[s][r].Get())
	}
	return numHits + numRef
}

//GetIntersections returns the slice of values that intersect with the Sphere for the give algebra.Ray
func (intersections *Intersections) GetIntersections(s Shape, r *algebra.Ray) []float64 {
	var hitHeap []float64
	var refHeap []float64
	if intersections.hits[s] == nil || intersections.hits[s][r] == nil {
		hitHeap = []float64{}
	} else {
		hitHeap = intersections.hits[s][r].Get()
	}
	if intersections.ref[s] == nil || intersections.ref[s][r] == nil {
		refHeap = []float64{}
	} else {
		refHeap = intersections.ref[s][r].Get()
	}
	return append(refHeap, hitHeap...)
}

//Hit returns the minimum positive value of a ray intersecting the given object
func (intersections *Intersections) Hit(s Shape, r *algebra.Ray) (float64, bool) {
	if intersections.hits[s] == nil || intersections.hits[s][r] == nil {
		return 0.0, false
	} else {
		return intersections.hits[s][r].GetMin(), true
	}
}
