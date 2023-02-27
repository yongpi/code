package stack

import (
	"fmt"
	"testing"
)

func TestVisibleMountain(t *testing.T) {
	nums := []int{5, 4, 3, 5, 4, 2, 4, 4, 5, 3, 2}

	fmt.Println(VisibleMountain(nums))
}
