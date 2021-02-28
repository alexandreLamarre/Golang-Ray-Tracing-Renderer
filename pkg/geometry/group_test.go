package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"testing"
)

func TestNewGroup(t *testing.T) {
	g := NewGroup(nil)
	testMatrixEquals(t, g.transform.Get(), algebra.IdentityMatrix(4).Get())
	assertEquals(t, float64(len(g.shapes)), 0)
	if g.GetParent() != nil{
		t.Errorf("Expected default group to not have a Parent shape")
	}
}


func TestGroup_GetMaterial(t *testing.T) {
	g := NewGroup(nil)
	if g.GetMaterial() != nil{
		t.Errorf("Group get material must be nil")
	}
}

func TestGroup_GetTransform(t *testing.T) {
	g := NewGroup(nil)
	testMatrixEquals(t, g.GetTransform().Get(), algebra.IdentityMatrix(4).Get())
	m := algebra.ScalingMatrix(2,2,2)
	g = NewGroup(m)
	testMatrixEquals(t, g.GetTransform().Get(), m.Get())
}

func TestGroup_GetParent(t *testing.T) {
	g := NewGroup(nil)
	if g.GetParent() != nil{
		t.Errorf("Expected group to not have a Parent Shape")
	}
}

func TestGroup_SetMaterial(t *testing.T) {
	g := NewGroup(nil)
	g.SetMaterial(canvas.NewDefaultMaterial())
	if g.GetMaterial() != nil{
		t.Errorf("Expected g to have no material")
	}
}

func TestGroup_SetTransform(t *testing.T) {
	g := NewGroup(nil)
	m := algebra.ScalingMatrix(2,2,2)
	g.SetTransform(m)
	testMatrixEquals(t, g.GetTransform().Get(), m.Get())
}

func TestGroup_SetParent(t *testing.T) {
	g1 := NewGroup(nil)
	g2 := NewGroup(nil)
	g1.SetParent(g2)
	if g1.GetParent() == nil{
		t.Errorf("Expected group to have a Parent Shape")
	}
}

func TestGroup_AddChild(t *testing.T) {
	g := NewGroup(nil)
	s := NewSphere(nil)
	s2 := NewSphere(nil)
	assertEquals(t, float64(len(g.shapes)), 0)
	if s.GetParent() != nil || s2.GetParent() != nil{
		t.Errorf("Expected spheres not part of a group to have no parent")
	}

	g.AddChild(s)
	assertEquals(t, float64(len(g.shapes)), 1)
	if s.GetParent() == nil{
		t.Errorf("Expected sphere1 to have group parent")
	}
	if s2.GetParent() != nil{
		t.Errorf("Expected sphere2 to not have a group parent")
	}

	g.AddChild(s2)
	assertEquals(t, float64(len(g.shapes)), 2)
	if s.GetParent() == nil{
		t.Errorf("Expected sphere1 to have group parent")
	}
	if s2.GetParent() == nil{
		t.Errorf("Expected sphere1 to have group parent")
	}
}

func TestGroup_LocalIntersect(t *testing.T) {
	g := NewGroup(nil)
	r := algebra.NewRay(0, 0, 0, 0, 0, 1)
	xs, hit := g.LocalIntersect(r)
	if hit{
		t.Errorf("Expected ray %v %v to not hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 0{
		t.Errorf("Expected %d hits, Got : %d", 0 ,len(xs))
	}

	s1 := NewSphere(nil)
	s2 := NewSphere(algebra.TranslationMatrix(0, 0, -3))
	s3 := NewSphere(algebra.TranslationMatrix(5, 0, 0))
	testMatrixEquals(t, s3.GetTransform().Get(), algebra.TranslationMatrix(5, 0, 0).Get())
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	r = algebra.NewRay(0, 0, -5, 0, 0, 1)
	_, hit = s3.LocalIntersect(r)
	xs, hit = g.LocalIntersect(r)
	if !hit{
		t.Errorf("Expected ray %v %v to hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 4{
		t.Errorf("Expected %d ray intersections, got %d", 4, len(xs))
	}
	if xs[0].Object != s1{
		t.Errorf("0, s1")
	}
	if xs[1].Object != s1{
		t.Errorf("1, s1")
	}
	if xs[2].Object != s2{
		t.Errorf("2, s2")
	}
	if xs[3].Object != s2{
		t.Errorf("3, s2")
	}

	g = NewGroup(algebra.ScalingMatrix(2,2,2))
	s := NewSphere(algebra.TranslationMatrix(5, 0 , 0))
	g.AddChild(s)
	r = algebra.NewRay(10, 0, -10, 0, 0, 1)

	is := NewIntersections()
	err := is.Intersect(g, r)
	if err != nil{
		t.Errorf("%s", err)
	}
	if is.Count() != 2{
		t.Errorf("Expected 2 ray intersections, got %d", is.Count())
	}

}

func TestGroup_LocalNormalAt(t *testing.T) {
	g := NewGroup(nil)
	_, gerr := g.LocalNormalAt(algebra.NewPoint(0,0,0))


	if gerr == nil{
		t.Errorf("Unexpected lack of error or unexpected error from group local normal")
	}
}

func TestGroup_GetBounds(t *testing.T) {
	g := NewGroup(nil)
	min, max := g.GetBounds()
	if min != nil{
		t.Errorf("Unexpected group bounding minimum %v", min)
	}
	if max != nil{
		t.Errorf("Unexpected group bounding maximum %v", max)
	}
}
