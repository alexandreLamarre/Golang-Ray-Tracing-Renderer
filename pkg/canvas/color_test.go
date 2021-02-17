package canvas

import (
	"testing"
)

func TestColor_Add(t *testing.T) {
	c1 := &Color{0.9, 0.6, 0.75}
	c2 := &Color{0.7, 0.1, 0.25}

	c3 := c1.Add(c2)
	res := &Color{1.6, 0.7, 1.0}

	testVectorEquals(t, c3, res)
}

func TestColor_Subtract(t *testing.T) {

	c1 := &Color{0.9, 0.6, 0.75}
	c2 := &Color{0.7, 0.1, 0.25}

	c3 := c1.Subtract(c2)
	res := &Color{0.2, 0.5, 0.5}

	testVectorEquals(t, c3, res)

}

func TestColor_ScalarMult(t *testing.T) {
	c1 := &Color{0.2, 0.3, 0.4}
	c := 2.0
	c1 = c1.ScalarMult(c)

	res := &Color{0.4, 0.6, 0.8}

	testVectorEquals(t, c1, res)
}

func TestColor_Multiply(t *testing.T) {
	c1 := &Color{1.0, 0.2, 0.4}
	c2 := &Color{0.9, 1, 0.1}

	c3 := Multiply(c1, c2)

	res := &Color{0.9, 0.2, 0.04}

	testVectorEquals(t, c3, res)
}

func equals(a, b float64) bool {
	EPSILON := 0.00001
	return a-b < EPSILON || b-a < EPSILON
}

func testVectorEquals(t *testing.T, values, results *Color) {
	for i, v := range results {
		if !equals(values[i], v) {
			t.Errorf("Expected %g, Got: %g", v, values[i])
		}
	}
}
