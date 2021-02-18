package algebra

import (
	"math"
)

type Matrix struct {
	tuples [][]float64
}

//NewMatrix returns a new matrix of the specified size and data
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

//NewEmptyMatrix returns a new zero matrix of the specified size
func NewEmptyMatrix(col, row int) *Matrix {
	tuples := make([][]float64, 0, 0)

	for i := 0; i < row; i++ {
		rows := make([]float64, col, col)

		tuples = append(tuples, rows)
	}

	return &Matrix{tuples: tuples}
}

//Get returns the slices that contain the matrix data
func (m *Matrix) Get() [][]float64 {
	return m.tuples
}

// At returns the value at position (col, row)
func (m *Matrix) At(col, row int) (float64, error) {
	if col < 0 || col > getNumCols(m) {
		return 0.0, InvalidMatrixIndex(col)
	}
	if row < 0 || row > getNumRows(m) {
		return 0.0, InvalidMatrixIndex(row)
	}
	return m.tuples[row][col], nil
}

//Equals returns whether or not two matrices are equals
func (m *Matrix) Equals(m2 *Matrix) bool {
	if getNumCols(m) != getNumCols(m2) || getNumRows(m) != getNumRows(m2) {
		return false
	}
	for i := 0; i < getNumRows(m); i++ {
		for j := 0; j < getNumCols(m); j++ {
			v, err := m.At(i, j)
			if err != nil {
				return false
			}
			v2, err := m2.At(i, j)
			if err != nil {
				return false
			}
			if !equals(v, v2) {
				return false
			}
		}
	}
	return true
}

//Multiply multiplies two matrices together
func Multiply(m1 *Matrix, m2 *Matrix) *Matrix {
	if getNumCols(m1) != getNumRows(m2) {
		panic(ExpectedDimension(len(m1.Get()[0])))
		return nil
	}

	res := make([]float64, 0, 0)
	for i := 0; i < getNumRows(m1); i++ {
		for j := 0; j < getNumCols(m2); j++ {
			v := 0.0
			for k := 0; k < getNumRows(m2); k++ {
				Aik, err := m1.At(k, i)
				if err != nil {
					panic(err)
					return nil
				}
				Bkj, err := m2.At(j, k)
				if err != nil {
					panic(err)
					return nil
				}
				v += Aik * Bkj
			}
			res = append(res, v)

		}
	}
	mMultiplied, err := NewMatrix(getNumRows(m1), getNumCols(m2), res...)
	if err != nil {
		panic(err)
		return nil
	}
	return mMultiplied
}

//MultiplyByVec returns left matrix multiplication with the given column vector
func (m *Matrix) MultiplyByVec(v *Vector) *Vector {
	if getNumCols(m) != len(v.tuple) {

		panic(ExpectedDimension(getNumRows(m)))
		return nil
	}

	res := make([]float64, 0, 0)

	for i := 0; i < getNumRows(m); i++ {
		val := 0.0
		for j := 0; j < getNumCols(m); j++ {
			a, err := m.At(j, i)
			if err != nil {
				panic(err)
				return nil
			}
			val += a * v.tuple[j]
		}
		res = append(res, val)
	}

	return &Vector{tuple: res}
}

//IdentityMatrix returns the identity matrix with the given size
func IdentityMatrix(size int) *Matrix {
	res := make([]float64, 0, 0)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j == i {
				res = append(res, 1.0)
			} else {
				res = append(res, 0)
			}
		}
	}
	iden, err := NewMatrix(size, size, res...)
	if err != nil {
		panic(err)
	}
	return iden
}

//Transpose returns the transpose of the matrix
func (m *Matrix) Transpose() *Matrix {
	rows := getNumRows(m)
	cols := getNumCols(m)
	res := make([]float64, rows*cols, rows*cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			val, err := m.At(j, i)
			if err != nil {
				panic(err)
				return nil
			}
			res[j*cols+i] = val
		}
	}
	mTranspose, err := NewMatrix(cols, rows, res...)
	if err != nil {
		panic(err)
		return nil
	}
	return mTranspose
}

//Determinant returns the determinant of a square matrix
func Determinant(m [][]float64) (float64, error) {
	rows := len(m)
	cols := len(m[0])
	if rows != cols {
		return 0.0, ExpectedSquareMatrix([2]int{rows, cols})
	}

	if rows == 1 {
		return m[0][0], nil
	}

	if rows == 2 {
		a := m[0][0]
		b := m[0][1]
		c := m[1][0]
		d := m[1][1]
		return a*d - b*c, nil
	}

	s := 0.0

	for i := 0; i < cols; i++ {
		sm := subMatrix(m[1:][:], i)
		z, err := Determinant(sm)

		if err == nil {
			if i%2 != 0 {
				s -= m[0][i] * z
			} else {
				s += m[0][i] * z
			}
		}
	}
	return s, nil
}

//subMatrix returns the subMatrix with the p'th row removed
func subMatrix(m [][]float64, p int) [][]float64 {
	stacks := make([]stack, len(m))
	for n := range m {
		stacks[n] = stack{}
		for j := range m[n] {
			if j != p {
				stacks[n].push(m[n][j])
			}
		}
	}
	out := make([][]float64, len(m))
	for k := range stacks {
		out[k] = stacks[k].ToSlice()
	}
	return out
}

//IsInvertible checks whether or not the matrix provided is invertible
func (m *Matrix) IsInvertible() bool {
	det, err := Determinant(m.Get())
	if err != nil {
		return false
	}
	if getNumCols(m) != getNumRows(m) || det == 0 {
		return false
	}
	return true
}

//Inverse returns the inverse of the matrix if applicable, otherwise nil
func (m *Matrix) Inverse() *Matrix {
	if !m.IsInvertible() {
		return nil
	}
	det, err := Determinant(m.Get())
	if err != nil {
		panic(err)
		return nil
	}
	if getNumRows(m) == 2 {
		a := m.Get()[0][0]
		b := m.Get()[0][1]
		c := m.Get()[1][0]
		d := m.Get()[1][1]
		mInverse, err := NewMatrix(getNumCols(m), getNumRows(m), d/det, -b/det, -c/det, a/det)
		if err != nil {
			panic(err)
		}
		return mInverse
	}

	res := make([]float64, 0, 0)
	for i := 0; i < getNumRows(m); i++ {
		for j := 0; j < getNumRows(m); j++ {
			preDeletedRow := m.Get()[0:i][:]
			postDeletedRow := m.Get()[i+1:][:]
			tempMatrix := make([][]float64, 0, 0)
			tempMatrix = append(tempMatrix, preDeletedRow...)
			tempMatrix = append(tempMatrix, postDeletedRow...)
			sm := subMatrix(tempMatrix, j)
			detSm, err := Determinant(sm)
			if err != nil {
				panic(err)
				return nil
			}
			d := 1.0
			if (i+j)%2 == 1 {
				d = -1.0
			}
			res = append(res, d*detSm/det)
		}
	}
	mInverse, err := NewMatrix(getNumCols(m), getNumRows(m), res...)
	if err != nil {
		panic(err)
		return nil
	}
	return mInverse.Transpose()
}

//TranslationMatrix returns a 4x4 translation matrix for 3d vectors/points
func TranslationMatrix(x, y, z float64) *Matrix {
	m, err := NewMatrix(4, 4,
		1, 0, 0, x,
		0, 1, 0, y,
		0, 0, 1, z,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//ScalingMatrix returns a 4x4 scaling matrix for 3d vectors/points
func ScalingMatrix(x, y, z float64) *Matrix {
	m, err := NewMatrix(4, 4,
		x, 0, 0, 0,
		0, y, 0, 0,
		0, 0, z, 0,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//RotationX returns a 4x4 matrix that rotates a 3d vector/point around the x-axis
func RotationX(r float64) *Matrix {
	m, err := NewMatrix(4, 4,
		1, 0, 0, 0,
		0, math.Cos(r), -math.Sin(r), 0,
		0, math.Sin(r), math.Cos(r), 0,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//RotationY returns a 4x4 matrix that rotates a 3d vector/point around the x-axis
func RotationY(r float64) *Matrix {
	m, err := NewMatrix(4, 4,
		math.Cos(r), 0, math.Sin(r), 0,
		0, 1, 0, 0,
		-math.Sin(r), 0, math.Cos(r), 0,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//RotationZ returns a 4x4 matrix that rotates a 3d vector/point around the x-axis
func RotationZ(r float64) *Matrix {
	m, err := NewMatrix(4, 4,
		math.Cos(r), -math.Sin(r), 0, 0,
		math.Sin(r), math.Cos(r), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//Shearing returns a 4x4 matrix that is used to shear a
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	m, err := NewMatrix(4, 4,
		1, xy, xz, 0,
		yx, 1, yz, 0,
		zx, zy, 1, 0,
		0, 0, 0, 1)
	if err != nil {
		panic(err)
	}
	return m
}

//stack datatype helper for matrix functions/methods
type stack []float64

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(n float64) {
	*s = append(*s, n)
}

func (s *stack) pop() (float64, bool) {
	if s.isEmpty() {
		return 0, false
	}
	i := len(*s) - 1
	n := (*s)[i]
	*s = (*s)[:i]
	return n, true
}

func (s *stack) ToSlice() []float64 {
	return *s
}

// other helper functions
func getNumRows(m *Matrix) int {
	return len(m.Get())
}

func getNumCols(m *Matrix) int {
	return len(m.Get()[0])
}

func equals(a, b float64) bool {
	EPSILON := 0.00001
	return math.Abs(a-b) < EPSILON
}
