package tree

func IsCBT(root *Tree) bool {
	var last bool
	stack := make([]*Tree, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		tl := len(stack)
		for i := 0; i < tl; i++ {
			node := stack[i]
			if node.Right != nil && node.Left == nil {
				return false
			}

			if last && (node.Right != nil || node.Left != nil) {
				return false
			}

			if node.Left == nil || node.Right == nil {
				last = true
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}

		stack = stack[tl:]
	}

	return true
}
