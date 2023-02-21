package algorithm

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	nums := []int{1, 5, 3, 4, 8, 3, 4, 10, 9}
	fmt.Println(MaxHeap(nums))
}

func TestNumMax(t *testing.T) {
	nums := []int{1, 5, 3, 4, 8, 9, 4, 10, 9}
	fmt.Println(NumMax(nums, 3))
}
