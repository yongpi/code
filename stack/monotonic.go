package stack

func MinClose(nums []int) [][]int {
	stack := make([]int, 0)
	ans := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[stack[len(stack)-1]] < nums[i] {
			stack = append(stack, i)
			continue
		}

		for nums[i] < nums[stack[len(stack)-1]] {
			vi := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				ans[vi] = []int{-1, i}
				break
			}

			ans[vi] = []int{stack[len(stack)-1], i}
		}

		stack = append(stack, i)
	}

	for {
		vi := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			ans[vi] = []int{-1, -1}
			break
		}

		ans[vi] = []int{stack[len(stack)-1], -1}
	}

	return ans
}
