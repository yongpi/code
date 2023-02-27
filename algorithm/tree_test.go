package algorithm

import "testing"

var orderTree = &TreeNode{
	Value: 1,
	Left: &TreeNode{
		Value: 2,
		Left: &TreeNode{
			Value: 3,
		},
		Right: &TreeNode{
			Value: 4,
		},
	},
	Right: &TreeNode{
		Value: 5,
		Left: &TreeNode{
			Value: 6,
		},
		Right: &TreeNode{
			Value: 7,
		},
	},
}

func TestPreOrder(t *testing.T) {
	PreOrder(orderTree)
}

func TestPreOrder2(t *testing.T) {
	PreOrder2(orderTree)
}

func TestInOrder(t *testing.T) {
	InOrder(orderTree)
}

func TestInOrder2(t *testing.T) {
	InOrder2(orderTree)
}

func TestPostOrder(t *testing.T) {
	PostOrder(orderTree)
}

func TestPostOrder2(t *testing.T) {
	PostOrder2(orderTree)
}

func TestLevelOrder(t *testing.T) {
	LevelOrder(orderTree)
}
