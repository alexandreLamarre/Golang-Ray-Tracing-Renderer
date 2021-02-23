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

	c := pattern.GetColor(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0, 1, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0, 2, 0))
	testVectorEquals(t, c, white)

	c = pattern.GetColor(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0, 0, 1))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0, 0, 2))
	testVectorEquals(t, c, white)

	c = pattern.GetColor(algebra.NewPoint(0, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0.9, 0, 0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(1.0, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(-0.1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(-1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(-1.1, 0, 0))
	testVectorEquals(t, c, white)
}

func TestGradientPattern(t *testing.T) {
	white := &Color{1,1,1}
	black := &Color{0,0,0}
	pattern := GradientPattern(white, black)
	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0.25, 0.0, 0.0))
	testVectorEquals(t, c, &Color{0.75, 0.75, 0.75})
	c = pattern.GetColor(algebra.NewPoint(0.5, 0, 0))
	testVectorEquals(t, c, &Color{0.5, 0.5, 0.5})
	c = pattern.GetColor(algebra.NewPoint(0.75, 0, 0))
	testVectorEquals(t, c, &Color{0.25, 0.25, 0.25})
}

func TestRingPattern(t *testing.T) {
	white := &Color{1,1,1}
	black := &Color{0,0,0}
	pattern := RingPattern(white, black)
	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0, 0, 1))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0.708, 0, 0.708))
	testVectorEquals(t, c, black)
}

func TestCheckerPattern(t *testing.T) {
	white := &Color{1,1,1}
	black := &Color{0,0,0}
	pattern := CheckerPattern(white, black)

	//repeating in X

	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0.99,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(1.01,0,0))
	testVectorEquals(t, c, black)
	//repeating in Y

	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0.99,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,1.01,0))
	testVectorEquals(t, c, black)
	//repeating in Z

	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0,0.99))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0,1.01))
	testVectorEquals(t, c, black)
}

func TestSolidPattern(t *testing.T) {
	red := &Color{1, 0, 0}
	pattern := SolidPattern(red)
	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, red)
	c = pattern.GetColor(algebra.NewPoint(1,3,12))
	testVectorEquals(t, c, red)
}

func TestNestedPattern(t *testing.T) {
	white := &Color{1,1,1}
	black := &Color{0,0,0}
	patternA := SolidPattern(white)
	patternB := SolidPattern(black)
	checkerPattern := CheckerPattern(nil, nil)
	ringPattern := RingPattern(nil, nil)

	// CHECKER PATTERN NESTED TEST
	pattern := NestedPattern(checkerPattern, patternA, patternB)
	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0.99,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(1.01,0,0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0.99,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,1.01,0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0,0.99))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(0,0,1.01))
	testVectorEquals(t, c, black)

	//RING PATTERN NESTED TEST
	pattern = NestedPattern(ringPattern, patternA, patternB)
	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, white)
	c = pattern.GetColor(algebra.NewPoint(1, 0, 0))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0, 0, 1))
	testVectorEquals(t, c, black)
	c = pattern.GetColor(algebra.NewPoint(0.708, 0, 0.708))
	testVectorEquals(t, c, black)

}

func  TestBlendedPattern(t *testing.T) {
	red := &Color{1, 0, 0}
	black := &Color{0,0,0}
	blue := &Color{0, 0, 1}
	patternA := CheckerPattern(red, black)
	patternB := CheckerPattern(blue, black)
	pattern := BlendedPattern(patternA, patternB, nil)
	c := pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, &Color{0.5, 0, 0.5})
	c = pattern.GetColor(algebra.NewPoint(1.01, 0, 0))
	testVectorEquals(t, c, &Color{0, 0, 0})

	blend := func(colorA *Color, colorB *Color) *Color{
		return colorA.ScalarMult(0.33).Add(colorB.ScalarMult(0.67))
	}
	pattern = BlendedPattern(patternA, patternB, blend)
	c = pattern.GetColor(algebra.NewPoint(0,0,0))
	testVectorEquals(t, c, &Color{0.33, 0, 0.67})
	c = pattern.GetColor(algebra.NewPoint(1.01, 0, 0))
	testVectorEquals(t, c, &Color{0,0,0})
}

func TestPerlinNoisePattern(t *testing.T) {
	//white := &Color{0,0,0}
	//black := &Color{1,1,1}
	//for i := 0; i < 100; i++ {
	//	pattern := PerlinNoisePattern(white, black)
	//	c := pattern.GetColor(algebra.NewPoint(0, 0, 0))
	//	t.Logf("%v", c)
	//	if c.Red() < 0 || c.Red() > 1 {
	//		t.Errorf("Red color %f out of bounds", c.Red())
	//	}
	//	if c.Green() < 0 || c.Green() > 1 {
	//		t.Errorf("Green color %f out of bounds", c.Green())
	//	}
	//	if c.Blue() < 0 || c.Blue() > 1 {
	//		t.Errorf("Blue color %f out of bounds", c.Blue())
	//	}
	//}
}