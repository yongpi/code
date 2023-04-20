package list

import "fmt"

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
