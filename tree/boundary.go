package tree

import "fmt"

type Tree struct {
	Left, Right, Parent *Tree
	Value               int
}

func PrintBoundary(root *Tree) {
	if root == nil {
		return
	}

	list := make([][2]*Tree, 0)
	stack := make([]*Tree, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		sl := len(stack)
		ld := [2]*Tree{}
		for i := 0; i < sl; i++ {
			node := stack[i]
			if i == 0 {
				ld[0] = node
			}
			if i == sl-1 {
				ld[1] = node
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		list = append(list, ld)
		stack = stack[sl:]
	}

	for _, item := range list {
		fmt.Println(item[0].Value)
	}

	printChild(root, 0, list)

	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		if item[0] != item[1] {
			fmt.Println(item[1].Value)
		}
	}
}

func printChild(node *Tree, height int, data [][2]*Tree) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil && data[height][0] != node && data[height][1] != node {
		fmt.Println(node.Value)
	}

	printChild(node.Left, height+1, data)
	printChild(node.Right, height+1, data)
}
