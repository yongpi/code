package algorithm

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nums := []int{1, 2, 3, 7, 8, 10, 34}

	fmt.Println(binary(nums, 1))
	fmt.Println(binary(nums, 34))
	fmt.Println(binary(nums, 7))
	fmt.Println(binary(nums, 9))
}

func TestBinaryMax(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 7, 7, 8, 10, 11, 34}

	fmt.Println(binaryMax(nums, 0))
	fmt.Println(binaryMax(nums, 34))
	fmt.Println(binaryMax(nums, 7))
	fmt.Println(binaryMax(nums, 9))
	fmt.Println(binaryMax(nums, 3))
	fmt.Println(binaryMax(nums, 6))
}

func TestBinaryMin(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 7, 7, 8, 10, 34}

	fmt.Println(binaryMin(nums, 1))
	fmt.Println(binaryMin(nums, 34))
	fmt.Println(binaryMin(nums, 7))
	fmt.Println(binaryMin(nums, 9))
	fmt.Println(binaryMin(nums, 3))
	fmt.Println(binaryMin(nums, 6))
}
