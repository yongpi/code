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
		if item == 1 && child[0] != nil {
			res += 1 << i
			child = child[0].Child
		} else if item == 1 && child[0] == nil {
			child = child[1].Child
		} else if item == 0 && child[1] != nil {
			res += 1 << i
			child = child[1].Child
		} else if item == 0 && child[1] == nil {
			child = child[0].Child
		}
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
