package str

func BinaryNullStr(list []string, val string) int {
	left, right := 0, len(list)-1
	res := -1

	for left <= right {
		mid := left + (right-left)/2

		if list[mid] == val {
			res = mid
			right = mid - 1
			continue
		}

		if list[mid] == "" {
			i := mid
			for i >= left && list[i] == "" {
				i--
			}

			if i == left-1 {
				left = mid + 1
				continue
			}

			if list[i] == val {
				res = i
				right = i - 1
				continue
			}

			if list[i] > val {
				right = i - 1
				continue
			}

			if list[i] < val {
				left = i + 1
				continue
			}
		}

		if list[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
