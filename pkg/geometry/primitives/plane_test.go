package primitives

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
	"testing"
)

func TestNewPlane(t *testing.T) {
	p := NewPlane(nil)
	testVectorEquals(t, p.origin.Get(), []float64{0, 0, 0, 1})
	testVectorEquals(t, p.direction.Get(), []float64{1, 0, 1, 0})
	testMatrixEquals(t, p.transform.Get(), algebra.IdentityMatrix(4).Get())
}

func TestPlane_GetBounds(t *testing.T) {
	p := NewPlane(nil)
	min, max := p.GetBounds()
	testVectorEquals(t, min.Get(), algebra.NewPoint(math.Inf(-1), 0, math.Inf(-1)).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(math.Inf(1), 0, math.Inf(1)).Get())
}

func TestPlane_GetMaterial(t *testing.T) {
	p := NewPlane(nil)
	m := p.GetMaterial()
	n := canvas.NewDefaultMaterial()
	testMaterialEquals(t, m, n)
}

func TestPlane_SetMaterial(t *testing.T) {
	m := canvas.NewDefaultMaterial()
	m.Ambient = 0.5
	p := NewPlane(nil)
	p.SetMaterial(m)
	testMaterialEquals(t, p.material, m)
}

func TestPlane_GetTransform(t *testing.T) {
	p := NewPlane(nil)
	testMatrixEquals(t, p.GetTransform().Get(), algebra.IdentityMatrix(4).Get())
}

func TestPlane_SetTransform(t *testing.T) {
	p := NewPlane(nil)
	p.SetTransform(algebra.TranslationMatrix(0, 2, 0))
	testMatrixEquals(t, p.transform.Get(), algebra.TranslationMatrix(0, 2, 0).Get())
}

func TestPlane_GetParent(t *testing.T) {
	p := NewPlane(nil)
	if p.GetParent() != nil {
		t.Errorf("Expected plane to have no Parent Shapes")
	}
}

func TestPlane_SetParent(t *testing.T) {
	p1 := NewPlane(nil)
	p2 := NewPlane(nil)
	p1.SetParent(p2)
	if p1.GetParent() == nil {
		t.Errorf("Expected plane to have a parent Shape")
	}
}

func TestPlane_GetPosition(t *testing.T) {
	p := NewPlane(nil)
	orig, direc := p.GetPosition()
	testVectorEquals(t, orig.Get(), []float64{0, 0, 0, 1})
	testVectorEquals(t, direc.Get(), []float64{1, 0, 1, 0})
}

func TestPlane_LocalIntersect(t *testing.T) {
	p := NewPlane(nil)
	r := algebra.NewRay(0, 10, 0, 0, 0, 1)
	xs, intersected := p.LocalIntersect(r)

	if intersected {
		fmt.Println(xs)
		t.Errorf("Expected no intersection for ray %v in the XY-plane", r)
	}

	p = NewPlane(nil)
	r = algebra.NewRay(0, 0, 0, 0, 0, 1)
	xs, intersected = p.LocalIntersect(r)
	if intersected {
		t.Errorf("Expected no intersections for ray %v in the XY plane", r)
	}

	r = algebra.NewRay(0, 1, 0, 0, -1, 0)
	xs, intersected = p.LocalIntersect(r)
	if !intersected {
		t.Errorf("Expected intersections for ray %v in the XY plane", r)
	}
	if len(xs) != 1 {
		t.Errorf("Invalid number of intersections")
	}
	if xs[0].T != 1.0 {
		t.Errorf("Expected 1.0, Got : %f", xs[0])
	}

	r = algebra.NewRay(0, -1, 0, 0, 1, 0)
	xs, intersected = p.LocalIntersect(r)
	if len(xs) != 1 {
		t.Errorf("Invalid number of intersections")
	}
	if xs[0].T != 1.0 {
		t.Errorf("Expected 1.0, Got : %f", xs[0])
	}

}

func TestPlane_LocalNormalAt(t *testing.T) {
	p := NewPlane(nil)
	n1, err := p.LocalNormalAt(algebra.NewPoint(0, 0, 0), nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	n2, err := p.LocalNormalAt(algebra.NewPoint(10, 0, -10), nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	n3, err := p.LocalNormalAt(algebra.NewPoint(-5, 0, 150), nil)
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	testVectorEquals(t, n1.Get(), []float64{0, 1, 0, 0})
	testVectorEquals(t, n2.Get(), []float64{0, 1, 0, 0})
	testVectorEquals(t, n3.Get(), []float64{0, 1, 0, 0})
}
