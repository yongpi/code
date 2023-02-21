package algorithm

func adjustMaxHeap(nums []int, index, size int) {
	left, right := index*2+1, index*2+2
	target := index

	if left < size && nums[left] > nums[target] {
		target = left
	}

	if right < size && nums[right] > nums[target] {
		target = right
	}

	if target != index {
		nums[target], nums[index] = nums[index], nums[target]
		adjustMaxHeap(nums, target, size)
	}
}

func initHeap(nums []int) {
	if len(nums) <= 1 {
		return
	}

	begin := len(nums)/2 - 1
	size := len(nums)
	for i := begin; i >= 0; i-- {
		adjustMaxHeap(nums, i, size)
	}
}

func MaxHeap(nums []int) int {
	initHeap(nums)
	return nums[0]
}

func NumMax(nums []int, k int) []int {
	initHeap(nums)
	size := len(nums)
	ans := make([]int, 0)

	for i := 0; i < k; i++ {
		value := nums[0]
		ans = append(ans, value)

		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		adjustMaxHeap(nums, 0, size)
	}

	return ans
}
