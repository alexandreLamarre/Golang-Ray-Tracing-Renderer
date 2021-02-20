package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"testing"
)

func TestNewDefaultWorld(t *testing.T) {
	w := NewDefaultWorld()
	if len(w.Lights) != 1 {
		t.Errorf("Expected default world to have one light source")
	}
	if len(w.Objects) != 2 {
		t.Errorf("Expected default world to have two light sources")
	}

}

func TestPrepareComputations(t *testing.T) {
	s := NewSphere(nil)
	r := algebra.NewRay(0, 0, -5, 0, 0, 1)
	i := NewIntersection(s, 4.0)
	comps := PrepareComputations(i, r)
	assertEquals(t, i.T, comps.T)
	resPoint := []float64{0, 0, -1, 1}
	resEye := []float64{0, 0, -1, 0}
	resNormal := []float64{0, 0, -1, 0}
	testVectorEquals(t, comps.Point.Get(), resPoint)
	testVectorEquals(t, comps.Eye.Get(), resEye)
	testVectorEquals(t, comps.Normal.Get(), resNormal)
	if comps.Inside == true {
		t.Errorf("Expected eye vector to not be inside shape")
	}

	r = algebra.NewRay(0, 0, 0, 0, 0, 1)
	i = NewIntersection(s, 1.0)
	comps = PrepareComputations(i, r)
	resPoint = []float64{0, 0, 1, 1}
	resEye = []float64{0, 0, -1, 0}
	resNormal = []float64{0, 0, -1, 0}
	testVectorEquals(t, comps.Point.Get(), resPoint)
	testVectorEquals(t, comps.Eye.Get(), resEye)
	testVectorEquals(t, comps.Normal.Get(), resNormal)
	if comps.Inside == false {
		t.Errorf("Expected eye vector to be inside shape")
	}
}

func TestWorld_Intersect(t *testing.T) {
	w := NewDefaultWorld()
	ray := algebra.NewRay(0, 0, -5, 0, 0, 1)
	xs := w.Intersect(ray)
	if xs.Count() != 4 {
		t.Errorf("Expected %d ray intersections instead got: %d", 4, xs.Count())
	}

	v := xs.GetIntersections()
	assertEquals(t, v[0].T, 4.0)
	assertEquals(t, v[1].T, 6.0)
	assertEquals(t, v[2].T, 4.5)
	assertEquals(t, v[3].T, 5.5)

}

func TestWorld_ShadeHit(t *testing.T) {
	w := NewDefaultWorld()
	r := algebra.NewRay(0, 0, -5, 0, 0, 1)
	s := w.Objects[0]
	i := NewIntersection(s, 4.0)
	comps := PrepareComputations(i, r)
	c := w.ShadeHit(*comps)
	if !equals(c.Red(), 0.38066) {
		t.Errorf("Expected %f, Got %f", 0.38066, c.Red())
	}
	if !equals(c.Green(), 0.47583) {
		t.Errorf("Expected %f, Got %f", 0.47583, c.Green())
	}
	if !equals(c.Blue(), 0.2855) {
		t.Errorf("Expected %f, Got %f", 0.2855, c.Blue())
	}

	w.Lights[0] = canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(0, 0.25, 0))
	r = algebra.NewRay(0, 0, 0, 0, 0, 1)
	s = w.Objects[1]
	i = NewIntersection(s, 0.5)
	comps = PrepareComputations(i, r)
	c = w.ShadeHit(*comps)
	res := 0.90498
	if !equals(c.Red(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Red())
	}
	if !equals(c.Green(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Green())
	}
	if !equals(c.Blue(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Blue())
	}
}

func TestWorld_ColorAt(t *testing.T) {
	w := NewDefaultWorld()
	r := algebra.NewRay(0, 0, -5, 0, 1, 0)
	c := w.ColorAt(r)
	res := 0.0
	if !equals(c.Red(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Red())
	}
	if !equals(c.Green(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Green())
	}
	if !equals(c.Blue(), res) {
		t.Errorf("Expected %f, Got %f", res, c.Blue())
	}

	r = algebra.NewRay(0, 0, -5, 0, 0, 1)
	c = w.ColorAt(r)
	if !equals(c.Red(), 0.38066) {
		t.Errorf("Expected %f, Got %f", 0.38066, c.Red())
	}
	if !equals(c.Green(), 0.47583) {
		t.Errorf("Expected %f, Got %f", 0.47583, c.Green())
	}
	if !equals(c.Blue(), 0.2855) {
		t.Errorf("Expected %f, Got %f", 0.2855, c.Blue())
	}

	w = NewDefaultWorld()
	outer := w.Objects[0]
	mat := outer.GetMaterial()
	mat.Ambient = 1.0
	outer.SetMaterial(mat)
	inner := w.Objects[1]
	mat = inner.GetMaterial()
	mat.Ambient = 1.0
	inner.SetMaterial(mat)

	r = algebra.NewRay(0, 0, 0.75, 0, 0, -1)
	c = w.ColorAt(r)
	for i := 0; i < 3; i++ {
		if !equals(c[i], inner.GetMaterial().Color[i]) {
			t.Errorf("Expected: %f, Got: %f", c[i], inner.GetMaterial().Color[i])
		}
	}

}
