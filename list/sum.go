package list

func PreSumMaxLen(nums []int, sum int) int {
	hm := make(map[int]int)
	hm[0] = -1

	var ans, count int
	for i := 0; i < len(nums); i++ {
		count += nums[i]

		if _, ok := hm[count]; !ok {
			hm[count] = i
		}

		if j, ok := hm[count-sum]; ok {
			ans = Max(ans, i-j)
		}
	}

	return ans
}
