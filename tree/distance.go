package tree

func MaxDistance(root *Tree) int {
	_, distance := distanceProcess(root)
	return distance
}

func distanceProcess(node *Tree) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftHeight, leftDistance := distanceProcess(node.Left)
	rightHeight, rightDistance := distanceProcess(node.Right)

	distance := max(leftDistance, rightDistance)
	height := max(leftHeight, rightHeight) + 1

	distance = max(distance, leftHeight+rightHeight+1)

	return height, distance
}
