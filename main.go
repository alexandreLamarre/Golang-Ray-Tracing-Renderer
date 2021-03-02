package main

import (
	"flag"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/examples"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/geometry/primitives"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/parser"
	"log"
)

func main() {
	useParser := flag.Bool("p", false, "Uses parser on file")
	fileNamePtr := flag.String("parsefile", "foo", "name of the file to open")
	exportPtr := flag.String("name", "", "name of the file to export to in pkg/examples")
	runExample := flag.Bool("e", false, "Runs specified example")
	flag.Parse()

	log.Println("==================== Golang ray tracer V 0.1 ====================")
	if *useParser{
		if *fileNamePtr == "foo"{
			log.Println("Incorrect -p usage: go run main.go -p parsefile=<string:filepath/filename.obj> export=<string:filename>")
			return
		}

		parseObj(*fileNamePtr,*exportPtr)

	} else if *runExample{
		log.Println("This should run an example")
	}
}

//parseObj is called on a fileName from the CLI if used with the -p tag
func parseObj(filePathName string, newName string){
	p := parser.ParseObjFile(filePathName)
	g := p.ToGeometry()

	w := &geometry.World{}
	objs := make([]primitives.Shape, 0, 0)
	objs = append(objs, g)
	lights := make([]*canvas.PointLight, 0, 0)
	lights = append(lights, canvas.NewPointLight(&canvas.Color{1, 1, 1}, algebra.NewPoint(-20, 10, -20)))
	w.Lights = lights
	w.Objects = objs

	if newName != ""{
		err := examples.CreateCustomScene(w, newName)
		if err != nil{
			log.Println(err)
		}
	} else {
		err := examples.CreateCustomScene(w, "example")
		if err != nil{
			log.Println(err)
		}
	}

}