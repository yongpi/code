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

func PickMaxSub(data string) string {
	hm := make(map[byte]int)
	var i, j, count int
	var ans string

	for j < len(data) {
		item := data[j]
		if hm[item] == 0 {
			hm[item]++
			j++
			continue
		}

		if j-i > count {
			ans = data[i:j]
			count = j - i
		}

		for hm[item] >= 1 && i < j {
			hm[data[i]]--
			i++
		}

		j++
	}

	return ans
}

func PickMinContainSub(data, sub string) string {
	sm := make(map[byte]int)
	for i := 0; i < len(sub); i++ {
		item := sub[i]
		sm[item]++
	}

	ans := data
	var i, j int
	hm := make(map[byte]int)
	count := len(sm)

	for j < len(data) {
		item := data[j]
		hm[item]++
		if hm[item] == sm[item] {
			count--
		}

		if count > 0 {
			j++
			continue
		}

		if len(ans) > j-i+1 {
			ans = data[i : j+1]
		}

		for i <= j {
			tmp := data[i]
			hm[tmp]--
			i++
			if hm[tmp] < sm[tmp] {
				count++
				break
			}

			if len(ans) > j-i+1 {
				ans = data[i : j+1]
			}
		}

		j++
	}

	return ans
}
