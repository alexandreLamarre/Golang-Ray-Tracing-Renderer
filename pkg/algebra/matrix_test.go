package algebra

import (
	"math"
	"testing"
)

func TestNewEmptyMatrix(t *testing.T) {
	m := NewEmptyMatrix(2, 2)

	rows := getNumRows(m)
	cols := getNumCols(m)

	if rows != cols || rows != 2 || cols != 2 {
		t.Errorf("Matrix has the wrong dimensions %d %d, Expected: %d %d", rows, cols, 2, 2)
	}

	res := [][]float64{{0, 0}, {0, 0}}
	testMatrixEquals(t, m.Get(), res)
}

func TestNewMatrix(t *testing.T) {
	m, err := NewMatrix(4, 4, 1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5)

	if err != nil {
		t.Errorf("%s", err)
	}

	if len(m.Get()) != 4 {
		t.Errorf("Expected length 4, Got: %d", len(m.Get()))
	}

	res := [][]float64{{1, 2, 3, 4}, {5.5, 6.5, 7.5, 8.5}, {9, 10, 11, 12}, {13.5, 14.5, 15.5, 16.5}}

	testMatrixEquals(t, m.Get(), res)

	m, err = NewMatrix(2, 2, -3, 5,
		1, -2)
	if err != nil {
		t.Errorf("%s", err)
	}

	res = [][]float64{{-3, 5}, {1, -2}}
	testMatrixEquals(t, m.Get(), res)

	m, err = NewMatrix(3, 3, -3, 5, 0,
		1, -2, -2,
		0, 1, 1)

	if err != nil {
		t.Errorf("%s", err)
	}

	res = [][]float64{{-3, 5, 0}, {1, -2, -2}, {0, 1, 1}}

	testMatrixEquals(t, m.Get(), res)
}

func TestMatrix_Equals(t *testing.T) {
	m, err := NewMatrix(2, 2, 1, 1,
		-1, 1)

	if err != nil {
		t.Errorf("%s", err)
	}

	m2, err := NewMatrix(3, 3, 1, 1, 1,
		-1, 1, 0,
		0, 0, 0)
	if err != nil {
		t.Errorf("%s", err)
	}

	if m.Equals(m2) {
		t.Errorf("Expected matrix %v to not be equal to matrix %v", m, m2)
	}

	m3, err := NewMatrix(2, 2, 1, 1,
		-1, 1)
	if !m.Equals(m3) {
		t.Errorf("Expected matrix %v to be equal to matrix %v", m, m3)
	}
}

func TestMultiply(t *testing.T) {
	m1, err := NewMatrix(4, 4, 1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2)
	if err != nil {
		t.Errorf("%s", err)
	}
	m2, err := NewMatrix(4, 4,
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8)

	if err != nil {
		t.Errorf("%s", err)
	}

	m3 := Multiply(m1, m2)

	if m3 == nil {
		t.Errorf("Something went wrong with matrix algebra")
	}

	res := [][]float64{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}
	testMatrixEquals(t, m3.Get(), res)
}

func TestMatrix_MultiplyByVec(t *testing.T) {
	m, err := NewMatrix(4, 4,
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1)

	if err != nil {
		t.Errorf("%s", err)
	}

	v := NewPoint(1, 2, 3)
	v = m.MultiplyByVec(v)
	if v == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res := []float64{18, 24, 33, 1}
	testVectorEquals(t, v.Get(), res)
}

func TestIdentityMatrix(t *testing.T) {
	m := IdentityMatrix(3)
	if m == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}

	res := [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	testMatrixEquals(t, m.Get(), res)
}

func TestDeterminant(t *testing.T) {
	m, err := NewMatrix(2, 2,
		1, 5,
		-3, 2)
	if err != nil {
		t.Errorf("%s", err)
	}

	det, err := Determinant(m.Get())
	res := 17.0

	if !equals(det, res) {
		t.Errorf("Expected: %f, Got: %f", res, det)
	}

	m, err = NewMatrix(3, 3, 1, 2, 6,
		-5, 8, -4,
		2, 6, 4)
	if err != nil {
		t.Errorf("%s", err)
	}
	det, err = Determinant(m.Get())

	if err != nil {
		t.Errorf("%s", err)
	}
	res = -196

	if !equals(det, res) {
		t.Errorf("Expected: %f, Got: %f", res, det)
	}

	m, err = NewMatrix(4, 4, -2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9)
	if err != nil {
		t.Errorf("%s", err)
	}

	det, err = Determinant(m.Get())

	if err != nil {
		t.Errorf("%s", err)
	}

	res = -4071
	if !equals(det, res) {
		t.Errorf("Expected: %f, Got: %f", res, det)
	}
}

func TestMatrix_IsInvertible(t *testing.T) {
	m, err := NewMatrix(4, 4,
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6)
	if err != nil {
		t.Errorf("%s", err)
	}
	if !m.IsInvertible() {
		t.Errorf("Expected %v to be invertible", m)
	}

	m2, err := NewMatrix(4, 4, -4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0)
	if err != nil {
		t.Errorf("%s", err)
	}
	if m2.IsInvertible() {
		t.Errorf("Expected %v to not be invertible", m2)
	}
}

func TestMatrix_Inverse(t *testing.T) {
	m, err := NewMatrix(2, 2, 4, 7, 2, 6)
	if err != nil {
		t.Errorf("%s", err)
	}
	m = m.Inverse()
	if m == nil {
		t.Errorf("%s", err)
		return
	}
	res := [][]float64{{0.6, -0.7}, {-0.2, 0.4}}
	testMatrixEquals(t, m.Get(), res)

	m, err = NewMatrix(4, 4,
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4)
	if err != nil {
		t.Errorf("%s", err)
	}
	m = m.Inverse()
	if m == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = [][]float64{{0.21805, 0.45113, 0.24060, -0.04511},
		{-0.80827, -1.45677, -0.44361, 0.52068},
		{-0.07895, -0.22368, -0.05263, 0.19737},
		{-0.52256, -0.81391, -0.30075, 0.30639}}
	testMatrixEquals(t, m.Get(), res)

	m, err = NewMatrix(4, 4,
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4)
	if err != nil {
		t.Errorf("%s", err)
	}
	m = m.Inverse()
	if m == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = [][]float64{{-0.15385, -0.15385, -0.28205, -0.53846},
		{-0.07692, 0.12308, 0.02564, 0.03077},
		{0.35897, 0.35897, 0.43590, 0.92308},
		{-0.69231, -0.69231, -0.76923, -1.92308}}
	testMatrixEquals(t, m.Get(), res)

	m, err = NewMatrix(4, 4, 9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2)

	m = m.Inverse()
	if m == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}

	res = [][]float64{{-0.04074, -0.07778, 0.14444, -0.22222}, {-0.07778, 0.03333, 0.36667, -0.33333},
		{-0.02901, -0.14630, -0.10926, 0.12963}, {0.177778, 0.06667, -0.26667, 0.33333}}
	testMatrixEquals(t, m.Get(), res)

	A, err := NewMatrix(4, 4, 3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1)
	if err != nil {
		t.Errorf("%s", err)
	}

	B, err := NewMatrix(4, 4,
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5)
	if err != nil {
		t.Errorf("%s", err)
	}
	C := Multiply(A, B)
	D := Multiply(C, B.Inverse())

	testMatrixEquals(t, D.Get(), A.Get())
}

func TestStack(t *testing.T) {
	s := &stack{}
	if !s.isEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	s.push(5.0)
	s.push(3.0)

	arr := s.ToSlice()
	res := []float64{5, 3}
	testVectorEquals(t, arr, res)

	b, removed := s.pop()
	if !removed {
		t.Errorf("Expected to remove an element from stack")
	}
	if !equals(b, 3.0) {
		t.Errorf("Expected to pop %f from stock, instead got: %f", 3.0, b)
	}
	b, removed = s.pop()
	if !removed {
		t.Errorf("Expected to remove an element from stack")
	}
	if !equals(b, 5.0) {
		t.Errorf("Expected to pop %f from stock, instead got: %f", 5.0, b)
	}

	b, removed = s.pop()
	if removed {
		t.Errorf("Expected to not remove anything from the stack")
	}

}

func TestMatrix_Transpose(t *testing.T) {
	m, err := NewMatrix(4, 4,
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8)
	if err != nil {
		t.Errorf("%s", err)
	}
	m = m.Transpose()
	if m == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}

	res := [][]float64{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}
	testMatrixEquals(t, m.Get(), res)
}

func TestTranslationMatrix(t *testing.T) {
	m := TranslationMatrix(5, -3, 2)
	p := NewPoint(-3, 4, 5)
	p = m.MultiplyByVec(p)
	if !p.IsPoint() {
		t.Errorf("Expected %v to be a point", p)
	}
	res := []float64{2, 1, 7, 1.0}
	testVectorEquals(t, p.Get(), res)

	m = m.Inverse()
	p = NewPoint(-3, 4, 5)
	p = m.MultiplyByVec(p)
	if !p.IsPoint() {
		t.Errorf("Expected %v to be a point", p)
	}
	res = []float64{-8, 7, 3, 1.0}
	testVectorEquals(t, p.Get(), res)

	v := NewVector(-3, 4, 5)
	m = TranslationMatrix(5, -3, 2)
	newV := m.MultiplyByVec(v)
	testVectorEquals(t, newV.Get(), v.Get())
}

func TestScalingMatrix(t *testing.T) {
	m := ScalingMatrix(2, 3, 4)
	p := NewPoint(-4, 6, 8)
	p = m.MultiplyByVec(p)

	if !p.IsPoint() {
		t.Errorf("Expected %v to be a point", p)
	}

	res := []float64{-8, 18, 32}
	testVectorEquals(t, p.Get(), res)

	v := NewVector(-4, 6, 8)
	v = m.MultiplyByVec(v)

	if !v.IsVector() {
		t.Errorf("Expected %v to be a vector", v)
	}
	testVectorEquals(t, v.Get(), res)

	inv := m.Inverse()
	v = NewVector(-4, 6, 8)
	v = inv.MultiplyByVec(v)

	if !v.IsVector() {
		t.Errorf("Expected %v to be a vector", v)
	}
	res = []float64{-2, 2, 2}
	testVectorEquals(t, v.Get(), res)
}

func TestMatrixRotation(t *testing.T) {

	// X-ROTATION

	p := NewPoint(0, 1, 0)
	quarter := RotationX(math.Pi / 4)
	half := RotationX(math.Pi / 2)

	p = quarter.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res := []float64{0, math.Sqrt(2) / 2, math.Sqrt(2) / 2, 1}
	testVectorEquals(t, p.Get(), res)

	p = NewPoint(0, 1, 0)
	p = half.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{0, 0, 1, 1}
	testVectorEquals(t, p.Get(), res)

	// Y - ROTATION

	p = NewPoint(0, 0, 1)
	quarter = RotationY(math.Pi / 4)
	half = RotationY(math.Pi / 2)

	p = quarter.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{math.Sqrt(2) / 2, 0, math.Sqrt(2) / 2, 1}
	testVectorEquals(t, p.Get(), res)

	p = NewPoint(0, 0, 1)
	p = half.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{1, 0, 0, 1}
	testVectorEquals(t, p.Get(), res)

	// Z - ROTATION

	p = NewPoint(0, 1, 0)
	quarter = RotationZ(math.Pi / 4)
	half = RotationZ(math.Pi / 2)

	p = quarter.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{-math.Sqrt(2) / 2, math.Sqrt(2) / 2, 0, 1}
	testVectorEquals(t, p.Get(), res)

	p = NewPoint(0, 1, 0)
	p = half.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{-1, 0, 0, 1}
	testVectorEquals(t, p.Get(), res)
}

func TestShearing(t *testing.T) {
	// =========================== X ========================
	m := Shearing(1, 0, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res := []float64{5, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)

	m = Shearing(0, 1, 0, 0, 0, 0)
	p = NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{6, 3, 4, 1}
	testVectorEquals(t, p.Get(), res)

	// =========================== y ============================

	m = Shearing(0, 0, 1, 0, 0, 0)
	p = NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{2, 5, 4, 1}
	testVectorEquals(t, p.Get(), res)

	m = Shearing(0, 0, 0, 1, 0, 0)
	p = NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{2, 7, 4, 1}
	testVectorEquals(t, p.Get(), res)

	// ============================ Z =============================

	m = Shearing(0, 0, 0, 0, 1, 0)
	p = NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{2, 3, 6, 1}
	testVectorEquals(t, p.Get(), res)

	m = Shearing(0, 0, 0, 0, 0, 1)
	p = NewPoint(2, 3, 4)
	p = m.MultiplyByVec(p)
	if p == nil {
		t.Errorf("Something went wrong with matrix algebra")
		return
	}
	res = []float64{2, 3, 7, 1}
	testVectorEquals(t, p.Get(), res)
}

func TestMatrixTransformationChaining(t *testing.T) {
	p1 := NewPoint(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := ScalingMatrix(5, 5, 5)
	C := TranslationMatrix(10, 5, 7)

	p1 = A.MultiplyByVec(p1)
	p1 = B.MultiplyByVec(p1)
	p1 = C.MultiplyByVec(p1)

	T := Multiply(C, Multiply(B, A))
	p := NewPoint(1, 0, 1)
	p2 := T.MultiplyByVec(p)

	res := []float64{15, 0, 7, 1}
	testVectorEquals(t, p1.Get(), res)
	testVectorEquals(t, p2.Get(), res)
}

func testMatrixEquals(t *testing.T, values [][]float64, expected [][]float64) {
	for i, _ := range values {
		for j, _ := range values[i] {
			if !equals(values[i][j], expected[i][j]) {
				t.Errorf("Expected: %f, Got: %f", expected[i][j], values[i][j])
			}
		}
	}
}
