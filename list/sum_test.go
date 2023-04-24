package list

import (
	"fmt"
	"testing"
)

func TestPreSumMaxLen(t *testing.T) {
	fmt.Println(PreSumMaxLen([]int{1, -2, 3, 1, 4, 5, 1, 1, 2}, 4))
}

func TestMinSum(t *testing.T) {
	fmt.Println(MinSum([]int{1, 3, 5, 2, 4, 6}))
}

func TestMaxSubSum(t *testing.T) {
	fmt.Println(MaxSubSum([]int{1, -2, 3, 5, -2, 6, -1}))
}