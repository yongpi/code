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
