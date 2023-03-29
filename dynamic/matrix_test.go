package dynamic

import (
	"fmt"
	"testing"
)

func TestMinSumMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 9},
		{8, 1, 3, 4},
		{5, 0, 6, 1},
		{8, 8, 4, 0},
	}

	fmt.Println(MinSumMatrix(matrix))
	fmt.Println(MinSumMatrix2(matrix))
}

func TestDragons(t *testing.T) {
	matrix := [][]int{
		{-2, -3, 3},
		{-5, -10, 1},
		{0, 30, -5},
	}

	fmt.Println(Dragons(matrix))
}
