package str

func IsValidBrackets(data string) bool {
	var count int
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			count++
		} else if data[i] == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

func LongValidBrackets(data string) string {
	n := len(data)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 && data[i] == ')' {
			long := dp[i-1]
			index := i - long - 1
			if index >= 0 && data[index] == '(' {
				dp[i] = dp[i-1] + 2

				if index >= 1 {
					dp[i] += dp[index-1]
				}
			}
		}
	}

	var long, index int
	for key, value := range dp {
		if value > long {
			long = value
			index = key
		}
	}

	if long == 0 {
		return ""
	}

	return data[index-long+1 : index+1]
}
