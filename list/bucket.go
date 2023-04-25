package list

import "math"

func MaxBucketValue(nums []int) int {
	min, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
		if nums[i] < min {
			min = nums[i]
		}
	}

	bucketSize := len(nums) + 1
	minList := make([]int, bucketSize)
	maxList := make([]int, bucketSize)
	existList := make([]bool, bucketSize)

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		bi := GetBucketIndex(max, min, value, len(nums))
		if !existList[bi] {
			minList[bi] = value
			maxList[bi] = value
			existList[bi] = true
			continue
		}

		if minList[bi] > value {
			minList[bi] = value
		}
		if maxList[bi] < value {
			maxList[bi] = value
		}
	}

	ans := math.MinInt32
	cur, next := 0, 1
	for existList[next] {
		ans = Max(minList[next]-maxList[cur], ans)
		cur = next
		next++
	}
	return ans
}

func GetBucketIndex(max, min, value, size int) int {
	return (value - min) * size / (max - min)
}
