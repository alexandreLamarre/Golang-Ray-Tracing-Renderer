package parser

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"testing"
)

func TestParseObjFile(t *testing.T) {
	ParseObjFile("./gibberish_test.obj")

	p := ParseObjFile("./vertex_test.obj")
	if len(p.Vertices) != 4 {
		t.Errorf("Expected to parse 4 vertices from %s, got : %d", "./vertex_test.obj", len(p.Vertices))
	}
	testVectorEquals(t, p.Vertices[0].Get(), algebra.NewPoint(-1, 1, 0).Get())
	testVectorEquals(t, p.Vertices[1].Get(), algebra.NewPoint(-1, 0.5, 0).Get())
	testVectorEquals(t, p.Vertices[2].Get(), algebra.NewPoint(1, 0, 0).Get())
	testVectorEquals(t, p.Vertices[3].Get(), algebra.NewPoint(1, 1, 0).Get())

	p = ParseObjFile("./triangle_test.obj")

	if len(p.DefaultGroup.GetShapes()) != 2{
		t.Errorf("Expected to parse 2 shapes in default group from %s, got : %d", "./triangle_test", len(p.DefaultGroup.GetShapes()))
	}

	p = ParseObjFile("./polygon_test.obj")

	if len(p.Vertices) != 5{
		t.Errorf("Expected to parse 5 vertices from %s, got : %d", "./polygon_test.obj", len(p.Vertices))
	}

	if len(p.DefaultGroup.GetShapes()) != 3{
		t.Errorf("Expected to parse 3 shapes in default group from %s, got : %d", "./polygon_test.obj", len(p.DefaultGroup.GetShapes()))
	}

	p = ParseObjFile("./trianglegroup_test.obj")

	if len(p.Vertices) != 4{
		t.Errorf("Expected to parse 4 vertices from %s, got : %d", "./trianglegroup_test.obj", len(p.Vertices))
	}

	if len(p.DefaultGroup.GetShapes()) != 0{
		t.Errorf("Expected no default group shapes in ./trianglegroup_test.obj")
	}

	if len(p.Groups) != 2{
		t.Errorf("Expected to parse 2 custom groups from %s, got : %d", "./trianglegroup_test.obj", len(p.Groups))
	}

	if p.Groups["FirstGroup"] == nil{
		t.Errorf("Expected 'FirstGroup' to be a name group in ./trianglegroup_test.obj")
	}

	if p.Groups["SecondGroup"] == nil{
		t.Errorf("Expected 'SecondGroup' to be a name group in ./trianglegroup_test.obj")
	}

}

func TestParser_ToGeometry(t *testing.T) {

}