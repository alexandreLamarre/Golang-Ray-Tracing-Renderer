package algebra

import(
	"math"
	)


type Matrix struct {
	tuples [][]float64
}

func NewMatrix(col, row int, data ...float64) (*Matrix, error) {
	if len(data) != col*row {
		return nil, ExpectedDimension(col * row)
	}

	tuples := make([][]float64, 0, 0)

	for i := 0; i < row; i++ {
		rows := make([]float64, col, col)
		for j := 0; j < col; j++ {
			rows[j] = data[i*col+j]
		}
		tuples = append(tuples, rows)
	}
	return &Matrix{tuples: tuples}, nil
}

func NewEmptyMatrix(col, row int) *Matrix {
	tuples := make([][]float64, 0, 0)

	for i := 0; i < row; i++ {
		rows := make([]float64, col, col)

		tuples = append(tuples, rows)
	}

	return &Matrix{tuples: tuples}
}

func (m *Matrix) Get() [][]float64 {
	return m.tuples
}


func (m *Matrix) At(col, row int) (float64, error) {
	if col < 0 || col > getNumCols(m) {
		return 0.0, InvalidMatrixIndex(col)
	}
	if row < 0 || row > getNumRows(m) {
		return 0.0, InvalidMatrixIndex(row)
	}
	return m.tuples[row][col], nil
}

func (m *Matrix) Equals(m2 * Matrix) bool{
	if getNumCols(m) != getNumCols(m2) || getNumRows(m) != getNumRows(m2){
		return false
	}
	for i:=0; i < getNumRows(m); i++{
		for j:=0; j < getNumCols(m); j++{
			v, err := m.At(i,j)
			if err != nil{
				return false
			}
			v2, err := m2.At(i,j)
			if err != nil{
				return false
			}
			if !equals(v, v2){
				return false
			}
		}
	}
	return true
}

func Multiply(m1 *Matrix, m2 *Matrix) (m *Matrix, err error){
	if getNumCols(m1) != getNumRows(m2){
		return nil, ExpectedDimension(len(m1.Get()[0]))
	}


	res := make([]float64, 0, 0)
	for i := 0; i < getNumRows(m1); i++{
		for j:=0; j < getNumCols(m2); j++{
			v := 0.0
			for k:=0; k < getNumRows(m2); k++{
				Aik, err := m1.At(k, i)
				if err != nil{
					return nil, err
				}
				Bkj, err := m2.At(j, k)
				if err !=nil{
					return nil, err
				}
				v += Aik * Bkj
			}
			res = append(res, v)

		}
	}
	return NewMatrix(getNumRows(m1), getNumCols(m2), res...)
}

func (m *Matrix) MultiplyByVec(v *Vector) (*Vector,error){
	if getNumCols(m) != len(v.tuple){
		return nil, ExpectedDimension(getNumRows(m))
	}

	res:= make([]float64, 0, 0)

	for i:=0; i < getNumRows(m); i++{
		val := 0.0
		for j:= 0; j < getNumCols(m); j++{
			a, err := m.At(j, i)
			if err != nil{
				return nil, err
			}
			val += a * v.tuple[j]
		}
		res = append(res, val)
	}

	return &Vector{tuple: res}, nil
}

func IdentityMatrix(size int) (*Matrix, error){
	res :=make([]float64, 0, 0)

	index := 0
	for i := 0; i < size; i++{
		for j:= 0; j < size; j++{
			if j == index {
				res = append(res, 1.0)
			} else{
				res = append(res, 0)
			}
		}
		index ++
	}

	return NewMatrix(size, size, res...)
}

func (m * Matrix) Transpose() (*Matrix, error){
	rows := getNumRows(m)
	cols := getNumCols(m)
	res := make([]float64, rows*cols, rows*cols)
	for i:= 0; i < rows; i++{
		for j:= 0; j < cols; j++{
			val, err := m.At(j,i)
			if err != nil{
				return nil, err
			}
			res[j*cols+i] = val
		}
	}
	return NewMatrix(cols, rows, res...)
}

func getNumRows(m *Matrix) int{
	return len(m.Get())
}

func getNumCols(m *Matrix) int{
	return len(m.Get()[0])
}

func equals(a,b float64) bool{
	EPSILON := 0.000001
	return math.Abs(a -b) < EPSILON
}