package tree

import (
	"fmt"
	"testing"
)

func TestBuildPost(t *testing.T) {
	root := BuildPost([]int{2, 1, 3, 6, 5, 7, 4})
	fmt.Println(root.Value)
}

func TestBuildSearchMid(t *testing.T) {
	root := BuildSearchMid([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(root.Value)
}
