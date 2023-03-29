package dynamic

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func MinSumMatrix(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + matrix[0][i]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}

	return dp[m-1][n-1]
}

func MinSumMatrix2(matrix [][]int) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}

	n := len(matrix[0])
	dp := make([]int, n)
	dp[0] = matrix[0][0]

	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] + matrix[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				dp[j] = dp[j] + matrix[i][j]
			} else {
				dp[j] = min(dp[j], dp[j-1]) + matrix[i][j]
			}
		}
	}

	return dp[n-1]
}

func Dragons(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[n-1][m-1] = max(1-matrix[n-1][m-1], 1)

	for i := n - 2; i >= 0; i-- {
		dp[i][m-1] = max(dp[i+1][m-1]-matrix[i][m-1], 1)
	}

	for i := m - 2; i >= 0; i-- {
		dp[n-1][i] = max(dp[n-1][i+1]-matrix[n-1][i], 1)
	}

	for i := n - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = max(dp[i+1][j]-matrix[i][j], 1)
			dp[i][j] = min(max(dp[i][j+1]-matrix[i][j], 1), dp[i][j])
		}
	}

	return dp[0][0]
}
