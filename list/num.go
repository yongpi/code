package list

func FindHalfCountValue(nums []int) int {
	ans := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if ans == nums[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			ans = nums[i]
			count = 1
		}
	}

	return ans
}
