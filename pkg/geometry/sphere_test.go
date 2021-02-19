package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	s := NewSphere(nil)
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

func TestSphere_SetTransform(t *testing.T) {
	s := NewSphere(nil)
	m := algebra.TranslationMatrix(2, 3, 4)
	s.SetTransform(m)

	if !s.transform.Equals(m) {
		t.Errorf("Expected %v, Got: %v", m, s.transform)
	}

	s2 := NewSphere(m)

	if !s2.transform.Equals(m) {
		t.Errorf("Expected %v, got: %v", m, s2.transform)
	}
}

func TestSphere_SetMaterial(t *testing.T) {
	s := NewSphere(nil)
	m := s.material
	d := canvas.NewDefaultMaterial()
	testColorEquals(t, m.Color, d.Color)
	assertEquals(t, m.Ambient, d.Ambient)
	assertEquals(t, m.Diffuse, d.Diffuse)
	assertEquals(t, m.Specular, d.Specular)
	assertEquals(t, m.Shininess, m.Shininess)

	newMaterial := canvas.NewMaterial(&canvas.Color{1, 0, 0}, 1, 1, 1, 1)
	s.SetMaterial(newMaterial)
	m = s.material
	testColorEquals(t, m.Color, newMaterial.Color)
	assertEquals(t, m.Ambient, newMaterial.Ambient)
	assertEquals(t, m.Diffuse, newMaterial.Diffuse)
	assertEquals(t, m.Specular, newMaterial.Specular)
	assertEquals(t, m.Shininess, newMaterial.Shininess)
}

func TestSphere_NormalAt(t *testing.T) {
	// x axis
	s := NewSphere(nil)
	n := s.NormalAt(algebra.NewPoint(1, 0, 0))
	res := []float64{1, 0, 0, 0}
	testVectorEquals(t, n.Get(), res)

	//y axis
	s = NewSphere(nil)
	n = s.NormalAt(algebra.NewPoint(0, 1, 0))
	res = []float64{0, 1, 0, 0}
	testVectorEquals(t, n.Get(), res)

	//z axis
	s = NewSphere(nil)
	n = s.NormalAt(algebra.NewPoint(0, 0, 1))
	res = []float64{0, 0, 1, 0}
	testVectorEquals(t, n.Get(), res)

	//non-axial point
	s = NewSphere(nil)
	n = s.NormalAt(algebra.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	res = []float64{math.Sqrt(3) / 3, math.Sqrt(3) / 3, math.Sqrt(3) / 3, 0}
	testVectorEquals(t, n.Get(), res)

	//on a translated sphere
	s = NewSphere(algebra.TranslationMatrix(0, 1, 0))
	n = s.NormalAt(algebra.NewPoint(0, 1.70711, -0.70711))
	res = []float64{0.0, 0.70711, -0.70711, 0.0}
	testVectorEquals(t, n.Get(), res)

	//on a transformed sphere
	transform := algebra.Multiply(algebra.ScalingMatrix(1, 0.5, 1), algebra.RotationZ(math.Pi/5))
	s = NewSphere(transform)
	n = s.NormalAt(algebra.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2))
	res = []float64{0, 0.97014, -0.24254}
}

func TestSphere_Intersect(t *testing.T) {
	//intersect at tangent
	is := NewIntersections()
	s := NewSphere(nil)
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

	val, success := is.Hit(s, r)

	if !success {
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

	val, success = is.Hit(s, r)

	if !success {
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

	val, success = is.Hit(s, r)

	if success {
		t.Errorf("Expected ray %v to not hit sphere %v, %f", r.Get(), s.origin, s.radius)
	}

	// scaled sphere ray test
	r = algebra.NewRay(0, 0, -5, 0, 0, 1)
	s2 := NewSphere(algebra.ScalingMatrix(2, 2, 2))
	err = is.Intersect(s2, r)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if is.Count(s2, r) != 2 {
		t.Errorf("Expected %d number of intersections, Got: %d", 2, is.Count(s2, r))
	}

	v = is.GetIntersections(s2, r)
	if v[0] != 3.0 {
		t.Errorf("Expected ray to intersect at %f, Got: %f", 3.0, v[0])
	}
	if v[1] != 7.0 {
		t.Errorf("Expected ray to intersect at %f, Got: %f", 7.0, v[1])
	}

	//translated sphere ray test

	r = algebra.NewRay(0, 0, -5, 0, 0, 1)
	s3 := NewSphere(algebra.TranslationMatrix(5, 0, 0))
	err = is.Intersect(s3, r)
	if err != nil {
		t.Errorf("%s", err)
	}
	if is.Count(s3, r) != 0 {
		t.Errorf("Expected %d number of intersections, Got: %d", 0, is.Count(s3, r))
	}
}

func testVectorEquals(t *testing.T, values, results []float64) {
	if len(values) != len(results) {
		t.Errorf("Mimatched lengths: Expected %d, got: %d", len(results), len(values))
	}
	for i, v := range results {
		if !equals(values[i], v) {
			t.Errorf("Expected %g, Got: %g", v, values[i])
		}
	}
}

func testColorEquals(t *testing.T, values, results *canvas.Color) {
	if len(values) != len(results) {
		t.Errorf("Mimatched lengths: Expected %d, got: %d", len(results), len(values))
	}
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

func assertEquals(t *testing.T, got, expected float64) {
	if got != expected {
		t.Errorf("Expected %f, Got: %f", expected, got)
	}
}
