package stack

func MaxWindow(nums []int, k int) []int {
	if k > len(nums) || len(nums) == 0 {
		return nil
	}

	var ans, stack []int
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
		if stack[0] == i-k {
			stack = stack[1:]
		}

		if i-k >= -1 {
			ans = append(ans, nums[stack[0]])
		}
	}

	return ans
}
