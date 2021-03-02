package triangulation

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"log"
)

func FanTriangulation(indices []int, vertices []*algebra.Vector) []*primitives.Triangle{
	triangles := make([]*primitives.Triangle, 0, 0)
	for i := 1; i + 1 < len(indices); i++{
		if indices[i + 1] -1 >= len(vertices){
			log.Println("Triangulation index out of bounds %d versus %d", i + 1, len(vertices))
		} else {
			tri := primitives.NewTriangle(vertices[indices[0] -1], vertices[indices[i] -1], vertices[indices[i+1] -1])
			triangles = append(triangles, tri)
		}
	}
	return triangles
}
