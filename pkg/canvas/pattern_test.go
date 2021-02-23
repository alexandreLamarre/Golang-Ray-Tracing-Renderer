package canvas

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"testing")

func TestStripePattern(t *testing.T) {
	white := &Color{1,1,1}
	black := &Color{0,0,0}
	pattern := StripePattern(white, black)
	testVectorEquals(t, pattern.a, &Color{1,1,1})
	testVectorEquals(t, pattern.b, &Color{0,0,0})

	c := pattern.GetPattern(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(0, 1, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(0, 2, 0))
	testVectorEquals(t, c, white)

	c = pattern.GetPattern(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(0, 0, 1))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(0, 0, 2))
	testVectorEquals(t, c, white)

	c = pattern.GetPattern(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(0.9, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetPattern(algebra.NewPoint(1.0, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetPattern(algebra.NewPoint(-0.1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetPattern(algebra.NewPoint(-1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetPattern(algebra.NewPoint(-1.1, 0, 0))
	testVectorEquals(t, c, white)
}

