package str

func ZeroLeftOne(n int) int {
	dp := make([][2]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = [2]int{}
	}

	dp[1][0] = 0
	dp[1][1] = 1

	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][1]
		dp[i][1] = dp[i-1][0] + dp[i-1][1]
	}

	return dp[n][0] + dp[n][1]
}
