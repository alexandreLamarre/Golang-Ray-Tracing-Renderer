package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
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
	comps := PrepareComputations(i, r, nil)
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
	comps = PrepareComputations(i, r, nil)
	resPoint = []float64{0, 0, 1, 1}
	resEye = []float64{0, 0, -1, 0}
	resNormal = []float64{0, 0, -1, 0}
	testVectorEquals(t, comps.Point.Get(), resPoint)
	testVectorEquals(t, comps.Eye.Get(), resEye)
	testVectorEquals(t, comps.Normal.Get(), resNormal)
	if comps.Inside == false {
		t.Errorf("Expected eye vector to be inside shape")
	}

	//Test reflection vector
	plane := NewPlane(nil)
	r = algebra.NewRay(0,1,-1,0, -math.Sqrt(2)/2, math.Sqrt(2)/2)
	i = NewIntersection(plane, math.Sqrt(2))
	comps = PrepareComputations(i, r, nil)
	testVectorEquals(t, comps.Reflect.Get(), []float64{0, math.Sqrt(2)/2, math.Sqrt(2)/2, 0.0})
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
	comps := PrepareComputations(i, r, nil)
	c := w.ShadeHit(*comps, 0)
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
	comps = PrepareComputations(i, r, nil)
	c = w.ShadeHit(*comps, 0)
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

	//shadow test
	lights := make([]*canvas.PointLight, 0 , 0)
	l := canvas.NewPointLight(&canvas.Color{1,1,1}, algebra.NewPoint(0, 0, -10))
	lights = append(lights, l)
	objs := make([]Shape, 0, 0)
	s1 := NewSphere(nil)
	s2 := NewSphere(algebra.TranslationMatrix(0,0, 10))
	objs = append(objs, s1, s2)

	w = &World{Lights: lights, Objects: objs}
	ray := algebra.NewRay(0,0,5,0,0,1)
	i = NewIntersection(s, 4)
	comps = PrepareComputations(i, ray, nil)
	c = w.ShadeHit(*comps, 0)
	res = 0.1
	if !equals(c.Red(), res){
		t.Errorf("Expected %f, Got %f", res, c.Red())
	}
	if !equals(c.Green(), res){
		t.Errorf("Expected %f, Got %f", res, c.Green())
	}
	if !equals(c.Blue(), res){
		t.Errorf("Expected %f, Got %f", res, c.Blue())
	}


	w = NewDefaultWorld()
	shape := NewPlane(algebra.TranslationMatrix(0, -1, 0))
	m := canvas.NewDefaultMaterial()
	m.Reflective = 0.5
	shape.SetMaterial(m)
	w.Objects = append(w.Objects, shape)
	r = algebra.NewRay(0, 0 , -3, 0, -math.Sqrt(2)/2, math.Sqrt(2)/2)
	i = NewIntersection(shape, math.Sqrt(2))
	comps = PrepareComputations(i, r, nil)
	color := w.ShadeHit(*comps, 1)
	testColorEquals(t, color, &canvas.Color{0.87677, 0.92436, 0.82918})
}

func TestWorld_ColorAt(t *testing.T) {
	w := NewDefaultWorld()
	r := algebra.NewRay(0, 0, -5, 0, 1, 0)
	c := w.ColorAt(r, 0)
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
	c = w.ColorAt(r, 0)
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
	c = w.ColorAt(r, 0)
	for i := 0; i < 3; i++ {
		if !equals(c[i], inner.GetMaterial().Color[i]) {
			t.Errorf("Expected: %f, Got: %f", c[i], inner.GetMaterial().Color[i])
		}
	}


	// test that an "infinite recursion" terminates
	lights := make([]*canvas.PointLight, 0 ,0)
	light1 := canvas.NewPointLight(&canvas.Color{1,1,1}, algebra.NewPoint(0,0,0))
	lower := NewPlane(algebra.TranslationMatrix(0, -1, 0))
	m := canvas.NewDefaultMaterial()
	m.Reflective = 1.0
	lower.SetMaterial(m)
	upper := NewPlane(algebra.TranslationMatrix(0, 1, 0))
	upper.SetMaterial(m)
	objs := make([]Shape, 0, 0)
	objs = append(objs, lower, upper)
	lights = append(lights, light1)
	w = &World{Lights: lights, Objects: objs}
	r = algebra.NewRay(0,0,0, 0, 1, 0)
	w.ColorAt(r, 10)


}

func TestWorld_ReflectedColor(t *testing.T) {
	w := NewDefaultWorld()
	r := algebra.NewRay(0,0,0, 0,0,1)
	shape := w.Objects[1]
	m := shape.GetMaterial()
	m.Ambient = 1.0
	shape.SetMaterial(m)
	i := NewIntersection(shape, 1)
	comps := PrepareComputations(i, r, nil)
	color := w.ReflectedColor(comps, 1)
	testColorEquals(t, color, &canvas.Color{0,0,0})

	w = NewDefaultWorld()
	shape = NewPlane(algebra.TranslationMatrix(0, -1, 0))
	m = canvas.NewDefaultMaterial()
	m.Reflective = 0.5
	shape.SetMaterial(m)
	w.Objects = append(w.Objects, shape)
	r = algebra.NewRay(0, 0 , -3, 0, -math.Sqrt(2)/2, math.Sqrt(2)/2)
	i = NewIntersection(shape, math.Sqrt(2))
	comps = PrepareComputations(i, r, nil)
	color = w.ReflectedColor(comps, 1)
	testColorEquals(t, color, &canvas.Color{0.19032, 0.2379, 0.14274})
}

//Test prepare computations for N1, N2 refractive indexes
func TestRefractiveComputations(t *testing.T){
	A := NewGlassSphere(algebra.ScalingMatrix(2,2,2), 1.5)
	B := NewGlassSphere(algebra.TranslationMatrix(0, 0, -0.25), 2.0)
	C := NewGlassSphere(algebra.TranslationMatrix(0, 0, 0.25), 2.5)
	r := algebra.NewRay(0, 0, -4, 0, 0, 1)

	xs := NewIntersections()
	i1 := NewIntersection(A, 2)
	i2 := NewIntersection(B, 2.75)
	i3 := NewIntersection(C, 3.25)
	i4 := NewIntersection(B, 4.75)
	i5 := NewIntersection(C, 5.25)
	i6 := NewIntersection(A, 6)
	xs.hits.PushAll(i1, i2, i3, i4, i5, i6)

	comps := PrepareComputations(i1, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 1.0, 1.5)

	comps = PrepareComputations(i2, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 1.5, 2.0)

	comps = PrepareComputations(i3, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 2.0, 2.5)

	comps = PrepareComputations(i4, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 2.5, 2.5)

	comps = PrepareComputations(i5, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 2.5, 1.5)

	comps = PrepareComputations(i6, r, xs)
	testRefractiveIndexes(t, comps.N1, comps.N2, 1.5, 1.0)


}

func TestCompsUnderPoint(t *testing.T){
	r := algebra.NewRay(0, 0, -5, 0, 0, 1)
	shape := NewGlassSphere(algebra.TranslationMatrix(0,0, 1), 1.5)
	i := NewIntersection(shape, 5)
	xs := NewIntersections()
	xs.hits.Push(i)
	comps := PrepareComputations(i, r, xs)
	EPSILON := 0.00001
	if !(comps.UnderPoint.Get()[2] > EPSILON/2){// z coord
		t.Errorf("Unexpected UnderPoint z-coordinate: %f", comps.UnderPoint.Get()[2])
	}
	if !(comps.Point.Get()[2] < comps.UnderPoint.Get()[2]){
		t.Errorf("Unexpected relative positioning of z-intersect with z-underpoint intersect: %f versus %f",
			comps.Point.Get()[2], comps.UnderPoint.Get()[2])
	}

}

func TestWorld_RefractedColor(t *testing.T) {
	w := NewDefaultWorld()
	shape := w.Objects[0]
	r := algebra.NewRay(0, 0, -5, 0, 0, 1)
	xs := NewIntersections()
	i1 := NewIntersection(shape, 4.0)
	i2 := NewIntersection(shape, 6.0)
	xs.hits.PushAll(i1, i2)
	comps := PrepareComputations(i1, r, xs)
	c := w.RefractedColor(comps, 5.0)
	testColorEquals(t, c, &canvas.Color{0,0,0})

	m := shape.GetMaterial()
	m.Transparency = 0.9
	w.Objects[0].SetMaterial(m)
	xs = NewIntersections()
	i1 = NewIntersection(shape, 4.0)
	i2 = NewIntersection(shape, 6.0)
	xs.hits.PushAll(i1, i2)
	comps = PrepareComputations(i1, r, xs)
	c = w.RefractedColor(comps, 0.0)
	testColorEquals(t, c, &canvas.Color{0,0,0})

	// test total internal refraction case
	w = NewDefaultWorld()
	mat := w.Objects[0].GetMaterial()
	mat.Transparency = 1.0
	mat.RefractiveIndex = 1.5
	w.Objects[0].SetMaterial(mat)
	r = algebra.NewRay(0, 0, math.Sqrt(2)/2, 0, 1, 0)
	xs = NewIntersections()
	i1 = NewIntersection(w.Objects[0], -math.Sqrt(2)/2)
	i2 = NewIntersection(w.Objects[0], math.Sqrt(2)/2)
	xs.ref.Push(i1)
	xs.hits.Push(i2)
	comps = PrepareComputations(i2, r, xs)
	c = w.RefractedColor(comps, 5)
	testColorEquals(t, c, &canvas.Color{0,0,0})

	//test actual refraction

	w = NewDefaultWorld()
	mat1 := w.Objects[0].GetMaterial()
	mat1.Ambient = 1.0
	mat1.Pattern = canvas.TestPattern()
	w.Objects[0].SetMaterial(mat1)
	mat2 := w.Objects[1].GetMaterial()
	mat2.Transparency = 1.0
	mat2.RefractiveIndex = 1.5
	w.Objects[1].SetMaterial(mat2)
	r = algebra.NewRay(0, 0, 0.1, 0, 1, 0)
	i1 = NewIntersection(w.Objects[0], -0.9899)
	i2 = NewIntersection(w.Objects[1], -0.4899)
	i3 := NewIntersection(w.Objects[1], 0.4899)
	i4 := NewIntersection(w.Objects[0], 0.9899)
	is := NewIntersections()
	is.ref.PushAll(i1, i2)
	is.hits.PushAll(i3, i4)
	comps = PrepareComputations(i3, r, is)
	c = w.RefractedColor(comps, 5)
	testColorEquals(t, c, &canvas.Color{0, 0.99888, 0.04725})

	//test ShadeHit with the stupid refraction
	w = NewDefaultWorld()
	floor := NewPlane(algebra.TranslationMatrix(0, -1, 0))
	matf := floor.GetMaterial()
	matf.Transparency = 0.5
	matf.RefractiveIndex = 1.5
	floor.SetMaterial(matf)
	w.Objects = append(w.Objects, floor)

	ball := NewSphere(algebra.TranslationMatrix(0, -3.5, -0.5))
	matb := ball.GetMaterial()
	matb.Color = &canvas.Color{1, 0, 0}
	matb.Ambient = 0.5
	ball.SetMaterial(matb)
	w.Objects = append(w.Objects, ball)

	r = algebra.NewRay(0, 0, -3, 0, -math.Sqrt(2)/2, math.Sqrt(2)/2)
	xs = NewIntersections()
	i := NewIntersection(floor, math.Sqrt(2))
	xs.hits.Push(i)
	comps = PrepareComputations(i, r, xs)
	color := w.ShadeHit(*comps, 5)
	testColorEquals(t, color, &canvas.Color{0.93642, 0.68642, 0.68642})
}

func TestSchlick(t *testing.T) {
	shape := NewGlassSphere(nil, 1.5)
	r := algebra.NewRay(0, 0, math.Sqrt(2)/2, 0, 1, 0)
	xs := NewIntersections()
	i1 := NewIntersection(shape, -math.Sqrt(2)/2)
	i2 := NewIntersection(shape, math.Sqrt(2)/2)
	xs.hits.PushAll(i1, i2)
	comps := PrepareComputations(i2, r, xs)
	reflectance := Schlick(comps)
	if !equals(reflectance, 1.0){
		t.Errorf("Expected reflectance %f, Got: %f", 1.0, reflectance)
	}

	r = algebra.NewRay(0, 0, 0, 0, 1, 0)
	xs = NewIntersections()
	i1 = NewIntersection(shape, -1)
	i2 = NewIntersection(shape, 1)
	xs.hits.PushAll(i1, i2)
	comps = PrepareComputations(i2, r, xs)
	reflectance = Schlick(comps)
	if !equals(reflectance, 0.04){
		t.Errorf("Expected reflectance %f, Got: %f", 0.04, reflectance)
	}

	r = algebra.NewRay(0, 0.99, -2, 0, 0, 1)
	xs = NewIntersections()
	i1 = NewIntersection(shape, 1.8589)
	xs.hits.Push(i1)
	comps = PrepareComputations(i1, r, xs)
	reflectance = Schlick(comps)
	if !equals(reflectance, 0.48873){
		t.Errorf("Expected reflectance %f, Got: %f", 0.48873, reflectance)
	}


	//test Shade Hit augmented with Schlick

	w := NewDefaultWorld()
	r = algebra.NewRay(0, 0, -3, 0, -math.Sqrt(2)/2, math.Sqrt(2)/2)
	floor := NewPlane(algebra.TranslationMatrix(0, -1, 0))
	matf := floor.GetMaterial()
	matf.Reflective = 0.5
	matf.Transparency = 0.5
	matf.RefractiveIndex = 1.5
	floor.SetMaterial(matf)
	w.Objects = append(w.Objects, floor)

	ball := NewSphere(algebra.TranslationMatrix(0, -3.5, -0.5))
	ballf := ball.GetMaterial()
	ballf.Color = &canvas.Color{1, 0, 0 }
	ballf.Ambient = 0.5
	ball.SetMaterial(ballf)
	w.Objects = append(w.Objects, ball)

	xs = NewIntersections()
	i1 = NewIntersection(floor, math.Sqrt(2))
	xs.hits.Push(i1)
	comps = PrepareComputations(i1, r, xs)
	color := w.ShadeHit(*comps, 5)
	testColorEquals(t, color, &canvas.Color{0.93391, 0.69643, 0.69243})
}

func TestWorld_PointIsShadowed(t *testing.T) {
	w := NewDefaultWorld()
	p := algebra.NewPoint(0, 10, 0)
	res := w.PointIsShadowed(p)
	if res{
		t.Errorf("Expected point %v to not be shadowed", p)
	}

	p = algebra.NewPoint(10, -10, 10)
	res = w.PointIsShadowed(p)
	if !res{
		t.Errorf("Expected point %v to be shadowed", p)
	}

	p = algebra.NewPoint(-20, 20, -20)
	res = w.PointIsShadowed(p)
	if res{
		t.Errorf("Expected point %v to not be shadowed", p)
	}

	p = algebra.NewPoint(-2, 2, -2)
	res = w.PointIsShadowed(p)
	if res{
		t.Errorf("Expected point %v to not be shadowed", p)
	}
}

func testRefractiveIndexes(t *testing.T, n1, n2, expected1, expected2 float64){
	if n1 != expected1 || n2 != expected2{
		t.Errorf("Expected  %f n1, Got: %f . Expected %f n2, Got: %f", expected1, n1, expected2, n2)
	}
}
