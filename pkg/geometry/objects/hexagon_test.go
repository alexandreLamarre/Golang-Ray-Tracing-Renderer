package objects

import "testing"

func TestNewHexagon(t *testing.T) {
	h := NewHexagon(nil)
	if len(h.GetShapes()) != 6{
		t.Errorf("Expected 6 sides for a hexagon, Got %d", len(h.GetShapes()))
	}
}