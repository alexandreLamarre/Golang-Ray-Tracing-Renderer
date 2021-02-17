package canvas

type Color [3]float64

//Add adds two colors together to return a new Color
func (c *Color) Add(c2 *Color) *Color{
	return &Color{c.Red()+c2.Red(), c.Green()+c2.Green(), c.Blue() + c2.Blue()}
}

//Subtract subtracts two colors together to return a new Color
func (c *Color) Subtract(c2 *Color) *Color{
	return &Color{c.Red()-c2.Red(), c.Green()-c2.Green(), c.Blue()-c2.Blue()}
}

//ScalarMult returns the scalar multiplication of a color
func (c *Color) ScalarMult(a float64) *Color{
	return &Color{a*c.Red(), a*c.Green(), a*c.Blue()}
}

//Multiply returns the hadamard product of a color
func Multiply(c *Color,c2 *Color) *Color{
	return &Color{c.Red()*c2.Red(), c.Green()*c2.Green(), c.Blue()*c2.Blue()}
}

//Red getter function for the red component of a Color
func (c * Color) Red() float64 {
	return c[0]
}

//Green getter function for the green component of a Color
func (c *Color) Green() float64{
	return c[1]
}

//Blue getter function for the blue component of a Color
func (c *Color) Blue() float64{
	return c[2]
}