package list

func SumWaterCapacity(nums []int) int {
	var sum int
	left, right := make([]int, len(nums)), make([]int, len(nums))

	leftMax := 0
	left[0] = leftMax
	for i := 1; i < len(nums); i++ {
		leftMax = Max(leftMax, nums[i-1])
		left[i] = leftMax
	}

	rightMax := 0
	right[len(nums)-1] = rightMax
	for i := len(nums) - 2; i >= 0; i-- {
		rightMax = Max(rightMax, nums[i+1])
		right[i] = rightMax
	}

	for i := 0; i < len(nums); i++ {
		if left[i] > nums[i] && right[i] > nums[i] {
			sum += Min(left[i], right[i]) - nums[i]
		}
	}

	return sum
}
