package stack

type Mountain struct {
	Value int
	Count int
}

func VisibleMountain(nums []int) int {
	if len(nums) < 2 {
		return 1
	}

	var mv, mi, ans int
	for key, value := range nums {
		if value > mv {
			mv = value
			mi = key
		}
	}

	stack := make([]*Mountain, 0)
	for i := mi; i < mi+len(nums); i++ {
		i = i % len(nums)

		for len(stack) > 0 && stack[len(stack)-1].Value < nums[i] {
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans += item.Count*2 + count(item.Count)
		}

		if len(stack) > 0 && stack[len(stack)-1].Value == nums[i] {
			stack[len(stack)-1].Count += 1
			continue
		}

		stack = append(stack, &Mountain{
			Value: nums[i],
			Count: 1,
		})
	}

	for len(stack) > 2 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans += item.Count*2 + count(item.Count)
	}

	if len(stack) == 2 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		last := stack[len(stack)-1]

		ans += last.Count*item.Count + count(item.Count)
	}

	item := stack[len(stack)-1]
	ans += count(item.Count)

	return ans

}

func count(c int) int {
	if c == 1 {
		return 0
	}

	return c * (c - 1) / 2
}
