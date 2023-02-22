package stack

func SubCount(nums []int, target int) int {
	maxStack := make([]int, 0)
	minStack := make([]int, 0)

	var ans, i, j int
	for i < len(nums) {
		for j < len(nums) {
			for len(maxStack) > 0 && nums[maxStack[len(maxStack)-1]] < nums[j] {
				maxStack = maxStack[:len(maxStack)-1]
			}
			maxStack = append(maxStack, j)

			for len(minStack) > 0 && nums[minStack[len(minStack)-1]] > nums[j] {
				minStack = minStack[:len(minStack)-1]
			}
			minStack = append(minStack, j)

			if nums[maxStack[0]]-nums[minStack[0]] > target {
				ans += j - i
				break
			}

			j++
		}

		if minStack[0] == i {
			minStack = minStack[1:]
		}
		if maxStack[0] == i {
			maxStack = maxStack[1:]
		}

		i++
	}

	return ans
}
