package examples

import (
	"bufio"
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"os"
)

func CreateSphere3DExample() error {

	rayOrigin := []float64{0,0,-5}

	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 400
	pixelSize := wallSize/float64(canvasPixels)
	half := wallSize/2

	c := canvas.NewCanvas(canvasPixels, canvasPixels)
	color := &canvas.Color{1, 0, 0}

	s := geometry.NewSphere(nil)
    m := canvas.NewDefaultMaterial()
    m.Color = &canvas.Color{1, 0.2, 1}
    s.SetMaterial(m)
    light := canvas.NewPointLight(&canvas.Color{1,1,1}, algebra.NewPoint(-10, 10, -10))
	for y:= 0; y < canvasPixels -1; y++{
		worldY := half- pixelSize*float64(y)
		for x := 0; x < canvasPixels -1; x++{
			worldX := -half+pixelSize*float64(x)

			position := algebra.NewPoint(worldX, worldY, wallZ)
			direction, err := position.Subtract(algebra.NewPoint(rayOrigin...))
			if err != nil{
				return err
			}
			direction, err = direction.Normalize()
			if err != nil{
				return err
			}
			vals := append(rayOrigin, direction.Get()...)
			r := algebra.NewRay(vals...)
			is := geometry.NewIntersections()
			err = is.Intersect(s, r)
			if err != nil{
				return err
			}
			hit := is.Hit(s,r)
			if hit != nil{
				p := r.Position(hit.T)
				n := hit.Object.NormalAt(p)
				eye := r.Get()["direction"].Negate()
				color = canvas.Lighting(s.GetMaterial(), light, p, eye, n)
				c.WritePixel(x,y,color)
			}
		}
	}
	stringRepr := c.ToPpmHeader(255) + c.ToPpmBody(255)

	err := writeToFile(stringRepr, "sphere")
	if err != nil {
		return err
	}
	return nil
}


func CreateProjectileExample() error {
	c := canvas.NewCanvas(900, 550)
	color := &canvas.Color{1, 0, 0}
	startVelocity, err := algebra.NewVector(1, 1.8, 0.0).Normalize()
	if err != nil {
		return err
	}
	startVelocity = startVelocity.MultScalar(11.25)
	p := &projectile{position: algebra.NewPoint(0, 1, 0), velocity: startVelocity}
	e := &environment{gravity: algebra.NewVector(0, -0.1, 0), wind: algebra.NewVector(-0.01, 0, 0)}

	fmt.Println(c.ToPpmHeader(255))

	for p.position.Get()[1] > 0 {
		err := tick(e, p)
		if err != nil {
			return err
		}
		x := int(p.position.Get()[0])
		y := c.Height - int(p.position.Get()[1])
		if x >= 0 && x < c.Width && y >= 0 && y < c.Height {
			c.WritePixel(x, y, color)
		}
	}
	s := c.ToPpmHeader(255)
	s += c.ToPpmBody(255)

	err = writeToFile(s, "projectile")
	if err != nil {
		return err
	}
	return nil
}

func writeToFile(toWrite, fileName string) error {
	f, err := os.Create("./pkg/examples/" + fileName + ".ppm")
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	n, err := w.WriteString(toWrite)
	fmt.Printf("Wrote %d bytes\n", n)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
