package tree

import (
	"fmt"
	"testing"
)

func TestFindSuccessorNode(t *testing.T) {
	r1 := &Tree{Value: 6}
	r2 := &Tree{Value: 3}
	r3 := &Tree{Value: 9}
	r4 := &Tree{Value: 1}
	r5 := &Tree{Value: 4}
	r6 := &Tree{Value: 8}
	r7 := &Tree{Value: 10}
	r8 := &Tree{Value: 2}
	r9 := &Tree{Value: 5}
	r10 := &Tree{Value: 7}

	r1.Left = r2
	r1.Right = r3

	r2.Parent = r1
	r2.Left = r4
	r2.Right = r5

	r3.Parent = r1
	r3.Left = r6
	r3.Right = r7

	r4.Parent = r2
	r4.Right = r8

	r5.Parent = r2
	r5.Right = r9

	r6.Parent = r3
	r6.Left = r10

	r7.Parent = r3

	r8.Parent = r4

	r9.Parent = r5

	r10.Parent = r6

	fmt.Println(FindSuccessorNode(r9))
	fmt.Println(FindSuccessorNode2(r9))
}
