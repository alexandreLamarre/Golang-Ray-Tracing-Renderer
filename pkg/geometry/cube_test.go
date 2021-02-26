package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"testing"
)

func TestNewCube(t *testing.T) {
	cube := NewCube()
	testMatrixEquals(t, cube.transform.Get(), algebra.IdentityMatrix(4).Get())
	testMaterialEquals(t, cube.material, canvas.NewDefaultMaterial())
}

func TestCube_GetMaterial(t *testing.T) {
	cube := NewCube()
	m := cube.GetMaterial()
	testMaterialEquals(t, m, canvas.NewDefaultMaterial())
}

func TestCube_GetTransform(t *testing.T) {
	cube := NewCube()
	m := cube.GetTransform()
	testMatrixEquals(t, m.Get(), cube.transform.Get())
}

func TestCube_SetMaterial(t *testing.T) {
	cube := NewCube()
	m := canvas.NewDefaultMaterial()
	m.Diffuse = 0.5
	cube.SetMaterial(m)
	testMaterialEquals(t, m, cube.GetMaterial())
}

func TestCube_SetTransform(t *testing.T) {
	cube := NewCube()
	m := algebra.ScalingMatrix(3,3,3)
	cube.SetTransform(m)
	testMatrixEquals(t, m.Get(), cube.transform.Get())
}

func TestCube_LocalIntersect(t *testing.T) {
	c := NewCube()

	rays := []*algebra.Ray{algebra.NewRay(5, 0.5, 0, -1, 0, 0),
		algebra.NewRay(-5, 0.5, 0, 1, 0, 0),
		algebra.NewRay(0.5, 5, 0, 0, -1, 0),
		algebra.NewRay(0.5, -5, 0, 0, 1, 0),
		algebra.NewRay(0.5, 0, 5, 0, 0, -1),
		algebra.NewRay(0.5, 0, -5, 0, 0, 1),
		algebra.NewRay(0, 0.5, 0, 0, 0, 1)}

	positions := [][2]float64{
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{4, 6},
		{-1, 1},
	}

	for i := 0; i < len(rays); i++{
		xs, hit := c.LocalIntersect(rays[i])
		if hit != true{
			t.Errorf("Expected ray %v, %v to hit default cube", rays[i].Get()["origin"], rays[i].Get()["direction"])
		}
		if len(xs) != 2{
			t.Errorf("Expected %d intersections, got: %d", 2, len(xs))
		}
		if !equals(xs[0], positions[i][0]){
			t.Errorf("Expected %f, got %f", positions[i][0], xs[0])
		}
		if !equals(xs[1], positions[i][1]){
			t.Errorf("Expected %f, got %f", positions[i][1], xs[1])
		}

	}

	//miss intersect
	rays2 := []*algebra.Ray{
		algebra.NewRay(-2, 0, 0, 0.2673, 0.5345, 0.8018),
		algebra.NewRay(0, -2, 0, 0.8018, 0.2673, 0.5345),
		algebra.NewRay(0, 0, -2, 0.5345, 0.8018, 0.2673),
		algebra.NewRay(2, 0, 2, 0, 0, -1),
		algebra.NewRay(0, 2, 2, 0, -1, 0),
		algebra.NewRay(2, 2, 0, -1, 0, 0)}

	for i := 0; i < len(rays2); i++{
		xs, hit := c.LocalIntersect(rays2[i])
		if hit == true{
			t.Errorf("Expected ray %v, %v to not hit default cube", rays2[i].Get()["origin"], rays2[i].Get()["direction"])
		}
		if len(xs) != 0{
			t.Errorf("Expected %d intersections, got: %d", 0, len(xs))
		}
	}
}

func TestCube_LocalNormalAt(t *testing.T) {
	c := NewCube()

	points := []*algebra.Vector{
		algebra.NewPoint(1, 0.5, -0.8),
		algebra.NewPoint(-1, -0.2, 0.9),
		algebra.NewPoint(-0.4, 1, -0.1),
		algebra.NewPoint(0.3, -1, -0.7),
		algebra.NewPoint(-0.6, 0.3, 1),
		algebra.NewPoint(0.4, 0.4, -1),
		algebra.NewPoint(1, 1, 1),
		algebra.NewPoint(-1, -1, -1),
	}

	normals := []*algebra.Vector{
		algebra.NewVector(1, 0, 0),
		algebra.NewVector(-1, 0, 0),
		algebra.NewVector(0, 1, 0),
		algebra.NewVector(0, -1, 0),
		algebra.NewVector(0, 0, 1),
		algebra.NewVector(0, 0, -1),
		algebra.NewVector(1, 0, 0),
		algebra.NewVector(-1, 0, 0),
	}

	for i := 0; i < len(points); i++{
		normal, err := c.LocalNormalAt(points[i])
		if err != nil{
			t.Errorf("Expected no error to be returned by cube normal")
		}
		testVectorEquals(t, normal.Get(), normals[i].Get())
	}
}