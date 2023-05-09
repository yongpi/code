package dynamic

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Balloon(balloon []int) int {
	help := make([]int, len(balloon)+2)
	help[0] = 1

	for i := 0; i < len(balloon); i++ {
		help[i+1] = balloon[i]
	}
	help[len(balloon)+1] = 1

	return balloonProcess(help, 1, len(balloon))
}

func balloonProcess(balloon []int, left, right int) int {
	if left == right {
		return balloon[left-1] * balloon[left] * balloon[right+1]
	}

	leftValue := balloon[left-1]*balloon[left]*balloon[right+1] + balloonProcess(balloon, left+1, right)
	rightValue := balloon[left-1]*balloon[right]*balloon[right+1] + balloonProcess(balloon, left, right-1)
	ans := max(leftValue, rightValue)

	for i := left + 1; i < right; i++ {
		ans = max(ans, balloon[left-1]*balloon[i]*balloon[right+1]+balloonProcess(balloon, left, i-1)+balloonProcess(balloon, i+1, right))
	}

	return ans
}

func BalloonDynamic(balloon []int) int {
	n := len(balloon)
	help := make([]int, n+2)
	help[0] = 1

	for i := 0; i < n; i++ {
		help[i+1] = balloon[i]
	}
	help[n+1] = 1

	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = help[i-1] * help[i] * help[i+1]
	}

	for i := n; i >= 1; i-- {
		for j := i + 1; j <= n; j++ {
			dp[i][j] = help[i-1]*help[i]*help[j+1] + dp[i+1][j]
			dp[i][j] = max(dp[i][j], help[i-1]*help[j]*help[j+1]+dp[i][j-1])

			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], help[i-1]*help[k]*help[j+1]+dp[k+1][j]+dp[i][k-1])
			}
		}
	}

	return dp[1][n]
}

func BalloonDynamic2(balloon []int) int {
	n := len(balloon)
	m := n + 2
	help := make([]int, m)
	help[0] = 1
	help[m-1] = 1

	for i := 0; i < n; i++ {
		help[i+1] = balloon[i]
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, m)
	}

	for i := 1; i <= n; i++ {
		dp[i][i] = help[i-1] * help[i] * help[i+1]
	}

	for j := 2; j <= n; j++ {
		for i := j; i >= 1; i-- {
			dp[i][j] = max(help[i-1]*help[i]*help[j+1]+dp[i+1][j], help[i-1]*help[j]*help[j+1]+dp[i][j-1])
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], help[i-1]*help[k]*help[j+1]+dp[i][k-1]+dp[k+1][j])
			}
		}
	}

	return dp[1][n]
}
