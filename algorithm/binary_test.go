package algorithm

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nums := []int{1, 2, 3, 7, 8, 10, 34}

	fmt.Println(Binary(nums, 1))
	fmt.Println(Binary(nums, 34))
	fmt.Println(Binary(nums, 7))
	fmt.Println(Binary(nums, 9))
}

func TestBinaryMax(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 7, 7, 8, 10, 11, 34}

	fmt.Println(BinaryMax(nums, 0))
	fmt.Println(BinaryMax(nums, 34))
	fmt.Println(BinaryMax(nums, 7))
	fmt.Println(BinaryMax(nums, 9))
	fmt.Println(BinaryMax(nums, 3))
	fmt.Println(BinaryMax(nums, 6))
}

func TestBinaryMin(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 7, 7, 8, 10, 34}

	fmt.Println(BinaryMin(nums, 1))
	fmt.Println(BinaryMin(nums, 34))
	fmt.Println(BinaryMin(nums, 7))
	fmt.Println(BinaryMin(nums, 9))
	fmt.Println(BinaryMin(nums, 3))
	fmt.Println(BinaryMin(nums, 6))
}
