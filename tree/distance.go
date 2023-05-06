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

func MaxDistance2(root *Tree) int {
	var getDistance func(node *Tree) (int, int)
	getDistance = func(node *Tree) (int, int) {
		if node == nil {
			return 0, 0
		}

		lh, ld := getDistance(node.Left)
		rh, rd := getDistance(node.Right)

		height := max(lh, rh) + 1
		distance := max(ld, rd)

		distance = max(distance, lh+rh+1)

		return height, distance
	}

	_, ans := getDistance(root)
	return ans
}
