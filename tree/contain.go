package tree

import "strings"

func Contain(node1, node2 *Tree) bool {
	if node2 == nil {
		return true
	}
	if node1 == nil {
		return false
	}

	return check(node1, node2) || Contain(node1.Left, node2) || Contain(node1.Right, node2)
}

func check(node1, node2 *Tree) bool {
	if node2 == nil {
		return true
	}

	if node1 == nil || node1.Value != node2.Value {
		return false
	}

	return check(node1.Left, node2.Left) && check(node1.Right, node2.Right)
}

func ContainTP(node1, node2 *Tree) bool {
	str1 := Serialization(node1)
	str2 := Serialization(node2)

	return strings.Contains(str1, str2)
}
