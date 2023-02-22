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

func MinCloseSame(nums []int) [][]int {
	stack := make([][]int, 0)
	ans := make([][]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[i] > nums[stack[len(stack)-1][0]] {
			stack = append(stack, []int{i})
			continue
		}

		if nums[i] == nums[stack[len(stack)-1][0]] {
			stack[len(stack)-1] = append(stack[len(stack)-1], i)
			continue
		}

		for nums[i] < nums[stack[len(stack)-1][0]] {
			vs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				for _, value := range vs {
					ans[value] = []int{-1, i}
				}
				break
			}

			pre := stack[len(stack)-1]
			end := pre[len(pre)-1]

			for _, value := range vs {
				ans[value] = []int{end, i}
			}
		}

		stack = append(stack, []int{i})
	}

	for {
		vs := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			for _, value := range vs {
				ans[value] = []int{-1, -1}
			}
			break
		}

		pre := stack[len(stack)-1]
		end := pre[len(pre)-1]

		for _, value := range vs {
			ans[value] = []int{end, -1}
		}
	}

	return ans
}
