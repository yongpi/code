package dynamic

import "sort"

func MaxLevelEnvelop(envelops [][2]int) [][2]int {
	sort.Slice(envelops, func(i, j int) bool {
		if envelops[i][0] < envelops[j][0] {
			return true
		}

		if envelops[i][0] == envelops[j][0] && envelops[i][1] > envelops[j][1] {
			return true
		}

		return false
	})

	nums := make([]int, len(envelops))
	for key, item := range envelops {
		nums[key] = item[1]
	}

	dp := getDp(nums)

	var mi, target int
	for key, value := range dp {
		if value > target {
			target = value
			mi = key
		}
	}

	ans := make([][2]int, target)
	ans[target-1] = envelops[mi]

	for i := mi - 1; i >= 0; i-- {
		if dp[i] == target-1 && nums[mi] > nums[i] {
			ans[dp[i]-1] = envelops[i]
			target = dp[i]
			mi = i
		}
	}

	return ans
}
