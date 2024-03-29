package list

import (
	"fmt"
)

func CirclePrint(matrix [][]int) {
	topVertical, downVertical := 0, len(matrix)-1
	leftLevel, rightLevel := 0, len(matrix[0])-1

	for leftLevel <= rightLevel && topVertical <= downVertical {
		for i := leftLevel; i < rightLevel; i++ {
			fmt.Println(matrix[topVertical][i])
		}

		for i := topVertical; i < downVertical; i++ {
			fmt.Println(matrix[i][rightLevel])
		}

		for i := rightLevel; i > leftLevel; i-- {
			fmt.Println(matrix[downVertical][i])
		}

		for i := downVertical; i > topVertical; i-- {
			fmt.Println(matrix[i][leftLevel])
		}

		leftLevel++
		rightLevel--
		topVertical++
		downVertical--
	}
}

func TupleNums(nums []int, k int, i, j int) {
	for i < j {
		if nums[i] >= k {
			return
		}
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		value := nums[i] + nums[j]
		if value > k {
			j--
		} else if value < k {
			i++
		} else {
			fmt.Println(nums[i], nums[j])
			i++
			j--
		}
	}
}

func TupleNumV2(nums []int, k int, i, j, start int) {
	for i < j {
		if nums[i] >= k {
			return
		}
		if i > 0 && nums[i] == nums[i-1] {
			i++
			continue
		}

		value := nums[i] + nums[j]
		if value > k {
			j--
		} else if value < k {
			i++
		} else {
			fmt.Println(nums[start], nums[i], nums[j])
			i++
			j--
		}
	}
}

func TriadNums(nums []int, k int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= k {
			return
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		TupleNumV2(nums, k-nums[i], i+1, len(nums)-1, i)
	}
}
