package tree

import "math"

type MaxSearchReturnType struct {
	MaxNode                     *Tree
	MaxSize, MaxValue, MinValue int
}

func MaxSearch(root *Tree) *Tree {
	return process(root).MaxNode
}

func process(root *Tree) MaxSearchReturnType {
	if root == nil {
		return MaxSearchReturnType{
			MaxNode:  nil,
			MaxValue: math.MinInt32,
			MinValue: math.MaxInt32,
		}
	}

	left := process(root.Left)
	right := process(root.Right)

	size := max(left.MaxSize, right.MaxSize)
	node := left.MaxNode
	if left.MaxSize < right.MaxSize {
		node = right.MaxNode
	}

	minValue := min(root.Value, min(left.MinValue, right.MinValue))
	maxValue := max(root.Value, max(left.MaxValue, right.MaxValue))

	if root.Value > left.MaxValue && root.Value < right.MinValue && root.Left == left.MaxNode && root.Right == right.MaxNode {
		node = root
		size = left.MaxSize + right.MaxSize + 1
	}

	return MaxSearchReturnType{
		MaxNode:  node,
		MaxSize:  size,
		MaxValue: maxValue,
		MinValue: minValue,
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func MaxLenSearch(root *Tree) int {
	if root == nil {
		return 0
	}

	maxLen := MaxTop(root, root)
	maxLen = max(MaxLenSearch(root.Left), maxLen)
	maxLen = max(MaxLenSearch(root.Right), maxLen)

	return maxLen
}

func MaxTop(head, node *Tree) int {
	if head != nil && node != nil && isBTS(head, node, node.Value) {
		return MaxTop(head, node.Left) + MaxTop(head, node.Right) + 1
	}

	return 0
}

func isBTS(head, node *Tree, value int) bool {
	if head == nil {
		return false
	}

	if head == node {
		return true
	}

	if value > head.Value {
		return isBTS(head.Right, node, value)
	}

	return isBTS(head.Left, node, value)
}

func IsBTSNormal(root *Tree) bool {
	var pre *Tree
	stack := make([]*Tree, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pre == nil {
			pre = node
		} else {
			if pre.Value > node.Value {
				return false
			}
		}

		cur = node.Right
	}

	return true
}

func IsBTSMorris(root *Tree) bool {
	cur := root
	var mostRight, pre *Tree

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

			mostRight.Right = nil
		}

		if pre == nil {
			pre = cur
		} else {
			if pre.Value > cur.Value {
				return false
			}
		}

		cur = cur.Right
	}

	return true
}

func MaxSearch2(root *Tree) *Tree {
	var getMaxNode func(node *Tree) *MaxSearchReturnType
	getMaxNode = func(node *Tree) *MaxSearchReturnType {
		if node == nil {
			return &MaxSearchReturnType{
				MaxNode:  nil,
				MaxSize:  0,
				MaxValue: math.MinInt32,
				MinValue: math.MaxInt32,
			}
		}

		left := getMaxNode(node.Left)
		right := getMaxNode(node.Right)

		minValue := min(min(left.MinValue, right.MinValue), node.Value)
		maxValue := max(max(left.MaxValue, right.MaxValue), node.Value)
		maxNode := left.MaxNode
		maxSize := left.MaxSize

		if right.MaxSize > left.MaxSize {
			maxNode = right.MaxNode
			maxSize = right.MaxSize
		}

		if node.Value > left.MaxValue && node.Value < right.MinValue && left.MaxNode == node.Left && right.MaxNode == node.Right {
			maxNode = node
			maxSize = left.MaxSize + right.MaxSize + 1
		}

		return &MaxSearchReturnType{
			MaxNode:  maxNode,
			MaxSize:  maxSize,
			MaxValue: maxValue,
			MinValue: minValue,
		}
	}

	return getMaxNode(root).MaxNode
}

func NodeInBST(root, node *Tree) bool {
	if root == nil {
		return false
	}

	if root == node {
		return true
	}

	if root.Value < node.Value {
		return NodeInBST(root.Right, node)
	}

	return NodeInBST(root.Left, node)
}

func BSTLen(root *Tree) int {
	if root == nil {
		return 0
	}

	var count int
	stack := make([]*Tree, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if NodeInBST(root, node) {
			count++
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}

	return count
}

func MaxLenSearch2(root *Tree) int {
	if root == nil {
		return 0
	}

	count := BSTLen(root)
	count = max(MaxLenSearch2(root.Left), count)
	count = max(MaxLenSearch2(root.Right), count)

	return count

}
