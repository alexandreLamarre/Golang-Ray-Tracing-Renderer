package datatypes

import (
	"fmt"
	"testing"
)

func TestMinHeap(t *testing.T) {

	h := NewMinHeap()
	if len(h.container) != 0{
		t.Errorf("invalid starting size of heap")
	}

	h.Push(5)

	if len(h.container) != 1{
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 1)
	}

	if h.GetMin() != 5.0{
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 5.0)
	}

	h.Push(3)
	h.Push(7)
	h.Push(8)
	h.Push(1)

	if len(h.container) != 5{
		t.Errorf("invalid size of heap %d, expected: %d", len(h.container), 5)
	}
	fmt.Println(h.container)
	if h.GetMin() != 1.0{
		t.Errorf("invalid minimum of heap %f, expected %f", h.GetMin(), 1.0)
	}
}

