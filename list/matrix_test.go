package list

import (
	"fmt"
	"testing"
)

func TestMinPath(t *testing.T) {
	matrix := [][]int{
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{1, 1, 1, 0, 1},
		{0, 0, 0, 0, 1},
	}

	fmt.Println(MinPath(matrix))
}
