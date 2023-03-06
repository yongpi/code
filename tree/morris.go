package tree

import "fmt"

func Morris(root *Tree) {
	if root == nil {
		return
	}

	cur := root
	fmt.Println(cur.Value)

	for cur != nil {
		mostRight := cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			}

			mostRight.Right = nil
		}

		cur = cur.Right
		fmt.Println(cur.Value)
	}
}
