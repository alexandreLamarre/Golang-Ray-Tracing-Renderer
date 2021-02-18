package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	s := NewSphere()
	origin := []float64{0, 0, 0, 1}
	testVectorEquals(t, s.origin.Get(), origin)
	radius := 1.0
	if !equals(radius, s.radius) {
		t.Errorf("Expected radius %f, Got: %f", radius, s.radius)
	}

}

func TestNewIntersections(t *testing.T) {
	is := NewIntersections()
	if len(is.hits) != 0 {
		t.Errorf("Expected size of default intersections to be 0")
	}
}

func TestSphere_Intersect(t *testing.T) {
	//intersect at tangent
	is := NewIntersections()
	s := NewSphere()
	r := algebra.NewRay(0, 1, -5, 0, 0, 1)
	err := is.Intersect(s, r)

	if err != nil {
		t.Errorf("%s", err)
	}

	if is.Count(s, r) != 2 {
		t.Errorf("Expected %d number of intersections, Got: %d", 2, is.Count(s, r))
	}

	v := is.GetIntersections(s, r)
	if !equals(v[0], 5.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", 5.0, v[0])
	}
	if !equals(v[1], 5.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", 5.0, v[0])
	}

	val, sucess := is.Hit(s, r)

	if !sucess {
		t.Errorf("Expected ray %v to hit sphere %v, %f", r.Get(), s.origin, s.radius)
	}

	if val != 5.0 {
		t.Errorf("Expected hit value to be %f, Got : %f", 5.0, val)
	}

	//misses sphere
	is = NewIntersections()
	r = algebra.NewRay(0, 2, -5, 0, 0, 1)
	err = is.Intersect(s, r)
	if err != nil {
		t.Errorf("%s", err)
	}
	if is.Count(s, r) != 0 {
		t.Errorf("Expected %d intersections, got : %d", 0, is.Count(s, r))
	}

	// ray inside a sphere
	r = algebra.NewRay(0, 0, 0, 0, 0, 1)
	err = is.Intersect(s, r)
	if err != nil {
		t.Errorf("%s", err)
	}
	if is.Count(s, r) != 2 {
		t.Errorf("Expected %d intersections, got: %d", 2, is.Count(s, r))
	}
	v = is.GetIntersections(s, r)
	if !equals(v[0], -1.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", -1.0, v[0])
	}
	if !equals(v[1], 1.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", 1.0, v[1])
	}

	val, sucess = is.Hit(s, r)

	if !sucess {
		t.Errorf("Expected ray %v to hit sphere %v, %f", r.Get(), s.origin, s.radius)
	}

	if val != 1.0 {
		t.Errorf("Expected hit value to be %f, Got : %f", 1.0, val)
	}

	// sphere behind ray

	is = NewIntersections()
	r = algebra.NewRay(0, 0, 5, 0, 0, 1)
	err = is.Intersect(s, r)
	if err != nil {
		t.Errorf("%s", err)
	}
	v = is.GetIntersections(s, r)
	if !equals(v[0], -6.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", -6.0, v[0])
	}
	if !equals(v[1], -4.0) {
		t.Errorf("Expected ray to intersect at %f, Got: %f", -4.0, v[1])
	}

	val, sucess = is.Hit(s, r)

	if sucess {
		t.Errorf("Expected ray %v to not hit sphere %v, %f", r.Get(), s.origin, s.radius)
	}

}

func testVectorEquals(t *testing.T, values, results []float64) {
	for i, v := range results {
		if !equals(values[i], v) {
			t.Errorf("Expected %g, Got: %g", v, values[i])
		}
	}
}

func equals(a, b float64) bool {
	EPSILON := 0.00001
	return math.Abs(a-b) < EPSILON
}
