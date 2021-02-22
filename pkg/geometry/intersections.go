package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
)

//Intersections data type keeps track of t values of the intersections of rays with a sphere
type Intersections struct {
	hits *MinHeap // hits on contours of objects
	ref  *MinHeap // used in ray reflections/refractions
}

//Intersection keeps track of a ray's position and the object it intersects
type Intersection struct {
	Object Shape
	T      float64
}

//NewIntersection creates a new intersection of an object with a ray's position value
func NewIntersection(s Shape, t float64) *Intersection {
	return &Intersection{Object: s, T: t}
}

//NewIntersections creates a new intersection data type
func NewIntersections() *Intersections {
	hits := NewMinHeap()
	ref := NewMinHeap()
	return &Intersections{hits: hits, ref: ref}
}

//Intersect Updates intersections of a Sphere with the given algebra.Ray
func (intersections *Intersections) Intersect(s Shape, r *algebra.Ray) error {
	m := s.GetTransform()
	r2 := r.Transform(m.Inverse())

	t1, t2, intersected := s.LocalIntersect(r2)
	if !intersected{
		return nil
	}
	i1 := NewIntersection(s, t1)
	i2 := NewIntersection(s, t2)

	if t1 >= 0 {
		intersections.hits.Push(i1)
	} else {
		intersections.ref.Push(i1)
	}
	if t2 >= 0 {
		intersections.hits.Push(i2)
	} else {
		intersections.ref.Push(i2)
	}
	return nil
}

//Count returns the number of intersections of a given algebra.Ray with the Sphere
func (intersections *Intersections) Count() int {
	numHits := len(intersections.hits.Get())
	numRef := len(intersections.ref.Get())
	return numHits + numRef
}

//GetIntersections returns the slice of values that intersect with the Sphere for the give algebra.Ray
func (intersections *Intersections) GetIntersections() []*Intersection {
	hitHeap := intersections.hits.Get()
	refHeap := intersections.ref.Get()

	return append(refHeap, hitHeap...)
}

//Hit returns the minimum positive value of a ray intersecting the given object
func (intersections *Intersections) Hit() *Intersection {
	if len(intersections.hits.Get()) == 0 {
		return nil
	} else {
		return intersections.hits.GetMin()
	}
}
