package algebra

type Ray struct {
	origin    *Vector
	direction *Vector
}

//NewRay returns a 3D ray composed of a origin point Vector and a direction vector Vector
func NewRay(values ...float64) *Ray {
	res := make([]float64, 6, 6)
	for i := 0; i < len(values); i++ {
		if i == 6 {
			break
		}
		res[i] = values[i]
	}
	vector := NewVector(res[3:]...)
	point := NewPoint(res[0:3]...)
	return &Ray{origin: point, direction: vector}
}

//Get returns a map of the attributes of the Ray struct
func (r *Ray) Get() map[string]*Vector {
	res := make(map[string]*Vector)
	res["origin"] = r.origin
	res["direction"] = r.direction
	return res
}

//Position returns the position of the ray after distance t
func (r *Ray) Position(t float64) *Vector {
	p, err := r.origin.Add((r.direction).MultScalar(t))
	if err != nil {
		panic(err)
	}
	return p
}
