package tree

import "fmt"

func PostOrder(root *Tree) {
	stack := make([]*Tree, 0)
	cur := root
	var pre *Tree

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			fmt.Println(cur.Value)
			stack = stack[:len(stack)-1]
			pre = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
}
