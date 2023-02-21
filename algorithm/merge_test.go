package algorithm

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{1, 8, 3, 4, 9, 3, 2, 3, 4, 6, 10}
	MergeSort(nums)
	fmt.Println(nums)
}
