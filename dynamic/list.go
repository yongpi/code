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

func IncrList(nums []int) []int {
	hm := make(map[int]struct{})
	for _, num := range nums {
		hm[num] = struct{}{}
	}

	var count int
	var ans []int
	for i := 0; i < len(nums); i++ {
		list := []int{nums[i]}
		num := nums[i]
		for {
			_, ok := hm[num+1]
			if !ok {
				break
			}
			list = append(list, num+1)
			num = num + 1
		}

		if count < len(list) {
			ans = list
			count = len(list)
		}
	}

	return ans
}

func MinIncrList2(nums []int) []int {
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

	index := 0
	value := dp[0]
	for i := 1; i < n; i++ {
		if dp[i] > value {
			value = dp[i]
			index = i
		}
	}

	end := dp[index]
	ans := make([]int, dp[index])
	for i := index; i >= 0; i-- {
		if dp[i] == end {
			end--
			ans[end] = nums[i]
		}
	}

	return ans
}

func MaxXORSubList2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	if nums[0] == 0 {
		dp[0] = 1
	}

	sum := make([]int, n)
	sum[0] = nums[0]

	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] ^ nums[i]
	}

	for i := 1; i < n; i++ {
		for k := i; k > 0; k-- {
			if sum[k-1] == sum[i] {
				dp[i] = max(dp[i], dp[k-1]+1)
			} else {
				dp[i] = max(dp[i], dp[k-1])
			}
		}

		if sum[i] == 0 {
			dp[i] = max(dp[i], 1)
		}
	}

	return dp[n-1]
}
