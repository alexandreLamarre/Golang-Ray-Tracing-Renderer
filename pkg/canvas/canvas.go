package canvas

import (
	"strconv"
)

type Canvas struct {
	Width, Height int
	Pixels        [][]*Color
}

func NewCanvas(width, height int) *Canvas {
	pixels := make([][]*Color, 0, 0)
	for i := 0; i < height; i++ {
		col := make([]*Color, 0, 0)
		for j := 0; j < width; j++ {
			defaultEl := &Color{0.0, 0.0, 0.0}
			col = append(col, defaultEl)
		}
		pixels = append(pixels, col[:])
	}

	return &Canvas{Width: width, Height: height, Pixels: pixels}
}

func (c *Canvas) WritePixel(width, height int, color *Color) {
	c.Pixels[height][width] = color
}

//ToPpmHeader returns the string header format for the image displayed by a ppm file
// maxColorValue dictates the upper bound that colors between 0 and 1 should be scaled to
func (c *Canvas) ToPpmHeader(maxColorValue int) string {
	s := "P3\n"
	s += strconv.Itoa(c.Width) + " "
	s += strconv.Itoa(c.Height) + "\n"
	s += strconv.Itoa(maxColorValue) + "\n"
	return s
}

//ToPomBody returns the string body format for image displayed by a ppm file
// maxColorValue dictates the upper bound that colors between 0 and 1 should be scaled to
func (c *Canvas) ToPpmBody(maxColorValue int) string {
	res := ""
	for i := 0; i < c.Height; i++ {
		s := ""
		for j := 0; j < c.Width; j++ {
			p := c.Pixels[i][j]
			red, green, blue := clampValuesAndScale(p.Red(), p.Green(), p.Blue(), maxColorValue)
			s = addValuesToPpm(s, red, green, blue)
		}
		res += s
		res += "\n"
	}
	return res
}

func addValuesToPpm(s string, red, green, blue int) string {
	prev := len(s) % 70
	redS := strconv.Itoa(red)
	greenS := strconv.Itoa(green)
	blueS := strconv.Itoa(blue)
	if (prev+len(redS))%70 < prev {
		s += "\n"
		s += redS + " "
	} else {
		s += redS + " "
	}
	prev = len(s) % 70
	if (prev+len(greenS))%70 < prev {
		s += "\n"
		s += greenS + " "
	} else {
		s += greenS + " "
	}
	prev = len(s) % 70

	if (prev+len(blueS))%70 < prev {
		s += "\n"
		s += blueS + " "
	} else {
		s += blueS + " "
	}
	return s
}

func clampValuesAndScale(red, green, blue float64, maxColorValue int) (int, int, int) {
	m := float64(maxColorValue)
	rred := red
	rgreen := green
	rblue := blue
	if red < 0 {
		rred = 0
	}
	if green < 0 {
		rgreen = 0
	}
	if blue < 0 {
		rblue = 0
	}

	if red > 1 {
		rred = 1
	}
	if green > 1 {
		rgreen = 1
	}
	if blue > 1 {
		rblue = 1
	}

	return int(rred * m), int(rgreen * m), int(rblue * m)
}
