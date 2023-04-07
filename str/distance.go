package str

func MinDistance(a, b string, list []string) int {
	ai, bi := -1, -1
	distance := -1

	for i := 0; i <= len(list)-1; i++ {
		if list[i] == a {
			ai = i
			if bi != -1 {
				if distance == -1 {
					distance = i - bi + 1
				} else {
					distance = min(distance, i-bi)

				}
			}
		}

		if list[i] == b {
			bi = i
			if ai != -1 {
				if distance == -1 {
					distance = i - ai + 1
				} else {
					distance = min(distance, i-ai)
				}
			}
		}
	}

	return distance
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
