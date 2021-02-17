package algebra

import (
	"testing"
)

func TestNewEmptyMatrix(t *testing.T) {
	m := NewEmptyMatrix(2,2)

	res := [][]float64{{0,0}, {0,0}}
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

	if err != nil{
		t.Errorf("%s", err)
	}

	res = [][]float64{{-3, 5, 0}, {1, -2, -2}, {0, 1, 1}}

	testMatrixEquals(t, m.Get(), res)
}

func TestMatrix_Equals(t *testing.T) {
	m, err := NewMatrix(2,2, 1, 1,
											-1, 1)

	if err != nil{
		t.Errorf("%s", err)
	}

	m2, err := NewMatrix(3,3, 1, 1, 1,
											-1, 1, 0,
											0, 0, 0)
	if err != nil{
		t.Errorf("%s", err)
	}

	if m.Equals(m2){
		t.Errorf("Expected matrix %v to not be equal to matrix %v", m ,m2)
	}

	m3, err := NewMatrix(2,2, 1, 1,
		  									-1, 1)
	if !m.Equals(m3){
		t.Errorf("Expected matrix %v to be equal to matrix %v", m, m3)
	}
}

func TestMultiply(t *testing.T) {
	m1, err := NewMatrix(4, 4, 1, 2, 3, 4,
											  5, 6, 7, 8,
											  9, 8, 7, 6,
											  5, 4, 3, 2)
	if err != nil{
		t.Errorf("%s",err)
	}
	m2, err := NewMatrix(4,4,
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4 ,3 ,6, 5,
		1, 2, 7, 8)

	if err != nil{
		t.Errorf("%s", err)
	}

	m3, err := Multiply(m1, m2)

	if err != nil{
		t.Errorf("%s", err)
	}

	res := [][]float64{ {20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}}
	testMatrixEquals(t, m3.Get(), res)
}

func TestMatrix_MultiplyByVec(t *testing.T) {
	m, err := NewMatrix(4, 4,
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1)

	if err != nil{
		t.Errorf("%s", err)
	}

	v := NewPoint(1, 2, 3)
	v,err = m.MultiplyByVec(v)
	if err != nil{
		t.Errorf("%s", err)
	}
	res := []float64{18,24,33,1}
	testVectorEquals(t, v.Get(), res)
}

func TestIdentityMatrix(t *testing.T) {
	m , err := IdentityMatrix(3)
	if err != nil{
		t.Errorf("%s", err)
	}

	res := [][]float64{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	testMatrixEquals(t, m.Get(), res)
}

func TestMatrix_Transpose(t *testing.T) {
	m, err := NewMatrix(4,4,
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8)
	if err != nil{
		t.Errorf("%s", err)
	}
	m, err = m.Transpose()

	res := [][]float64{{0, 9, 1, 0}, {9, 8, 8, 0}, {3, 0, 5, 5}, {0, 8, 3, 8}}
	testMatrixEquals(t, m.Get(), res)
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
