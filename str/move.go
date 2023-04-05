package str

func Move(data string) string {
	j := len(data) - 1
	list := []byte(data)
	for i := len(list) - 1; i >= 0; i-- {
		if list[i] != '*' {
			list[j] = list[i]
			j--
		}
	}

	for j >= 0 {
		list[j] = '*'
		j--
	}

	return string(list)
}

func Reverse(data string) string {
	list := []byte(data)
	start, end := 0, len(list)-1
	for start < end {
		list[end], list[start] = list[start], list[end]
		end--
		start++
	}

	return string(list)
}
