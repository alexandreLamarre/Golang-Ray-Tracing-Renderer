package triangulation

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"log"
)

//FanTriangulation triangulates polygon faces parsed from .obj files
func FanTriangulation(indices []int, vertices []*algebra.Vector) []primitives.Shape {
	triangles := make([]primitives.Shape, 0, 0)
	for i := 1; i+1 < len(indices); i++ {
		if indices[i+1]-1 >= len(vertices) {
			log.Printf("Warning: Triangulation index out of bounds %d versus %d", indices[i + 1]-1, len(vertices))
		} else {
			tri := primitives.NewTriangle(vertices[indices[0]-1], vertices[indices[i]-1], vertices[indices[i+1]-1])
			triangles = append(triangles, tri)
		}
	}
	return triangles
}

//SmoothFanTriangulation triangulates smooth polygon faces from .obj files
func SmoothFanTriangulation(indices []int, normalIndices []int,
	vertices []*algebra.Vector, normalVertices []*algebra.Vector) []primitives.Shape {
	striangles := make([]primitives.Shape, 0, 0)

	if len(indices) != len(normalIndices){
		log.Printf("Mismatched number of vertex indices %d and normal vertex indices %d",
			len(indices), len(normalIndices))
		return striangles
	}

	for i := 1; i+1< len(indices); i++{
		if indices[i + 1] -1 >= len(vertices) || normalIndices[i + 1] -1 >= len(normalVertices){
			log.Printf("Warning: Triangulation index out of bounds %d versus %d", indices[i + 1] -1, len(vertices)-1)
			log.Printf("Warning: Triangulation normal index out of bounds %d versus %d", normalIndices[i + 1] -1,
				len(normalVertices)-1)
		} else{
			tri := primitives.NewSmoothTriangle(vertices[indices[0] -1], vertices[indices[i] - 1], vertices[indices[i + 1] -1],
				normalVertices[normalIndices[0] -1], normalVertices[normalIndices[i] -1], normalVertices[normalIndices[i +1] -1])
			striangles = append(striangles, tri)
		}
	}
	return striangles
}