package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"log"
	"math"
	"reflect"
)

//World manages the world space of the Shape(s) inside of it and the light sources illuminating it
type World struct {
	Objects []primitives.Shape
	Lights  []*canvas.PointLight
}

//NewDefaultWorld creates a new default world with one light source and 2 spheres
func NewDefaultWorld() *World {
	lights := make([]*canvas.PointLight, 0, 0)
	light := canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10))
	lights = append(lights, light)

	objects := make([]primitives.Shape, 0, 0)
	s1 := primitives.NewSphere(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.8, 1.0, 0.6}
	m.Diffuse = 0.7
	m.Specular = 0.2
	s1.SetMaterial(m)

	s2 := primitives.NewSphere(algebra.ScalingMatrix(0.5, 0.5, 0.5))
	objects = append(objects, s1, s2)
	return &World{Objects: objects, Lights: lights}
}

//Intersect intersects all objects in the World space with the given ray
func (w *World) Intersect(r *algebra.Ray) *primitives.Intersections {
	is := primitives.NewIntersections()

	for _, s := range w.Objects {
		err := is.Intersect(s, r)
		if err != nil {
			panic(err)
		}
	}
	return is
}

//ShadeHit Determines the color at some valid ray intersection (hit)
func (w World) ShadeHit(comps Comps, depth int) *canvas.Color {
	color := &canvas.Color{0, 0, 0}
	inShadow := w.PointIsShadowed(comps.OverPoint)
	for _, l := range w.Lights {
		pattern := comps.Object.GetMaterial().Pattern
		var patternColor *canvas.Color
		if pattern != nil {
			patternColor = primitives.PatternAtObject(comps.Object, pattern, comps.Point)
		} else {
			patternColor = nil
		}
		lightingColor := canvas.Lighting(comps.Object.GetMaterial(), patternColor, l, comps.Point, comps.Eye, comps.Normal, inShadow)
		color = color.Add(lightingColor)

		reflected := w.ReflectedColor(&comps, depth)
		refracted := w.RefractedColor(&comps, depth)
		material := comps.Object.GetMaterial()
		if material.Reflective > 0 && material.Transparency > 0 {
			reflectance := Schlick(&comps)
			color = color.Add(reflected.ScalarMult(reflectance))
			color = color.Add(refracted.ScalarMult(1 - reflectance))
		} else {
			color = color.Add(refracted)
			color = color.Add(reflected)
		}
	}
	return color
}

//ColorAt returns the color where the ray intersects (if at all), with a maximum recursive depth of depth
func (w World) ColorAt(ray *algebra.Ray, depth int) *canvas.Color {
	intersections := w.Intersect(ray)
	if h := intersections.Hit(); h == nil {
		return &canvas.Color{0, 0, 0}
	} else {
		c := PrepareComputations(h, ray, intersections)
		return w.ShadeHit(*c, depth)
	}
}

//PointIsShadowed returns whether or not the point in question is in the shadow of some other object
func (w World) PointIsShadowed(p *algebra.Vector) bool {
	for i := 0; i < len(w.Lights); i++ {
		v, err := w.Lights[i].Position.Subtract(p)
		if err != nil {
			panic(err)
			return false
		}

		dist := v.Magnitude()
		direction, err := v.Normalize()
		if err != nil {
			panic(err)
			return false
		}
		res := append(p.Get()[:3], direction.Get()[:3]...)
		r := algebra.NewRay(res...)
		is := w.Intersect(r)
		if h := is.Hit(); h != nil && h.T < dist {
			return true
		}
	}
	return false
}

//ReflectedColor determines if there is a reflected color being emitted at some ray intersection
// Takes the pre-computed computations at the ray intersection (struct Comps)
func (w *World) ReflectedColor(comps *Comps, depth int) *canvas.Color {
	if comps.Object.GetMaterial().Reflective == 0.0 || depth <= 0 {
		return &canvas.Color{0, 0, 0}
	}
	p := []float64{comps.OverPoint.Get()[0], comps.OverPoint.Get()[1], comps.OverPoint.Get()[2]}
	d := []float64{comps.Reflect.Get()[0], comps.Reflect.Get()[1], comps.Reflect.Get()[2]}
	res := append(p, d...)
	reflectRay := algebra.NewRay(res...)
	color := w.ColorAt(reflectRay, depth-1)
	return color.ScalarMult(comps.Object.GetMaterial().Reflective)
}

//RefractedColor determines if there is a refracted color being emitted at some ray intersection
// Takes the pre-computed computations at the ray intersection (struct Comps)
func (w *World) RefractedColor(comps *Comps, depth int) *canvas.Color {
	//completely opaque object
	if comps.Object.GetMaterial().Transparency == 0.0 || depth == 0 {
		return &canvas.Color{0, 0, 0}
	}
	refractiveRatio := comps.N1 / comps.N2
	cosI, err := algebra.DotProduct(comps.Eye, comps.Normal)
	if err != nil {
		panic(err)
	}
	sin2T := refractiveRatio * refractiveRatio * (1 - cosI*cosI)
	// Total reflection occurs
	if sin2T > 1 {
		return &canvas.Color{0, 0, 0}
	}

	cosT := math.Sqrt(1.0 - sin2T)
	direction, err := comps.Normal.MultScalar(refractiveRatio*cosI - cosT).Subtract(comps.Eye.MultScalar(refractiveRatio))
	if err != nil {
		panic(err)
	}
	point := comps.UnderPoint.Get()[:3]
	res := append(point, direction.Get()[:3]...)
	refractRay := algebra.NewRay(res...)
	color := w.ColorAt(refractRay, depth-1).ScalarMult(comps.Object.GetMaterial().Transparency)
	return color
}

//Schlick returns the reflectance at a pre-computed ray intersection based on the Schlick model
func Schlick(comps *Comps) float64 {
	cos, err := algebra.DotProduct(comps.Eye, comps.Normal)
	if err != nil {
		panic(err)
	}

	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2
		sin2T := n * n * (1 - cos*cos)
		if sin2T > 1.0 {
			return 1.0
		}

		cosT := math.Sqrt(1.0 - sin2T)
		cos = cosT
	}

	r0 := (comps.N1 - comps.N2) / (comps.N1 + comps.N2)
	r0 *= r0
	return r0 + (1-r0)*math.Pow((1-cos), 5)
}

//Comps manages the precomputed state of the necessary vectors for lighting
type Comps struct {
	T          float64
	Object     primitives.Shape
	Point      *algebra.Vector
	OverPoint  *algebra.Vector
	UnderPoint *algebra.Vector
	N1         float64
	N2         float64
	Eye        *algebra.Vector
	Normal     *algebra.Vector
	Reflect    *algebra.Vector
	Inside     bool
}

func PrepareComputations(intersection *primitives.Intersection, ray *algebra.Ray, is *primitives.Intersections) *Comps {
	position := ray.Position(intersection.T)
	c := &Comps{T: intersection.T, Object: intersection.Object, Point: position,
		Eye: ray.Get()["direction"].Negate(), Normal: primitives.NormalAt(intersection.Object, position, intersection)}

	if d, err := algebra.DotProduct(c.Normal, c.Eye); err != nil {
		panic(err)
	} else if d < 0 {
		c.Inside = true
		c.Normal = c.Normal.Negate()
	} else {
		c.Inside = false
	}
	EPSILON := 0.0001
	overPoint, err := c.Point.Add(c.Normal.MultScalar(EPSILON))
	if err != nil {
		panic(err)
	}
	c.OverPoint = overPoint

	underPoint, err := c.Point.Subtract(c.Normal.MultScalar(EPSILON))
	if err != nil {
		panic(err)
	}
	c.UnderPoint = underPoint

	direction := ray.Get()["direction"]
	c.Reflect = direction.Reflect(c.Normal)
	determineRefractiveIndexes(c, intersection, is)
	return c
}

// Helper functions

func determineRefractiveIndexes(comps *Comps, hit *primitives.Intersection, is *primitives.Intersections) {
	if is == nil || len(is.GetHits().Get()) == 0 {
		log.Print("Warning: no intersections provided, this should only occur during unit testing")
		comps.N1 = 1.0
		comps.N2 = 1.0
		return
	} // this is only possible if calling the PrepareComputations method directly with nil in testing
	containers := make([]*primitives.Intersection, 0, 0)
	allIntersections := getSortedIntersections(is)

	for i := 0; i < len(allIntersections); i++ {
		if intersectionEquals(allIntersections[i], hit) {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = containers[len(containers)-1].Object.GetMaterial().RefractiveIndex
			}
		}
		if index, found := has(containers, allIntersections[i]); found {
			containers = append(containers[:index], containers[index+1:]...) //remove object
		} else {
			containers = append(containers, allIntersections[i])
		}
		if intersectionEquals(allIntersections[i], hit) {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].Object.GetMaterial().RefractiveIndex
			}
			break

		}
	}
}

func intersectionEquals(a *primitives.Intersection, b *primitives.Intersection) bool {
	if a.T == b.T && reflect.TypeOf(a.Object) == reflect.TypeOf(b.Object) && a.Object == b.Object {
		return true
	}
	return false
}

func has(container []*primitives.Intersection, intersect *primitives.Intersection) (int, bool) {
	for i := 0; i < len(container); i++ {
		if reflect.TypeOf(container[i]) == reflect.TypeOf(intersect) && container[i].Object == intersect.Object {
			return i, true
		}
	}
	return -1, false
}

func getSortedIntersections(is *primitives.Intersections) []*primitives.Intersection {
	res := make([]*primitives.Intersection, 0, 0)
	ref := is.GetRef().Copy()
	hits := is.GetHits().Copy()
	intersect := ref.ExtractMin()
	for intersect != nil {
		res = append(res, intersect)
		intersect = ref.ExtractMin()
	}
	intersect = hits.ExtractMin()
	for intersect != nil {
		res = append(res, intersect)
		intersect = hits.ExtractMin()
	}
	return res
}
