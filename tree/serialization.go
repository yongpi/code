package tree

import (
	"fmt"
	"strconv"
	"strings"
)

func Serialization(root *Tree) string {
	if root == nil {
		return "#!"
	}

	var ans string
	ans = fmt.Sprintf("%d!", root.Value)
	ans += Serialization(root.Left)
	ans += Serialization(root.Right)

	return ans
}

func Deserialization(ans string) *Tree {
	if ans == "" {
		return nil
	}

	stack := make([]string, 0)
	for _, value := range strings.Split(ans, "!") {
		stack = append(stack, value)
	}

	root, _ := makeTree(stack)
	return root
}

func makeTree(stack []string) (*Tree, []string) {
	item := stack[0]
	stack = stack[1:]

	if item == "#" {
		return nil, stack
	}

	value, err := strconv.Atoi(item)
	if err != nil {
		panic(err)
	}
	node := &Tree{Value: value}

	node.Left, stack = makeTree(stack)
	node.Right, stack = makeTree(stack)

	return node, stack
}
