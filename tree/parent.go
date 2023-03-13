package tree

func LastParent(head, node1, node2 *Tree) *Tree {
	hm := make(map[*Tree]*Tree)

	stack := make([]*Tree, 0)
	stack = append(stack, head)

	for len(stack) > 0 {
		tl := len(stack)
		for i := 0; i < tl; i++ {
			node := stack[i]
			if node.Left != nil {
				stack = append(stack, node.Left)
				hm[node.Left] = node
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
				hm[node.Right] = node
			}
		}

		stack = stack[tl:]
	}

	h1 := make(map[*Tree]struct{})
	for node1 != nil {
		h1[node1] = struct{}{}
		node1 = hm[node1]
	}

	for node2 != nil {
		if _, ok := h1[node2]; ok {
			return node2
		}

		node2 = hm[node2]
	}

	return nil
}
