package dynamic

func MinCash(cash []int, money int) int {
	if len(cash) == 0 || money <= -1 {
		return -1
	}

	return processCash(cash, 0, money)
}

func processCash(cash []int, i int, money int) int {
	if i == len(cash) {
		if money == 0 {
			return 0
		}

		return -1
	}

	value := cash[i]
	ans := -1
	for k := 0; k <= money/value; k++ {
		next := processCash(cash, i+1, money-k*value)
		if next == -1 {
			continue
		}

		if ans == -1 {
			ans = next + k
		} else {
			ans = min(ans, next+k)
		}
	}

	return ans
}

func MinCashDynamic(cash []int, money int) int {
	n := len(cash)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, money+1)
	}

	dp[n-1][0] = 0
	for i := 1; i <= money; i++ {
		dp[n-1][i] = -1
		if i >= cash[n-1] && dp[n-1][i-cash[n-1]] != -1 {
			dp[n-1][i] = dp[n-1][i-cash[n-1]] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := 1; j <= money; j++ {
			dp[i][j] = -1
			if i < n-1 && dp[i+1][j] != -1 {
				dp[i][j] = dp[i+1][j]
			}

			nj := j - cash[i]
			if j >= cash[i] && dp[i][nj] != -1 {
				if dp[i][j] != -1 {
					dp[i][j] = min(dp[i][j], dp[i][j-cash[i]]+1)
				} else {
					dp[i][j] = dp[i][j-cash[i]] + 1
				}
			}
		}
	}

	return dp[0][money]
}

func AllCash(cash []int, money int) int {
	return processAllCash(cash, 0, money)
}

func processAllCash(cash []int, i, money int) int {
	if money < 0 {
		return 0
	}

	if money == 0 {
		return 1
	}

	if i == len(cash) {
		return 0
	}

	size := money / cash[i]
	var ans int
	for k := 0; k <= size; k++ {
		ans += processCash(cash, i+1, money-k*cash[i])
	}

	return ans
}

func AllCashDynamic(cash []int, money int) int {
	n := len(cash)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, money+1)
		dp[i][0] = 1
	}

	for i := 0; i <= money; i++ {
		if i%cash[0] == 0 {
			dp[0][i] = 1
		}
	}

	for j := 1; j <= money; j++ {
		for i := n - 1; i > 0; i-- {
			dp[i][j] = -1
			if i < n-1 && dp[i+1][j] != -1 {
				dp[i][j] = dp[i+1][j]
			}

			rest := j - cash[i]
			if j >= cash[i] && dp[i][rest] != -1 {
				if dp[i][j] != -1 {
					dp[i][j] += dp[i][rest] + 1
				}
			}
		}
	}

	return dp[0][money]
}
