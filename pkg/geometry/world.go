package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
)

//World manages the world space of the Shape(s) inside of it and the light sources illuminating it
type World struct {
	Objects []Shape
	Lights  []*canvas.PointLight
}

//NewDefaultWorld creates a new default world with one light source and 2 spheres
func NewDefaultWorld() *World {
	lights := make([]*canvas.PointLight, 0, 0)
	light := canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10))
	lights = append(lights, light)

	objects := make([]Shape, 0, 0)
	s1 := NewSphere(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.8, 1.0, 0.6}
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)

	s2 := NewSphere(algebra.ScalingMatrix(0.5, 0.5, 0.5))
	objects = append(objects, s1, s2)
	return &World{Objects: objects, Lights: lights}
}

//Intersect intersects all objects in the World space with the given ray
func (w *World) Intersect(r *algebra.Ray) *Intersections {
	is := NewIntersections()

	for _, s := range w.Objects {
		err := is.Intersect(s, r)
		if err != nil {
			panic(err)
		}
	}
	return is
}

func (w World) ShadeHit(comps Comps) *canvas.Color {
	color := &canvas.Color{0, 0, 0}
	for _, l := range w.Lights {
		lightingColor := canvas.Lighting(comps.Object.GetMaterial(), l, comps.Point, comps.Eye, comps.Normal)
		color = color.Add(lightingColor)
	}
	return color
}

func (w World) ColorAt(ray *algebra.Ray) *canvas.Color {
	intersections := w.Intersect(ray)
	if h := intersections.Hit(); h == nil {
		return &canvas.Color{0, 0, 0}
	} else {
		c := PrepareComputations(h, ray)
		return w.ShadeHit(*c)
	}
}

//Comps manages the precomputed state of the necessary vectors for lighting
type Comps struct {
	T      float64
	Object Shape
	Point  *algebra.Vector
	Eye    *algebra.Vector
	Normal *algebra.Vector
	Inside bool
}

func PrepareComputations(intersection *Intersection, ray *algebra.Ray) *Comps {
	position := ray.Position(intersection.T)
	c := &Comps{T: intersection.T, Object: intersection.Object, Point: position,
		Eye: ray.Get()["direction"].Negate(), Normal: intersection.Object.NormalAt(position)}

	if d, err := algebra.DotProduct(c.Normal, c.Eye); err != nil {
		panic(err)
	} else if d < 0 {
		c.Inside = true
		c.Normal = c.Normal.Negate()
	} else {
		c.Inside = false
	}
	return c
}
