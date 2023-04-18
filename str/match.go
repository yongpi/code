package str

func MatchStr(data, sub string) bool {
	n, m := len(data), len(sub)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	if sub[0] == '.' {
		dp[0][1] = true
	}

	for j := 2; j <= m; j++ {
		if sub[j-1] == '*' && sub[j-2] == '.' {
			dp[0][j] = dp[0][j] || dp[0][j-2]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if isEq(data[i-1], sub[j-1]) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}

			if sub[j-1] == '*' && j >= 2 {
				dp[i][j] = dp[i][j-2]
				if isEq(data[i-1], sub[j-2]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[n][m]
}

func isEq(a, b byte) bool {
	if b == '.' {
		return true
	}

	if a == b {
		return true
	}

	return false
}
