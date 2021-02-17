package examples

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"bufio"
	"fmt"
	"os"
)

func CreateProjectileExample() error{
	c := canvas.NewCanvas(900, 550)
	color := &canvas.Color{1, 0, 0}
	startVelocity, err := algebra.NewVector([]float64{1, 1.8, 0.0}).Normalize()
	if err != nil {
		return err
	}
	startVelocity = startVelocity.MultScalar(11.25)
	p := &projectile{position: algebra.NewPoint([]float64{0, 1, 0}), velocity: startVelocity}
	e := &environment{gravity: algebra.NewVector([]float64{0, -0.1, 0}), wind: algebra.NewVector([]float64{-0.01, 0, 0})}

	fmt.Println(c.ToPpmHeader(255))

	for p.position.Get()[1] > 0{
		err := tick(e, p)
		if err != nil{
			return err
		}
		x := int(p.position.Get()[0])
		y:= c.Height - int(p.position.Get()[1])
		if x >= 0 && x < c.Width && y >=0 && y < c.Height{
			c.WritePixel(x, y, color)
		}
	}
	 s := c.ToPpmHeader(255)
	 s += c.ToPpmBody(255)

	 err = writeToFile(s, "projectile")
	 if err != nil{
	 	return err
	 }
	return nil
}



func writeToFile(toWrite, fileName string) error{
	f, err := os.Create("./pkg/examples/"+fileName+".ppm")
	if err != nil{
		return err
	}

	w := bufio.NewWriter(f)

	n, err := w.WriteString(toWrite)
	fmt.Printf("Wrote %d bytes\n", n)
	if err != nil{
		return err
	}
	err = f.Close()
	if err != nil{
		return err
	}
	return nil
}

