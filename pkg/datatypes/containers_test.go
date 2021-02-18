package datatypes

import (
	"testing"
)

func TestMinHeap(t *testing.T) {

	h := NewMinHeap()
	if len(h.container) != 0 {
		t.Errorf("invalid starting size of heap")
	}

	h.Push(5)

	if len(h.container) != 1 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 1)
	}

	if h.GetMin() != 5.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 5.0)
	}

	h.Push(3)
	h.Push(7)
	h.Push(8)
	h.Push(1)

	if len(h.container) != 5 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 5)
	}
	if h.GetMin() != 1.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 1.0)
	}

	v := h.Get()
	res := []float64{1, 3, 7, 8, 5}
	testSliceEquals(t, v, res)

	h.PushAll(-1, 1, 3, 4)

	if len(h.container) != 9 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 9)
	}

	if h.GetMin() != -1.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), -1.0)
	}

	res = []float64{-1, 1, 1, 4, 5, 3, 7, 3, 8}

	testSliceEquals(t, h.Get(), res)
}

func testSliceEquals(t *testing.T, values []float64, expected []float64) {
	if len(values) != len(expected) {
		t.Errorf("Mismatches lengths: Expected: %d, Got: %d", len(expected), len(values))
	}
	for i := 0; i < len(values); i++ {
		if values[i] != expected[i] {
			t.Errorf("Expected value at %d : %f, Got: %f", i, expected[i], values[i])
		}
	}
}
