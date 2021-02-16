package algebra

type Vector struct {
	tuple []float64
}

//NewPoint creates a tuple with the data and appends a 1.0 to signify it is a point
func NewPoint(data []float64) *Vector {
	newData := append(data, 1.0)
	v := &Vector{tuple: newData}
	return v
}

//NewVector creates a tuple with the data and appends a 0.0 to signify it is point
func NewVector(data []float64) *Vector {
	newData := append(data, 0.0)
	v := &Vector{tuple: newData}
	return v
}

//IsPoint returns true iff the Vector is a point
func (v *Vector) IsPoint() bool {
	return v.tuple[len(v.tuple)-1] == 1.0
}

//IsVector returns true iff the Vector is a vector and returns that new vector
func (v *Vector) IsVector() bool {
	return v.tuple[len(v.tuple)-1] == 0.0
}

//Add adds two vectors together
func (v *Vector) Add(v2 *Vector) (*Vector, error) {

	if len(v.tuple) != len(v2.tuple) {
		return nil, MismatchedLength([2]int{len(v.tuple), len(v2.tuple)})
	}

	res := make([]float64, len(v.tuple), len(v.tuple))

	for i, val := range v.tuple {
		res[i] = val + v2.tuple[i]
	}


	return &Vector{res}, nil
}

//Subtract subtracts two vectors together and returns that new vector
func (v *Vector) Subtract(v2 *Vector) (*Vector, error) {

	if len(v.tuple) != len(v2.tuple) {
		return nil, MismatchedLength([2]int{len(v.tuple), len(v2.tuple)})
	}

	res := make([]float64, len(v.tuple), len(v.tuple))

	for i, val := range v.tuple {
		res[i] = val - v2.tuple[i]
	}

	return &Vector{res}, nil

}
