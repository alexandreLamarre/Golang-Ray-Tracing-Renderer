package canvas

import (
	"strings"
	"testing"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)

	testZeroCanvas(t, c)
}

func TestWritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	color := &Color{1, 0, 0}
	c.WritePixel(2, 3, color)

	testVectorEquals(t, c.Pixels[3][2], color)
}

func TestCanvas_ToPpmHeader(t *testing.T) {
	c := NewCanvas(10, 20)
	header := c.ToPpmHeader(255)
	testStringEquals(t, header, "P3\n10 20\n255\n")
}

func TestCanvasHelpers(t *testing.T) {
	red, green, blue := -1.0, -1.0, -1.0

	rred, rgreen, rblue := clampValuesAndScale(red, green, blue, 255)
	if !(rred == 0 && rgreen == 0 && rblue == 0) {
		t.Errorf("Expected all to be 0, Got: %d %d %d", rred, rgreen, rblue)
	}

	red, green, blue = 2.0, 2.0, 2.0
	rred, rgreen, rblue = clampValuesAndScale(red, green, blue, 255)
	if !(rred == 255 && rgreen == 255 && rblue == 255) {
		t.Errorf("Expected all to be 255, Got: %d %d %d", rred, rgreen, rblue)
	}
}

func TestCanvas_ToPpmBody(t *testing.T) {
	c := NewCanvas(5, 3)
	c1 := &Color{1.5, 0.0, 0.0}
	c2 := &Color{0, 0.5, 0}
	c3 := &Color{-0.5, 0, 1}

	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)

	body := c.ToPpmBody(255)
	testStringEquals(t, body, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 127 0 0 0 0 0 0 0 \n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255 \n")

	c = NewCanvas(10, 2)
	c1 = &Color{1, 0.8, 0.6}

	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			c.WritePixel(j, i, c1)
		}
	}
	body = c.ToPpmBody(255)
	testStringEquals(t, body, "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 \n153 255 204 153 255 204 153 255 204 153 255 204 153 \n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204 \n153 255 204 153 255 204 153 255 204 153 255 204 153 \n")
}

func testZeroCanvas(t *testing.T, c *Canvas) {
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			zeroColor := &Color{0, 0, 0}
			testVectorEquals(t, c.Pixels[i][j], zeroColor)
		}
	}
}

func testStringEquals(t *testing.T, value string, expected string) {
	if len(value) != len(expected) {
		t.Errorf("Mismatched strings length %d, %d", len(value), len(expected))
	}

	if strings.Trim(value, "\r\n") != strings.Trim(expected, "\r\n") {
		t.Errorf("Expected: \n %s Got: \n %s", expected, value)
	}

}
