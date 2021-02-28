package geometry

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
	"testing"
)

func testVectorEquals(t *testing.T, values, results []float64) {
	if len(values) != len(results) {
		fmt.Println(values, results)
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
	EPSILON := 0.0001
	return math.Abs(a-b) < EPSILON || (math.IsInf(a, 1) && math.IsInf(b, 1)) || (math.IsInf(a, -1) && math.IsInf(b, -1))
}

func assertEquals(t *testing.T, got, expected float64) {
	if got != expected {
		t.Errorf("Expected %f, Got: %f", expected, got)
	}
}

func testMaterialEquals(t *testing.T, m *canvas.Material, expected *canvas.Material){
	if m.Diffuse != expected.Diffuse{
		t.Errorf("mistmatched diffuse")
	}
	if m.Specular != expected.Specular{
		t.Errorf("mismatched specular")
	}

	if m.Shininess != expected.Shininess{
		t.Errorf("mismatched shininess")
	}

	if m .RefractiveIndex != expected.RefractiveIndex{
		t.Errorf("mismatched refractive index")
	}

	if m.Reflective != expected.Reflective{
		t.Errorf("mismtached reflective")
	}

	if m.Transparency != expected.Transparency{
		t.Errorf("mismtached transparency")
	}

	if m.Ambient != expected.Ambient{
		t.Errorf("mismatched ambient")
	}
	if m.Color.Red() != expected.Color.Red(){
		t.Errorf("mismatched red colors")
	}
	if m.Color.Green() != expected.Color.Green(){
		t.Errorf("mismatched green colors")
	}
	if m.Color.Blue() != expected.Color.Blue(){
		t.Errorf("mismatched blue colors")
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