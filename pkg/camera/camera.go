package camera

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"math"
)

//Camera describes a camera object that renders pixels from the setup scene
type Camera struct {
	hSize      float64 // horizontal size of picture to be rendered in pixels
	vSize      float64 // vertical size of picture to be rendered in pixels
	fov        float64 //determines the zoom of the camera
	halfWidth  float64
	halfHeight float64
	pixelSize  float64
	transform  *algebra.Matrix
}

//NewDefaultCamera returns a new camera with the given size and fov, that has 4x4 identity matrix as
// the camera transformation
func NewDefaultCamera(hSize, vSize, fov float64) *Camera {
	halfView := math.Tan(fov / 2)
	aspect := hSize / vSize
	var halfWidth float64
	var halfHeight float64
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	pixelSize := (halfWidth * 2) / hSize
	return &Camera{hSize: hSize, vSize: vSize, fov: fov,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		pixelSize:  pixelSize,
		transform:  algebra.IdentityMatrix(4)}
}

//NewCamera return a new camera with the given size and fov and the provided 4x4 transform matrix
func NewCamera(hSize, vSize, fov float64, transform *algebra.Matrix) (*Camera, error) {
	if len(transform.Get()) != 4 || len(transform.Get()[0]) != 4 {
		return nil, algebra.ExpectedDimension(4)
	}
	halfView := math.Tan(fov / 2)
	aspect := hSize / vSize
	var halfWidth float64
	var halfHeight float64
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}
	pixelSize := (halfWidth * 2) / hSize
	return &Camera{hSize: hSize, vSize: vSize, fov: fov,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
		pixelSize:  pixelSize,
		transform:  transform}, nil
}

func (c Camera) RayForPixel(px, py float64) *algebra.Ray {
	xOffset := (px + 0.5) * c.pixelSize
	yOffset := (py + 0.5) * c.pixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	intermediate := c.transform.Inverse()
	pixel := intermediate.MultiplyByVec(algebra.NewPoint(worldX, worldY, -1))
	origin := intermediate.MultiplyByVec(algebra.NewPoint(0, 0, 0))
	direction, err := pixel.Subtract(origin)
	if err != nil {
		panic(err)
		return nil
	}
	direction, err = direction.Normalize()
	if err != nil {
		panic(err)
		return nil
	}
	rDirection := direction.Get()[:3]
	rOrigin := origin.Get()[:3]
	res := append(rOrigin, rDirection...)
	return algebra.NewRay(res...)
}

func (c Camera) Render(w *geometry.World) *canvas.Canvas{
	image := canvas.NewCanvas(int(c.hSize), int(c.vSize))

	for y := 0.0; y < c.vSize; y++ {
		for x := 0.0; x < c.hSize; x++ {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray)
			image.WritePixel(int(x), int(y), color)
		}
	}
	return image
}
