package dynamic

func Robot(n, m, k, p int) int {
	return robotProcess(n, m, k, p)
}

func robotProcess(n, m, k, p int) int {
	if k < 0 {
		return 0
	}
	if k == 0 && m == p {
		return 1
	}

	var ans int
	if m == 1 {
		ans += robotProcess(n, 2, k-1, p)
		return ans
	} else if m == n {
		ans += robotProcess(n, n-1, k-1, p)
		return ans
	}

	return robotProcess(n, m-1, k-1, p) + robotProcess(n, m+1, k-1, p)
}

func RobotDynamic(n, m, k, p int) int {
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	for i := 1; i <= n; i++ {
		if i == p {
			dp[i][0] = 1
		}
	}

	for j := 1; j <= k; j++ {
		for i := 1; i <= n; i++ {
			if i == 1 {
				dp[i][j] = dp[2][j-1]
			} else if i == n {
				dp[i][j] = dp[n-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + dp[i+1][j-1]
			}
		}
	}

	return dp[m][k]
}
