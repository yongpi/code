package algorithm

import "fmt"

type TreeNode struct {
	Left, Right *TreeNode
	Value       int
}

func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}

	stack := make([]*TreeNode, 0)
	node := root

	for node != nil || len(stack) > 0 {
		if node != nil {
			fmt.Println(node.Value)
			stack = append(stack, node)
			node = node.Left
			continue
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
}

func PreOrder2(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Value)
	PreOrder2(root.Left)
	PreOrder2(root.Right)
}

func InOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	node := root

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(node.Value)
		node = node.Right
	}
}

func InOrder2(root *TreeNode) {
	if root == nil {
		return
	}
	InOrder2(root.Left)
	fmt.Println(root.Value)
	InOrder2(root.Right)
}

func PostOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	node := root
	var pre *TreeNode

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
			continue
		}

		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == pre {
			stack = stack[:len(stack)-1]
			fmt.Println(node.Value)
			pre = node
			node = nil
		} else {
			node = node.Right
		}
	}
}

func PostOrder2(root *TreeNode) {
	if root == nil {
		return
	}

	PostOrder2(root.Left)
	PostOrder2(root.Right)
	fmt.Println(root.Value)
}

func LevelOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		sl := len(stack)
		for i := 0; i < sl; i++ {
			node := stack[i]
			fmt.Println(node.Value)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		stack = stack[sl:]
	}
}
