package tree

import (
	"bytes"
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

func Serialization2(root *Tree) string {
	if root == nil {
		return "#!"
	}

	var ans bytes.Buffer
	ans.WriteString(fmt.Sprintf("%d!", root.Value))
	ans.WriteString(Serialization2(root.Left))
	ans.WriteString(Serialization2(root.Right))

	return ans.String()
}

func Deserialization2(ans string) *Tree {
	if ans == "" {
		return nil
	}

	var buildTree func(list []string) (*Tree, []string)
	buildTree = func(list []string) (*Tree, []string) {
		if len(list) == 0 {
			return nil, list
		}

		item := list[0]
		list = list[1:]
		if item == "#" {
			return nil, list
		}

		value, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		root := &Tree{Value: value}
		root.Left, list = buildTree(list)
		root.Right, list = buildTree(list)

		return root, list
	}

	list := strings.Split(ans, "!")
	root, _ := buildTree(list)
	return root
}
