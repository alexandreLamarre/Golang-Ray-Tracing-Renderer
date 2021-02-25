package geometry

import "testing"

func TestMinHeap(t *testing.T) {

	h := NewMinHeap()
	if len(h.container) != 0 {
		t.Errorf("invalid starting size of heap")
	}
	s := NewSphere(nil)
	i1 := NewIntersection(s, 5)
	i2 := NewIntersection(s, 3)
	i3 := NewIntersection(s, 7)
	i4 := NewIntersection(s, 8)
	i5 := NewIntersection(s, 1)
	h.Push(i1)

	if len(h.container) != 1 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 1)
	}

	if h.GetMin().T != 5.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 5.0)
	}

	h.Push(i2)
	h.Push(i3)
	h.Push(i4)
	h.Push(i5)

	if len(h.container) != 5 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 5)
	}
	if h.GetMin().T != 1.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 1.0)
	}

	i6 := NewIntersection(s, -1)
	i7 := NewIntersection(s, 1)
	i8 := NewIntersection(s, 3)
	i9 := NewIntersection(s, 4)

	h.PushAll(i6, i7, i8, i9)

	if len(h.container) != 9 {
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 9)
	}

	if h.GetMin().T != -1.0 {
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), -1.0)
	}

	l := len(h.container)
	v := h.ExtractMin()
	if v.T != -1.0{
		t.Errorf("invalid extracted minimum of heap %f, expected %f", h.GetMin(), -1.0)
	}
	if len(h.container) != l -1{
		t.Errorf("Heap did not shrink in size after extract min")
	}

	l = len(h.container)
	v = h.ExtractMin()
	if v.T != 1.0 {
		t.Errorf("invalid extracted minimum of heap %f, expected %f", h.GetMin(), 1.0)
	}
	if len(h.container) != l -1{
		t.Errorf("Heap did not shrink in size after extract min")
	}
}

func TestMinHeap_Copy(t *testing.T) {
	s := NewSphere(nil)
	h := NewMinHeap()
	h.Push(NewIntersection(s, 5 ))

	h1 := h.Copy()
	if h1.GetMin().T != h.GetMin().T{
		t.Errorf("Expected copies to share the same minimum")
	}

	h1.Push(NewIntersection(s, 4))
	if h.GetMin().T == h1.GetMin().T{
		t.Errorf("Expected copes to not share the same minimum")
	}
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
