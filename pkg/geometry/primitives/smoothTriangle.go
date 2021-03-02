package primitives

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/canvas"
	"math"
)

//SmoothTriangle datastructure that handles calculations for SmoothTriangle Shape
type SmoothTriangle struct{
	Parent Shape
	transform *algebra.Matrix
	material *canvas.Material
	p1 *algebra.Vector
	p2 *algebra.Vector
	p3 *algebra.Vector
	e1 *algebra.Vector
	e2 *algebra.Vector
	n1 *algebra.Vector
	n2 *algebra.Vector
	n3 *algebra.Vector
}

//NewSmoothTriangle Initializer for fully specificed Smooth Triangle Shape
// n1, n2, n3 are the normals used for the vertices p1, p2, p3
func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 *algebra.Vector) *SmoothTriangle{
	e1, err := p2.Subtract(p1)
	if err != nil{
		panic(err)
	}
	e2, err := p3.Subtract(p1)
	if err != nil{
		panic(err)
	}
	normal, err := algebra.CrossProduct(e2, e1)

	if err != nil{
		panic(err)
	}
	normal, err = normal.Normalize()
	if err != nil{
		panic(err)
	}

	return &SmoothTriangle{
		p1:p1, p2: p2, p3:p3,
		e1: e1, e2: e2,
		material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4),
		Parent: nil,
		n1:n1,
		n2:n2,
		n3:n3,
	}
}


//NewDefaultSmoothTriangle Initializer for Smooth Triangle Shape and Infer Normals at vertices
// based on cross products
func NewDefaultSmoothTriangle(p1, p2, p3 *algebra.Vector) *SmoothTriangle{
	a := p1
	b := p2
	c := p3

	ab, err := b.Subtract(a)
	if err != nil{
		panic(err)
	}
	ac, err := c.Subtract(a)
	if err != nil{
		panic(err)
	}


	abXac, err := algebra.CrossProduct(ab, ac)
	if err != nil{
		panic(err)
	}
	cross1, err := algebra.CrossProduct(abXac, ab)
	if err != nil{
		panic(err)
	}
	cross1 = cross1.MultScalar(ac.Magnitude2())

	cross2, err := algebra.CrossProduct(ac, abXac)
	cross2 = cross2.MultScalar(ab.Magnitude2())

	toCircumCenter, err := cross1.Add(cross2)
	if err != nil{
		panic(err)
	}

	toCircumCenter = toCircumCenter.DivideScalar(2*abXac.Magnitude2())

	circumcenter, err := a.Add(toCircumCenter)
	if err != nil{
		panic(err)
	}
	//next we calculate the vector normals of this sphere passing through each vertex
	n1, err := p1.Subtract(circumcenter)
	if err != nil{
		panic(err)
	}
	n1, err = n1.Normalize()
	if err != nil{
		panic(err)
	}

	n2, err := p2.Subtract(circumcenter)
	if err != nil{
		panic(err)
	}
	n2, err = n2.Normalize()
	if err != nil{
		panic(err)
	}

	n3, err := p3.Subtract(circumcenter)
	if err != nil{
		panic(err)
	}
	n3, err = n3.Normalize()
	if err != nil{
		panic(err)
	}



	return &SmoothTriangle{
		p1:a, p2: b, p3:c,
		e1: ab, e2: ac,
		material: canvas.NewDefaultMaterial(),
		transform: algebra.IdentityMatrix(4),
		Parent: nil,
		n1:n1,
		n2:n2,
		n3:n3,
	}
}

//Shape Interface Methods

//GetTransform Getter for transform of SmoothTriangle Shape, interface method
func (t *SmoothTriangle) GetTransform() *algebra.Matrix{
	return t.transform
}

//GetMaterial Getter for material of SmoothTriangle Shape, interface method
func (t *SmoothTriangle) GetMaterial() *canvas.Material{
	return t.material
}

//GetParent Getter for parent of SmoothTriangle Shape, interface method
func (t *SmoothTriangle) GetParent() Shape{
	return t.Parent
}

//GetBounds Getter for bounds of SmoothTriangle Shape, interface method
func (t *SmoothTriangle) GetBounds() (*algebra.Vector, *algebra.Vector){
	var xMin = math.Inf(1); var yMin = math.Inf(1); var zMin = math.Inf(1)
	var xMax = math.Inf(-1); var yMax = math.Inf(-1); var zMax = math.Inf(-1)
	points := []*algebra.Vector{t.p1, t.p2, t.p3}
	for _, p := range points{
		xMin = math.Min(p.Get()[0], xMin); yMin = math.Min(p.Get()[1], yMin); zMin = math.Min(p.Get()[2], zMin)
		xMax = math.Max(p.Get()[0], xMax); yMax = math.Max(p.Get()[1], yMax); zMax = math.Max(p.Get()[2], zMax)
	}
	return algebra.NewPoint(xMin, yMin, zMin), algebra.NewPoint(xMax, yMax, zMax)
}

//SetTransform Setter for SmoothTriangle Shape transform, interface method
func (t *SmoothTriangle) SetTransform(m *algebra.Matrix){
	if len(m.Get()) != 4 || len(m.Get()[0]) != 4{
		panic(algebra.ExpectedDimension(4))
	}
	t.transform = m
}

//SetMaterial Setter for SmoothTriangle Shape material, interface method
func (t *SmoothTriangle) SetMaterial(m *canvas.Material){
	t.material = m
}

//SetParent Setter for SmoothTriangle Shape parent, interface method
func (t *SmoothTriangle) SetParent(s Shape){
	t.Parent = s
}

//LocalIntersect Intersect implementation for SmoothTriangle Shape, interface method
func (t *SmoothTriangle) LocalIntersect(r *algebra.Ray) ([]*Intersection, bool){
	xs := make([]*Intersection, 0, 0)
	direction := r.Get()["direction"]
	dirCrossProduct, err := algebra.CrossProduct(direction, t.e2)
	if err != nil {panic(err)}
	det, err := algebra.DotProduct(t.e1, dirCrossProduct)
	if err != nil{panic(err)}
	if math.Abs(det) <= algebra.EPSILON{
		return xs, false
	}

	origin := r.Get()["origin"]
	f := 1.0/det
	p1ToOrigin, err := origin.Subtract(t.p1)
	if err != nil{
		panic(err)
	}
	u, err := algebra.DotProduct(p1ToOrigin, dirCrossProduct)
	if err != nil{panic(err)}
	u *= f
	if u < 0 || u > 1{
		return xs, false
	}

	originCross, err := algebra.CrossProduct(p1ToOrigin, t.e1)
	if err != nil{
		panic(err)
	}
	v, err:= algebra.DotProduct(direction, originCross)
	v *= f
	if err != nil{panic(err)}
	if v < 0 || (u+v) > 1{
		return xs, false
	}

	pos, err := algebra.DotProduct(t.e2, originCross)
	if err != nil{
		panic(err)
	}

	pos *= f
	i := NewIntersection(t, pos)
	i.SetUV(u, v)
	xs = append(xs, i)
	return xs, true
}

//LocalNormalAt normal interpolation method for SmoothTriangle Shape, interface method
func (t *SmoothTriangle) LocalNormalAt(p *algebra.Vector, hit *Intersection) (*algebra.Vector, error){
	temp ,err := t.n2.MultScalar(hit.U).Add(t.n3.MultScalar(hit.V))
	if err != nil{
		panic(err)
	}
	temp, err = temp.Add(t.n1.MultScalar(1- hit.U -hit.V))
	return temp , nil
}