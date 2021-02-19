package canvas

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
)

//PointLight defines a light without size described by an Intensity color and a Position vector(point)
type PointLight struct{
	Intensity *Color
	Position *algebra.Vector
}

//NewPointLight returns a new PointLight with attributes provided as parameters
func NewPointLight(color *Color, position *algebra.Vector) *PointLight{
	return &PointLight{Intensity: color, Position: position}
}

//Lighting computes the lighting from the PointLight onto the Material at the illuminatedPoint with its normal Vector
// from the point of view of the eye vector
func Lighting(material *Material, light *PointLight, illuminatedPoint, eyeVector, normalVector *algebra.Vector) *Color{

	effectiveColor := Multiply(material.Color, light.Intensity)

	lightVector, err := light.Position.Subtract(illuminatedPoint)
	if err != nil{panic(err); return nil}
	lightVector, err = lightVector.Normalize()
	if err != nil{panic(err); return nil}

	ambient := effectiveColor.ScalarMult(material.Ambient)
	lightDotNormal, err := algebra.DotProduct(lightVector, normalVector)
	if err != nil{panic(err); return nil}
	var diffuse *Color
	var specular *Color
	if lightDotNormal < 0{
		diffuse = &Color{0,0,0}
		specular = &Color{0,0,0}
	} else{
		diffuse = effectiveColor.ScalarMult(material.Diffuse).ScalarMult(lightDotNormal)

		reflectVector := lightVector.Negate().Reflect(normalVector)
		reflectDotEye, err := algebra.DotProduct(reflectVector, eyeVector)
		if err != nil{panic(err); return nil}

		if reflectDotEye < 0 {
			specular = &Color{0,0,0}
		} else{
			factor := math.Pow(reflectDotEye, material.Shininess)
			specular = light.Intensity.ScalarMult(material.Specular).ScalarMult(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}