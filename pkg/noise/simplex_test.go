package noise

import (
	"math/rand"
	"testing"
)

func TestSimplex(t *testing.T) {
	for i := 0; i < 100; i++ {
		noise := Simplex(rand.Float64(), rand.Float64(), rand.Float64(), 435817348970)
		if noise < -1 || noise > 1 {
			t.Errorf("Invalid noise range %f", noise)
		}
	}
}
