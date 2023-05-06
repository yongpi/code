package tree

func FindExchangeNode(root *Tree) []*Tree {
	cur := root
	stack := make([]*Tree, 0)
	ans := make([]*Tree, 2)
	var pre *Tree

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil && pre.Value > node.Value {
			if ans[0] == nil {
				ans[0] = pre
			}

			ans[1] = node
		}

		pre = node
		cur = node.Right
	}

	return ans
}
