package str

func BecomeHW(data string) string {
	list := []byte(data)
	n := len(list)
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if list[i] == list[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}

	rl := dp[0][n-1]
	ans := make([]byte, rl+n)
	start, end := 0, len(ans)-1
	i, j := 0, n-1

	for i <= j {
		if list[i] == list[j] {
			ans[end] = list[i]
			ans[start] = list[i]

			i++
			j--
		} else if dp[i][j] == dp[i+1][j]+1 {
			ans[end] = list[i]
			ans[start] = list[i]
			i++
		} else {
			ans[end] = list[j]
			ans[start] = list[j]
			j--
		}

		end--
		start++
	}

	return string(ans)
}
