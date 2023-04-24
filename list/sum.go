package list

func PreSumMaxLen(nums []int, sum int) int {
	hm := make(map[int]int)
	hm[0] = -1

	var ans, count int
	for i := 0; i < len(nums); i++ {
		count += nums[i]

		if _, ok := hm[count]; !ok {
			hm[count] = i
		}

		if j, ok := hm[count-sum]; ok {
			ans = Max(ans, i-j)
		}
	}

	return ans
}

func MinSum(nums []int) int {
	return Split(nums, 0, len(nums)-1)
}

func Split(nums []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	return Split(nums, left, mid) + Split(nums, mid+1, right) + Merge(nums, left, mid, right)
}

func Merge(nums []int, left, mid, right int) int {
	ls, le, rs, re := left, mid, mid+1, right
	tmp := make([]int, 0)
	var ans int

	for ls <= le && rs <= re {
		if nums[ls] < nums[rs] {
			tmp = append(tmp, nums[ls])
			ans += nums[ls] * (re - rs + 1)
			ls++
		} else {
			tmp = append(tmp, nums[rs])
			rs++
		}
	}

	for ls <= le {
		tmp = append(tmp, nums[ls])
		ls++
	}

	for rs <= re {
		tmp = append(tmp, nums[rs])
		rs++
	}

	for i := left; i <= right; i++ {
		nums[i] = tmp[i-left]
	}

	return ans
}

func MaxSubSum(nums []int) int {
	cur := nums[0]
	ans := cur

	for i := 1; i < len(nums); i++ {
		cur += nums[i]
		ans = Max(ans, cur)
		if cur < 0 {
			cur = 0
		}
	}

	return ans
}
