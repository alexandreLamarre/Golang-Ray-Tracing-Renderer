package camera

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"math"
	"testing"
)

func TestNewDefaultCamera(t *testing.T) {
	c := NewDefaultCamera(160, 120, math.Pi/2)
	assertEquals(t, c.hSize, 160)
	assertEquals(t, c.vSize, 120)
	assertEquals(t, c.fov, math.Pi/2)
	testMatrixEquals(t, c.transform.Get(), algebra.IdentityMatrix(4).Get())

	c = NewDefaultCamera(200, 125, math.Pi/2)
	assertEquals(t, c.pixelSize, 0.01)

	c = NewDefaultCamera(125, 200, math.Pi/2)
	assertEquals(t, c.pixelSize, 0.01)
}

func TestNewCamera(t *testing.T) {
	m := algebra.TranslationMatrix(2, 3, 4)
	c, err := NewCamera(160, 120, math.Pi/2, m)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	assertEquals(t, c.hSize, 160)
	assertEquals(t, c.vSize, 120)
	assertEquals(t, c.fov, math.Pi/2)
	testMatrixEquals(t, c.transform.Get(), m.Get())
}

func TestCamera_RayForPixel(t *testing.T) {
	c := NewDefaultCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	testVectorEquals(t, r.Get()["origin"].Get(), []float64{0, 0, 0, 1})
	testVectorEquals(t, r.Get()["direction"].Get(), []float64{0, 0, -1, 0})

	r = c.RayForPixel(0, 0)
	testVectorEquals(t, r.Get()["origin"].Get(), []float64{0, 0, 0, 1})
	testVectorEquals(t, r.Get()["direction"].Get(), []float64{0.66519, 0.33259, -0.66851})

	c.transform = algebra.Multiply(algebra.RotationY(math.Pi/4), algebra.TranslationMatrix(0, -2, 5))
	r = c.RayForPixel(100, 50)
	testVectorEquals(t, r.Get()["origin"].Get(), []float64{0, 2, -5})
	testVectorEquals(t, r.Get()["direction"].Get(), []float64{math.Sqrt(2) / 2, 0, -math.Sqrt(2) / 2})
}

func TestCamera_Render(t *testing.T) {
	w := geometry.NewDefaultWorld()
	c, err := NewCamera(11, 11, math.Pi/2,
		algebra.ViewTransform(0, 0, -5, 0, 0, 0, 0, 1, 0))
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	image := c.Render(w)
	color := image.Pixels[5][5]
	if !equals(color.Red(), 0.38066) {
		t.Errorf("Incorrect red color %f, wanted %f", color.Red(), 0.38066)
	}
	if !equals(color.Green(), 0.47583) {
		t.Errorf("Incorrect green color %f, wanted %f", color.Green(), 0.47583)
	}
	if !equals(color.Blue(), 0.2855) {
		t.Errorf("Incorrect blue color %f, wanted %f", color.Blue(), 0.2855)
	}
}

func equals(a, b float64) bool {
	EPSILON := 0.0001
	return math.Abs(a-b) < EPSILON
}

func assertEquals(t *testing.T, got, expected float64) {
	if got != expected {
		t.Errorf("Expected %f, Got: %f", expected, got)
	}
}

func testMatrixEquals(t *testing.T, values [][]float64, expected [][]float64) {
	for i, _ := range values {
		for j, _ := range values[i] {
			if !equals(values[i][j], expected[i][j]) {
				t.Errorf("Expected: %f, Got: %f", expected[i][j], values[i][j])
			}
		}
	}
}

func testVectorEquals(t *testing.T, values, results []float64) {
	for i, v := range results {
		if !equals(values[i], v) {
			t.Errorf("Expected %g, Got: %g", v, values[i])
		}
	}
}
