package algorithm

func split(nums []int, begin, end int) {
	if begin >= end {
		return
	}

	mid := begin + (end-begin)/2
	split(nums, begin, mid)
	split(nums, mid+1, end)
	merge(nums, begin, mid, end)
}

func merge(nums []int, begin, mid, end int) {
	ls, le, rs, re := begin, mid, mid+1, end

	tmp := make([]int, end-begin+1)
	var ti int
	for ls <= le && rs <= re {
		if nums[ls] >= nums[rs] {
			tmp[ti] = nums[rs]
			rs++
		} else {
			tmp[ti] = nums[ls]
			ls++
		}
		ti++
	}

	for ls <= le {
		tmp[ti] = nums[ls]
		ls++
		ti++
	}

	for rs <= re {
		tmp[ti] = nums[rs]
		rs++
		ti++
	}

	for i := 0; i < ti; i++ {
		nums[begin+i] = tmp[i]
	}
}

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	split(nums, 0, len(nums)-1)
}
