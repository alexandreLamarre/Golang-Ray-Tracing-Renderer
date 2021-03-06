package examples

import (
	"bufio"
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	camera2 "github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/camera"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/objects"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"log"
	"math"
	"os"
	"time"
)

func CreateCustomScene(w *geometry.World, name string, rotate bool) error {
	log.Println("Rendering scene...")
	start := time.Now()
	cam, err := camera2.NewCamera(1400, 1000, math.Pi/3,
		algebra.ViewTransform(0, 30, -50,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}
	image := cam.Render(w)
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Done (%s)!", elapsed)
	start = time.Now()
	log.Println("Writing results to file ./pkg/examples/" + name + ".ppm...")
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, name)
	if err != nil {
		panic(err)
		return err
	}
	t = time.Now()
	elapsed = t.Sub(start)
	log.Printf("Done (%s)!", elapsed)
	return nil
}

func writeToFile(toWrite, fileName string) error {
	f, err := os.Create("./pkg/examples/" + fileName + ".ppm")
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	n, err := w.WriteString(toWrite)
	log.Printf("Wrote %d bytes", n)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func TestReflectiveScene() error {
	white := &canvas.Color{0.9, 0.9, 0.9}
	black := &canvas.Color{0, 0, 0}
	cam, err := camera2.NewCamera(400, 200, math.Pi/3,
		algebra.ViewTransform(0, 3.0, -15,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}

	floor := primitives.NewPlane(algebra.TranslationMatrix(0, -1, 0))
	m := canvas.NewDefaultMaterial()
	pat := canvas.CheckerPattern(white, black)
	m.Pattern = pat
	m.Reflective = 0.5
	m.Specular = 0
	m.Diffuse = 0
	m.Shininess = 0
	m.Ambient = 1
	floor.SetMaterial(m)

	s := primitives.NewSphere(algebra.Multiply(algebra.ScalingMatrix(4, 4, 4), algebra.TranslationMatrix(0.0, 1.2, 0)))
	m = canvas.NewDefaultMaterial()
	m.Color = white
	m.RefractiveIndex = 1.5
	m.Reflective = 1.0
	m.Shininess = 300
	m.Specular = 0.9
	m.Ambient = 0.1
	m.Diffuse = 0.4
	s.SetMaterial(m)

	light := canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-100, 0, -50))
	lights := []*canvas.PointLight{light}
	objs := []primitives.Shape{floor, s}
	w := &geometry.World{Objects: objs, Lights: lights}
	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "reflectiveTestScene")
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CreateSimpleReflectiveScene() error {
	floor := primitives.NewPlane(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.1, 0.1, 0.3}
	m.Specular = 0
	m.Ambient = 0.5
	m.Reflective = 0.1
	pat := canvas.CheckerPattern(&canvas.Color{0, 0, 0}, &canvas.Color{1, 1, 1})
	m.Pattern = pat
	floor.SetMaterial(m)

	lwall := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.Multiply(algebra.RotationY(-math.Pi/4), algebra.RotationX(math.Pi/2))))
	m = canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.1, 0.1, 0.3}
	m.Specular = 0.0
	m.Reflective = 0.9
	m.RefractiveIndex = 1.5
	m.Transparency = 0.9
	m.Diffuse = 0.4
	//pat = canvas.CheckerPattern(&canvas.Color{0,0,0}, &canvas.Color{1,1,1})
	//pat.SetTransform(algebra.TranslationMatrix(10,10,10))
	//m.Pattern = pat
	lwall.SetMaterial(m)

	rwall := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.Multiply(algebra.RotationY(math.Pi/4), algebra.RotationX(math.Pi/2))))
	m = canvas.NewDefaultMaterial()
	//pat = canvas.CheckerPattern(&canvas.Color{0,0,0}, &canvas.Color{1,1,1})
	//pat.SetTransform(algebra.TranslationMatrix(10,10,10))
	m.Color = &canvas.Color{1.0, 1.0, 1.0}
	m.Specular = 0.0
	m.Reflective = 0.9
	m.RefractiveIndex = 1.5
	m.Transparency = 0.9
	m.Diffuse = 0.4
	rwall.SetMaterial(m)

	middle := primitives.NewSphere(algebra.Multiply(algebra.TranslationMatrix(0.0, 2, -3.0), algebra.ScalingMatrix(2, 2, 2)))
	middleMat := canvas.NewDefaultMaterial()
	middleMat.Color = &canvas.Color{1.0, 1.0, 1.0}
	middleMat.Diffuse = 0.4
	middleMat.Specular = 0.3
	middleMat.RefractiveIndex = 1.5
	middleMat.Transparency = 1.0
	middleMat.Reflective = 0.9
	middle.SetMaterial(middleMat)

	middle2 := primitives.NewSphere(algebra.TranslationMatrix(0.0, 2, -3.0))
	middleMat2 := canvas.NewDefaultMaterial()
	middleMat2.Color = &canvas.Color{1.0, 1.0, 1.0}
	middleMat2.Diffuse = 0.4
	middleMat2.Specular = 0.3
	middleMat2.RefractiveIndex = 1.0
	middleMat2.Transparency = 1.0
	middleMat2.Reflective = 0.9
	middle2.SetMaterial(middleMat2)

	objs := make([]primitives.Shape, 0, 0)
	objs = append(objs, floor, middle, middle2, lwall, rwall)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10)))
	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(1400, 1000, math.Pi/3,
		algebra.ViewTransform(0, 3.0, -15,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}

	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "reflectiveScene")
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CreateRefractiveReflectiveScene() error {

	lights := make([]*canvas.PointLight, 0, 0)
	objs := make([]primitives.Shape, 0, 0)

	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-2.5, 2.5, -4)))

	bottomMirror := primitives.NewPlane(algebra.TranslationMatrix(0, -1, 0))
	matb := bottomMirror.GetMaterial()
	matb.Diffuse = 0.3
	matb.Color = &canvas.Color{0.9, 0.9, 0.9}
	matb.RefractiveIndex = 1.5
	matb.Reflective = 0.9
	matb.Transparency = 0.9
	bottomMirror.SetMaterial(matb)

	upperMirror := primitives.NewPlane(algebra.TranslationMatrix(0, 2, 0))
	matu := upperMirror.GetMaterial()
	matu.Diffuse = 0.3
	matu.Color = &canvas.Color{0.9, 0.9, 0.9}
	matu.RefractiveIndex = 1.5
	matu.Reflective = 0.9
	matu.Transparency = 0.9
	upperMirror.SetMaterial(matu)

	backgroundMirror := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5), algebra.RotationX(math.Pi/2)))
	matback := backgroundMirror.GetMaterial()
	matback.Diffuse = 0.3
	matback.Color = &canvas.Color{0.9, 0.9, 0.9}
	matback.RefractiveIndex = 1.5
	matback.Reflective = 0.9
	matback.Transparency = 0.9
	backgroundMirror.SetMaterial(matback)

	foregroundMirror := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(0, 0, -5), algebra.RotationX(math.Pi/2)))
	matfore := foregroundMirror.GetMaterial()
	matfore.Diffuse = 0.3
	matfore.Color = &canvas.Color{0.9, 0.9, 0.9}
	matfore.RefractiveIndex = 1.5
	matfore.Reflective = 0.9
	matfore.Transparency = 0.9
	foregroundMirror.SetMaterial(matfore)

	leftMirror := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(-3, 0, 0), algebra.RotationY(math.Pi/2)))
	matleft := leftMirror.GetMaterial()
	matleft.Diffuse = 0.3
	matleft.Color = &canvas.Color{0.9, 0.9, 0.9}
	matleft.RefractiveIndex = 1.5
	matleft.Reflective = 0.9
	matleft.Transparency = 0.9
	leftMirror.SetMaterial(matleft)

	rightMirror := primitives.NewPlane(algebra.Multiply(algebra.TranslationMatrix(3, 0, 0), algebra.RotationY(math.Pi/2)))
	matright := leftMirror.GetMaterial()
	matright.Diffuse = 0.3
	matright.Color = &canvas.Color{0.9, 0.9, 0.9}
	matright.RefractiveIndex = 1.5
	matright.Reflective = 0.9
	matright.Transparency = 0.9
	rightMirror.SetMaterial(matright)

	ball := primitives.NewSphere(nil)
	matball := ball.GetMaterial()
	matball.Color = &canvas.Color{1, 0, 0}
	ball.SetMaterial(matball)

	objs = append(objs, bottomMirror, upperMirror, backgroundMirror, foregroundMirror, leftMirror, rightMirror, ball)

	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(100, 50, math.Pi/3,
		algebra.ViewTransform(0, 1.0, -3,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}

	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "reflectiveScene2")
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CreateSimpleScene2() error {
	white := &canvas.Color{1, 1, 1}
	black := &canvas.Color{0, 0, 0}
	floor := primitives.NewPlane(nil)
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.7, 0.7, 0.7}
	m.Specular = 0
	m.Ambient = 0.5
	m.Pattern = canvas.CheckerPattern(white, black)
	floor.SetMaterial(m)

	lw := algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.Multiply(algebra.RotationY(-math.Pi/4), algebra.RotationX(math.Pi/2)))
	leftWall := primitives.NewPlane(lw)
	m = canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{0.9, 0.9, 0.9}
	m.Specular = 0
	pat := canvas.CheckerPattern(white, black)
	m.Pattern = pat
	leftWall.SetMaterial(m)

	rw := algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.Multiply(algebra.RotationY(math.Pi/4), algebra.RotationX(math.Pi/2)))
	rightWall := primitives.NewPlane(rw)
	m = canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{1.0, 1.0, 1.0}
	m.Specular = 0
	m.Pattern = canvas.CheckerPattern(white, black)
	rightWall.SetMaterial(m)

	middle := primitives.NewSphere(algebra.TranslationMatrix(-0.5, 1, 0.5))
	middleMat := canvas.NewDefaultMaterial()
	middleMat.Color = &canvas.Color{0.1, 1, 0.5}
	middleMat.Diffuse = 0.7
	middleMat.Specular = 0.3
	pat = canvas.CheckerPattern(&canvas.Color{1, 0, 0}, &canvas.Color{0, 0, 1})
	pat.SetTransform(algebra.ScalingMatrix(0.1, 0.1, 0.1))
	middleMat.Pattern = pat
	middle.SetMaterial(middleMat)

	right := primitives.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(1.5, 0.5, -0.5),
			algebra.ScalingMatrix(0.5, 0.5, 0.5)))
	rightMat := canvas.NewDefaultMaterial()
	rightMat.Color = &canvas.Color{0.5, 1, 0.1}
	rightMat.Diffuse = 0.7
	rightMat.Specular = 0.3
	right.SetMaterial(rightMat)

	left := primitives.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(-1.5, 0.33, -0.75),
			algebra.ScalingMatrix(0.33, 0.33, 0.33)))
	leftMat := canvas.NewDefaultMaterial()
	leftMat.Color = &canvas.Color{1, 0.8, 0.1}
	leftMat.Diffuse = 0.7
	leftMat.Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]primitives.Shape, 0, 0)
	objs = append(objs, floor, leftWall, rightWall, middle, left, right)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10)))
	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(100, 50, math.Pi/3,
		algebra.ViewTransform(0, 1.5, -10,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}
	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "simpleScene2")
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func CreateSimpleScene() error {
	floor := primitives.NewSphere(algebra.ScalingMatrix(10, 0.001, 10))
	m := canvas.NewDefaultMaterial()
	m.Color = &canvas.Color{1, 0.9, 0.9}
	m.Specular = 0
	floor.SetMaterial(m)

	lW := algebra.Multiply(algebra.Multiply(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.RotationY(-math.Pi/4)), algebra.RotationX(math.Pi/2)), algebra.ScalingMatrix(10, 0.001, 10))
	leftWall := primitives.NewSphere(lW)
	leftWall.SetMaterial(floor.GetMaterial())

	rW := algebra.Multiply(algebra.Multiply(algebra.Multiply(algebra.TranslationMatrix(0, 0, 5),
		algebra.RotationY(math.Pi/4)), algebra.RotationX(math.Pi/2)), algebra.ScalingMatrix(10, 0.01, 10))
	rightWall := primitives.NewSphere(rW)
	rightWall.SetMaterial(floor.GetMaterial())

	middle := objects.NewHexagon(algebra.RotationX(-math.Pi / 2))

	min, max := middle.GetBounds()
	fmt.Println(primitives.GetBoundsTransform(min, max, middle.GetTransform()).Get())

	right := primitives.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(1.5, 0.5, -0.5),
			algebra.ScalingMatrix(0.5, 0.5, 0.5)))
	rightMat := canvas.NewDefaultMaterial()
	rightMat.Color = &canvas.Color{0.5, 1, 0.1}
	rightMat.Diffuse = 0.7
	rightMat.Specular = 0.3
	right.SetMaterial(rightMat)

	left := primitives.NewSphere(
		algebra.Multiply(algebra.TranslationMatrix(-1.5, 0.33, -0.75),
			algebra.ScalingMatrix(0.33, 0.33, 0.33)))
	leftMat := canvas.NewDefaultMaterial()
	leftMat.Color = &canvas.Color{1, 0.8, 0.1}
	leftMat.Diffuse = 0.7
	leftMat.Specular = 0.3
	left.SetMaterial(leftMat)

	objs := make([]primitives.Shape, 0, 0)
	objs = append(objs, middle)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-10, 10, -10)))
	w := &geometry.World{Objects: objs, Lights: lights}

	cam, err := camera2.NewCamera(400, 200, math.Pi/3,
		algebra.ViewTransform(0, 0, -10,
			0, 1, 0,
			0, 1, 0))
	if err != nil {
		panic(err)
		return err
	}
	image := cam.Render(w)
	imageToStr := image.ToPpmHeader(255) + image.ToPpmBody(255)
	err = writeToFile(imageToStr, "simpleScene")
	if err != nil {
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

	s := primitives.NewSphere(nil)
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
			is := primitives.NewIntersections()
			err = is.Intersect(s, r)
			if err != nil {
				return err
			}
			hit := is.Hit()
			if hit != nil {
				p := r.Position(hit.T)
				n := primitives.NormalAt(hit.Object, p, nil)
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
