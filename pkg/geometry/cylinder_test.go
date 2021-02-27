package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
	"testing"
	)

func TestNewCylinder(t *testing.T) {
	c := NewCylinder(nil)
	testMatrixEquals(t, c.transform.Get(), algebra.IdentityMatrix(4).Get())
	testMaterialEquals(t, c.material, canvas.NewDefaultMaterial())
	assertEquals(t, c.minimum, math.Inf(-1))
	assertEquals(t, c.maximum, math.Inf(1))
	if c.closed {
		t.Errorf("New cylinders are expected to not be closed by default")
	}
}

func TestCylinder_SetMinimum(t *testing.T) {
	c := NewCylinder(nil)
	c.SetMinimum(2)
	assertEquals(t, c.minimum, 2)
}

func TestCylinder_SetMaximum(t *testing.T) {
	c := NewCylinder(nil)
	c.SetMaximum(1)
	assertEquals(t, c.maximum, 1)
}

func TestCylinder_SetClosed(t *testing.T) {
	c := NewCylinder(nil)
	c.SetClosed(true)
	if !c.closed{
		t.Errorf("Expected cylinder to be closed")
	}
	c.SetClosed(false)
	if c.closed{
		t.Errorf("Expected cylinder to not be closed")
	}
}

func TestCylinder_GetMaterial(t *testing.T) {
	c := NewCylinder(nil)
	m := c.GetMaterial()
	testMaterialEquals(t, m, canvas.NewDefaultMaterial())
}

func TestCylinder_GetTransform(t *testing.T) {
	c := NewCylinder(nil)
	m := c.GetTransform()
	testMatrixEquals(t, m.Get(), algebra.IdentityMatrix(4).Get())
}

func TestCylinder_SetMaterial(t *testing.T) {
	c := NewCylinder(nil)
	m := canvas.NewDefaultMaterial()
	m.Specular = 10
	c.SetMaterial(m)
	testMaterialEquals(t, c.material, m)
}

func TestCylinder_SetTransform(t *testing.T) {
	c := NewCylinder(nil)
	m := algebra.ScalingMatrix(2,2,2)
	c.SetTransform(m)
	testMatrixEquals(t, m.Get(), c.transform.Get())
}

func TestCylinder_GetParent(t *testing.T) {
	c := NewCylinder(nil)
	if c.GetParent() != nil{
		t.Errorf("Expected cylinder to have no parent Shapes")
	}
}

func TestCylinder_SetParent(t *testing.T) {
	c1 := NewCylinder(nil)
	c2 := NewCylinder(nil)
	c1.SetParent(c2)
	if c1.GetParent() == nil{
		t.Errorf("Expected cylinder to have a parent Shape")
	}
}

func TestCylinder_LocalIntersect(t *testing.T) {
	c := NewCylinder(nil)
	rays := []*algebra.Ray{
		algebra.NewRay(1, 0, 0, 0, 1, 0),
		algebra.NewRay(0, 0, 0, 0, 1, 0),
		algebra.NewRay(0, 0, -5, 0.333, 0.333, 0.333),
	}

	for i := 0; i < len(rays); i++{
		xs, hit := c.LocalIntersect(rays[i])
		if hit{
			t.Errorf("Expected ray %v, %v to not hit the cylinder",
				rays[i].Get()["origin"], rays[i].Get()["direction"])
		}
		if len(xs) != 0{
			t.Errorf("Expected %d hits, got: %d", 0, len(xs))
		}
	}

	norm, err := algebra.NewVector(0.1, 1, 1).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}
	rays2 := []*algebra.Ray{
		algebra.NewRay(1, 0, -5, 0, 0, 1),
		algebra.NewRay(0, 0, -5, 0, 0, 1),
		algebra.NewRay(0.5, 0, -5, norm.Get()[0], norm.Get()[1], norm.Get()[2]),
	}

	positions := [][2]float64{
		{5, 5},
		{4, 6},
		{6.80789, 7.08872},
	}

	for i:= 0; i < len(rays2); i++{
		xs, hit := c.LocalIntersect(rays2[i])

		if !hit{
			t.Errorf("Expected ray %v, %v to hit the cylinder",
				rays2[i].Get()["origin"], rays2[i].Get()["direction"])
		}

		if len(xs) != 2{
			t.Errorf("Expected %d hits, got: %d", 2, len(xs))
		}

		if !equals(xs[0],positions[i][0]){
			t.Errorf("Expected intersection position: %f, Got: %f", positions[i][0], xs[0])
		}

		if !equals(xs[1], positions[i][1]){
			t.Errorf("Expected intersection position: %f, Got: %f", positions[i][1], xs[1])
		}
	}

	//intersect constrained cylinders
	c = NewCylinder(nil)
	c.SetMaximum(2)
	c.SetMinimum(1)
	norm, err = algebra.NewVector(0.1, 1, 0).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}

	rays3 := []*algebra.Ray{
		algebra.NewRay(0, 1.5, 0, norm.Get()[0], norm.Get()[1], 0),
		algebra.NewRay(0, 3, -5, 0, 0, 1),
		algebra.NewRay(0, 0, -5, 0, 0, 1),
		algebra.NewRay(0, 2, -5, 0, 0, 1),
		algebra.NewRay(0, 1, -5, 0, 0, 1),
		algebra.NewRay(0, 1.5, -2, 0, 0, 1),
	}

	for i:= 0; i < len(rays3); i++{
		xs, hit := c.LocalIntersect(rays3[i])

		if i == 5{
			if !hit{
				t.Errorf("Expected ray %v, %v to hit the cylinder",
					rays2[i].Get()["origin"], rays3[i].Get()["direction"])
			}

			if len(xs) != 2{
				t.Errorf("Expected %d hits, got: %d", 2, len(xs))
			}
		} else{
			if hit{
				t.Errorf("Expected ray %v, %v to not hit the cylinder",
					rays3[i].Get()["origin"], rays3[i].Get()["direction"])
			}
			if len(xs) != 0{
				t.Errorf("Expected %d hits, got: %d", 0, len(xs))
			}
		}
	}

	//capped cylinder
	c = NewCylinder(nil)
	c.SetMinimum(1)
	c.SetMaximum(2)
	c.SetClosed(true)
	norm0 := algebra.NewVector(0, -1, 0)
	norm1 := algebra.NewVector(0, -1, 2)
	norm2 := algebra.NewVector(0, -1, 1)
	norm3 := algebra.NewVector(0, 1, 2)
	norm4 := algebra.NewVector(0, 1, 1)

	rays4 := []*algebra.Ray{
		algebra.NewRay(0, 3, 0, norm0.Get()[0], norm0.Get()[1], norm0.Get()[2]),
		algebra.NewRay(0, 3, -2, norm1.Get()[0], norm1.Get()[1], norm1.Get()[2]),
		algebra.NewRay(0, 4, -2, norm2.Get()[0], norm2.Get()[1], norm2.Get()[2]),
		algebra.NewRay(0, 0, -2, norm3.Get()[0], norm3.Get()[1], norm3.Get()[2]),
		algebra.NewRay(0, -1, -2, norm4.Get()[0], norm4.Get()[1], norm4.Get()[2]),
	}

	for i := 0; i < len(rays4); i++{
		xs, hit := c.LocalIntersect(rays4[i])

		if !hit{
			t.Errorf("Expected ray %v, %v to hit the cylinder",
				rays4[i].Get()["origin"], rays4[i].Get()["direction"])
		}

		if len(xs) != 2{
			t.Errorf("Expected %d hits, got: %d", 2, len(xs))
		}
	}
}

func TestCylinder_LocalNormalAt(t *testing.T) {
	c := NewCylinder(nil)

	points := []*algebra.Vector{
		algebra.NewPoint(1, 0, 0),
		algebra.NewPoint(0, 5, -1),
		algebra.NewPoint(0, -2, 1),
		algebra.NewPoint(-1, 1, 0),
	}

	normals := []*algebra.Vector{
		algebra.NewVector(1, 0, 0),
		algebra.NewVector(0, 0, -1),
		algebra.NewVector(0, 0, 1),
		algebra.NewVector(-1, 0, 0),
	}

	for i := 0; i < len(points); i++{
		n, err := c.LocalNormalAt(points[i])
		if err != nil{
			t.Errorf("%s", err)
		}
		testVectorEquals(t, n.Get(), normals[i].Get())
	}

	c = NewCylinder(nil)
	c.SetMinimum(1)
	c.SetMaximum(2)
	c.SetClosed(true)


	points2 := []*algebra.Vector{
		algebra.NewPoint(0, 1, 0),
		algebra.NewPoint(0.5, 1, 0),
		algebra.NewPoint(0, 1, 0.5),
		algebra.NewPoint(0, 2, 0),
		algebra.NewPoint(0.5, 2, 0),
		algebra.NewPoint(0, 2, 0.5),
	}

	normals2 := []*algebra.Vector{
		algebra.NewVector(0, -1, 0),
		algebra.NewVector(0, -1, 0),
		algebra.NewVector(0, -1, 0),
		algebra.NewVector(0, 1, 0),
		algebra.NewVector(0, 1, 0),
		algebra.NewVector(0, 1, 0),
	}

	for i := 0; i < len(points2); i++{
		n, err := c.LocalNormalAt(points2[i])
		if err != nil{
			t.Errorf("%s", err)
		}
		testVectorEquals(t, n.Get(), normals2[i].Get())
	}
}