package primitives

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"testing"
)

func TestNewSmoothTriangle(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	testVectorEquals(t, tri.p1.Get(), algebra.NewPoint(0, 1, 0).Get())
	testVectorEquals(t, tri.p2.Get(), algebra.NewPoint(-1, 0, 0).Get())
	testVectorEquals(t, tri.p3.Get(), algebra.NewPoint(1, 0, 0).Get())

	testVectorEquals(t, tri.e1.Get(), algebra.NewVector(-1, -1, 0).Get())
	testVectorEquals(t, tri.e2.Get(), algebra.NewVector(1, -1, 0).Get())

	testVectorEquals(t, tri.n1.Get(), algebra.NewPoint(0, 1, 0).Get())
	testVectorEquals(t, tri.n2.Get(), algebra.NewPoint(-1, 0, 0).Get())
	testVectorEquals(t, tri.n3.Get(), algebra.NewPoint(1, 0, 0).Get())

	tri2 := NewDefaultSmoothTriangle(algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0),
		algebra.NewPoint(0, 0, 1))

	testVectorEquals(t, tri2.p1.Get(), algebra.NewPoint(1, 0, 0).Get())
	testVectorEquals(t, tri2.p2.Get(), algebra.NewPoint(0, 1, 0).Get())
	testVectorEquals(t, tri2.p3.Get(), algebra.NewPoint(0, 0, 1).Get())
	testVectorEquals(t, tri2.e1.Get(), algebra.NewVector(-1, 1, 0).Get())
	testVectorEquals(t, tri2.e2.Get(), algebra.NewVector(-1, 0, 1).Get())

	testVectorEquals(t, tri2.n1.Get(), algebra.NewVector(0.81649, -0.40824, -0.40824).Get())
	testVectorEquals(t, tri2.n2.Get(), algebra.NewVector(-0.40824, 0.81649, -0.40824).Get())
	testVectorEquals(t, tri2.n3.Get(), algebra.NewVector(-0.40824, -0.40824, 0.81649).Get())
}


func TestSmoothTriangle_GetBounds(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 0, 0), algebra.NewPoint(1, 0, 0),
		algebra.NewPoint(0, 1, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	min, max := tri.GetBounds()
	testVectorEquals(t, min.Get(), algebra.NewPoint(0,0,0).Get())
	testVectorEquals(t, max.Get(), algebra.NewPoint(1, 1, 0).Get())
}

func TestSmoothTriangle_GetMaterial(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	testMaterialEquals(t, tri.GetMaterial(), canvas.NewDefaultMaterial())

	m := canvas.NewDefaultMaterial()
	m.Diffuse =  5.0

}

func TestSmoothTriangle_GetTransform(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	testMatrixEquals(t, tri.GetTransform().Get(), algebra.IdentityMatrix(4).Get())
}

func TestSmoothTriangle_GetParent(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	if tri.GetParent() != nil{
		t.Errorf("Expected default triangle to have no Parent Shape")
	}


}

func TestSmoothTriangle_SetMaterial(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	m := canvas.NewDefaultMaterial()
	m.Diffuse =  5.0
	tri.SetMaterial(m)
	testMaterialEquals(t, m, tri.GetMaterial())
}

func TestSmoothTriangle_SetTransform(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	m := algebra.ScalingMatrix(1,2,3)
	tri.SetTransform(m)
	testMatrixEquals(t, tri.GetTransform().Get(), m.Get())
}

func TestSmoothTriangle_SetParent(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	if tri.GetParent() != nil{
		t.Errorf("Expected default triangle to have no Parent Shape")
	}

	tri2 := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	tri.SetParent(tri2)
	if tri.GetParent() == nil{
		t.Errorf("Expected  triangle to have Parent Shape")
	}
}

func TestSmoothTriangle_LocalIntersect(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0),
		algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
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


	tri2 := NewDefaultSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	r = algebra.NewRay(-0.2, 0.3, -2, 0, 0, 1)
	xs, hit = tri2.LocalIntersect(r)
	if !equals(xs[0].U, 0.45){
		t.Errorf("Expected U 0.45, Got %f", xs[0].U)
	}
	if !equals(xs[0].V, 0.25){
		t.Errorf("Expected V 0.25, Got %f", xs[1].U)
	}

}

func TestSmoothTriangle_LocalNormalAt(t *testing.T) {
	tri := NewSmoothTriangle(algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0),
		algebra.NewPoint(1, 0, 0), algebra.NewPoint(0, 1, 0), algebra.NewPoint(-1, 0, 0), algebra.NewPoint(1, 0, 0))
	r := algebra.NewRay(-0.2, 0.3, -2, 0, 0, 1)
	xs, hit := tri.LocalIntersect(r)
	if !hit{
		t.Errorf("Expected smooth triangle ray to hit")
	}
	n := NormalAt(tri, algebra.NewPoint(0,0,0), xs[0])
	testVectorEquals(t, n.Get(), algebra.NewVector(-0.5547, 0.83205, 0).Get())
}


