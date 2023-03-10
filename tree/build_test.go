package tree

import (
	"fmt"
	"testing"
)

func TestBuildPost(t *testing.T) {
	root := BuildPost([]int{2, 1, 3, 6, 5, 7, 4})
	fmt.Println(root.Value)
}
