package str

func PickMin(data string) string {
	hm := make(map[byte]int)
	for i := 0; i < len(data); i++ {
		hm[data[i]]++
	}

	var ans []byte
	var left, right int
	for right < len(data) {
		value := data[right]
		hm[value]--
		if hm[value] == 0 {
			min := left
			for i := left; i <= right; i++ {
				if data[i] > data[min] {
					min = i
				}
				hm[data[i]]++
			}

			ans = append(ans, data[min])
			for i := left; i <= min; i++ {
				hm[data[i]]--
			}

			left = min
		}
		right = left
	}

	return string(ans)
}
