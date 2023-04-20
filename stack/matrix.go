package stack

func MaxMatrixArea(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	height := make([]int, len(matrix[0]))
	var area int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				height[j] = height[j] + 1
			} else {
				height[j] = 0
			}
		}

		area = max(area, maxArea(height))
	}

	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxArea(nums []int) int {
	stack := make([]int, 0)
	var area int

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[i] >= nums[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}

		for nums[i] < nums[stack[len(stack)-1]] {
			vi := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(stack) == 0 {
				area = max(area, i*nums[vi])
				break
			}

			pre := stack[len(stack)-1]
			area = max(area, (i-pre-1)*nums[vi])
		}
		stack = append(stack, i)
	}

	for {
		vi := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) == 0 {
			area = max(area, len(nums)*nums[vi])
			break
		}

		pre := stack[len(stack)-1]
		area = max(area, (len(nums)-pre-1)*nums[vi])
	}

	return area
}

func IslandsNum(islands [][]int) int {
	var ans int
	for i := 0; i < len(islands); i++ {
		for j := 0; j < len(islands[i]); j++ {
			if islands[i][j] == 1 {
				ans++
				Influence(i, j, islands)
			}
		}
	}

	return ans
}

func Influence(i, j int, islands [][]int) {
	if i >= len(islands) || j >= len(islands[0]) || i < 0 || j < 0 || islands[i][j] != 1 {
		return
	}

	islands[i][j] = 2
	Influence(i+1, j, islands)
	Influence(i-1, j, islands)
	Influence(i, j+1, islands)
	Influence(i, j-1, islands)
}
