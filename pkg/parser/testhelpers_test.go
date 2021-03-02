package parser

import (
	"fmt"
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

func equals(a, b float64) bool {
	EPSILON := 0.0001
	return math.Abs(a-b) < EPSILON ||
		(math.IsInf(a, 1) && math.IsInf(b, 1)) || (math.IsInf(a, -1) && math.IsInf(b, -1))
}