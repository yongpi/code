package dynamic

func MinIncrList(nums []int) []int {
	dp := getDp(nums)

	var mi, target int
	for key, value := range dp {
		if value > target {
			target = value
			mi = key
		}
	}

	ans := make([]int, target)
	ans[target-1] = nums[mi]

	for i := mi - 1; i >= 0; i-- {
		if dp[i] == target-1 && nums[mi] > nums[i] {
			ans[dp[i]-1] = nums[i]
			target = dp[i]
			mi = i
		}
	}

	return ans
}

func getDp(nums []int) []int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1

	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return dp
}

func MaxXORSubList(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	if nums[0] == 0 {
		dp[0] = 1
	}

	hm := make(map[int]int)
	hm[0] = -1

	value := nums[0]
	for i := 1; i < n; i++ {
		value ^= nums[i]
		k, ok := hm[value]
		if ok {
			if k == -1 {
				dp[i] = 1
			} else {
				dp[i] = dp[k] + 1
			}
		}

		dp[i] = max(dp[i], dp[i-1])
		hm[value] = i
	}

	return dp[n-1]
}

func FirstOrSecond(nums []int) int {
	n := len(nums)
	f := make([][]int, n)
	s := make([][]int, n)

	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		s[i] = make([]int, n)
		f[i][i] = nums[i]
	}

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			f[i][j] = max(nums[i]+s[i+1][j], nums[j]+s[i][j-1])
			s[i][j] = max(f[i+1][j], f[i][j-1])
		}
	}

	return max(f[0][n-1], s[0][n-1])
}
