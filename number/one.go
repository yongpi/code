package number

func OneCount(num int) int {
	var ans int
	for num&num-1 != 0 {
		ans++
	}

	return ans
}

func FindTwo(nums []int32) (int32, int32) {
	var ans int32
	for _, value := range nums {
		ans ^= value
	}

	var flag int32
	for i := 0; i <= 31; i++ {
		if (ans>>i)&1 == 1 {
			flag = 1 << i
			break
		}
	}

	var l1, l2 []int32
	for _, value := range nums {
		if value&flag == 0 {
			l1 = append(l1, value)
		} else {
			l2 = append(l2, value)
		}
	}

	a1, a2 := l1[0], l2[0]
	for i := 1; i < len(l1); i++ {
		a1 ^= l1[i]
	}

	for i := 1; i < len(l2); i++ {
		a2 ^= l2[i]
	}

	return a1, a2
}
