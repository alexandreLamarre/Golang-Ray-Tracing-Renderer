package geometry


import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
	"testing"
)

func TestNewCone(t *testing.T) {
	c := NewCone(nil)
	testMatrixEquals(t, c.transform.Get(), algebra.IdentityMatrix(4).Get())
	testMaterialEquals(t, c.material, canvas.NewDefaultMaterial())
	assertEquals(t, c.minimum, math.Inf(-1))
	assertEquals(t, c.maximum, math.Inf(1))
	if c.closed {
		t.Errorf("New cylinders are expected to not be closed by default")
	}
}

func TestCone_SetMinimum(t *testing.T) {
	c := NewCone(nil)
	c.SetMinimum(2)
	assertEquals(t, c.minimum, 2)
}

func TestCone_SetMaximum(t *testing.T) {
	c := NewCone(nil)
	c.SetMaximum(1)
	assertEquals(t, c.maximum, 1)
}

func TestCone_SetClosed(t *testing.T) {
	c := NewCone(nil)
	c.SetClosed(true)
	if !c.closed{
		t.Errorf("Expected cylinder to be closed")
	}
	c.SetClosed(false)
	if c.closed{
		t.Errorf("Expected cylinder to not be closed")
	}
}

func TestCone_GetMaterial(t *testing.T) {
	c := NewCone(nil)
	m := c.GetMaterial()
	testMaterialEquals(t, m, canvas.NewDefaultMaterial())
}

func TestCone_GetTransform(t *testing.T) {
	c := NewCone(nil)
	m := c.GetTransform()
	testMatrixEquals(t, m.Get(), algebra.IdentityMatrix(4).Get())
}

func TestCone_SetMaterial(t *testing.T) {
	c := NewCone(nil)
	m := canvas.NewDefaultMaterial()
	m.Specular = 10
	c.SetMaterial(m)
	testMaterialEquals(t, c.material, m)
}

func TestCone_SetTransform(t *testing.T) {
	c := NewCone(nil)
	m := algebra.ScalingMatrix(2,2,2)
	c.SetTransform(m)
	testMatrixEquals(t, m.Get(), c.transform.Get())
}

func TestCone_GetParent(t *testing.T) {
	c := NewCone(nil)
	if c.GetParent() != nil{
		t.Errorf("Expected cone to have no parent Shapes")
	}
}

func TestCone_SetParent(t *testing.T) {
	c1 := NewCone(nil)
	c2 := NewCone(nil)
	c1.SetParent(c2)
	if c1.GetParent() == nil{
		t.Errorf("Expected cone to have a parent Shape")
	}
}

func TestCone_LocalIntersect(t *testing.T) {
	c := NewCone(nil)
	norm, err := algebra.NewVector(0, 0, 1).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}
	norm2, err := algebra.NewVector(1, 1, 1).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}
	norm3 , err:= algebra.NewVector(-0.5, -1, 1).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}

	rays := []*algebra.Ray{
		algebra.NewRay(0, 0, -5, norm.Get()[0], norm.Get()[1], norm.Get()[2]),
		algebra.NewRay(0, 0, -5, norm2.Get()[0], norm2.Get()[1], norm2.Get()[2]),
		algebra.NewRay(1, 1, -5, norm3.Get()[0], norm3.Get()[1], norm3.Get()[2]),
	}

	positions := [][2]float64{
		{5, 5},
		{8.66025, 8.66025},
		{4.55006, 49.44994},
	}

	for i := 0; i < len(rays); i++{
		xs, hit := c.LocalIntersect(rays[i])
		if !hit{
			t.Errorf("Expected ray %v, %v to not hit the cone",
				rays[i].Get()["origin"], rays[i].Get()["direction"])
		}
		if len(xs) != 2{
			t.Errorf("Expected %d hits, got: %d", 2, len(xs))
		}
		if !equals(xs[0].T,positions[i][0]){
			t.Errorf("Expected intersection position: %f, Got: %f", positions[i][0], xs[0])
		}

		if !equals(xs[1].T, positions[i][1]){
			t.Errorf("Expected intersection position: %f, Got: %f", positions[i][1], xs[1])
		}
	}

	//ray intersect parallel to one of the halves
	norm,err = algebra.NewVector(0, 1, 1).Normalize()
	if err != nil{
		t.Errorf("%s", err)
	}
	r := algebra.NewRay(0, 0, -1, norm.Get()[0], norm.Get()[1], norm.Get()[2])
	xs, hit := c.LocalIntersect(r)

	if !hit{
		t.Errorf("Expected ray %v, %v to hit the cylinder",
			r.Get()["origin"], r.Get()["direction"])
	}

	if len(xs) != 1{
		t.Errorf("Expected %d hits, got: %d", 1, len(xs))
	}

	if !equals(xs[0].T, 0.35355){
		t.Errorf("Expected intersection at %f, Got: %f", 0.35355, xs[0])
	}

	//intersect the end caps

	c = NewCone(nil)
	c.minimum = -0.5
	c.maximum = 0.5
	c.closed = true

	norm, err = algebra.NewVector(0, 1, 0).Normalize()
	norm2, err = algebra.NewVector(0, 1, 1).Normalize()
	norm3, err = algebra.NewVector(0 , 1, 0).Normalize()

	rays = []*algebra.Ray{
		algebra.NewRay(0, 0, -5, norm.Get()[0], norm.Get()[1], norm.Get()[2]),
		algebra.NewRay(0, 0, -0.25, norm2.Get()[0], norm2.Get()[1], norm2.Get()[2]),
		algebra.NewRay(0, 0, -0.25, norm3.Get()[0], norm3.Get()[1], norm3.Get()[2]),
	}

	for i := 0; i < len(rays); i++{
		xs, hit := c.LocalIntersect(rays[i])
		if i == 0{
			if hit{
				t.Errorf("Expected ray %v, %v to not hit the cylinder",
					rays[i].Get()["origin"], rays[i].Get()["direction"])
			}
			if len(xs) != 0{
				t.Errorf("Expected %d hits, got: %d", 0, len(xs))
			}

		} else{
			if !hit{
				t.Errorf("Expected ray %v, %v to hit the cylinder",
					rays[i].Get()["origin"], rays[i].Get()["direction"])
			}
			if i == 1 && len(xs) != 2{
				t.Errorf("Expected %d hits, got: %d", 2, len(xs))
				for _, v := range xs{
					fmt.Println(v)
				}
			}
			if i ==2 && len(xs) != 4{
				t.Errorf("Expected %d hits, got: %d", 4, len(xs))
			}
		}
	}
}

func TestCone_LocalNormalAt(t *testing.T) {
	c := NewCone(nil)

	points := []*algebra.Vector{
		algebra.NewPoint(0, 0, 0),
		algebra.NewPoint(1, 1, 1),
		algebra.NewPoint(-1, -1, 0),
	}

	normals := []*algebra.Vector{
		algebra.NewVector(0 , 0, 0),
		algebra.NewVector(1, -math.Sqrt(2), 1),
		algebra.NewVector(-1, 1, 0),
	}

	for i := 0; i < len(points); i++{
		n,err := c.LocalNormalAt(points[i])
		if err != nil{
			t.Errorf("%s", err)
		}
		testVectorEquals(t, n.Get(), normals[i].Get())
	}
}
