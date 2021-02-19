package algebra

import (
	"testing"
)

func TestNewRay(t *testing.T) {
	r := NewRay(1, 2, 3, 4, 5, 6)
	testRayEquals(t, r, []float64{1, 2, 3, 1}, []float64{4, 5, 6, 0})

}

func testRayEquals(t *testing.T, r *Ray, origin []float64, direction []float64) {
	i := r.Get()
	gotOrigin := i["origin"]
	gotDirection := i["direction"]
	testVectorEquals(t, gotOrigin.Get(), origin)
	testVectorEquals(t, gotDirection.Get(), direction)
}

func TestRay_Position(t *testing.T) {
	r := NewRay(2, 3, 4, 1, 0, 0)
	p := r.Position(0)
	res := []float64{2, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)

	p = r.Position(1)
	res = []float64{3, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)

	p = r.Position(-1)
	res = []float64{1, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)

	p = r.Position(2.5)
	res = []float64{4.5, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)
}

func TestRay_Transform(t *testing.T) {
	r := NewRay(1,2,3,0,1,0)
	m := TranslationMatrix(3,4,5)
	r2 := r.Transform(m)

	v:= r2.Get()
	origin := []float64{4, 6, 8, 1}
	testVectorEquals(t, v["origin"].Get(), origin)
	direction := []float64{0, 1, 0, 0}
	testVectorEquals(t, v["direction"].Get(), direction)

	r = NewRay(1, 2, 3, 0, 1, 0)
	m = ScalingMatrix(2, 3, 4)
	r2 = r.Transform(m)
	v = r2.Get()
	origin = []float64{2, 6, 12}
	testVectorEquals(t, v["origin"].Get(), origin)
	direction = []float64{0, 3, 0}
	testVectorEquals(t, v["direction"].Get(), direction)
}