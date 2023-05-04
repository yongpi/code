package tree

func MaxSumLevel(root *Tree, target int) int {
	hm := make(map[int]int)
	hm[0] = 0

	return GetSumLevel(root, hm, 1, 0, target, 0)
}

func GetSumLevel(node *Tree, hm map[int]int, level, sum, target, maxLen int) int {
	if node == nil {
		return maxLen
	}

	sum += node.Value
	if _, ok := hm[sum]; !ok {
		hm[sum] = level
	}

	if data, ok := hm[sum-target]; ok {
		maxLen = max(maxLen, level-data)
	}

	maxLen = GetSumLevel(node.Left, hm, level+1, sum, target, maxLen)
	maxLen = GetSumLevel(node.Right, hm, level+1, sum, target, maxLen)

	if data, ok := hm[sum]; ok && data == level {
		delete(hm, sum)
	}

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func MaxSumLevel2(root *Tree, target int) int {
	hm := make(map[int]int)
	hm[0] = 0
	var maxLen int

	var getMaxLen func(root *Tree, level, sum, target int)
	getMaxLen = func(root *Tree, level, sum, target int) {
		if root == nil {
			return
		}

		sum += root.Value
		if _, ok := hm[sum]; !ok {
			hm[sum] = level
		}

		if data, ok := hm[sum-target]; ok {
			maxLen = max(maxLen, level-data)
		}

		getMaxLen(root.Left, level+1, sum, target)
		getMaxLen(root.Right, level+1, sum, target)

		if data, ok := hm[sum]; ok && data == level {
			delete(hm, sum)
		}
	}

	getMaxLen(root, 1, 0, target)

	return maxLen

}
