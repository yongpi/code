package list

import "math"

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

func MaxSubMatrixSum(matrix [][]int) int {
	ans := math.MinInt32
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < n; i++ {
		tmp := make([]int, m)
		for j := i; j < n; j++ {
			var cur int
			for k := 0; k < m; k++ {
				tmp[k] += matrix[j][k]
				cur += tmp[k]
				ans = Max(ans, cur)
				if cur < 0 {
					cur = 0
				}
			}
		}
	}

	return ans
}

func MaxMultiSub(nums []int) int {
	ans, maxValue, minValue := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxValue = Max(Max(maxValue*nums[i], minValue*nums[i]), nums[i])
		minValue = Min(Min(minValue*nums[i], maxValue*nums[i]), nums[i])
		ans = Max(maxValue, ans)
	}

	return ans
}

func NotContainCurMulti(nums []int) []int {
	n := len(nums)
	left := 1
	right := 1
	for i := n - 1; i > 0; i-- {
		right *= nums[i]
	}

	var ans []int
	for i := 0; i < n-1; i++ {
		ans = append(ans, left*right)
		left = left * nums[i]
		right = right / nums[i+1]
	}

	ans = append(ans, left*right)

	return ans
}
