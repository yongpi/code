package stack

import (
	"fmt"
	"testing"
)

func TestMaxWindow(t *testing.T) {
	nums := []int{1, 8, 6, 10, 2, 5, 9, 3}

	fmt.Println(MaxWindow(nums, 3))
}
