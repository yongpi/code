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
			break
		}

		if ans == -1 {
			ans = next + k
		} else {
			ans = min(ans, next+k)
		}
	}

	return ans
}
