package algebra

import (
	"fmt"
)
//MismatchedLength is an error interface returned when two objects should have the same length
type MismatchedLength [2]int


func (e MismatchedLength) Error() string {
	return fmt.Sprintf("Expected %d and %d to match", e[0], e[1])
}
