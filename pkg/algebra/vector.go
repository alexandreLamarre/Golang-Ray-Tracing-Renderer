package algebra

import "math"

type Vector struct {
	tuple []float64
}

//NewPoint creates a tuple with the data and appends a 1.0 to signify it is a point
func NewPoint(data ...float64) *Vector {
	newData := append(data, 1.0)
	v := &Vector{tuple: newData}
	return v
}

//NewVector creates a tuple with the data and appends a 0.0 to signify it is point
func NewVector(data ...float64) *Vector {
	newData := append(data, 0.0)
	v := &Vector{tuple: newData}
	return v
}

//Get getter function for the vector tuple
func (v *Vector) Get() []float64 {
	return v.tuple
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

//Negate returns -vector
func (v *Vector) Negate() *Vector {

	res := make([]float64, len(v.tuple), len(v.tuple))

	for i := 0; i < len(v.tuple); i++ {
		if i == len(v.tuple)-1 {
			res[i] = v.tuple[i] // keep attribute that keeps track of pointedness/vectoredness the same
		} else {
			res[i] = -v.tuple[i]
		}
	}

	return &Vector{res}
}

//MultScalar performs vector scalar multiplication
func (v *Vector) MultScalar(c float64) *Vector {

	res := make([]float64, len(v.tuple), len(v.tuple))

	for i, val := range v.tuple {
		if i == len(v.tuple)-1 {
			res[i] = val // keep attribute that keeps track of pointedness/vectoredness the same
		} else {
			res[i] = c * val
		}
	}

	return &Vector{res}
}

//DivideScalar performs vector scalar division. (Notationally more convenient sometimes than MultScalar method)
func (v *Vector) DivideScalar(c float64) *Vector {
	res := make([]float64, len(v.tuple), len(v.tuple))

	for i, val := range v.tuple {
		if i == len(v.tuple)-1 {
			res[i] = val
		} else {
			res[i] = (1 / c) * val
		}
	}
	return &Vector{res}
}

//Magnitude returns the provided vectors magnitude
func (v *Vector) Magnitude() float64 {
	sum := 0.0
	for i, val := range v.tuple {
		if i == len(v.tuple)-1 {
			continue
		} else {
			sum += math.Pow(val, 2)
		}

	}
	return math.Sqrt(sum)
}

//Normalize returns a normalized version of the original vector
func (v *Vector) Normalize() (*Vector, error) {
	norm := v.Magnitude()
	if norm == 0 {
		return nil, ZeroDivide(0)
	}

	res := make([]float64, len(v.tuple), len(v.tuple))

	for i, val := range v.tuple {
		if i == len(v.tuple)-1 {
			res[i] = val
		} else {
			res[i] = val / norm
		}
	}

	return &Vector{res}, nil
}

//DotProduct returns the dot product of two vectors
func DotProduct(v *Vector, v2 *Vector) (float64, error) {
	if len(v.tuple) != len(v2.tuple) {
		return 0.0, MismatchedLength([2]int{len(v.tuple), len(v2.tuple)})
	}

	sum := 0.0
	for i, val := range v.tuple {
		sum += val * v2.tuple[i]
	}
	return sum, nil
}

func CrossProduct(v1 *Vector, v2 *Vector) (*Vector, error) {
	if len(v1.tuple) != len(v2.tuple) {
		return nil, MismatchedLength([2]int{len(v1.tuple), len(v2.tuple)})
	}

	if len(v1.tuple) != 4 {
		return nil, ExpectedDimension(3)
	}
	a := v1.tuple
	b := v2.tuple
	return NewVector(a[1]*b[2]-a[2]*b[1], a[2]*b[0]-a[0]*b[2], a[0]*b[1]-a[1]*b[0]), nil
}
