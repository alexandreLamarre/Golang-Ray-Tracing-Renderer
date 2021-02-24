package examples

import (
	"bufio"
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	camera2 "github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/camera"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"math"
	"os"
)

func CreateSimpleScene2() error{
	floor := geometry.NewPlane(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0, 1, 0.0}
	m.Specular = 0
	m.Ambient = 0.5
	pat1 := canvas.GradientPattern(&canvas.Color{0, 0, 1}, &canvas.Color{1, 1, 1})
	m.Pattern = canvas.PerlinNoisePattern(pat1)
	floor.SetMaterial(m)


	wall := geometry.NewPlane(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5), algebra.RotationX(math.Pi/2) ))
	m = canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{1.0, 1.0, 1.0}
	m.Specular = 0
	wall.SetMaterial(m)

	middle := geometry.NewSphere(algebra.TranslationMatrix(-0.5,1, 0.5))
	middleMat := canvas.NewDefaultMaterial()
	middleMat.Color = &canvas.Color{0.1, 1, 0.5}
	middleMat.Diffuse = 0.7
	middleMat.Specular = 0.3
	pat :=  canvas.CheckerPattern(&canvas.Color{1,0,0}, &canvas.Color{0,0,1})
	pat.SetTransform(algebra.ScalingMatrix(0.25, 0.25, 0.25))
	middleMat.Pattern = canvas.PerlinNoisePattern(pat)
	middle.SetMaterial(middleMat)

	right := geometry.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(1.5, 0.5, -0.5),
			algebra.ScalingMatrix(0.5, 0.5, 0.5)))
	rightMat := canvas.NewDefaultMaterial()
	rightMat.Color = &canvas.Color{0.5, 1, 0.1}
	rightMat.Diffuse = 0.7
	rightMat.Specular = 0.3
	right.SetMaterial(rightMat)

	left := geometry.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(-1.5,0.33,-0.75),
			algebra.ScalingMatrix(0.33,0.33,0.33)))
	leftMat := canvas.NewDefaultMaterial()
	leftMat.Color = &canvas.Color{1, 0.8, 0.1}
	leftMat.Diffuse = 0.7
	leftMat.Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]geometry.Shape, 0, 0)
	objs = append(objs, floor, wall, middle, left, right)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10)))
	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(400, 200, math.Pi/3,
		algebra.ViewTransform(0, 1.5, -5,
			0, 1, 0,
			0, 1, 0))
	if err != nil{
		panic(err)
		return err
	}
	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "simpleScene2")
	if err != nil{
		panic(err)
		return err
	}
	return nil
}

func CreateSimpleScene() error{
	floor := geometry.NewSphere(algebra.ScalingMatrix(10, 0.001, 10))
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{1, 0.9, 0.9}
	m.Specular = 0
	floor.SetMaterial(m)

	lW := algebra.Multiply(algebra.Multiply(algebra.Multiply(algebra.TranslationMatrix(0,0,5),
		algebra.RotationY(-math.Pi/4)), algebra.RotationX(math.Pi/2)), algebra.ScalingMatrix(10, 0.001, 10))
	leftWall := geometry.NewSphere(lW)
	leftWall.SetMaterial(floor.GetMaterial())


	rW := algebra.Multiply(algebra.Multiply(algebra.Multiply(algebra.TranslationMatrix(0,0,5),
		algebra.RotationY(math.Pi/4)), algebra.RotationX(math.Pi/2)), algebra.ScalingMatrix(10, 0.01, 10))
	rightWall := geometry.NewSphere(rW)
	rightWall.SetMaterial(floor.GetMaterial())

	middle := geometry.NewSphere(algebra.TranslationMatrix(-0.5,1, 0.5))
	middleMat := canvas.NewDefaultMaterial()
	middleMat.Color = &canvas.Color{0.1, 1, 0.5}
	middleMat.Diffuse = 0.7
	middleMat.Specular = 0.3
	middle.SetMaterial(middleMat)

	right := geometry.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(1.5, 0.5, -0.5),
		algebra.ScalingMatrix(0.5, 0.5, 0.5)))
	rightMat := canvas.NewDefaultMaterial()
	rightMat.Color = &canvas.Color{0.5, 1, 0.1}
	rightMat.Diffuse = 0.7
	rightMat.Specular = 0.3
	right.SetMaterial(rightMat)

	left := geometry.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(-1.5,0.33,-0.75),
		algebra.ScalingMatrix(0.33,0.33,0.33)))
	leftMat := canvas.NewDefaultMaterial()
	leftMat.Color = &canvas.Color{1, 0.8, 0.1}
	leftMat.Diffuse = 0.7
	leftMat.Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]geometry.Shape, 0, 0)
	objs = append(objs, floor, leftWall, rightWall, middle, left, right)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10)))
	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(400, 200, math.Pi/3,
		algebra.ViewTransform(0, 1.5, -5,
			0, 1, 0,
			0, 1, 0))
	if err != nil{
		panic(err)
		return err
	}
	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "simpleScene")
	if err != nil{
		panic(err)
		return err
	}
	return nil
}

func CreateSphere3DExample() error {

	rayOrigin := []float64{0, 0, -5}

	wallZ := 10.0
	wallSize := 7.0

	canvasPixels := 400
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2

	c := canvas.NewCanvas(canvasPixels, canvasPixels)
	color := &canvas.Color{1, 0, 0}

	s := geometry.NewSphere(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{1, 0.2, 1}
	s.SetMaterial(m)
	light := canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10))
	for y := 0; y < canvasPixels-1; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels-1; x++ {
			worldX := -half + pixelSize*float64(x)

			position := algebra.NewPoint(worldX, worldY, wallZ)
			direction, err := position.Subtract(algebra.NewPoint(rayOrigin...))
			if err != nil {
				return err
			}
			direction, err = direction.Normalize()
			if err != nil {
				return err
			}
			vals := append(rayOrigin, direction.Get()...)
			r := algebra.NewRay(vals...)
			is := geometry.NewIntersections()
			err = is.Intersect(s, r)
			if err != nil {
				return err
			}
			hit := is.Hit()
			if hit != nil {
				p := r.Position(hit.T)
				n := geometry.NormalAt(hit.Object, p)
				eye := r.Get()["direction"].Negate()
				color = canvas.Lighting(s.GetMaterial(), nil, light, p, eye, n, false)
				c.WritePixel(x, y, color)
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
