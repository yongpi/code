package dynamic

func JumpCount(nums []int) int {
	var count int
	var faster int
	var cur int

	for i := 0; i < len(nums); i++ {
		if cur < i {
			count++
			cur = faster
		}

		faster = max(faster, nums[i]+i)
	}

	return count
}
