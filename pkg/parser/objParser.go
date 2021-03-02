package parser

import (
	"bufio"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/triangulation"
	"log"
	"os"
	"strconv"
	"strings"
)

type Parser struct{
	DefaultGroup *primitives.Group
	setGroup string
	Groups map[string] *primitives.Group
	Vertices []*algebra.Vector
}

//ToGeometry "exports" the parser to a single primitives.Group Shape
func (p *Parser) ToGeometry() *primitives.Group{
	g := primitives.NewGroup(nil)
	g.AddChild(p.DefaultGroup)
	for _, namedGroup := range p.Groups{
		g.AddChild(namedGroup)
	}
	return g
}

//ParseObjFile opens a file with the given path/name from the root directory (main.go)
func ParseObjFile(filePathName string) *Parser{
	parser := &Parser{Vertices: make([]*algebra.Vector, 0,0), DefaultGroup: primitives.NewGroup(nil),
						Groups: make(map[string]*primitives.Group), setGroup: ""}
	file, err := os.Open(filePathName)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		parseObjLine(scanner.Text(), parser)
	}
	if err := scanner.Err(); err != nil{
		log.Fatal(err)
	}
	return parser
}

func parseObjLine(line string, parser *Parser){
	if strings.HasPrefix(line, "v"){
		line = strings.TrimPrefix(line, "v")
		line = strings.TrimSpace(line)
		v := strings.Split(line, " ")
		createVertex(v, parser)
	} else if strings.HasPrefix(line, "f"){
		line = strings.TrimPrefix(line, "f")
		line = strings.TrimSpace(line)
		v := strings.Split(line, " ")
		if len(v) <= 3{
			createTriangle(v, parser)
		} else{
			createPolygon(v, parser)
		}
	} else if strings.HasPrefix(line, "g"){
		line = strings.TrimPrefix(line, "g")
		line = strings.TrimSpace(line)
		parser.setGroup = line
	}
}

func createVertex(v []string, parser *Parser){
	res := []float64{}
	for i := 0; i < 3; i++{
		if i >= len(v){
			res = append(res, 0.0)
		} else {

			if f, err := strconv.ParseFloat(v[i], 64); err != nil{
				log.Println("Could not parse float :", v[i])
				res = append(res, 0.0)
			} else{
				res = append(res, f)
			}
		}
	}
	parser.Vertices = append(parser.Vertices, algebra.NewPoint(res...))
}

func createTriangle(v []string, parser *Parser){
	vertices := []*algebra.Vector{}

	for i := 0; i < 3; i++{
		if index, err := strconv.Atoi(v[i]); err != nil{
			log.Println(err)
			p := algebra.NewPoint(0,0,0)
			vertices = append(vertices, p)

		} else{
			if index-1 >= len(parser.Vertices) || index -1 < 0{
				log.Printf("Parsed vertex index out of bounds %d versus %d", index - 1, len(parser.Vertices))
				p := algebra.NewPoint(0,0,0)
				vertices = append(vertices, p)
			} else{
				vertices = append(vertices, parser.Vertices[index-1])
			}
		}
	}
	tri := primitives.NewTriangle(vertices[0], vertices[1], vertices[2])
	if parser.setGroup == ""{
		parser.DefaultGroup.AddChild(tri)
	} else{
		if parser.Groups[parser.setGroup] == nil{
			parser.Groups[parser.setGroup] = primitives.NewGroup(nil)
			parser.Groups[parser.setGroup].AddChild(tri)
		}
	}
}

func createPolygon(v []string, parser *Parser){
	indices := make([]int, 0, 0)
	for _, val := range v{
		index, err := strconv.Atoi(val)
		if err != nil{
			log.Println(err)

		} else{
			indices = append(indices, index)
		}
	}

	triangles := triangulation.FanTriangulation(indices, parser.Vertices)
	if parser.setGroup == ""{
		for i := 0; i < len(triangles); i++{
			parser.DefaultGroup.AddChild(triangles[i])
		}
	} else {
		if parser.Groups[parser.setGroup] == nil{
			parser.Groups[parser.setGroup] = primitives.NewGroup(nil)
		}
		for i := 0; i < len(triangles); i++{
			parser.Groups[parser.setGroup].AddChild(triangles[i])
		}
	}
}