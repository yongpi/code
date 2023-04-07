package str

func PickMin(data string) string {
	hm := make(map[byte]int)
	for i := 0; i < len(data); i++ {
		hm[data[i]]++
	}

	list := []byte(data)
	var ans []byte
	var left, right int
	for right < len(list) {
		value := list[right]
		hm[value]--

		if hm[value] == 0 {
			min := left
			for i := left; i <= right; i++ {
				if list[i] < list[min] {
					min = i
				}
			}

			delete(hm, list[min])
			ans = append(ans, list[min])

			for i := min + 1; i <= right; i++ {
				hm[list[i]]++
			}

			tmp := make([]byte, 0)
			for i := min + 1; i <= len(list)-1; i++ {
				if list[i] != list[min] {
					tmp = append(tmp, list[i])
				} else if i <= right {
					right--
				}
			}

			right = 0
			list = tmp
			left = 0
			continue
		}

		right++
	}

	return string(ans)
}
