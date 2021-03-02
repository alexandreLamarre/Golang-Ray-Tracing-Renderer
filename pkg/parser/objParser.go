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
	"time"
)

//MAXGROUPSIZE determines when to split groups into smaller groups
var MAXGROUPSIZE int = 10

type Parser struct {
	DefaultGroup   *primitives.Group
	setGroup       string
	Groups         map[string]*primitives.Group
	Vertices       []*algebra.Vector
	NormalVertices []*algebra.Vector
}

//ToGeometry "exports" the parser to a single primitives.Group Shape
func (p *Parser) ToGeometry() *primitives.Group {
	numShapes := p.DefaultGroup.NumShapes()
	for _, namedGroup := range p.Groups {
		numShapes += namedGroup.NumShapes()
	}
	log.Printf("Optimizing parsed Shapes (%d)...", numShapes)
	start := time.Now()
	g := primitives.NewGroup(nil)
	optimizedDefaultGroup := optimize(p.DefaultGroup)
	g.AddChild(optimizedDefaultGroup)

	for _, namedGroup := range p.Groups {
		optimizedNamedGroup := optimize(namedGroup)
		g.AddChild(optimizedNamedGroup)
	}

	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Done (%s)!", elapsed)
	return g
}

func optimize(g *primitives.Group) *primitives.Group {
	if g.NumShapes() < MAXGROUPSIZE {
		return g
	}

	numShapes := 0
	for g.NumShapes() > MAXGROUPSIZE {
		optimizedG := primitives.NewGroup(nil)
		tempGroup := primitives.NewGroup(nil)
		for _, s := range g.GetShapes() {
			tempGroup.AddChild(s)
			numShapes++
			if numShapes%MAXGROUPSIZE == 0 {
				optimizedG.AddChild(tempGroup)
				tempGroup = primitives.NewGroup(nil)
			}
		}
		if tempGroup.NumShapes() != 0 {
			optimizedG.AddChild(tempGroup)
		}
		g = optimizedG
	}

	return g
}

//ParseObjFile opens a file with the given path/name from the root directory (main.go)
func ParseObjFile(filePathName string) *Parser {
	log.Println("Opening file : " + filePathName + "...")
	start := time.Now()
	parser := &Parser{Vertices: make([]*algebra.Vector, 0, 0), DefaultGroup: primitives.NewGroup(nil),
		Groups: make(map[string]*primitives.Group), setGroup: ""}
	file, err := os.Open(filePathName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	log.Println("Parsing Object file...")
	for scanner.Scan() {
		parseObjLine(scanner.Text(), parser)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	log.Printf("Done(%s)!", elapsed)
	return parser
}

func parseObjLine(line string, parser *Parser) {
	words := strings.Fields(line)
	if len(words) != 0 {
		if words[0] == "v" {
			//parse a vertex
			v := words[1:]
			if len(v) == 3 {
				createVertex(words[1:], parser)
			} else {
				log.Println("Warning could not parse vertex from :", words)
			}
		} else if words[0] == "vn" {
			//parse a normal to a vertex

			v := words[1:]
			if len(v) == 3 {
				createNormalVertex(words[1:], parser)
			} else {
				log.Println("Warning could not parse vertex normal from :", words)
			}

		} else if words[0] == "vt" {
			//Parse textures, do nothing for now since we dont have abstract textures implemented

		} else if words[0] == "f" {
			// parse a triangle/polygon
			v := words[1:]
			if len(v) > 0 {
				if strings.Contains(words[1], "/") {
					//normalized vertex face
					if len(v) == 3 {
						createSmoothTriangle(v, parser)
					} else if len(v) > 3 {
						createSmoothPolygon(v, parser)
					} else {
						log.Printf("Warning: could not parse triangle from : %s", words)
					}
				} else {
					//non-normalized face
					if len(v) == 3 {
						createTriangle(v, parser)
					} else if len(v) > 3 {
						createPolygon(v, parser)
					} else {
						log.Printf("Warning: could not parse triangle from : %s", words)
					}
				}
			} else {
				log.Printf("Warning: could not parse triangle from : %s", words)
			}
		} else if words[0] == "g" { //parse
			if len(words) > 1 {
				parser.setGroup = words[1]
			}
		}
	}
}

func createVertex(v []string, parser *Parser) {
	res := []float64{}
	for i := 0; i < 3; i++ {

		if f, err := strconv.ParseFloat(v[i], 64); err != nil {
			log.Println("Warning : Could not parse float :", v[i])
			res = append(res, 0.0)
		} else {
			res = append(res, f)
		}

	}
	parser.Vertices = append(parser.Vertices, algebra.NewPoint(res...))
}

func createNormalVertex(v []string, parser *Parser) {
	res := []float64{}
	for i := 0; i < 3; i++ {

		if f, err := strconv.ParseFloat(v[i], 64); err != nil {
			log.Println("Warning : Could not parse float :", v[i])
			res = append(res, 0.0)
		} else {
			res = append(res, f)
		}

	}
	parser.NormalVertices = append(parser.NormalVertices, algebra.NewPoint(res...))
}

func createTriangle(v []string, parser *Parser) {
	vertices := []*algebra.Vector{}

	for i := 0; i < 3; i++ {
		if index, err := strconv.Atoi(v[i]); err != nil {
			log.Println(err)
			p := algebra.NewPoint(0, 0, 0)
			vertices = append(vertices, p)

		} else {
			if index-1 >= len(parser.Vertices) || index-1 < 0 {
				log.Printf("Warning: Parsed vertex index out of bounds %d versus %d", index-1, len(parser.Vertices))
				p := algebra.NewPoint(0, 0, 0)
				vertices = append(vertices, p)
			} else {
				vertices = append(vertices, parser.Vertices[index-1])
			}
		}
	}
	tri := primitives.NewTriangle(vertices[0], vertices[1], vertices[2])
	if parser.setGroup == "" {
		parser.DefaultGroup.AddChild(tri)
	} else {
		if parser.Groups[parser.setGroup] == nil {
			parser.Groups[parser.setGroup] = primitives.NewGroup(nil)
			parser.Groups[parser.setGroup].AddChild(tri)
		}
	}
}

func createSmoothTriangle(v []string, parser *Parser) {
	//TODO: implement smooth triangle parsing
}

func createPolygon(v []string, parser *Parser) {
	indices := make([]int, 0, 0)
	for _, val := range v {
		index, err := strconv.Atoi(val)
		if err != nil {
			log.Println(err)

		} else {
			indices = append(indices, index)
		}
	}

	triangles := triangulation.FanTriangulation(indices, parser.Vertices)
	if parser.setGroup == "" {
		for i := 0; i < len(triangles); i++ {
			parser.DefaultGroup.AddChild(triangles[i])
		}
	} else {
		if parser.Groups[parser.setGroup] == nil {
			parser.Groups[parser.setGroup] = primitives.NewGroup(nil)
		}
		for i := 0; i < len(triangles); i++ {
			parser.Groups[parser.setGroup].AddChild(triangles[i])
		}
	}
}

func createSmoothPolygon(v []string, parser *Parser) {
	//TODO: implement smooth polygon parsing
}
