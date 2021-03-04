package objects

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"testing"
)

func TestContains(t *testing.T) {
	g := primitives.NewGroup(nil)
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	s3 := primitives.NewSphere(nil)

	g.AddChild(s1)
	g.AddChild(s2)

	if !Contains(s1, s1) {
		t.Errorf("Expected s1 to trivially contain s1")
	}

	if Contains(s2, s1) {
		t.Errorf("Expected s2 to not trivially contain s1")
	}

	if !Contains(g, s1) {
		t.Errorf("Expected group g to contain s1")
	}

	if Contains(g, s3) {
		t.Errorf("Expected group g to not contain s3")
	}

	g1 := primitives.NewGroup(nil)
	g1.AddChild(g)

	if !Contains(g1, s1) {
		t.Errorf("Expected group g to contain s1")
	}

	if Contains(g1, s3) {
		t.Errorf("Expected group g to not contain s3")
	}

	g1.AddChild(s3)

	if !Contains(g1, s3) {
		t.Errorf("Expected group g1 to contain s3")
	}

	s4 := primitives.NewSphere(nil)
	s5 := primitives.NewSphere(nil)
	csg := UnionCSG(s4, s5)

	if Contains(csg, s1) {
		t.Errorf("Expected csg")
	}

	if !Contains(csg, s4) {
		t.Errorf("Expected csg to contain s4")
	}

	if !Contains(csg, s5) {
		t.Errorf("Expected csg to contain s5")
	}

	g.AddChild(csg)

	if !Contains(g1, s4) {
		t.Errorf("Expected g1 to contain s5")
	}

	if !Contains(g1, s5) {
		t.Errorf("Expected g1 to contain s5")
	}
	s6 := primitives.NewSphere(nil)
	csg2 := IntersectCSG(csg, s6)
	if !Contains(csg2, s6) {
		t.Errorf("Expected csg2 to contain s6")
	}

	s1 = primitives.NewSphere(nil)
	s2 = primitives.NewSphere(algebra.TranslationMatrix(0, 0, 0.5))
	csg = UnionCSG(s1, s2)
	if Contains(csg.Left(), s2) {
		t.Errorf("Left side of union should not contain right side")
	}
}

func TestUnionCSG(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)

	if csg.left != s1 {
		t.Errorf("Expected csg.left to be s1")
	}
	if csg.right != s2 {
		t.Errorf("Expected csg.right to be s2")
	}
	if csg.action != "union" {
		t.Errorf("Expected union action")
	}
}

func TestDifferenceCSG(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := DifferenceCSG(s1, s2)

	if csg.left != s1 {
		t.Errorf("Expected csg.left to be s1")
	}
	if csg.right != s2 {
		t.Errorf("Expected csg.right to be s2")
	}
	if csg.action != "difference" {
		t.Errorf("Expected union action")
	}
}

func TestIntersectCSG(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := IntersectCSG(s1, s2)

	if csg.left != s1 {
		t.Errorf("Expected csg.left to be s1")
	}
	if csg.right != s2 {
		t.Errorf("Expected csg.right to be s2")
	}
	if csg.action != "intersect" {
		t.Errorf("Expected union action")
	}
}

func TestCSGShape_Left(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)

	if csg.Left() != csg.left {
		t.Errorf("Left doesnt behave appropriately")
	}
}

func TestCSGShape_Right(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)

	if csg.Right() != csg.right {
		t.Errorf("Left doesnt behave appropriately")
	}
}

func TestCSGShape_GetBounds(t *testing.T) {
	s := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(algebra.TranslationMatrix(2, 0, 0).Multiply(algebra.ScalingMatrix(2, 2, 2)))

	csg := UnionCSG(s, s2)
	min, max := csg.GetBounds()
	testVectorEquals(t, min.Get(), algebra.NewVector(-1, -2, -2).Get())
	testVectorEquals(t, max.Get(), algebra.NewVector(4, 2, 2).Get())
}

func TestCSGShape_GetMaterial(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	testMaterialEquals(t, csg.GetMaterial(), canvas.NewDefaultMaterial())
}

func TestCSGShape_GetTransform(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	testMatrixEquals(t, csg.GetTransform().Get(), algebra.IdentityMatrix(4).Get())
}

func TestCSGShape_GetParent(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	if csg.Parent != nil {
		t.Errorf("Expected default parent to be nil")
	}
}

func TestCSGShape_SetMaterial(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	m := canvas.NewDefaultMaterial()
	m.Diffuse = 1.0
	csg.SetMaterial(m)
	testMaterialEquals(t, csg.GetMaterial(), m)
}

func TestCSGShape_SetParent(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	csg.SetParent(primitives.NewGroup(nil))
	if csg.GetParent() == nil {
		t.Errorf("Expected parent to not be nil")
	}
}

func TestCSGShape_SetTransform(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := UnionCSG(s1, s2)
	csg.SetTransform(algebra.ScalingMatrix(2, 2, 2))
	testMatrixEquals(t, csg.GetTransform().Get(), algebra.ScalingMatrix(2, 2, 2).Get())
}

func Test_intersectionAllowed(t *testing.T) {
	truthTable := [][]bool{
		{true, true, true},
		{true, true, false},
		{true, false, true},
		{true, false, false},
		{false, true, true},
		{false, true, false},
		{false, false, true},
		{false, false, false}}

	unionResult := []bool{false, true, false, true, false, false, true, true}

	intersectionResult := []bool{true, false, true, false, true, true, false, false}

	differenceResult := []bool{false, true, false, true, true, true, false, false}

	for i, truthVals := range truthTable {

		if intersectionAllowed("union", truthVals[0], truthVals[1], truthVals[2]) != unionResult[i] {

			t.Errorf("Expected union intersection to be %v, for lhit:%v, inl :%v, inr : %v", unionResult[i],
				truthVals[0], truthVals[1], truthVals[2])
		}
		if intersectionAllowed("intersect", truthVals[0], truthVals[1], truthVals[2]) != intersectionResult[i] {
			t.Errorf("Expected intersect intersection to be %v, for lhit:%v, inl :%v, inr : %v", intersectionResult[i],
				truthVals[0], truthVals[1], truthVals[2])
		}
		if intersectionAllowed("difference", truthVals[0], truthVals[1], truthVals[2]) != differenceResult[i] {
			t.Errorf("Expected difference intersection to be %v, for lhit:%v, inl :%v, inr : %v", differenceResult[i],
				truthVals[0], truthVals[1], truthVals[2])
		}
	}

}

func Test_filterIntersections(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewCube(nil)
	xs := []*primitives.Intersection{
		primitives.NewIntersection(s1, 1),
		primitives.NewIntersection(s2, 2),
		primitives.NewIntersection(s1, 3),
		primitives.NewIntersection(s2, 4),
	}

	action := "union"
	csg := UnionCSG(s1, s2)
	res := csg.filterIntersections(xs)
	if len(res) != 2 {
		t.Errorf("Expected only 2 intersections for %s shape", action)
		return
	}
	if res[0] != xs[0] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[0], action, res[0])
	}
	if res[1] != xs[3] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[3], action, res[1])
	}

	action = "intersect"
	csg = IntersectCSG(s1, s2)

	res = csg.filterIntersections(xs)
	if len(res) != 2 {
		t.Errorf("Expected only 2 intersections for %s shape", action)
		return
	}
	if res[0] != xs[1] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[1], action, res[0])
	}
	if res[1] != xs[2] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[2], action, res[1])
	}

	action = "difference"
	csg = DifferenceCSG(s1, s2)

	res = csg.filterIntersections(xs)
	if len(res) != 2 {
		t.Errorf("Expected only 2 intersections for %s shape", action)
		return
	}
	if res[0] != xs[0] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[0], action, res[0])
	}
	if res[1] != xs[1] {
		t.Errorf("Expected intersection at %v for %s shape, got: %v", xs[1], action, res[1])
	}
}

func TestCSGShape_LocalIntersect(t *testing.T) {
	c := UnionCSG(primitives.NewSphere(nil), primitives.NewCube(nil))
	r := algebra.NewRay(0, 2, -5, 0, 0, 1)
	xs, hit := c.LocalIntersect(r)
	if hit {
		t.Errorf("Expected ray %v %v to not hit", r.Get()["origin"], r.Get()["direction"])
	}
	if len(xs) != 0 {
		t.Errorf("Expected ray %v %v to have %d intersections, got : %d",
			r.Get()["origin"], r.Get()["direction"], 0, len(xs))
	}

	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(algebra.TranslationMatrix(0, 0, 0.5))
	csg := UnionCSG(s1, s2)
	r = algebra.NewRay(0, 0, -5, 0, 0, 1)
	xs, hit = csg.LocalIntersect(r)
	if !hit {
		t.Errorf("Expected ray %v %v to hit", r.Get()["origin"], r.Get()["direction"])
		return
	}
	if len(xs) != 2 {
		t.Errorf("Expected ray %v %v to have %d intersections, got : %d",
			r.Get()["origin"], r.Get()["direction"], 2, len(xs))
	}
}

func TestCSGShape_LocalNormalAt(t *testing.T) {
	s1 := primitives.NewSphere(nil)
	s2 := primitives.NewSphere(nil)
	csg := IntersectCSG(s1, s2)
	n, err := csg.LocalNormalAt(algebra.NewPoint(-1, 0, 1), primitives.NewIntersection(s1, 1))
	if n != nil {
		t.Errorf("Expected local normal at csg to always return nil")
	}
	if err == nil {
		t.Errorf("Expected error to be no nil")
	}
	if serr, _ := err.(*ErrorCSGNormal); serr != nil {
		t.Errorf("Expected custom Error 'ErrorCSGNormal', instead got : %s", serr)
	}
}
