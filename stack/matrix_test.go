package stack

import (
	"fmt"
	"testing"
)

func TestMaxMatrixArea(t *testing.T) {
	matrix := [][]int{
		{1, 0, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
	}

	fmt.Println(MaxMatrixArea(matrix))
}
