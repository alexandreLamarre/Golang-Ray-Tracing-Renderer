package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"reflect"
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

func (w World) ShadeHit(comps Comps, depth int) *canvas.Color {
	color := &canvas.Color{0, 0, 0}
	inShadow := w.PointIsShadowed(comps.OverPoint)
	for _, l := range w.Lights {
		pattern := comps.Object.GetMaterial().Pattern
		var patternColor *canvas.Color
		if pattern != nil{
			patternColor = PatternAtObject(comps.Object, pattern, comps.Point)
		} else {
			patternColor = nil
		}
		lightingColor := canvas.Lighting(comps.Object.GetMaterial(), patternColor, l, comps.Point, comps.Eye, comps.Normal, inShadow)
		color = color.Add(lightingColor)

		reflected := w.ReflectedColor(&comps, depth)
		color = color.Add(reflected)
	}
	return color
}

func (w World) ColorAt(ray *algebra.Ray, depth int) *canvas.Color {
	intersections := w.Intersect(ray)
	if h := intersections.Hit(); h == nil {
		return &canvas.Color{0, 0, 0}
	} else {
		c := PrepareComputations(h, ray, intersections)
		return w.ShadeHit(*c, depth)
	}
}

func (w World) PointIsShadowed(p *algebra.Vector) bool{
	for i := 0; i < len(w.Lights); i++{
		v, err := w.Lights[i].Position.Subtract(p)
		if err != nil{
			panic(err)
			return false
		}

		dist := v.Magnitude()
		direction, err := v.Normalize()
		if err != nil{
			panic(err)
			return false
		}
		res := append(p.Get()[:3], direction.Get()[:3]...)
		r := algebra.NewRay(res...)
		is := w.Intersect(r)
		if h:=is.Hit(); h != nil && h.T < dist{
				return true
		}
	}
	return false
}

func (w *World) ReflectedColor(comps *Comps, depth int) *canvas.Color {
	if comps.Object.GetMaterial().Reflective == 0.0 || depth <= 0{
		return &canvas.Color{0,0,0}
	}
	p := []float64{comps.OverPoint.Get()[0], comps.OverPoint.Get()[1], comps.OverPoint.Get()[2]}
	d := []float64{comps.Reflect.Get()[0], comps.Reflect.Get()[1], comps.Reflect.Get()[2] }
	res := append(p, d...)
	reflectRay := algebra.NewRay(res...)
	color := w.ColorAt(reflectRay, depth -1)
	return color.ScalarMult(comps.Object.GetMaterial().Reflective)
}

//Comps manages the precomputed state of the necessary vectors for lighting
type Comps struct {
	T      float64
	Object Shape
	Point  *algebra.Vector
	OverPoint *algebra.Vector
	UnderPoint *algebra.Vector
	N1 float64
	N2 float64
	Eye    *algebra.Vector
	Normal *algebra.Vector
	Reflect *algebra.Vector
	Inside bool
}

func PrepareComputations(intersection *Intersection, ray *algebra.Ray, is *Intersections) *Comps {
	position := ray.Position(intersection.T)
	c := &Comps{T: intersection.T, Object: intersection.Object, Point: position,
		Eye: ray.Get()["direction"].Negate(), Normal: NormalAt(intersection.Object,position)}

	if d, err := algebra.DotProduct(c.Normal, c.Eye); err != nil {
		panic(err)
	} else if d < 0 {
		c.Inside = true
		c.Normal = c.Normal.Negate()
	} else {
		c.Inside = false
	}
	EPSILON := 0.00001
	overPoint,err := c.Point.Add(c.Normal.MultScalar(EPSILON))
	if err != nil{
		panic(err)
	}
	c.OverPoint = overPoint

	direction := ray.Get()["direction"]
	c.Reflect = direction.Reflect(c.Normal)
	determineRefractiveIndexes(c, intersection, is)
	return c
}

func determineRefractiveIndexes(comps *Comps, hit *Intersection, is *Intersections){
	if is == nil {
		comps.N1 = 1.0
		comps.N2 = 1.0
		return
	} // this is only possible if calling the PrepareComputations method directly with nil in testing
	containers := make([]*Intersection, 0, 0)
	allIntersections := append(is.ref.Get(), is.hits.Get()...)

	for i := 0; i < len(allIntersections); i++{
		if intersectionEquals(allIntersections[i], hit){
			if len(containers) == 0{
				comps.N1 = 1.0
			} else{
				comps.N1 = containers[len(containers) -1].Object.GetMaterial().RefractiveIndex
			}
		}
		if index, found := has(containers, allIntersections[i]); found{
			containers = append(containers[:index], containers[index+1:]...)
		} else{
			containers = append(containers, allIntersections[i])
		}
		if intersectionEquals(allIntersections[i], hit){
			if len(containers) == 0{
				comps.N2 = 1.0
			} else{
				comps.N2 = containers[len(containers) -1].Object.GetMaterial().RefractiveIndex
			}
			break
		}
	}
}

func intersectionEquals(a *Intersection, b *Intersection) bool{
	if a.T == b.T && reflect.TypeOf(a.Object) == reflect.TypeOf(b.Object) && a.Object == b.Object{
		return true
	}
	return false
}

func has(container []*Intersection, intersect *Intersection) (int, bool){
	for i:= 0; i < len(container); i++{
		if reflect.TypeOf(container[i]) == reflect.TypeOf(intersect) && container[i].Object == intersect.Object{
			return i, true
		}
	}
	return -1, false
}
