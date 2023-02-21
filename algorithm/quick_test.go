package algorithm

import (
	"fmt"
	"testing"
)

func TestQuick(t *testing.T) {
	nums := []int{1, 1, 1, 6, 3, 4, 8, 23, 0}
	Quick(nums)
	fmt.Println(nums)
}
