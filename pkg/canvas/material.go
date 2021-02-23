package canvas

//Material encapsulates surface color but also lighting parameters of a surface
type Material struct {
	Color     *Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
	Pattern *Pattern
}

//NewDefaultMaterial creates a material with preset default values
func NewDefaultMaterial() *Material {
	return &Material{Color: &Color{1, 1, 1}, Ambient: 0.1, Diffuse: 0.9, Specular: 0.9, Shininess: 200, Pattern : nil}
}

//NewMaterial creates a new material from the provided parameters
func NewMaterial(color *Color, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{Color: color, Ambient: ambient, Diffuse: diffuse, Specular: specular, Shininess: shininess, Pattern: nil}
}
