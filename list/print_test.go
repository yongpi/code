package list

import "testing"

func TestCirclePrint(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}

	CirclePrint(matrix)
}
