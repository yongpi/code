package list

import (
	"fmt"
	"testing"
)

func TestMiniNotSortSubList(t *testing.T) {
	fmt.Println(MiniNotSortSubList([]int{6, 3, 2, 7, 1, 8, 9, 10}))
}

func TestMaxCanMergeSub(t *testing.T) {
	fmt.Println(MaxCanMergeSub([]int{5, 5, 3, 2, 6, 4, 3}))
}

func TestSortN(t *testing.T) {
	fmt.Println(SortN([]int{2, 3, 1, 4, 5}))
}

func TestSortOddEven(t *testing.T) {
	fmt.Println(SortOddEven([]int{1, 3, 2, 4, 5}))
}

func TestSortZeroOneTwo(t *testing.T) {
	fmt.Println(SortZeroOneTwo([]int{1, 2, 1, 0, 1, 2, 0}))
}
