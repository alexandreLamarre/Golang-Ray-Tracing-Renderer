package noise

import (
	"bufio"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPerlin(t *testing.T) {
	f, err := os.Open("perlinResults.txt")

	if err != nil{
		t.Errorf("%s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		s := scanner.Text()
		if  v, err := parsePerlinResults(s); err != nil{
			t.Errorf("%s", err)
		} else{
			x := v[0]
			y := v[1]
			z := v[2]
			val := v[3]
			noise := Perlin(x,y,z)
			if !equals(val, noise){
				t.Errorf("Expected noise at %f, %f, %f to be %f, instead got %f", x, y, z, val, noise)
			}
		}
	}

	if err := scanner.Err(); err != nil{
		t.Errorf("%s", err)
	}
}

func parsePerlinResults(s string) ([4]float64, error){
	vals :=strings.Split(s, ",")
	if len(vals) != 4{
		return [4]float64{}, algebra.ExpectedDimension(4)
	}
	x, err := strconv.ParseFloat(vals[0], 64)
	if err != nil{return [4]float64{}, err}
	y, err := strconv.ParseFloat(vals[1], 64)
	if err != nil{return [4]float64{}, err}
	z, err := strconv.ParseFloat(vals[2], 64)
	if err != nil{return [4]float64{}, err}
	value, err := strconv.ParseFloat(vals[3], 64)
	if err != nil{return [4]float64{}, err}

	return [4]float64{x, y, z, value}, nil
}

func equals(a, b float64) bool {
	EPSILON := 0.00001
	return math.Abs(a-b) < EPSILON
}