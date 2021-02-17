package algebra

import (
	"fmt"
)

//MismatchedLength is an error interface returned when two objects should have the same length
type MismatchedLength [2]int

//ZeroDivide is an error returned if a zero division occurs in Vector methods
type ZeroDivide int

//ExpectedDimension is an error returned when a data strcutre does not match the expected input dimension
type ExpectedDimension int

//InvalidMatrixIndex is an error returned when an index provided is out of bounds of one of the matrix slices
type InvalidMatrixIndex int

func (e MismatchedLength) Error() string {
	return fmt.Sprintf("Expected %d and %d to match", e[0], e[1])
}

func (e ZeroDivide) Error() string {
	return fmt.Sprintf("Zero divide error")
}

func (e ExpectedDimension) Error() string {
	return fmt.Sprintf("Expected vector input dimension to be %d", e)
}

func (e InvalidMatrixIndex) Error() string {
	return fmt.Sprintf("Index %d out of matrix bounds", e)
}
