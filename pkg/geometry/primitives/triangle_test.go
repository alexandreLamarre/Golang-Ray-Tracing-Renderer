package primitives

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"testing"
)

func TestNewTriangle(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	testVectorEquals(t, tri.p1.Get(), algebra.NewPoint(0, 1, 0).Get())
	testVectorEquals(t, tri.p2.Get(), algebra.NewPoint(-1, 0, 0).Get())
	testVectorEquals(t, tri.p3.Get(), algebra.NewPoint(1, 0, 0).Get())

	testVectorEquals(t, tri.e1.Get(), algebra.NewVector(-1, -1, 0).Get())
	testVectorEquals(t, tri.e2.Get(), algebra.NewVector(1, -1, 0).Get())
	testVectorEquals(t, tri.normal.Get(), algebra.NewVector(0, 0, -1).Get())
}


func TestTriangle_GetBounds(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	min, max := tri.GetBounds()
	testVectorEquals(t, min.Get(), algebra.NewPoint(0,0,0).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(1, 1, 0).Get())
}

func TestTriangle_GetMaterial(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	testMaterialEquals(t, tri.GetMaterial(), canvas.NewDefaultMaterial())

	m := canvas.NewDefaultMaterial()
	m.Diffuse =  5.0

}

func TestTriangle_GetTransform(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	testMatrixEquals(t, tri.GetTransform().Get(), algebra.IdentityMatrix(4).Get())
}

func TestTriangle_GetParent(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	if tri.GetParent() != nil{
		t.Errorf("Expected default triangle to have no Parent Shape")
	}


}

func TestTriangle_SetMaterial(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	m := canvas.NewDefaultMaterial()
	m.Diffuse =  5.0
	tri.SetMaterial(m)
	testMaterialEquals(t, m, tri.GetMaterial())
}

func TestTriangle_SetTransform(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	m := algebra.ScalingMatrix(1,2,3)
	tri.SetTransform(m)
	testMatrixEquals(t, tri.GetTransform().Get(), m.Get())
}

func TestTriangle_SetParent(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 0, 0), algebra.NewPoint(0, 1, 0))
	if tri.GetParent() != nil{
		t.Errorf("Expected default triangle to have no Parent Shape")
	}

	tri2 := NewTriangle(algebra.NewPoint(1, 1, 1), algebra.NewPoint(0 ,1, 0), algebra.NewPoint(0, 0, 1))
	tri.SetParent(tri2)
	if tri.GetParent() == nil{
		t.Errorf("Expected  triangle to have Parent Shape")
	}
}

func TestTriangle_LocalNormalAt(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	n1, err := tri.LocalNormalAt(algebra.NewPoint(0, 0.5, 0))
	if err != nil{
		t.Errorf("%s", err)
	}
	testVectorEquals(t, n1.Get(), tri.normal.Get())

	n2, err := tri.LocalNormalAt(algebra.NewPoint(-0.5, 0.75));
	if err != nil{
		t.Errorf("%s", err)
	}
	testVectorEquals(t, n2.Get(), tri.normal.Get())

	n3, err := tri.LocalNormalAt(algebra.NewPoint(0.5, 0.25, 0));
	if err != nil{
		t.Errorf("%s", err)
	}
	testVectorEquals(t, n3.Get(), tri.normal.Get())
}

func TestTriangle_LocalIntersect(t *testing.T) {
	tri := NewTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	r := algebra.NewRay(0, -1, -2, 0, 1, 0)
	xs, hit := tri.LocalIntersect(r)
	//parallel
	if hit{
		t.Errorf("Expected ray %v %v to not hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 0{
		t.Errorf("Expected %d hits, Got : %d", 0, len(xs))
	}

	r = algebra.NewRay(1, 1, -2, 0, 0, 1)
	xs, hit = tri.LocalIntersect(r)
	//over edge 1
	if hit{
		t.Errorf("Expected ray %v %v to not hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 0{
		t.Errorf("Expected %d hits, Got : %d", 0, len(xs))
	}

	r = algebra.NewRay(-1, 1, -2, 0, 0, 1)
	xs, hit = tri.LocalIntersect(r)
	//over edge 2
	if hit{
		t.Errorf("Expected ray %v %v to not hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 0{
		t.Errorf("Expected %d hits, Got : %d", 0, len(xs))
	}

	r = algebra.NewRay(0, 0.5, -2, 0, 0, 1)
	xs, hit = tri.LocalIntersect(r)
	//intersect
	if !hit{
		t.Errorf("Expected ray %v %v to hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 1{
		t.Errorf("Expected %d hits, Got : %d", 1, len(xs))
	}

	if !equals(xs[0].T, 2){
		t.Errorf("Expected ray to intersect at %f, Got: %f", 2.0, xs[0].T )
	}
}

