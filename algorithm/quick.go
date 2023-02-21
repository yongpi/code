package algorithm

func position(nums []int, left, right int) int {
	begin := left + 1
	end := right
	value := nums[left]

	for {
		for ; begin < right && nums[begin] <= value; begin++ {
		}

		for ; end > left && nums[end] > value; end-- {
		}

		if begin >= end {
			break
		}

		nums[begin], nums[end] = nums[end], nums[begin]
	}

	nums[end], nums[left] = nums[left], nums[end]

	return end
}

func sort(nums []int, left, right int) {
	if left >= right {
		return
	}

	pos := position(nums, left, right)
	sort(nums, left, pos-1)
	sort(nums, pos+1, right)
}

func Quick(nums []int) {
	if len(nums) <= 1 {
		return
	}

	sort(nums, 0, len(nums)-1)
}
