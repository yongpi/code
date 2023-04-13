package str

func ConvertMinDistance(start, end string, list []string) int {
	next := make(map[string][]string)
	convert := []byte{'a', 'b', 'c'}

	cd := []string{start, end}
	cd = append(cd, list...)

	for _, item := range cd {
		for i := 0; i < len(item); i++ {
			data := []byte(item)
			for _, value := range convert {
				if data[i] != value {
					data[i] = value
					next[item] = append(next[item], string(data))
				}
			}
		}
	}

	exist := make(map[string]struct{})
	exist[start] = struct{}{}
	stack := []string{start}

	distance := make(map[string]int)
	distance[start] = 0

	for len(stack) > 0 {
		value := stack[0]
		stack = stack[1:]
		for _, n := range next[value] {
			if _, ok := exist[n]; ok {
				continue
			}

			distance[n] = distance[value] + 1
			stack = append(stack, n)
		}
	}

	count := 0
	stack = stack[:0]

	for len(stack) > 0 {
		ls := len(stack)
		for i := 0; i < ls; i++ {
			value := stack[i]
			if value == end {
				return count
			}

			for _, n := range next[value] {
				if d, ok := distance[n]; ok && d == count+1 {
					stack = append(stack, n)
				}
			}
		}

		stack = stack[ls:]
	}

	return count
}
