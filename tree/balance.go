package tree

func Balance(root *Tree) bool {
	_, balance := balanceProcess(root)
	return balance
}

func balanceProcess(root *Tree) (int, bool) {
	if root == nil {
		return 0, true
	}

	lh, lb := balanceProcess(root.Left)
	rh, rb := balanceProcess(root.Right)

	height := max(lh, rh) + 1
	balance := lb && rb && (lh-rh <= 1 && lh-rh >= -1)

	return height, balance
}

func Balance2(root *Tree) bool {
	var isBalance func(node *Tree) (bool, int)
	isBalance = func(node *Tree) (b bool, i int) {
		if node == nil {
			return true, 0
		}

		lb, ll := isBalance(node.Left)
		rb, rl := isBalance(node.Right)

		var value int
		if ll > rl {
			value = ll - rl
		} else {
			value = rl - ll
		}

		nl := max(ll, rl) + 1
		if lb && rb && value <= 1 {
			return true, nl
		}

		return false, nl
	}

	ans, _ := isBalance(root)
	return ans
}
