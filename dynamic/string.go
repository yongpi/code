package dynamic

func MaxSubSeq(a, b string) string {
	dp := subSeqDP(a, b)

	var mi, mj int
	value := dp[mi][mj]

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if dp[i][j] > value {
				value = dp[i][j]
				mi = i
				mj = j
			}
		}
	}

	ans := make([]byte, value)
	for value > 0 {
		if mi > 0 && dp[mi][mj] == dp[mi-1][mj] {
			mi = mi - 1
		} else if mj > 0 && dp[mi][mj] == dp[mi][mj-1] {
			mj = mj - 1
		} else {
			ans[value-1] = a[mi]
			mi--
			mj--
			value--
		}
	}

	return string(ans)
}

func subSeqDP(a, b string) [][]int {
	n, m := len(a), len(b)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	if a[0] == b[0] {
		dp[0][0] = 1
	}

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0]
		if a[i] == b[0] {
			dp[i][0] = max(dp[i][0], 1)
		}
	}

	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1]
		if a[0] == b[i] {
			dp[0][i] = max(dp[0][i], 1)
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if a[i] == b[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}

	return dp
}

func MaxSubStr(a, b string) string {
	dp := GetMaxSubStrDP(a, b)

	var mi, value int
	value = dp[mi][0]

	for i, item := range dp {
		for _, iv := range item {
			if iv > value {
				value = iv
				mi = i
			}
		}
	}

	return a[mi-value+1 : mi+1]
}

func GetMaxSubStrDP(a, b string) [][]int {
	n, m := len(a), len(b)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
		}
	}

	for i := 0; i < m; i++ {
		if a[0] == b[i] {
			dp[0][i] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if a[i] == b[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return dp
}

func OperateStrCost(a, b string, iv, dv, rv int) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = dv * i
	}

	for i := 1; i <= m; i++ {
		dp[0][i] = iv * i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rv
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j]+dv)
			dp[i][j] = min(dp[i][j], dp[i][j-1]+iv)
		}
	}

	return dp[n][m]
}

func IsMix(a, b, c string) bool {
	n, m := len(a), len(b)
	if len(c) != n+m {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		dp[i][0] = a[:i] == c[:i]
	}

	for i := 1; i <= m; i++ {
		dp[0][i] = b[:i] == c[:i]
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if dp[i-1][j] && a[i-1] == c[i+j-1] {
				dp[i][j] = true
			} else if dp[i][j-1] && b[j-1] == c[i+j-1] {
				dp[i][j] = true
			}
		}
	}

	return dp[n][m]
}

func ConvertLetter(a string) int {
	n := len(a)
	dp := make([]int, n+1)
	dp[n] = 1
	if a[n-1] != '0' {
		dp[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if a[i] >= '1' && a[i] <= '9' {
			dp[i] = dp[i+1]
		}

		if a[i] == '1' || (a[i] == '2' && a[i+1] >= '0' && a[i+1] <= '6') {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}

func MaxSubSeq2(a, b string) string {
	n, m := len(a), len(b)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	if a[0] == b[0] {
		dp[0][0] = 1
	}

	for i := 1; i < n; i++ {
		if a[i] == b[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		if a[0] == b[i] {
			dp[0][i] = 1
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	ai, bi, value := 0, 0, dp[0][0]
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if a[i] == b[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}

			if dp[i][j] > value {
				value = dp[i][j]
				ai = i
				bi = j
			}
		}
	}

	ans := make([]byte, value)
	for value > 0 {
		if ai > 0 && dp[ai-1][bi] == dp[ai][bi] {
			ai--
		} else if bi > 0 && dp[ai][bi-1] == dp[ai][bi] {
			bi--
		} else {
			ans[value-1] = a[ai]
			value--
			ai--
			bi--
		}
	}

	return string(ans)
}

func OperateStrCost2(a, b string, iv, dv, rv int) int {
	n, m := len(a), len(b)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for j := 1; j <= m; j++ {
		dp[0][j] = iv * j
	}

	for i := 1; i <= n; i++ {
		dp[i][0] = dv * i
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + rv
			}

			dp[i][j] = min(dp[i][j], dp[i-1][j]+dv)
			dp[i][j] = min(dp[i][j], dp[i][j-1]+iv)
		}
	}

	return dp[n][m]
}
