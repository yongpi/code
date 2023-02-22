package stack

import (
	"fmt"
	"testing"
)

func TestMinStack_GetMin(t *testing.T) {
	stack := &MinStack{}
	stack.Push(10)
	stack.Push(8)
	stack.Push(9)
	stack.Push(3)
	stack.Push(2)
	stack.Push(10)

	fmt.Println(stack.GetMin())
	fmt.Println(stack.Pop())
	fmt.Println(stack.GetMin())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.GetMin())
}
