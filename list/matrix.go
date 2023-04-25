package list

func MaxSideArea(matrix [][]int) int {
	n, m := len(matrix), len(matrix[0])
	down, right := make([][]int, n), make([][]int, n)

	for i := 0; i < n; i++ {
		down[i] = make([]int, m)
		right[i] = make([]int, m)
	}

	if matrix[n-1][m-1] == 1 {
		right[n-1][m-1] = 1
		down[n-1][m-1] = 1

	}

	for i := m - 2; i >= 0; i-- {
		if matrix[n-1][i] == 1 {
			right[n-1][i] = right[n-1][i+1] + 1
			down[n-1][i] = 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if matrix[i][m-1] == 1 {
			right[i][m-1] = 1
			down[i][m-1] = down[i+1][m-1]
		}
	}

	for j := m - 2; j >= 0; j-- {
		for i := n - 2; i >= 0; i-- {
			if matrix[i][j] == 1 {
				right[i][j] = right[i][j+1] + 1
				down[i][j] = down[i+1][j] + 1
			}
		}
	}

	var ans int
	for k := Min(n, m); k > 0; k-- {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if right[i][j] == k && down[i][j+k-1] == k && down[i][j] == k && right[i+k-1][j] == k {
					ans = Max(k*k, ans)
				}
			}
		}
	}

	return ans
}

func MinPath(matrix [][]int) int {
	if matrix[0][0] != 1 {
		return -1
	}

	return next(matrix, 0, 0, 1)
}

func next(matrix [][]int, i, j int, path int) int {
	n, m := len(matrix), len(matrix[0])

	if i < 0 || i >= n || j < 0 || j >= m || matrix[i][j] != 1 {
		return -1
	}

	matrix[i][j] = 2
	if i == n-1 && j == m-1 {
		return path
	}

	list := make([]int, 0)
	list = append(list, next(matrix, i+1, j, path+1))
	list = append(list, next(matrix, i-1, j, path+1))
	list = append(list, next(matrix, i, j+1, path+1))
	list = append(list, next(matrix, i, j-1, path+1))

	ans := -1
	for _, value := range list {
		if value == -1 {
			continue
		}

		if ans == -1 {
			ans = value
		} else {
			ans = Min(ans, value)
		}
	}

	return ans
}
