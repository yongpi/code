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

func MaxLevelEnvelop2(envelops [][2]int) [][2]int {
	sort.Slice(envelops, func(i, j int) bool {
		if envelops[i][0] < envelops[j][0] {
			return true
		}

		if envelops[i][0] == envelops[j][0] {
			return envelops[i][1] > envelops[j][1]
		}

		return false
	})

	n := len(envelops)
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelops[i][1] > envelops[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	index := 0
	for i := 1; i < n; i++ {
		if dp[i] > dp[index] {
			index = i
		}
	}

	ans := make([][2]int, dp[index])
	end := dp[index]

	for i := index; i >= 0; i-- {
		if dp[i] == end {
			ans[end-1] = envelops[i]
			end--
		}
	}

	return ans
}
