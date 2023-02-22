package stack

import (
	"fmt"
	"testing"
)

func TestMinClose(t *testing.T) {
	nums := []int{1, 5, 6, 8, 0, 4, 9, 2, 10}
	fmt.Println(MinClose(nums))
}

func TestMinCloseSame(t *testing.T) {
	nums := []int{1, 5, 6, 0, 0, 4, 9, 2, 2}
	fmt.Println(MinCloseSame(nums))
}
