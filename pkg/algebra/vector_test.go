package algebra

import (
	"testing"
)

func TestVector_IsPoint(t *testing.T) {
	v1 := NewPoint([]float64{0.0, 0.0, 1.0})
	if !v1.IsPoint() {
		t.Errorf("Expected %v to be a point", v1.tuple)
	}

	if v1.IsVector() {
		t.Errorf("Expected %v to not be a vector", v1.tuple)
	}
}

func TestVector_IsVector(t *testing.T) {
	v := NewVector([]float64{0.0, 1.0, 0.0})

	if v.IsPoint() {
		t.Errorf("Expected %v to not be a point", v.tuple)
	}

	if !v.IsVector() {
		t.Errorf("Expected %v to be a vector", v.tuple)
	}
}

func TestVector_Add(t *testing.T) {
	v1 := NewPoint([]float64{3, -2, 5})
	v2 := NewVector([]float64{-2, 3, 1})

	v3, err := v1.Add(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	res := []float64{1.0, 1.0, 6.0, 1.0}

	for i, v := range res {
		if v3.tuple[i] != v {
			t.Errorf("Expected %g, got: %g", v, v3.tuple[i])
		}
	}

	if !v3.IsPoint() {
		t.Errorf("Expected v3 to be a point")
	}
}

func TestVector_Subtract(t *testing.T) {
	v1 := NewPoint([]float64{3.0, 2.0, 1.0})
	v2 := NewPoint([]float64{5.0, 6.0, 7.0})

	res := []float64{-2, -4, -6, 0.0}

	v3, err := v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}
	t.Logf("%v", v3)
	for i, v := range res {
		if v3.tuple[i] != v {
			t.Errorf("Expected %g, got: %g", v, v3.tuple[i])
		}
	}

	if !v3.IsVector() {
		t.Errorf("Expected v3 to be a vector")
	}

	v1 = NewPoint([]float64{3.0, 2.0, 1.0})
	v2 = NewVector([]float64{5.0, 6.0, 7.0})

	res = []float64{-2, -4, -6, 1.0}

	v3, err = v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	for i, v := range res {
		if v3.tuple[i] != v {
			t.Errorf("Expected %g, got: %g", v, v3.tuple[i])
		}
	}

	if !v3.IsPoint() {
		t.Errorf("Expected v3 to be a vector")
	}

	v1 = NewVector([]float64{3.0, 2.0, 1.0})
	v2 = NewVector([]float64{5.0, 6.0, 7.0})

	res = []float64{-2, -4, -6, 0.0}

	v3, err = v1.Subtract(v2)

	if err != nil {
		t.Errorf("%s", err)
	}

	for i, v := range res {
		if v3.tuple[i] != v {
			t.Errorf("Expected %g, got: %g", v, v3.tuple[i])
		}
	}

	if !v3.IsVector() {
		t.Errorf("Expected v3 to be a vector")
	}

}
