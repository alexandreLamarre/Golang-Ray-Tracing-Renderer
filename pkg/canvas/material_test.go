package canvas

import ("testing")

func TestNewDefaultMaterial(t *testing.T) {
	m := NewDefaultMaterial()
	testVectorEquals(t, m.Color, &Color{1,1,1})

	assertEquals(t, m.Ambient, 0.1)
	assertEquals(t, m.Diffuse, 0.9)
	assertEquals(t, m.Specular, 0.9)
	assertEquals(t, m.Shininess, 200.0)
}

func TestNewMaterial(t *testing.T) {
	m := NewMaterial(&Color{1,0,0}, 4.0, 5.0, 6.0, 1.0)
	testVectorEquals(t, m.Color, &Color{1,1,1})

	assertEquals(t, m.Ambient, 4.0)
	assertEquals(t, m.Diffuse,5.0)
	assertEquals(t, m.Specular, 6.0)
	assertEquals(t, m.Shininess, 1.0)
}

func assertEquals(t *testing.T, got, expected float64){
	if got != expected{
		t.Errorf("Expected %f, Got: %f", expected, got)
	}
}