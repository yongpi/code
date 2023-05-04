package tree

import (
	"fmt"
	"testing"
)

func TestMaxSearch(t *testing.T) {
	r1 := &Tree{Value: 6}
	r2 := &Tree{Value: 1}
	r3 := &Tree{Value: 12}
	r4 := &Tree{Value: 0}
	r5 := &Tree{Value: 3}
	r6 := &Tree{Value: 10}
	r7 := &Tree{Value: 13}

	r12 := &Tree{Value: 4}
	r13 := &Tree{Value: 14}
	r14 := &Tree{Value: 20}
	r15 := &Tree{Value: 16}

	r17 := &Tree{Value: 2}
	r18 := &Tree{Value: 5}
	r19 := &Tree{Value: 11}
	r20 := &Tree{Value: 15}

	r1.Left = r2
	r1.Right = r3

	r2.Left = r4
	r2.Right = r5

	r3.Left = r6
	r3.Right = r7

	r6.Left = r12
	r6.Right = r13

	r7.Left = r14
	r7.Right = r15

	r12.Left = r17
	r12.Right = r18

	r13.Left = r19
	r13.Right = r20

	node := MaxSearch(r1)
	fmt.Println(node.Value)

	fmt.Println(MaxSearch2(r1).Value)
}

func TestMaxLenSearch(t *testing.T) {
	r1 := &Tree{Value: 6}
	r2 := &Tree{Value: 1}
	r3 := &Tree{Value: 12}
	r4 := &Tree{Value: 0}
	r5 := &Tree{Value: 3}
	r6 := &Tree{Value: 10}
	r7 := &Tree{Value: 13}

	r12 := &Tree{Value: 4}
	r13 := &Tree{Value: 14}
	r14 := &Tree{Value: 20}
	r15 := &Tree{Value: 16}

	r17 := &Tree{Value: 2}
	r18 := &Tree{Value: 5}
	r19 := &Tree{Value: 11}
	r20 := &Tree{Value: 15}

	r1.Left = r2
	r1.Right = r3

	r2.Left = r4
	r2.Right = r5

	r3.Left = r6
	r3.Right = r7

	r6.Left = r12
	r6.Right = r13

	r7.Left = r14
	r7.Right = r15

	r12.Left = r17
	r12.Right = r18

	r13.Left = r19
	r13.Right = r20

	fmt.Println(MaxLenSearch(r1))
}

func TestIsBTSNormal(t *testing.T) {
	root := BuildPost([]int{2, 1, 3, 6, 5, 7, 4})
	fmt.Println(IsBTSNormal(root))
	fmt.Println(IsBTSMorris(root))

	node2 := &Tree{
		Left: &Tree{
			Left: nil,
			Right: &Tree{
				Value: 8,
			},
			Value: 4,
		},
		Right: &Tree{
			Left: &Tree{
				Value: 9,
			},
			Value: 5,
		},
		Value: 2,
	}

	fmt.Println(IsBTSNormal(node2))
	fmt.Println(IsBTSMorris(node2))
}
