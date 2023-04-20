package str

type NumNode struct {
	Child map[int32]*NumNode
}

func InsertNum(root *NumNode, num int32) {
	if root == nil {
		return
	}

	child := root.Child
	for i := 31; i >= 0; i-- {
		item := (num >> i) & 1
		if child[item] == nil {
			child[item] = &NumNode{Child: make(map[int32]*NumNode)}
		}

		child = child[item].Child
	}
}

func MaxXOR(root *NumNode, num int32) int32 {
	if root == nil {
		return 0
	}

	var res int32
	child := root.Child
	for i := 31; i >= 0; i-- {
		item := (num >> i) & 1
		best := item ^ 1
		if i == 31 {
			best = item
		}
		if _, ok := child[best]; !ok {
			best = best ^ 1
		}

		res |= (item ^ best) << i
		child = child[best].Child
	}

	return res
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}

	return b
}

func SubMaxXOR(nums []int32) int32 {
	root := &NumNode{Child: make(map[int32]*NumNode)}
	InsertNum(root, 0)

	var ans, xor int32
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]

		ans = max(ans, MaxXOR(root, xor))
		InsertNum(root, xor)
	}

	return ans
}
