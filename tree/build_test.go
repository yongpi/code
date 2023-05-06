package tree

import (
	"fmt"
	"testing"
)

func TestBuildPost(t *testing.T) {
	list := []int{2, 1, 3, 6, 5, 7, 4}
	root := BuildPost(list)
	fmt.Println(root.Value)

	root2 := BuildPost2(list)
	fmt.Println(root2.Value)

	fmt.Println(IsPostArray(list))
}

func TestBuildSearchMid(t *testing.T) {
	root := BuildSearchMid([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(root.Value)
}

func TestBuildPostList(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	mid := []int{4, 2, 5, 1, 6, 3, 7}

	fmt.Println(BuildPostList(pre, mid))
}
