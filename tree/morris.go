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
			for mostRight.Right != nil && mostRight.Right != cur {
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
		if cur != nil {
			fmt.Println(cur.Value)
		}
	}
}

func MorrisPre(root *Tree) {
	if root == nil {
		return
	}

	cur := root
	var mostRight *Tree

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				fmt.Println(cur.Value)
				cur = cur.Left
				continue
			}

			mostRight.Right = nil
		} else {
			fmt.Println(cur.Value)
		}

		cur = cur.Right
	}
}

func MorrisMid(root *Tree) {
	if root == nil {
		return
	}

	cur := root
	var mostRight *Tree

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			}

			fmt.Println(cur.Value)
			mostRight.Right = nil
		} else {
			fmt.Println(cur.Value)
		}

		cur = cur.Right
	}
}
