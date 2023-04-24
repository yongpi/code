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
