package canvas

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
	"testing"
)

func TestNewPointLight(t *testing.T) {
	l := NewPointLight(&Color{1, 1, 1}, algebra.NewVector(0, 0, 0))

	testRealVectorEquals(t, l.Position.Get(), []float64{0, 0, 0, 0})
	testVectorEquals(t, l.Intensity, &Color{1, 1, 1})
}

func testRealVectorEquals(t *testing.T, vector []float64, expected []float64) {
	if len(vector) != len(expected) {
		t.Errorf("Mismatched vector lengths, Expected : %d, Got: %d", len(expected), len(vector))
	}

	for i := 0; i < len(vector); i++ {
		if !equals(vector[i], expected[i]) {
			t.Errorf("Expected value at %d : %f, Got: %f", i, expected[i], vector[i])
		}
	}
}

func TestLighting(t *testing.T) {
	p := algebra.NewPoint(0, 0, 0)
	m := NewDefaultMaterial()

	eyeVector := algebra.NewVector(0, 0, -1)
	normalVector := algebra.NewVector(0, 0, -1)
	light := NewPointLight(&Color{1, 1, 1}, algebra.NewPoint(0, 0, -10))

	color := Lighting(m, light, p, eyeVector, normalVector)
	testVectorEquals(t, color, &Color{1.9, 1.9, 1.9})

	eyeVector = algebra.NewVector(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	color = Lighting(m, light, p, eyeVector, normalVector)
	testVectorEquals(t, color, &Color{1.0, 1.0, 1.0})

	eyeVector = algebra.NewVector(0, 0, -1)
	normalVector = algebra.NewVector(0, 0, -1)
	light = NewPointLight(&Color{1.0, 1.0, 1.0}, algebra.NewPoint(0, 10, -10))

	color = Lighting(m, light, p, eyeVector, normalVector)
	testVectorEquals(t, color, &Color{0.7364, 0.7634})

	eyeVector = algebra.NewVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2)
	color = Lighting(m, light, p, eyeVector, normalVector)
	testVectorEquals(t, color, &Color{1.6364, 1.6364, 1.6364})

	eyeVector = algebra.NewVector(0, 0, -1)
	normalVector = algebra.NewVector(0, 0, -1)
	light = NewPointLight(&Color{1, 1, 1}, algebra.NewPoint(0, 0, 10))

	color = Lighting(m, light, p, eyeVector, normalVector)
	testVectorEquals(t, color, &Color{0.1, 0.1, 0.1})
}
