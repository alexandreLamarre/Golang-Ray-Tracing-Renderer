package primitives

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
	"testing")


func TestNewBounds(t *testing.T) {
	b := NewBounds(algebra.NewPoint(1, 1, 1), algebra.NewPoint(2, 2, 2))
	testVectorEquals(t, b.minimum.Get(), algebra.NewPoint(1,1,1).Get())
	testVectorEquals(t, b.maximum.Get(), algebra.NewPoint(2,2,2).Get())
}

func TestBounds_Get(t *testing.T) {
	b := NewBounds(algebra.NewPoint(1, 1, 1), algebra.NewPoint(2, 2, 2))
	min, max := b.Get()
	testVectorEquals(t, min.Get(), algebra.NewPoint(1,1,1).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(2,2,2).Get())
}

func TestGetBoundsTransform(t *testing.T) {
	s := NewCube(nil)
	min, max := s.GetBounds()
	testVectorEquals(t, min.Get(), algebra.NewPoint(-1,-1,-1).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(1,1,1).Get())

	b := GetBoundsTransform(min, max, algebra.RotationZ(-math.Pi/2))
	min, max = b.Get()
	testVectorEquals(t, min.Get(), algebra.NewPoint(-1, -1, -1).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(1,1,1).Get())

	b = GetBoundsTransform(min, max, algebra.RotationZ(math.Pi/4))
	min, max = b.Get()
	testVectorEquals(t, min.Get(), algebra.NewPoint(-math.Sqrt(2), -math.Sqrt(2), -1).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(math.Sqrt(2), math.Sqrt(2), 1).Get())
}

func TestBounds_Intersect(t *testing.T) {
	s := NewCube(nil)
	min, max := s.GetBounds()


	b := GetBoundsTransform(min, max, algebra.RotationZ(math.Pi/4))
	assert(t, b.Intersect(algebra.NewRay(0,0,10,0,0,-1).Transform(s.GetTransform())), fmt.Sprintf("Ray intersect"))
	assert(t, b.Intersect(algebra.NewRay(0,-1.1,10,0,0,-1).Transform(s.GetTransform())), fmt.Sprintf("Ray intersect"))
	assert(t, b.Intersect(algebra.NewRay(0,-1.1,10,0,0,-1).Transform(s.GetTransform())), fmt.Sprintf("Ray  intersect"))

	assertFalse(t, b.Intersect(algebra.NewRay(0,-1.6,10,0,-1,-1).Transform(s.GetTransform())), fmt.Sprintf("Ray intersect"))

}

func assert(t *testing.T, statement bool, clarification string){
	if !statement{
		t.Errorf("Expected true from statement %s", clarification)
	}
}

func assertFalse(t *testing.T, statement bool, clarification string){
	if statement{
		t.Errorf("Expected false from statement %s", clarification)
	}
}