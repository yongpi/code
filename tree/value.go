package tree

func MaxValue(root *Tree) int {
	if root == nil {
		return 0
	}

	sv, nv := maxValueProcess(root)

	return max(sv+root.Value, nv)
}

func maxValueProcess(node *Tree) (int, int) {
	if node == nil {
		return 0, 0
	}

	lsv, lnv := maxValueProcess(node.Left)
	rsv, rnv := maxValueProcess(node.Right)

	selectValue := node.Value + lnv + rnv
	notSelectValue := lsv + rsv

	return selectValue, notSelectValue
}
