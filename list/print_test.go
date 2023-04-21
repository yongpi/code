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

func TestTupleNums(t *testing.T) {
	nums := []int{-8, -4, -3, 0, 1, 2, 2, 4, 5, 8, 8, 9}
	TupleNums(nums, 10, 0, len(nums)-1)
}

func TestTriadNums(t *testing.T) {
	TriadNums([]int{-8, -4, -3, 0, 1, 2, 2, 4, 5, 8, 8, 9}, 10)
}
