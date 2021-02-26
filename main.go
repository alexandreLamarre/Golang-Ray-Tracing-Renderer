package main

import (
	"fmt"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/examples"
	"log"
)

func main() {
	fmt.Println("Golang ray tracer V 0.1")
	//err := examples.CreateProjectileExample()
	//if err != nil {
	//	log.Println(err)
	//}
	//err := examples.CreateSphere3DExample()
	//if err != nil {
	//	log.Println(err)
	//}
	err := examples.TestReflectiveScene()
	if err != nil {
		log.Println(err)
	}
}
