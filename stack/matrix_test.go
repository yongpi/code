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

func TestIslandsNum(t *testing.T) {
	matrix := [][]int{
		{1, 0, 0, 1},
		{1, 0, 1, 1},
		{1, 0, 1, 1},
	}

	fmt.Println(IslandsNum(matrix))
}
