package algebra

import (
	"math"
	"testing"
)

func TestVector_IsPoint(t *testing.T) {
	v1 := NewPoint(0.0, 0.0, 1.0)
	if !v1.IsPoint() {
		t.Errorf("Expected %v to be a point", v1.tuple)
	}

	if v1.IsVector() {
		t.Errorf("Expected %v to not be a vector", v1.tuple)
	}
}

func TestVector_IsVector(t *testing.T) {
	v := NewVector(0.0, 1.0, 0.0)

	if v.IsPoint() {
		t.Errorf("Expected %v to not be a point", v.tuple)
	}

	if !v.IsVector() {
		t.Errorf("Expected %v to be a vector", v.tuple)
	}
}

func TestVector_Get(t *testing.T) {
	v := NewVector(1.0, 2.0, 3.0)

	testVectorEquals(t, v.Get(), []float64{1.0, 2.0, 3.0, 0.0})
}

func TestVector_Add(t *testing.T) {
	v1 := NewPoint(3, -2, 5)
	v2 := NewVector(-2, 3, 1)

	v3, err := v1.Add(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	res := []float64{1.0, 1.0, 6.0, 1.0}

	testVectorEquals(t, v3.Get(), res)

	if !v3.IsPoint() {
		t.Errorf("Expected v3 to be a point")
	}
}

func TestVector_Subtract(t *testing.T) {
	v1 := NewPoint(3.0, 2.0, 1.0)
	v2 := NewPoint(5.0, 6.0, 7.0)

	res := []float64{-2, -4, -6, 0.0}

	v3, err := v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	testVectorEquals(t, v3.Get(), res)
	if !v3.IsVector() {
		t.Errorf("Expected v3 to be a vector")
	}

	v1 = NewPoint(3.0, 2.0, 1.0)
	v2 = NewVector(5.0, 6.0, 7.0)

	res = []float64{-2, -4, -6, 1.0}

	v3, err = v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	testVectorEquals(t, v3.Get(), res)

	if !v3.IsPoint() {
		t.Errorf("Expected v3 to be a vector")
	}

	v1 = NewVector(3.0, 2.0, 1.0)
	v2 = NewVector(5.0, 6.0, 7.0)

	res = []float64{-2, -4, -6, 0.0}

	v3, err = v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	testVectorEquals(t, v3.Get(), res)

	if !v3.IsVector() {
		t.Errorf("Expected v3 to be a vector")
	}

}

func TestVector_Negate(t *testing.T) {
	v1 := &Vector{[]float64{1, -2, 3, 1}}
	v1 = v1.Negate()
	res := []float64{-1, 2, -3, 1}

	testVectorEquals(t, v1.Get(), res)
}

func TestVector_MultScalar(t *testing.T) {
	v1 := NewVector(1, -2, 3)
	c := 3.5
	res := []float64{3.5, -7, 10.5, 0}

	v1 = v1.MultScalar(c)

	testVectorEquals(t, v1.Get(), res)

	v1 = NewVector(1, -2, 3)
	c = 0.5
	res = []float64{0.5, -1, 1.5, 0}

	v1 = v1.MultScalar(c)

	testVectorEquals(t, v1.Get(), res)
}

func TestVector_DivideScalar(t *testing.T) {

	v1 := NewVector(1, -2, 3)
	c := 2.0

	res := []float64{0.5, -1, 1.5, 0}

	v1 = v1.DivideScalar(c)
	for i, v := range res {
		if !equals(v1.Get()[i], v) {
			t.Errorf("Expected %g, got %g", v, v1.Get()[i])
		}
	}

}

func TestVector_Magnitude(t *testing.T) {

	v1 := NewVector(1, 0, 0)

	res := 1.0

	if !equals(v1.Magnitude(), res) {
		t.Errorf("Expected %g, got:%g", res, v1.Magnitude())
	}

	v1 = NewVector(0, 1, 0)

	if !equals(v1.Magnitude(), res) {
		t.Errorf("Expected %g, got:%g", res, v1.Magnitude())
	}

	v1 = NewVector(0, 0, 1)

	if !equals(v1.Magnitude(), res) {
		t.Errorf("Expected %g, got: %g", res, v1.Magnitude())
	}

	v1 = NewVector(-1, -2, -3)
	res = math.Sqrt(14)
	if !equals(v1.Magnitude(), res) {
		t.Errorf("Expected %g, got: %g", res, v1.Magnitude())
	}
}

func TestVector_Normalize(t *testing.T) {
	v1 := NewVector(4, 0, 0)
	v2, err := v1.Normalize()
	if err != nil {
		t.Errorf("Zero divide in Normalize method")
	}
	res := []float64{1.0, 0.0, 0.0, 0.0}

	testVectorEquals(t, v2.Get(), res)

	v1 = NewVector(1, 2, 3)
	v2, err = v1.Normalize()

	if err != nil {
		t.Errorf("Zero divide in Normalize method")
		return
	}

	res = []float64{1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14), 0.0}
	testVectorEquals(t, v2.Get(), res)

	if !equals(v2.Magnitude(), 1) {
		t.Errorf("Expected vector's magnitude to be %g, instead got: %g", 1.0, v2.Magnitude())
	}
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1.0, 2.0, 3.0)
	v2 := NewVector(2.0, 3.0, 4.0)

	res := 20.0
	dot, err := DotProduct(v1, v2)

	if err != nil {
		t.Logf("%s", err)
		t.Errorf("")
		return
	}
	if !equals(dot, res) {
		t.Errorf("Expected %g, Got: %g", dot, res)
	}
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector(1.0, 2.0, 3.0)
	v2 := NewVector(2.0, 3.0, 4.0)

	cross, err := CrossProduct(v1, v2)

	if err != nil {
		t.Logf("%s", err)
		t.Errorf("")
		return
	}

	res := []float64{-1, 2, -1}

	testVectorEquals(t, cross.Get(), res)

	v1 = NewVector(1.0, 2.0, 3.0)
	v2 = NewVector(2.0, 3.0, 4.0)

	cross, err = CrossProduct(v2, v1)

	if err != nil {
		t.Logf("%s", err)
		t.Errorf("")
		return
	}

	res = []float64{1, -2, 1}

	testVectorEquals(t, cross.Get(), res)
}


func testVectorEquals(t *testing.T, values, results []float64) {
	for i, v := range results {
		if !equals(values[i], v) {
			t.Errorf("Expected %g, Got: %g", v, values[i])
		}
	}
}
